package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Cleric struct {
	ClassToken     shared.NamedToken     `json:"class-token"`
	Domain         string                `json:"domain"`
	PreparedSpells []string              `json:"prepared-spells"`
	OtherFeatures  []models.ClassFeature `json:"other-features"`
}

const channelDivinityToken string = "channel-divinity"

func LoadCleric(data []byte) (*Cleric, error) {
	var cleric Cleric
	if err := json.Unmarshal(data, &cleric); err != nil {
		return &cleric, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &cleric, nil
}

func (cl *Cleric) ValidateMethods(c *models.Character) {
	isValid := cl.validateCantripVersatility(c)
	if isValid {
		logger.HandleInfo("Cantrip Versatility: You have too many cantrips or ability score improvement bonuss for your level")
	}
}

func (cl *Cleric) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (cl *Cleric) ExecutePostCalculateMethods(c *models.Character) {
	cl.executeSpellCastingAbility(c)
	cl.executePreparedSpells(c)
	cl.executeChannelDiversity(c)
}

func (cl *Cleric) ExecutePreCalculateMethods(c *models.Character) {
}

func (cl *Cleric) executeChannelDiversity(c *models.Character) {
	if cl.ClassToken.Name == "" {
		return
	} else if cl.ClassToken.Name != channelDivinityToken {
		logger.HandleInfo("Invalid Class Token Name")
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

func (cl *Cleric) validateCantripVersatility(c *models.Character) bool {
	// Even though this is functionally the same as the Druid version, the switch table is different,
	// so we're going to have these exist separately
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
		return false
	}

	return true
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
			logger.HandleInfo(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(cl.PreparedSpells), preparedSpellsMax))
		} else if len(cl.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(cl.PreparedSpells)
			logger.HandleInfo(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, cl.PreparedSpells)
}

func (cl *Cleric) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if cl.Domain != "" && c.Level > 3 {
		domainHeader := fmt.Sprintf("Domain: *%s*\n\n", cl.Domain)
		s = append(s, domainHeader)
	}

	if cl.ClassToken.Maximum != 0 && cl.ClassToken.Name == channelDivinityToken {
		channelDivinitySlots := c.GetSlots(cl.ClassToken.Available, cl.ClassToken.Maximum)
		biLine := fmt.Sprintf("**Channel Divinity**: %s\n\n", channelDivinitySlots)
		s = append(s, biLine)
	}

	if len(cl.OtherFeatures) > 0 {
		for _, detail := range cl.OtherFeatures {
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

func (cl *Cleric) UseClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since bard only has channel divinity, we won't check the slot name value
	if cl.ClassToken.Available <= 0 {
		logger.HandleInfo("Channel Divinity had no uses left")
		return
	}

	cl.ClassToken.Available -= quantity
}

func (cl *Cleric) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since bard only has channel divinity, we won't check the slot name value
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
