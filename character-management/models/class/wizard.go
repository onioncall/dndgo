package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Wizard struct {
	SignatureSpells []string              `json:"signature-spells"`
	ArcaneTradition string                `json:"arcane-tradition"`
	PreparedSpells  []string              `json:"prepared-spells"`
	OtherFeatures   []models.ClassFeature `json:"other-features"`
}

func LoadWizard(data []byte) (*Wizard, error) {
	var wizard Wizard
	if err := json.Unmarshal(data, &wizard); err != nil {
		return &wizard, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &wizard, nil
}

func (w *Wizard) ExecutePostCalculateMethods(c *models.Character) {
	w.executeSpellCastingAbility(c)
	w.executePreparedSpells(c)
}

func (w *Wizard) ExecutePreCalculateMethods(c *models.Character) {
}

func (w *Wizard) ValidateMethods(c *models.Character) {
	w.validateSignatureSpells(c)
}

func (w *Wizard) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd6", level)
}

func (w *Wizard) executePreparedSpells(c *models.Character) {
	intMod := c.GetMod(types.AbilityIntelligence)
	preparedSpellsMax := intMod + c.Level

	if !c.ValidationDisabled {
		if len(w.PreparedSpells) > preparedSpellsMax {
			logger.HandleInfo(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(w.PreparedSpells), preparedSpellsMax))
		} else if len(w.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(w.PreparedSpells)
			logger.HandleInfo(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, w.PreparedSpells)
}

func (w *Wizard) executeSpellCastingAbility(c *models.Character) {
	intMod := c.GetMod(types.AbilityIntelligence)

	executeSpellSaveDC(c, intMod)
	executeSpellAttackMod(c, intMod)
}

func (w *Wizard) validateSignatureSpells(c *models.Character) {
	if c.ValidationDisabled {
		return
	}

	for _, sigSpell := range w.SignatureSpells {
		spellFound := false
		for _, spell := range c.Spells {
			if strings.ToLower(spell.Name) == strings.ToLower(sigSpell) {
				spellFound = true

				if spell.SlotLevel > 3 {
					logger.HandleInfo(fmt.Sprintf("Signature Spell '%s' is an invalid level", sigSpell))
				}

				break
			}
		}

		if !spellFound {
			logger.HandleInfo(fmt.Sprintf("Signature Spell '%s' was not found in your list of learned spells", sigSpell))
		}
	}
}

func (w *Wizard) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if w.ArcaneTradition != "" && c.Level > 3 {
		s = append(s, fmt.Sprintf("Arcane Tradition: *School of %s*\n\n", w.ArcaneTradition))
	}

	if c.Level >= 20 {
		s = append(s, fmt.Sprintf("Signature Spells:\n"))
		for _, spell := range w.SignatureSpells {
			s = append(s, fmt.Sprintf("- %s\n", spell))
		}
	}

	if len(w.OtherFeatures) > 0 {
		for _, detail := range w.OtherFeatures {
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

// CLI

func (w *Wizard) UseClassTokens(tokenName string, quantity int) {
	// Not sure Wizards have a token like system to implement
	logger.HandleInfo("No token set up for Wizard class")
}

func (w *Wizard) RecoverClassTokens(tokenName string, quantity int) {
	// Not sure Wizards have a token like system to implement
	logger.HandleInfo("No token set up for Wizard class")
}

func (w *Wizard) GetTokens() []string {
	return []string{}
}
