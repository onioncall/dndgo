package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Ranger struct {
	Archetype      string                 `json:"archetype"`
	FightingStyle  string                 `json:"fighting-style"`
	FavoredEnemies []string               `json:"favored-enemies"`
	OtherFeatures  []models.ClassFeatures `json:"other-features"`
}

func LoadRanger(data []byte) (*Ranger, error) {
	var ranger Ranger
	if err := json.Unmarshal(data, &ranger); err != nil {
		return &ranger, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &ranger, nil
}

func (r *Ranger) ValidateMethods(c *models.Character) {
}

// func (r *Ranger) ExecutePostCalculateMethods(c *models.Character) {
// 	r.PostCalculateFightingStyle(c)
// }
//
// func (r *Ranger) ExecutePreCalculateMethods(c *models.Character) {
// }

func (r *Ranger) PostCalculateSpellAttackMod(c *models.Character) {
	wisMod := c.GetMod(types.AbilityWisdom)

	executeSpellSaveDC(c, wisMod)
	executeSpellAttackMod(c, wisMod)
}

func (r *Ranger) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if r.Archetype != "" && c.Level > 3 {
		archetypeHeader := fmt.Sprintf("Archetype: *%s*\n\n", r.Archetype)
		s = append(s, archetypeHeader)
	}

	if r.FightingStyle != "" && c.Level >= 2 {
		fightingStyleHeader := fmt.Sprintf("Fighting Style: *%s*\n\n", r.FightingStyle)
		s = append(s, fightingStyleHeader)
	}

	if len(r.FavoredEnemies) > 0 {
		favoredEnemyHeader := fmt.Sprintf("Favored Enemies:\n")
		s = append(s, favoredEnemyHeader)

		for _, enemy := range r.FavoredEnemies {
			enemyLine := fmt.Sprintf("- %s\n", enemy)
			s = append(s, enemyLine)
		}
		s = append(s, "\n")
	}

	if len(r.OtherFeatures) > 0 {
		for _, detail := range r.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			name := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, name)
			detail := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, detail)
		}
	}

	return s
}

// At level 2, Rangers adopt a fighting style as their specialty
// only one of these styles can be selected
func (r *Ranger) PostCalculateFightingStyle(c *models.Character) {
	if c.Level < 2 {
		return
	}

	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles, %s, %s, %s, %s",
		r.FightingStyle,
		types.FightingStyleArchery,
		types.FightingStyleDefense,
		types.FightingStyleDueling,
		types.FightingStyleTwoWeaponFighting)

	fightingStyleApplied := false
	// There are more fighting styles, but these are the four available to rangers
	switch r.FightingStyle {
	case types.FightingStyleArchery:
		fightingStyleApplied = applyArchery(c)
	case types.FightingStyleDefense:
		fightingStyleApplied = applyDefense(c)
	case types.FightingStyleDueling:
		fightingStyleApplied = applyDueling(c)
	case types.FightingStyleTwoWeaponFighting:
		fightingStyleApplied = applyTwoWeaponFighting(c)
	default:
		logger.HandleInfo(invalidMsg)
	}

	if !fightingStyleApplied {
		// TODO: in the methods for each fighting style, log more specific details for why
		// the fight style bonus was not appied to the class
		logger.HandleInfo(fmt.Sprintf("Fighting Style bonus '%s' was not applied", r.FightingStyle))
	}
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
	if c.WornEquipment.Armor == "" {
		c.AC += 1
		return true
	}

	return false
}

// Only to be called from executeFightingStyle
func applyDueling(c *models.Character) bool {
	// this is less defined, since it depends on us knowing what weapons are currently
	// in the characters hand. We'll assume that they have which ever weapon they want
	// this applied to to be the first one they have, and that it will be equipped in combat.
	// Maybe later we'll come up with a flag for the weapon being in use or something
	for i, weapon := range c.Weapons {
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

	secondaryWeaponIndex := -1
	primaryWeaponIndex := -1

	for i, weapon := range c.Weapons {
		isLight := false
		isOneHanded := true

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

		// We'll take the first secondary weapon that meets these criteria
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

// CLI

func (r *Ranger) UseClassTokens(tokenName string) {
	// Not sure Rangers have a token like system to implement
	logger.HandleInfo("No token set up for Ranger class")
}

func (r *Ranger) RecoverClassTokens(tokenName string, quantity int) {
	// Not sure Rangers have a token like system to implement
	logger.HandleInfo("No token set up for Ranger class")
}
