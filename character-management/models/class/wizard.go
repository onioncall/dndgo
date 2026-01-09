package class

import (
	"encoding/json"
	"fmt"
	"slices"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Wizard struct {
	models.BaseClass
	SignatureSpells []string `json:"signature-spells" clover:"signature-spells"`
	PreparedSpells  []string `json:"prepared-spells" clover:"prepared-spells"`
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
	w.executeSignatureSpellValidation(c)
}

func (w *Wizard) CalculateHitDice() string {
	return fmt.Sprintf("%dd6", w.Level)
}

func (w *Wizard) executePreparedSpells(c *models.Character) {
	executePreparedSpellsShared(c, w.PreparedSpells)
}

func (w *Wizard) executeSpellCastingAbility(c *models.Character) {
	intMod := c.GetMod(shared.AbilityIntelligence)

	executeSpellSaveDC(c, intMod)
	executeSpellAttackMod(c, intMod)
}

func (w *Wizard) executeSignatureSpellValidation(c *models.Character) {
	if c.ValidationDisabled {
		return
	}

	for _, sigSpell := range w.SignatureSpells {
		spellFound := false
		for _, spell := range c.Spells {
			if strings.ToLower(spell.Name) == strings.ToLower(sigSpell) {
				spellFound = true

				if spell.SlotLevel > 3 {
					logger.Info(fmt.Sprintf("Signature Spell '%s' is an invalid level", sigSpell))
				}

				break
			}
		}

		if !spellFound {
			logger.Info(fmt.Sprintf("Signature Spell '%s' was not found in your list of learned spells", sigSpell))
		}
	}
}

func (w *Wizard) ClassDetails() string {
	var s string

	if w.Level >= 20 {
		s += fmt.Sprintf("Signature Spells:\n")
		for _, spell := range w.SignatureSpells {
			s += fmt.Sprintf("- %s\n", spell)
		}
	}

	return s
}

func (w *Wizard) AddPreparedSpell(spell string) error {
	for _, ps := range w.PreparedSpells {
		if strings.EqualFold(ps, spell) {
			return fmt.Errorf("Spell 's' already exists as a prepared spell")
		}
	}

	w.PreparedSpells = append(w.PreparedSpells, spell)

	return nil
}

func (w *Wizard) RemovePreparedSpell(spell string) error {
	for i, ps := range w.PreparedSpells {
		if strings.EqualFold(ps, spell) {
			w.PreparedSpells = slices.Delete(w.PreparedSpells, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("Failed to find spell '%s' in list of prepared spells to remove", spell)
}

func (w *Wizard) GetPreparedSpells() []string {
	return w.PreparedSpells
}
