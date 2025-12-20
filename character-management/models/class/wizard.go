package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Wizard struct {
	models.BaseClass
	SignatureSpells []string `json:"signature-spells"`
	PreparedSpells  []string `json:"prepared-spells"`
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

func (w *Wizard) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd6", level)
}

func (w *Wizard) executePreparedSpells(c *models.Character) {
	intMod := c.GetMod(shared.AbilityIntelligence)
	preparedSpellsMax := intMod + c.Level

	if !c.ValidationDisabled {
		if len(w.PreparedSpells) > preparedSpellsMax {
			logger.Info(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(w.PreparedSpells), preparedSpellsMax))
		} else if len(w.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(w.PreparedSpells)
			logger.Info(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

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

func (w *Wizard) ClassDetails(level int) string {
	var s string

	if level >= 20 {
		s += fmt.Sprintf("Signature Spells:\n")
		for _, spell := range w.SignatureSpells {
			s += fmt.Sprintf("- %s\n", spell)
		}
	}

	return s
}
