package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Fighter struct {
	Archetype            string                `json:"archetype"`
	FightingStyle        string                `json:"fighting-style"`
	FightingStyleFeature FightingStyleFeature  `json:"-"`
	OtherFeatures        []models.ClassFeature `json:"other-features"`
	ActionSurge          types.NamedToken      `json:"action-surge"`
	SecondWind           types.NamedToken      `json:"second-wind"`
	Indomitable          types.NamedToken      `json:"indomitable"`
}

type FightingStyleFeature struct {
	Name      string `json:"name"`
	IsApplied bool   `json:"is-applied"`
	Details   string `json:"details"`
}

func LoadFighter(data []byte) (*Fighter, error) {
	var fighter Fighter
	if err := json.Unmarshal(data, &fighter); err != nil {
		return &fighter, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &fighter, nil
}

func (f *Fighter) ValidateMethods(c *models.Character) {
}

func (f *Fighter) ExecutePostCalculateMethods(c *models.Character) {
	f.executerFightingStyle(c)
}

func (f *Fighter) ExecutePreCalculateMethods(c *models.Character) {
}

func (f *Fighter) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd10", level)
}

func (f *Fighter) executerFightingStyle(c *models.Character) {
	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles, %s, %s, %s, %s",
		f.FightingStyle,
		types.FightingStyleArchery,
		types.FightingStyleDefense,
		types.FightingStyleDueling,
		types.FightingStyleTwoWeaponFighting)

	switch strings.ToLower(f.FightingStyle) {
	case types.FightingStyleArchery:
		f.FightingStyleFeature = applyArchery(c)
	case types.FightingStyleDefense:
		f.FightingStyleFeature = applyDefense(c)
	case types.FightingStyleDueling:
		f.FightingStyleFeature = applyDueling(c)
	case types.FightingStyleTwoWeaponFighting:
		f.FightingStyleFeature = applyTwoWeaponFighting(c)
	case types.FightingtStyleGreatWeaponFighting:
		f.FightingStyleFeature = applyGreatWeaponFighting(c)
	case types.FightingSyleProtection:
		f.FightingStyleFeature = applyProtection(c)
	default:
		logger.HandleInfo(invalidMsg)
	}
}

func (f *Fighter) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if f.Archetype != "" && c.Level > 3 {
		archetypeHeader := fmt.Sprintf("Archetype: *%s*\n\n", f.Archetype)
		s = append(s, archetypeHeader)
	}

	if f.FightingStyleFeature.Name != "" && c.Level >= 2 {
		appliedText := "Requirements for fighting style not met."
		if f.FightingStyleFeature.IsApplied {
			appliedText = "Requirements for this fighting style are met, and any bonuses to armor or weapons have been applied to your character."
		}

		fightingStyleHeader := fmt.Sprintf("**Fighting Style**: *%s*\n", f.FightingStyleFeature.Name)
		fightingStyleDetail := fmt.Sprintf("%s\n%s\n\n", f.FightingStyleFeature.Details, appliedText)
		s = append(s, fightingStyleHeader)
		s = append(s, fightingStyleDetail)
	}

	if f.ActionSurge.Available != 0 && f.ActionSurge.Maximum != 0 && c.Level >= 2 {
		actionSurgeSlots := c.GetSlots(f.ActionSurge.Available, f.ActionSurge.Maximum)
		line := fmt.Sprintf("**Action Surge**: %s\n\n", actionSurgeSlots)
		s = append(s, line)
	}

	if f.SecondWind.Available != 0 && f.SecondWind.Maximum != 0 {
		secondWindSlots := c.GetSlots(f.SecondWind.Available, f.SecondWind.Maximum)
		line := fmt.Sprintf("**Second Wind**: %s\n\n", secondWindSlots)
		s = append(s, line)
	}

	if f.Indomitable.Available != 0 && f.Indomitable.Maximum != 0 && c.Level >= 9 {
		indomitableSlots := c.GetSlots(f.Indomitable.Available, f.Indomitable.Maximum)
		line := fmt.Sprintf("**Indomitable**: %s\n\n", indomitableSlots)
		s = append(s, line)
	}

	if len(f.OtherFeatures) > 0 {
		for _, detail := range f.OtherFeatures {
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

func (f *Fighter) AddFightingStyleFeature(feature models.ClassFeature) {

}

func (f *Fighter) RemoveFightingStyleFeature(feature models.ClassFeature) {

}

// CLI

func (f *Fighter) UseClassTokens(tokenName string) {
	token := f.getToken(tokenName)

	if token == nil {
		logger.HandleInfo(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	if token.Available <= 0 {
		logger.HandleInfo(fmt.Sprintf("%s had no uses left", tokenName))
		return
	}

	token.Available--
}

func (f *Fighter) RecoverClassTokens(tokenName string, quantity int) {
	token := f.getToken(tokenName)

	if token == nil {
		logger.HandleInfo(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || token.Available > token.Maximum {
		token.Available = token.Maximum
	}
}

func (f *Fighter) getToken(tokenName string) *types.NamedToken {
	switch strings.ToLower(tokenName) {
	case f.SecondWind.Name:
		return &f.SecondWind
	case f.ActionSurge.Name:
		return &f.ActionSurge
	case f.Indomitable.Name:
		return &f.Indomitable
	default:
		return nil
	}
}
