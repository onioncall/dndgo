package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Sorcerer struct {
	SorcerousOrigin string                `json:"sorcerous-origin"`
	SorceryPoints   types.Token           `json:"sorcery-points"`
	OtherFeatures   []models.ClassFeature `json:"other-features"`
}

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

func (s *Sorcerer) ExecutePreCalculateMethods(c *models.Character) {
}

func (s *Sorcerer) ValidateMethods(c *models.Character) {
}

func (s *Sorcerer) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd6", level)
}

func (s *Sorcerer) executeSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(types.AbilityCharisma)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

func (s *Sorcerer) executeSorceryPoints(c *models.Character) {
	s.SorceryPoints.Maximum = 2
	s.SorceryPoints.Maximum += c.Level
}

// CLI

func (s *Sorcerer) UseClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since sorcerer only has sorcery points, we won't check the slot name value
	if s.SorceryPoints.Available <= 0 {
		logger.HandleInfo("There were no sorcery points left")
		return
	}

	s.SorceryPoints.Available -= quantity
}

func (s *Sorcerer) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since sorcerer only has sorcery points, we won't check the slot name value
	s.SorceryPoints.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || s.SorceryPoints.Available > s.SorceryPoints.Maximum {
		s.SorceryPoints.Available = s.SorceryPoints.Maximum
	}
}
