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

type Cleric struct {
	models.BaseClass
	ClassToken     shared.NamedToken `json:"class-token" clover:"class-token"`
	PreparedSpells []string          `json:"prepared-spells" clover:"prepared-spells"`
}

const channelDivinityToken string = "channel-divinity"

func LoadCleric(data []byte) (*Cleric, error) {
	var cleric Cleric
	if err := json.Unmarshal(data, &cleric); err != nil {
		return &cleric, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &cleric, nil
}

func (cl *Cleric) CalculateHitDice() string {
	return fmt.Sprintf("%dd8", cl.Level)
}

func (cl *Cleric) ExecutePostCalculateMethods(c *models.Character) {
	cl.executeSpellCastingAbility(c)
	cl.executePreparedSpells(c)
	cl.executeChannelDiversity(c)
	cl.executeCantripVersatility(c)
}

func (cl *Cleric) executeChannelDiversity(c *models.Character) {
	if cl.ClassToken.Name == "" {
		return
	} else if cl.ClassToken.Name != channelDivinityToken {
		logger.Info("Invalid Class Token Name")
		return
	}

	switch {
	case c.Level < 2:
		cl.ClassToken.Maximum = 0
	case c.Level < 6:
		cl.ClassToken.Maximum = 1
	case c.Level < 18:
		cl.ClassToken.Maximum = 2
	case c.Level >= 18:
		cl.ClassToken.Maximum = 3
	}
}

func (cl *Cleric) executeCantripVersatility(c *models.Character) {
	// Even though this is functionally the same as the Druid version, the switch table is different,
	// so we're going to have these exist separately
	if c.ValidationDisabled {
		return
	}

	cantripCount := 0
	for _, spell := range c.Spells {
		if spell.SlotLevel == 0 {
			cantripCount++
		}
	}

	abilityImprovementTotal := 0
	for _, ability := range c.AbilityScoreImprovement {
		abilityImprovementTotal += ability.Bonus
	}

	cantripVersatilityMax := 0
	switch {
	case c.Level < 4:
		cantripVersatilityMax = 3 // +3 cantrip
	case c.Level < 8:
		cantripVersatilityMax = 6 // +1 cantrip, +2 ASI
	case c.Level < 10:
		cantripVersatilityMax = 7 // +1 cantrip
	case c.Level < 12:
		cantripVersatilityMax = 10 // +2 ASI
	case c.Level < 16:
		cantripVersatilityMax = 12 // +2 ASI
	case c.Level < 20:
		cantripVersatilityMax = 14 // +2 ASI
	case c.Level >= 20:
		cantripVersatilityMax = 16 // +2 ASI
	}

	if cantripVersatilityMax < cantripCount+abilityImprovementTotal {
		logger.Info("Cantrip Versatility: You have too many cantrips or ability score improvement bonuses for your level")
	}
}

func (cl *Cleric) executeSpellCastingAbility(c *models.Character) {
	wisMod := c.GetMod(shared.AbilityWisdom)

	executeSpellSaveDC(c, wisMod)
	executeSpellAttackMod(c, wisMod)
}

func (cl *Cleric) executePreparedSpells(c *models.Character) {
	wisMod := c.GetMod(shared.AbilityWisdom)
	preparedSpellsMax := wisMod + c.Level

	if !c.ValidationDisabled {
		if len(cl.PreparedSpells) > preparedSpellsMax {
			logger.Info(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(cl.PreparedSpells), preparedSpellsMax))
		} else if len(cl.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(cl.PreparedSpells)
			logger.Info(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, cl.PreparedSpells)
}

func (cl *Cleric) ClassDetails() string {
	var s string
	s += formatTokens(cl.ClassToken, channelDivinityToken, cl.Level)

	return s
}

// CLI

func (cl *Cleric) UseClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != channelDivinityToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, cl.ClassType))
		return
	}

	if cl.ClassToken.Available <= 0 {
		logger.Info("Channel Divinity had no uses left")
		return
	}

	cl.ClassToken.Available -= quantity
}

func (cl *Cleric) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != channelDivinityToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, cl.ClassType))
		return
	}

	cl.ClassToken.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || cl.ClassToken.Available > cl.ClassToken.Maximum {
		cl.ClassToken.Available = cl.ClassToken.Maximum
	}
}

func (cl *Cleric) GetTokens() []string {
	return []string{
		channelDivinityToken,
	}
}

func (cl *Cleric) AddPreparedSpell(spell string) error {
	for _, ps := range cl.PreparedSpells {
		if strings.EqualFold(ps, spell) {
			return fmt.Errorf("Spell 's' already exists as a prepared spell")
		}
	}

	cl.PreparedSpells = append(cl.PreparedSpells, spell)

	return nil
}

func (cl *Cleric) RemovePreparedSpell(spell string) error {
	for i, ps := range cl.PreparedSpells {
		if strings.EqualFold(ps, spell) {
			cl.PreparedSpells = slices.Delete(cl.PreparedSpells, i, i+1)
			return nil
		}
	}

	return fmt.Errorf("Failed to find spell '%s' in list of prepared spells to remove", spell)
}

func (cl *Cleric) GetPreparedSpells() []string {
	return cl.PreparedSpells
}
