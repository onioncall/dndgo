package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Fighter struct {
	Archetype            string                `json:"archetype"`
	FightingStyle        string                `json:"fighting-style"`
	FightingStyleFeature FightingStyleFeature  `json:"-"`
	OtherFeatures        []models.ClassFeature `json:"other-features"`
	ClassTokens          []shared.NamedToken   `json:"class-tokens"`
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

func (f *Fighter) ExecutePostCalculateMethods(c *models.Character) {
	f.executeFightingStyle(c)
	f.executeClassTokens()
}

func (f *Fighter) ExecutePreCalculateMethods(c *models.Character) {
}

func (f *Fighter) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd10", level)
}

func (f *Fighter) executeClassTokens() {
	for i := range f.ClassTokens {
		f.ClassTokens[i].Maximum = 1
	}
}

func (f *Fighter) executeFightingStyle(c *models.Character) {
	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles, %s, %s, %s, %s",
		f.FightingStyle,
		shared.FightingStyleArchery,
		shared.FightingStyleDefense,
		shared.FightingStyleDueling,
		shared.FightingStyleTwoWeaponFighting)

	switch strings.ToLower(f.FightingStyle) {
	case shared.FightingStyleArchery:
		f.FightingStyleFeature = applyArchery(c)
	case shared.FightingStyleDefense:
		f.FightingStyleFeature = applyDefense(c)
	case shared.FightingStyleDueling:
		f.FightingStyleFeature = applyDueling(c)
	case shared.FightingStyleTwoWeaponFighting:
		f.FightingStyleFeature = applyTwoWeaponFighting(c)
	case shared.FightingStyleGreatWeaponFighting:
		f.FightingStyleFeature = applyGreatWeaponFighting(c)
	case shared.FightingStyleProtection:
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

	for _, token := range f.ClassTokens {
		tokenHeader := ""

		switch token.Name {
		case "action-surge":
			tokenHeader = "Action Surge"
		case "second-wind":
			tokenHeader = "Second Wind"
		case "indomitable":
			tokenHeader = "Indomitable"
		default:
			logger.HandleInfo(fmt.Sprintf("Invalid token name: %s", token.Name))
			continue
		}

		if token.Maximum != 0 && c.Level >= token.Level {
			actionSurgeSlots := c.GetSlots(token.Available, token.Maximum)
			line := fmt.Sprintf("**%s**: %s\n\n", tokenHeader, actionSurgeSlots)
			s = append(s, line)
		}
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

func (f *Fighter) UseClassTokens(tokenName string, quantity int) {
	token := getToken(tokenName, f.ClassTokens)

	if token == nil {
		logger.HandleInfo(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	if token.Available <= 0 {
		logger.HandleInfo(fmt.Sprintf("%s had no uses left", tokenName))
		return
	}

	token.Available -= quantity
}

func (f *Fighter) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName == "all" {
		fullTokenRecovery(f.ClassTokens)
		return
	}

	token := getToken(tokenName, f.ClassTokens)

	if token == nil {
		logger.HandleInfo(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || token.Available > token.Maximum {
		token.Available = token.Maximum
	}
}

func (f *Fighter) GetTokens() []string {
	s := []string{}

	for _, token := range f.ClassTokens {
		s = append(s, token.Name)
	}

	return s
}
