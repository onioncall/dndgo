package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Ranger struct {
	Archetype      string                `json:"archetype"`
	FightingStyle  string                `json:"fighting-style"`
	FavoredEnemies []string              `json:"favored-enemies"`
	OtherFeatures  []models.ClassFeature `json:"other-features"`
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

func (r *Ranger) ExecutePostCalculateMethods(c *models.Character) {
	r.executeFightingStyle(c)
}

func (r *Ranger) ExecutePreCalculateMethods(c *models.Character) {
}

func (r *Ranger) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd10", level)
}

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
func (r *Ranger) executeFightingStyle(c *models.Character) {
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

// CLI

func (r *Ranger) UseClassTokens(tokenName string) {
	// Not sure Rangers have a token like system to implement
	logger.HandleInfo("No token set up for Ranger class")
}

func (r *Ranger) RecoverClassTokens(tokenName string, quantity int) {
	// Not sure Rangers have a token like system to implement
	logger.HandleInfo("No token set up for Ranger class")
}
