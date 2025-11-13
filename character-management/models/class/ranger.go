package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Ranger struct {
	Archetype            string                `json:"archetype"`
	FightingStyle        string                `json:"fighting-style"`
	FightingStyleFeature FightingStyleFeature  `json:"-"`
	FavoredEnemies       []string              `json:"favored-enemies"`
	OtherFeatures        []models.ClassFeature `json:"other-features"`
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

	if r.FightingStyleFeature.Name != "" && c.Level >= 2 {
		appliedText := "Requirements for fighting style not met."
		if r.FightingStyleFeature.IsApplied {
			appliedText = "Requirements for this fighting style are met, and any bonuses to armor or weapons have been applied to your character."
		}

		fightingStyleHeader := fmt.Sprintf("**Fighting Style**: *%s*\n", r.FightingStyleFeature.Name)
		fightingStyleDetail := fmt.Sprintf("%s\n%s\n\n", r.FightingStyleFeature.Details, appliedText)
		s = append(s, fightingStyleHeader)
		s = append(s, fightingStyleDetail)
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

	switch r.FightingStyle {
	case types.FightingStyleArchery:
		r.FightingStyleFeature = applyArchery(c)
	case types.FightingStyleDefense:
		r.FightingStyleFeature = applyDefense(c)
	case types.FightingStyleDueling:
		r.FightingStyleFeature = applyDueling(c)
	case types.FightingStyleTwoWeaponFighting:
		r.FightingStyleFeature = applyTwoWeaponFighting(c)
	default:
		logger.HandleInfo(invalidMsg)
	}
}

func (r *Ranger) AddFightingStyleFeature(feature models.ClassFeature) {

}

func (r *Ranger) RemoveFightingStyleFeature(feature models.ClassFeature) {

}

// CLI

func (r *Ranger) UseClassTokens(tokenName string, quantity int) {
	// Not sure Rangers have a token like system to implement
	logger.HandleInfo("No token set up for Ranger class")
}

func (r *Ranger) RecoverClassTokens(tokenName string, quantity int) {
	// Not sure Rangers have a token like system to implement
	logger.HandleInfo("No token set up for Ranger class")
}

func (r *Ranger) GetTokens() []string {
	return []string{}
}
