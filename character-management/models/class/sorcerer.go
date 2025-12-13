package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Sorcerer struct {
	SorcerousOrigin string                `json:"sorcerous-origin"`
	ClassToken      shared.NamedToken     `json:"class-token"`
	MetaMagicSpells []models.ClassFeature `json:"meta-magic-spells"`
	OtherFeatures   []models.ClassFeature `json:"other-features"`
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

func (s *Sorcerer) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd6", level)
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

func (s *Sorcerer) SubClass(level int) string {
	if level <= 2 {
		return ""
	}

	return s.SorcerousOrigin
}

func (s *Sorcerer) ClassDetails(level int) string {
	var str string

	if level >= 2 && s.ClassToken.Name == sorceryPointsToken {
		str += fmt.Sprintf("*Sorcery Points*: %d/%d\n\n", s.ClassToken.Available, s.ClassToken.Maximum)
	}

	if len(s.MetaMagicSpells) > 0 && level >= 3 {
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

func (s *Sorcerer) ClassFeatures(level int) string {
	var str string
	str += formatOtherFeatures(s.OtherFeatures, level)

	return str
}

// CLI

func (s *Sorcerer) UseClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since sorcerer only has sorcery points, we won't check the slot name value
	if s.ClassToken.Available <= 0 {
		logger.Info("There were no sorcery points left")
		return
	}

	s.ClassToken.Available -= quantity
}

func (s *Sorcerer) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since sorcerer only has sorcery points, we won't check the slot name value
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
