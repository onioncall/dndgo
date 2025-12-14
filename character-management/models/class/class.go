package class

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

func executeExpertiseShared(c *models.Character, expertiseSkills []string) {
	if c.Level < 3 {
		return
	}

	if c.Level < 10 && len(expertiseSkills) > 2 {
		// We'll allow the user to specify more, but only the first two get taken for it to be legal
		expertiseSkills = expertiseSkills[:2]
	}

	if c.Level > 10 && len(expertiseSkills) > 4 {
		// We'll allow the user to specify more, but only the first four get taken for it to be legal
		expertiseSkills = expertiseSkills[:4]
	}

	seen := make(map[string]bool)
	for _, profToDouble := range expertiseSkills {
		if seen[profToDouble] == true {
			logger.Info("Bard Config Error - Expertise can not have dupliate proficiencies")
			return
		}
		seen[profToDouble] = true

		for i, cs := range c.Skills {
			if strings.ToLower(cs.Name) == strings.ToLower(profToDouble) {
				c.Skills[i].SkillModifier += c.Proficiency
			}
		}
	}
}

// If not wearing armor, Armor Class is boosted to 10 + modifiers outlined by implementing class
func executeUnarmoredDefenseShared(c *models.Character, abilities []string) {
	if c.WornEquipment.Armor.Name != "" {
		return
	}

	c.AC = 0

	for _, charAbility := range c.Abilities {
		for _, classAbility := range abilities {
			if charAbility.Name == classAbility {
				c.AC += charAbility.AbilityModifier
			}
		}
	}

	c.AC += 10
}

func executePreparedSpellsShared(c *models.Character, preparedSpells []string) {
	for i := range c.Spells {
		for _, ps := range preparedSpells {
			if strings.ToLower(ps) == strings.ToLower(c.Spells[i].Name) {
				c.Spells[i].IsPrepared = true
			}
		}
	}
}

func executeSpellSaveDC(c *models.Character, abilityMod int) {
	c.SpellSaveDC = 8 + c.Proficiency + abilityMod
}

func executeSpellAttackMod(c *models.Character, abilityMod int) {
	c.SpellAttackMod = c.Proficiency + abilityMod
}

// Applies bonus for fighting style, and returns feature with details and weather or not the feature was applied
func applyArchery(c *models.Character) FightingStyleFeature {
	feature := FightingStyleFeature{
		Name:      "Archery",
		Details:   "You gain a +2 bonus to attack rolls you make with ranged weapons.\n",
		IsApplied: false,
	}

	for i, weapon := range c.Weapons {
		if weapon.Ranged {
			c.Weapons[i].Bonus += 2
			feature.IsApplied = true
			break
		}
	}

	return feature
}

// Applies bonus for fighting style, and returns feature with details and weather or not the feature was applied
func applyDefense(c *models.Character) FightingStyleFeature {
	feature := FightingStyleFeature{
		Name:      "Defense",
		Details:   "While you are wearing armor, you gain a +1 bonus to AC.\n",
		IsApplied: false,
	}

	if c.WornEquipment.Armor.Name == "" {
		c.AC += 1
		feature.IsApplied = true
	}

	return feature
}

// Applies bonus for fighting style, and returns feature with details and weather or not the feature was applied
func applyDueling(c *models.Character) FightingStyleFeature {
	feature := FightingStyleFeature{
		Name:      "Dueling",
		Details:   "When you are wielding a melee weapon in one hand and no other weapons, you gain a +2 bonus to damage rolls with that weapon.\n",
		IsApplied: false,
	}

	// If both primary and secondary are equipped, or neither are equipped, this won't apply
	if c.PrimaryEquipped != "" && c.SecondaryEquipped != "" {
		return feature
	} else if c.PrimaryEquipped == "" && c.SecondaryEquipped == "" {
		return feature
	}

	for i, weapon := range c.Weapons {
		if strings.ToLower(c.PrimaryEquipped+c.SecondaryEquipped) != strings.ToLower(weapon.Name) {
			continue
		}

		if weapon.Ranged {
			continue
		}

		isTwoHanded := false

		for _, prop := range weapon.Properties {
			if prop == shared.WeaponPropertyTwoHanded {
				isTwoHanded = true
			}
		}

		if !isTwoHanded {
			c.Weapons[i].Bonus += 2
			feature.IsApplied = true
			break
		}
	}

	return feature
}

