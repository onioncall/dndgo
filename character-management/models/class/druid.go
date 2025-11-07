package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Druid struct {
	WildShape      WildShape             `json:"wild-shape"`
	Circle         string                `json:"circle"`
	PreparedSpells []string              `json:"prepared-spells"`
	OtherFeatures  []models.ClassFeature `json:"other-features"`
}

type WildShape struct {
	Available int `json:"available"`
	Maximum   int `json:"maximum"`
}

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
}

func (d *Druid) ExecutePreCalculateMethods(c *models.Character) {
}

func (d *Druid) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (d *Druid) ValidateMethods(c *models.Character) {
	isValid := d.validateCantripVersatility(c)
	if isValid {
		logger.HandleInfo("Cantrip Versatility: You have too many cantrips or ability score improvement bonuss for your level")
	}
}

func (d *Druid) validateCantripVersatility(c *models.Character) bool {
	if c.ValidationDisabled {
		return true
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
		return false
	}

	return true
}

func (d *Druid) executeSpellCastingAbility(c *models.Character) {
	wisMod := c.GetMod(types.AbilityWisdom)

	executeSpellSaveDC(c, wisMod)
	executeSpellAttackMod(c, wisMod)
}

func (d *Druid) executePreparedSpells(c *models.Character) {
	wisMod := c.GetMod(types.AbilityWisdom)
	preparedSpellsMax := wisMod + c.Level

	if !c.ValidationDisabled {
		if len(d.PreparedSpells) > preparedSpellsMax {
			logger.HandleInfo(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(d.PreparedSpells), preparedSpellsMax))
		} else if len(d.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(d.PreparedSpells)
			logger.HandleInfo(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, d.PreparedSpells)
}

func (d *Druid) executeArchDruid(c *models.Character) {
	if c.Level < 20 {
		return
	}

	// These are now unlimited, no need to track them anymore
	d.WildShape.Available = 0
	d.WildShape.Maximum = 0
}

func (d *Druid) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if d.Circle != "" && c.Level > 3 {
		collegeHeader := fmt.Sprintf("Circle: *%s*\n\n", d.Circle)
		s = append(s, collegeHeader)
	}

	if d.WildShape.Available != 0 && d.WildShape.Maximum != 0 {
		wildShapeSlots := c.GetSlots(d.WildShape.Available, d.WildShape.Maximum)
		biLine := fmt.Sprintf("**Wild Shape Transformations**: %s\n\n", wildShapeSlots)
		s = append(s, biLine)
	}

	if len(d.OtherFeatures) > 0 {
		for _, detail := range d.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			detailName := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, detailName)
			details := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, details)
		}
	}

	return s
}

// CLI

func (d *Druid) UseClassTokens(tokenName string) {
	// We only really need slot name for classes that have multiple slots
	// since druid only has wild shape, we won't check the slot name value
	if d.WildShape.Available <= 0 {
		logger.HandleInfo("Wild Shape had no uses left")
		return
	}

	d.WildShape.Available--
}

func (d *Druid) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since druid only has wild shape, we won't check the slot name value
	d.WildShape.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || d.WildShape.Available > d.WildShape.Maximum {
		d.WildShape.Available = d.WildShape.Maximum
	}
}
