package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Sorcerer struct {
	models.BaseClass
	ClassToken      shared.NamedToken     `json:"class-token" clover:"class-token"`
	MetaMagicSpells []models.ClassFeature `json:"meta-magic-spells" clover:"meta-magic-spells"`
}

const sorceryPointsToken string = "sorcery-points"

func LoadSorcerer(data []byte) (*Sorcerer, error) {
	var sorcerer Sorcerer
	if err := json.Unmarshal(data, &sorcerer); err != nil {
		return &sorcerer, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &sorcerer, nil
}

func (s *Sorcerer) ExecutePostCalculateMethods(c *models.Character) {
	s.executeSpellCastingAbility(c)
	s.executeSorceryPoints(c)
}

func (s *Sorcerer) CalculateHitDice() string {
	return fmt.Sprintf("%dd6", s.Level)
}

func (s *Sorcerer) executeSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(shared.AbilityCharisma)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

func (s *Sorcerer) executeSorceryPoints(c *models.Character) {
	s.ClassToken.Maximum = 2
	s.ClassToken.Maximum += c.Level
}

func (s *Sorcerer) ClassDetails() string {
	var str string

	str += fmt.Sprintf("Level: %d\n", s.Level)

	if s.Level >= 2 && s.ClassToken.Name == sorceryPointsToken {
		str += fmt.Sprintf("*Sorcery Points*: %d/%d\n\n", s.ClassToken.Available, s.ClassToken.Maximum)
	}

	if len(s.MetaMagicSpells) > 0 && s.Level >= 3 {
		mmHeader := fmt.Sprintf("Meta Magic Spells:\n")
		str += mmHeader

		for _, spell := range s.MetaMagicSpells {
			spellLine := fmt.Sprintf("*%s*\n", spell.Name)
			str += fmt.Sprintf("*%s*\n", spell.Name)
			str += fmt.Sprintf("%s\n\n", spell.Details)
			str += spellLine
		}
		str += "\n"
	}

	return str
}

// CLI

func (s *Sorcerer) UseClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != sorceryPointsToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, s.ClassType))
		return
	}

	if s.ClassToken.Available <= 0 {
		logger.Info("There were no sorcery points left")
		return
	}

	s.ClassToken.Available -= quantity
}

func (s *Sorcerer) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != sorceryPointsToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, s.ClassType))
		return
	}

	s.ClassToken.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || s.ClassToken.Available > s.ClassToken.Maximum {
		s.ClassToken.Available = s.ClassToken.Maximum
	}
}

func (s *Sorcerer) GetTokens() []string {
	return []string{
		sorceryPointsToken,
	}
}
