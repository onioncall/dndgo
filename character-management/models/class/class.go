package class

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
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
			logger.HandleInfo("Bard Config Error - Expertise can not have dupliate proficiencies")
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

func buildClassDetailsHeader() []string {
	s := make([]string, 0, 100)
	header := fmt.Sprintf("Class Details\n")
	spacer := fmt.Sprintf("---\n")
	s = append(s, header)
	s = append(s, spacer)

	return s
}

// Only to be called from executeFightingStyle
func applyArchery(c *models.Character) bool {
	for i, weapon := range c.Weapons {
		if weapon.Ranged {
			c.Weapons[i].Bonus += 2
			return true
		}
	}

	return false
}

// Only to be called from executeFightingStyle
func applyDefense(c *models.Character) bool {
	if c.WornEquipment.Armor.Name == "" {
		c.AC += 1
		return true
	}

	return false
}

// Only to be called from executeFightingStyle
func applyDueling(c *models.Character) bool {
	// If both primary and secondary are equipped, or neither are equipped, this won't apply
	if c.PrimaryEquipped != "" && c.SecondaryEquipped != "" {
		return false
	} else if c.PrimaryEquipped == "" && c.SecondaryEquipped == "" {
		return false
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
			if prop == types.WeaponPropertyTwoHanded {
				isTwoHanded = true
			}
		}

		if !isTwoHanded {
			c.Weapons[i].Bonus += 2
			return true
		}
	}

	return false
}

// Only to be called from executeFightingStyle
func applyTwoWeaponFighting(c *models.Character) bool {
	// This is somewhat nuanced, so I'm just going to document how this works for clarity
	// When dual wielding, there is a primary weapon that must be one handed, and an off hand
	// weapon that must be one handed and have the "light" property. The primary weapon can be
	// light, but does not have to be. The dexterity bonus while dual weilding applies to the
	// off hand weapon. For our purposes, we'll consider the first weapon that meets both
	// criteria the "secondary" weapon, and the the next one to meet just the one handed criteria
	// the primary. We'll only apply the bonus if both primary and secondary weapons are found

	// If one of these isn't equipped no need to add the bonus
	if c.PrimaryEquipped == "" || c.SecondaryEquipped == "" {
		return false
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
			if strings.ToLower(prop) == types.WeaponPropertyTwoHanded {
				isOneHanded = false
				break
			}
			if strings.ToLower(prop) == types.WeaponPropertyLight {
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
		return false
	}

	for _, mod := range c.Abilities {
		if strings.ToLower(mod.Name) == types.AbilityDexterity {
			c.Weapons[secondaryWeaponIndex].Bonus += mod.AbilityModifier
			return true
		}
	}

	return false
}

// Add one to AC if the character does not have heavy armor or a shield equipped
func applyMariner(c *models.Character) bool {
	if c.WornEquipment.Armor.Type == types.HeavyArmor ||
		c.WornEquipment.Shield == c.PrimaryEquipped ||
		c.WornEquipment.Shield == c.SecondaryEquipped {
		return false
	}

	c.AC += 1

	return true
}
