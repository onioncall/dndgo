package class

import (
	"encoding/json"
	"fmt"
	"math"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Bard struct {
	ExpertiseSkills []string              `json:"expertise"`
	College         string                `json:"college"`
	OtherFeatures   []models.ClassFeature `json:"other-features"`
	ClassToken      shared.NamedToken     `json:"class-token"`
}

const bardicInspirationToken string = "bardic-inspiration"

const bardSpellCastingAbility string = shared.AbilityCharisma

func LoadBard(data []byte) (*Bard, error) {
	var bard Bard
	if err := json.Unmarshal(data, &bard); err != nil {
		return &bard, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &bard, nil
}

func (b *Bard) ExecutePostCalculateMethods(c *models.Character) {
	b.executeJackOfAllTrades(c)
	b.executeExpertise(c)
	b.executeSpellCastingAbility(c)
	b.executeBardicInspiration(c)
}

func (b *Bard) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (b *Bard) executeSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(bardSpellCastingAbility)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

func (b *Bard) executeBardicInspiration(c *models.Character) {
	if b.ClassToken.Name == "" {
		return
	} else if b.ClassToken.Name != bardicInspirationToken {
		logger.Info("Invalid Class Token Name")
		return
	}

	b.ClassToken.Maximum = c.GetMod(shared.AbilityCharisma)
}

// At level 3, bards can pick two skills they are proficient in, and double the modifier.
// They select two more at level 10
func (b *Bard) executeExpertise(c *models.Character) {
	if c.Level < 3 {
		return
	}

	if c.Level < 10 && len(b.ExpertiseSkills) > 2 {
		// We'll allow the user to specify more, but only the first two get taken for it to be ExpertiseSkills
		b.ExpertiseSkills = b.ExpertiseSkills[:2]
	}

	if c.Level >= 10 && len(b.ExpertiseSkills) > 4 {
		// We'll allow the user to specify more, but only the first two get taken for it to be ExpertiseSkills
		b.ExpertiseSkills = b.ExpertiseSkills[:4]
	}

	executeExpertiseShared(c, b.ExpertiseSkills)
}

// At level 2, bards can add half their proficiency bonus (rounded down) to any ability check
// that doesn't already use their proficiency bonus.
func (b *Bard) executeJackOfAllTrades(c *models.Character) {
	if c.Level < 2 {
		return
	}

	for i, skill := range c.Skills {
		if !skill.Proficient {
			c.Skills[i].SkillModifier += int(math.Floor(float64(c.Proficiency / 2)))
		}
	}
}

func (b *Bard) SubClass(level int) string {
	if level <= 2 {
		return ""
	}

	return b.College
}

func (b *Bard) ClassDetails(level int) string {
	var s string
	s += formatTokens(b.ClassToken, bardicInspirationToken, level) + "\n"

	if len(b.ExpertiseSkills) > 0 && level >= 3 {
		expertiseHeader := fmt.Sprintf("Expertise:\n")
		s += expertiseHeader

		for _, exp := range b.ExpertiseSkills {
			expLine := fmt.Sprintf("- %s\n", exp)
			s += expLine
		}

		s += "\n"
	}

	return s
}

func (b *Bard) ClassFeatures(level int) string {
	var s string
	s += formatOtherFeatures(b.OtherFeatures, level)

	return s
}

// CLI

func (b *Bard) UseClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since bard only has bardic inspiration, we won't check the slot name value
	if b.ClassToken.Available <= 0 {
		logger.Info("Bardic Inpsiration had no uses left")
		return
	}

	b.ClassToken.Available -= quantity
}

func (b *Bard) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since bard only has bardic inspiration, we won't check the slot name value
	b.ClassToken.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || b.ClassToken.Available > b.ClassToken.Maximum {
		b.ClassToken.Available = b.ClassToken.Maximum
	}
}

func (b *Bard) GetTokens() []string {
	return []string{
		bardicInspirationToken,
	}
}
