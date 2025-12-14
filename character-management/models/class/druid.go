package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Druid struct {
	ClassToken     shared.NamedToken     `json:"class-token"`
	Circle         string                `json:"circle"`
	PreparedSpells []string              `json:"prepared-spells"`
	OtherFeatures  []models.ClassFeature `json:"other-features"`
}

const wildShapeToken string = "wild-shape"

func LoadDruid(data []byte) (*Druid, error) {
	var druid Druid
	if err := json.Unmarshal(data, &druid); err != nil {
		return &druid, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &druid, nil
}

func (d *Druid) ExecutePostCalculateMethods(c *models.Character) {
	d.executeSpellCastingAbility(c)
	d.executePreparedSpells(c)
	d.executeArchDruid(c)
	d.executeWildShape(c)
	d.executeCantripVersatility(c)
}

func (d *Druid) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (d *Druid) executeWildShape(c *models.Character) {
	if c.Level < 2 || d.ClassToken.Name == "" {
		return
	} else if d.ClassToken.Name != wildShapeToken {
		logger.Info("Invalid Class Token Name")
		return
	}

	d.ClassToken.Maximum = 2
}

func (d *Druid) executeCantripVersatility(c *models.Character) {
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
		cantripVersatilityMax = 2 // +2 cantrip
	case c.Level < 8:
		cantripVersatilityMax = 5 // +1 cantrip, +2 ASI
	case c.Level < 10:
		cantripVersatilityMax = 6 // +1 cantrip
	case c.Level < 12:
		cantripVersatilityMax = 8 // +2 ASI
	case c.Level < 16:
		cantripVersatilityMax = 10 // +2 ASI
	case c.Level < 20:
		cantripVersatilityMax = 12 // +2 ASI
	case c.Level >= 20:
		cantripVersatilityMax = 14 // +2 ASI
	}

	if cantripVersatilityMax < cantripCount+abilityImprovementTotal {
		logger.Info("Cantrip Versatility: You have too many cantrips or ability score improvement bonuss for your level")
	}
}

func (d *Druid) executeSpellCastingAbility(c *models.Character) {
	wisMod := c.GetMod(shared.AbilityWisdom)

	executeSpellSaveDC(c, wisMod)
	executeSpellAttackMod(c, wisMod)
}

func (d *Druid) executePreparedSpells(c *models.Character) {
	wisMod := c.GetMod(shared.AbilityWisdom)
	preparedSpellsMax := wisMod + c.Level

	if !c.ValidationDisabled {
		if len(d.PreparedSpells) > preparedSpellsMax {
			logger.Info(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(d.PreparedSpells), preparedSpellsMax))
		} else if len(d.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(d.PreparedSpells)
			logger.Info(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, d.PreparedSpells)
}

func (d *Druid) executeArchDruid(c *models.Character) {
	if c.Level < 20 {
		return
	}

	// These are now unlimited, no need to track them anymore
	d.ClassToken.Available = 0
	d.ClassToken.Maximum = 0
}

func (d *Druid) SubClass(level int) string {
	if level <= 2 {
		return ""
	}

	return d.Circle
}

func (d *Druid) ClassDetails(level int) string {
	var s string
	s += formatTokens(d.ClassToken, wildShapeToken, level)

	return s
}

func (d *Druid) ClassFeatures(level int) string {
	var s string
	s += formatOtherFeatures(d.OtherFeatures, level)

	return s
}

// CLI

func (d *Druid) UseClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since druid only has wild shape, we won't check the slot name value
	if d.ClassToken.Available <= 0 {
		logger.Info("Wild Shape had no uses left")
		return
	}

	d.ClassToken.Available -= quantity
}

func (d *Druid) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since druid only has wild shape, we won't check the slot name value
	d.ClassToken.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || d.ClassToken.Available > d.ClassToken.Maximum {
		d.ClassToken.Available = d.ClassToken.Maximum
	}
}

func (d *Druid) GetTokens() []string {
	return []string{
		wildShapeToken,
	}
}
