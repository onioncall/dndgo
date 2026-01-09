package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

var fightingStyles = []string{
	shared.FightingStyleArchery,
	shared.FightingStyleDefense,
	shared.FightingStyleDueling,
	shared.FightingStyleTwoWeaponFighting,
	shared.FightingStyleGreatWeaponFighting,
	shared.FightingStyleProtection,
}

type Fighter struct {
	models.BaseClass
	FightingStyle        string               `json:"fighting-style" clover:"fighting-style"`
	FightingStyleFeature FightingStyleFeature `json:"-" clover:"-"`
	ClassTokens          []shared.NamedToken  `json:"class-tokens" clover:"class-tokens"`
}

type FightingStyleFeature struct {
	Name      string `json:"name" clover:"name"`
	IsApplied bool   `json:"is-applied" clover:"is-applied"`
	Details   string `json:"details" clover:"details"`
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

func (f *Fighter) CalculateHitDice() string {
	return fmt.Sprintf("%dd10", f.Level)
}

func (f *Fighter) executeClassTokens() {
	for i := range f.ClassTokens {
		f.ClassTokens[i].Maximum = 1
	}
}

func (f *Fighter) executeFightingStyle(c *models.Character) {
	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles, %s, %s, %s, %s, %s, %s",
		f.FightingStyle,
		shared.FightingStyleArchery,
		shared.FightingStyleDefense,
		shared.FightingStyleDueling,
		shared.FightingStyleTwoWeaponFighting,
		shared.FightingStyleGreatWeaponFighting,
		shared.FightingStyleProtection)

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
		logger.Info(invalidMsg)
	}
}

func (f *Fighter) ClassDetails() string {
	var s string

	s += fmt.Sprintf("Level: %d\n", f.Level)

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
			logger.Info(fmt.Sprintf("Invalid token name: %s", token.Name))
			continue
		}

		s += formatTokens(token, tokenHeader, f.Level)
	}

	if f.FightingStyleFeature.Name != "" && f.Level >= 2 {
		appliedText := "Requirements for fighting style not met."
		if f.FightingStyleFeature.IsApplied {
			appliedText = "Requirements for this fighting style are met, and any bonuses to armor or weapons have been applied to your character."
		}

		fightingStyleHeader := fmt.Sprintf("**Fighting Style**: *%s*\n", f.FightingStyleFeature.Name)
		fightingStyleDetail := fmt.Sprintf("%s\n%s\n\n", f.FightingStyleFeature.Details, appliedText)
		s += fightingStyleHeader
		s += fightingStyleDetail
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
		logger.Info(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	if token.Available <= 0 {
		logger.Info(fmt.Sprintf("%s had no uses left", tokenName))
		return
	}

	token.Available -= quantity
}

func (f *Fighter) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName == "" {
		fullTokenRecovery(f.ClassTokens)
		return
	}

	token := getToken(tokenName, f.ClassTokens)

	if token == nil {
		logger.Info(fmt.Sprintf("Invalid token name: %s", tokenName))
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

func (f *Fighter) ModifyFightingStyle(fightingStyle string) error {
	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles", fightingStyle)
	for _, fs := range fightingStyles {
		if strings.EqualFold(fs, fightingStyle) {
			f.FightingStyle = fightingStyle
			return nil
		}

		invalidMsg += fmt.Sprintf(", %s", fs)
	}

	return fmt.Errorf("%s", invalidMsg)
}
