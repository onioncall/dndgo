package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Sorcerer struct {
	BaseClass
	SorcerousOrigin string                `json:"sorcerous-origin" clover:"sorcerous-origin"`
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

func (s *Sorcerer) PrintClassDetails(c *models.Character) []string {
	sb := buildClassDetailsHeader()

	if s.SorcerousOrigin != "" {
		sb = append(sb, fmt.Sprintf("Sorcerous Origin: *%s*\n\n", s.SorcerousOrigin))
	}

	if c.Level >= 2 && s.ClassToken.Name == sorceryPointsToken {
		sb = append(sb, fmt.Sprintf("*Sorcery Points*: %d/%d\n\n", s.ClassToken.Available, s.ClassToken.Maximum))
	}

	if len(s.MetaMagicSpells) > 0 && c.Level >= 3 {
		mmHeader := fmt.Sprintf("Meta Magic Spells:\n")
		sb = append(sb, mmHeader)

		for _, spell := range s.MetaMagicSpells {
			spellLine := fmt.Sprintf("*%s*\n", spell.Name)
			sb = append(sb, fmt.Sprintf("*%s*\n", spell.Name))
			sb = append(sb, fmt.Sprintf("%s\n\n", spell.Details))
			sb = append(sb, spellLine)
		}
		sb = append(sb, "\n")
	}

	if len(s.OtherFeatures) > 0 {
		for _, detail := range s.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			detailName := fmt.Sprintf("---\n**%s**\n", detail.Name)
			sb = append(sb, detailName)
			details := fmt.Sprintf("%s\n", detail.Details)
			sb = append(sb, details)
		}
	}

	return sb
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
