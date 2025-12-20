package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Ranger struct {
	models.BaseClass
	FightingStyle        string               `json:"fighting-style"`
	FightingStyleFeature FightingStyleFeature `json:"-"`
	FavoredEnemies       []string             `json:"favored-enemies"`
}

func LoadRanger(data []byte) (*Ranger, error) {
	var ranger Ranger
	if err := json.Unmarshal(data, &ranger); err != nil {
		return &ranger, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &ranger, nil
}

func (r *Ranger) ExecutePostCalculateMethods(c *models.Character) {
	r.executeFightingStyle(c)
}

func (r *Ranger) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd10", level)
}

func (r *Ranger) PostCalculateSpellAttackMod(c *models.Character) {
	wisMod := c.GetMod(shared.AbilityWisdom)

	executeSpellSaveDC(c, wisMod)
	executeSpellAttackMod(c, wisMod)
}

func (r *Ranger) ClassDetails(level int) string {
	var s string
	if r.FightingStyleFeature.Name != "" && level >= 2 {
		appliedText := "Requirements for fighting style not met."
		if r.FightingStyleFeature.IsApplied {
			appliedText = "Requirements for this fighting style are met, and any bonuses to armor or weapons have been applied to your character."
		}

		fightingStyleHeader := fmt.Sprintf("**Fighting Style**: *%s*\n", r.FightingStyleFeature.Name)
		fightingStyleDetail := fmt.Sprintf("%s\n%s\n\n", r.FightingStyleFeature.Details, appliedText)
		s += fightingStyleHeader
		s += fightingStyleDetail
	}

	if len(r.FavoredEnemies) > 0 {
		favoredEnemyHeader := fmt.Sprintf("Favored Enemies:\n")
		s += favoredEnemyHeader

		for _, enemy := range r.FavoredEnemies {
			enemyLine := fmt.Sprintf("- %s\n", enemy)
			s += enemyLine
		}
		s += "\n"
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
		shared.FightingStyleArchery,
		shared.FightingStyleDefense,
		shared.FightingStyleDueling,
		shared.FightingStyleTwoWeaponFighting)

	switch r.FightingStyle {
	case shared.FightingStyleArchery:
		r.FightingStyleFeature = applyArchery(c)
	case shared.FightingStyleDefense:
		r.FightingStyleFeature = applyDefense(c)
	case shared.FightingStyleDueling:
		r.FightingStyleFeature = applyDueling(c)
	case shared.FightingStyleTwoWeaponFighting:
		r.FightingStyleFeature = applyTwoWeaponFighting(c)
	default:
		logger.Info(invalidMsg)
	}
}