// Applies bonus for fighting style, and returns feature with details and weather or not the feature was applied
func applyTwoWeaponFighting(c *models.Character) FightingStyleFeature {
	// This is somewhat nuanced, so I'm just going to document how this works for clarity
	// When dual wielding, there is a primary weapon that must be one handed, and an off hand
	// weapon that must be one handed and have the "light" property. The primary weapon can be
	// light, but does not have to be. The dexterity bonus while dual weilding applies to the
	// off hand weapon. For our purposes, we'll consider the first weapon that meets both
	// criteria the "secondary" weapon, and the the next one to meet just the one handed criteria
	// the primary. We'll only apply the bonus if both primary and secondary weapons are found
	feature := FightingStyleFeature{
		Name:      "Two Weapon Fighting",
		Details:   "When you engage in two-weapon fighting, you can add your ability modifier to the damage of the second attack.\n",
		IsApplied: false,
	}

	// If one of these isn't equipped no need to add the bonus
	if c.PrimaryEquipped == "" || c.SecondaryEquipped == "" {
		return feature
	}

	secondaryWeaponIndex := -1
	primaryWeaponIndex := -1

	for i, weapon := range c.Weapons {
		isLight := false
		isOneHanded := true

		if strings.ToLower(c.PrimaryEquipped) != strings.ToLower(weapon.Name) &&
			strings.ToLower(c.SecondaryEquipped) != strings.ToLower(weapon.Name) {
			continue
		}

		for _, prop := range weapon.Properties {
			if strings.ToLower(prop) == shared.WeaponPropertyTwoHanded {
				isOneHanded = false
				break
			}
			if strings.ToLower(prop) == shared.WeaponPropertyLight {
				isLight = true
			}
		}

		if !isOneHanded {
			continue
		}

		// We'll take the first secondary weapon that meets these criteria since I'm not sure
		// that primary vs secondary weapons really matter for this distinction as long
		// as the requirements are met
		if isOneHanded && isLight && secondaryWeaponIndex == -1 {
			secondaryWeaponIndex = i
			if primaryWeaponIndex != -1 {
				break
			}
			continue
		}

		// The first weapon that meets this criteria will be considered the primary
		if isOneHanded {
			primaryWeaponIndex = i
		}

		// Once both are set we don't need to continue looping over weapons
		if primaryWeaponIndex != -1 && secondaryWeaponIndex != -1 {
			break
		}
	}

	if primaryWeaponIndex == -1 || secondaryWeaponIndex == -1 {
		return feature
	}

	for _, mod := range c.Abilities {
		if strings.ToLower(mod.Name) == shared.AbilityDexterity {
			c.Weapons[secondaryWeaponIndex].Bonus += mod.AbilityModifier
			feature.IsApplied = true
			break
		}
	}

	return feature
}

// Applies bonus for fighting style, and returns feature with details and weather or not the feature was applied
func applyGreatWeaponFighting(c *models.Character) FightingStyleFeature {
	feature := FightingStyleFeature{
		Name:      "Great Weapon Fighting",
		Details:   "When you roll a 1 or 2 on a damage die for an attack you make with a melee weapon that you are wielding with two hands, you can reroll the die and must use the new roll, even if the new roll is a 1 or a 2. The weapon must have the two-handed or versatile property for you to gain this benefit.\n",
		IsApplied: false,
	}

	for _, weapon := range c.Weapons {
		if weapon.Name == c.PrimaryEquipped || weapon.Name == c.SecondaryEquipped {
			for _, prop := range weapon.Properties {
				if prop == shared.WeaponPropertyTwoHanded {
					feature.IsApplied = true
				}
			}
		}
	}

	return feature
}

func applyProtection(c *models.Character) FightingStyleFeature {
	feature := FightingStyleFeature{
		Name:      "Protection",
		Details:   "When a creature you can see attacks a target other than you that is within 5 feet of you, you can use your reaction to impose disadvantage on the attack roll. You must be wielding a shield.\n",
		IsApplied: false,
	}

	if c.WornEquipment.Shield == c.PrimaryEquipped || c.WornEquipment.Shield == c.SecondaryEquipped {
		feature.IsApplied = true
	}

	return feature
}

// this is only used for classes that have mutliple tokens types to implement
func fullTokenRecovery(tokens []shared.NamedToken) {
	for i := range tokens {
		tokens[i].Available = tokens[i].Maximum
	}
}

func formatTokens(token shared.NamedToken, tokenName string, level int) string {
	var s string

	if token.Maximum > 0 && token.Name == tokenName && level >= token.Level {
		slots := models.GetSlots(token.Available, token.Maximum)
		s += fmt.Sprintf("%s: %s\n", tokenName, slots)
	}

	return s
}

func formatOtherFeatures(features []models.ClassFeature, level int) string {
	var s string
	if len(features) > 0 {
		for _, feature := range features {
			if feature.Level > level {
				continue
			}

			featureName := fmt.Sprintf("---\n**%s**\n", feature.Name)
			s += featureName
			features := fmt.Sprintf("%s\n", feature.Details)
			s += features
		}
	}

	return s
}

// this is only used for classes that have mutliple tokens types to implement
func getToken(tokenName string, tokens []shared.NamedToken) *shared.NamedToken {
	for i := range tokens {
		if tokens[i].Name == tokenName {
			return &tokens[i]
		}
	}

	return nil
}
