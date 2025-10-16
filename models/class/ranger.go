package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
)

type Ranger struct {
	Archetype					string							`json:"archetype"`
	FightingStyle				string							`json:"fighting-style"`
	FavoredEnemies				[]string						`json:"favored-enemies"`
	OtherFeatures 				[]models.ClassFeatures			`json:"other-features"`
}

// Fighting Styles
const (
	Archery				string = "archery"
	Defense				string = "defense"
	Dueling				string = "dueling"
	TwoWeaponFighting	string = "two-weapon-fighting"
)

// Weapon Ranges
const (
	Ranged		string = "ranged"
	Melee		string = "melee"
)

func LoadRanger(data []byte) (*Ranger, error) {
	var ranger Ranger
	if err := json.Unmarshal(data, &ranger); err != nil {
		errLog := fmt.Errorf("Failed to parse class data: %w", err)
		logger.HandleError(errLog, err)

		return nil, err
	}

	return &ranger, nil
}

func (r *Ranger) LoadMethods() {
}

func (r *Ranger) ExecutePostCalculateMethods(c *models.Character) {
	models.PostCalculateMethods = append(models.PostCalculateMethods, r.executeFightingStyle)
	for _, m := range models.PostCalculateMethods {
		m(c)
	}
}

func (r *Ranger) ExecutePreCalculateMethods(c *models.Character) {
	for _, m := range models.PreCalculateMethods {
		m(c)
	}
}

func (r *Ranger) PrintClassDetails(c *models.Character) []string { 
	s := c.BuildClassDetailsHeader()
	
	if r.Archetype != "" && c.Level > 3 {
		archetypeHeader := fmt.Sprintf("Archetype: *%s*\n\n", r.Archetype)
		s = append(s, archetypeHeader)
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
func (r *Ranger) executeFightingStyle(c *models.Character) {
	if c.Level < 2 {
		return 
	}

	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles, %s, %s, %s, %s", 
		r.FightingStyle,
		Archery,
		Defense,
		Dueling,
		TwoWeaponFighting)

	switch r.FightingStyle {
	case Archery: executeArchery(c)
	case Defense: executeDefense(c)
	case Dueling: executeDueling(c)
	case TwoWeaponFighting: executeTwoWeaponFighting(c)
	default: logger.HandleInfo(invalidMsg)
	}
}

func executeArchery(c *models.Character) {
	for i, weapon := range c.Weapons {
		if strings.ToLower(weapon.Range) == Ranged {
			c.Weapons[i].Bonus += 2
		}
	}
}

func executeDefense(c *models.Character) {
	if c.BodyEquipment.Armour != "" {
		c.AC += 1
	}
}

func executeDueling(c *models.Character) {
	// this is less defined, since it depends on character actions we can't know, but we will 
	// assume that if someone has this fighting style that they aren't dual weilding or something
	for i, weapon := range c.Weapons {
		if strings.ToLower(weapon.Range) == Melee {
			c.Weapons[i].Bonus += 2
		}
	}
}

func executeTwoWeaponFighting(c *models.Character) {
	// This is a little wonky to implement becuase we don't have the concept of multiple attacks
	// so as a middle ground the solution will be to add a duplicate weapon if you have two, and then 
	// this bonus will only be applied to one while duel weilding
    weaponCounts := make(map[string]int)
    
	for _, weapon := range c.Weapons {
        weaponCounts[weapon.Name]++
	}

	for i, weapon := range c.Weapons {
		for _, prop := range weapon.Properties {
			// weapon must be light to dual weild
			if strings.ToLower(prop) == "light" && weaponCounts[weapon.Name] > 1 {
				for _, mod := range c.Abilities {
					if strings.ToLower(mod.Name) == models.Dexterity {
						c.Weapons[i].Bonus += mod.AbilityModifier
						// we only want to apply this once. In text,
						// its's the second weapon, but we'll do it for the first one
						return
					}
				}
			}
		}
	}
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
