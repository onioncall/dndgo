package class

import (
	"encoding/json"
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Bard struct {
	models.BaseClass
	ExpertiseSkills []string          `json:"expertise" clover:"expertise"`
	ClassToken      shared.NamedToken `json:"class-token" clover:"class-token"`
}

const (
	bardicInspirationToken  string = "bardic-inspiration"
	bardSpellCastingAbility string = shared.AbilityCharisma
)

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

func (b *Bard) CalculateHitDice() string {
	return fmt.Sprintf("%dd8", b.Level)
}

func (b *Bard) executeSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(bardSpellCastingAbility)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

func (b *Bard) executeBardicInspiration(c *models.Character) {
	if b.ClassToken.Name != bardicInspirationToken {
		logger.Info("Invalid Class Token Name")
		return
	}

	b.ClassToken.Maximum = c.GetMod(shared.AbilityCharisma)
}

// At b.Level 3, bards can pick two skills they are proficient in, and double the proficiency.
// They select two more at b.Level 10
func (b *Bard) executeExpertise(c *models.Character) {
	if c.Level < 3 {
		return
	}

	if c.Level < 10 && len(b.ExpertiseSkills) > 2 {
		logger.Warn("Only two expertise skills should be configured for your class b.Level")
	}

	if c.Level >= 10 && len(b.ExpertiseSkills) > 4 {
		logger.Warn("Only four expertise skills should be configured for your class b.Level")
	}

	executeExpertiseShared(c, b.ExpertiseSkills)
}

// At b.Level 2, bards can add half their proficiency bonus (rounded down) to any ability check
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

func (b *Bard) ClassDetails() string {
	var s string
	s += fmt.Sprintf("Level: %d\n", b.Level)
	s += formatTokens(b.ClassToken, bardicInspirationToken, b.Level) + "\n"

	if len(b.ExpertiseSkills) > 0 && b.Level >= 3 {
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

// CLI

func (b *Bard) UseClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != bardicInspirationToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, b.ClassType))
		return
	}

	if b.ClassToken.Available <= 0 {
		logger.Info("Bardic Inpsiration had no uses left")
		return
	}

	b.ClassToken.Available -= quantity
}

func (b *Bard) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != bardicInspirationToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, b.ClassType))
		return
	}

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

func (b *Bard) AddExpertiseSkill(skill string) error {
	if !slices.Contains(shared.Skills, strings.ToLower(skill)) {
		return fmt.Errorf("Skill '%s' does not exist, check spelling.", skill)
	}

	if slices.Contains(b.ExpertiseSkills, strings.ToLower(skill)) {
		return fmt.Errorf("Duplicate skill '%s' cannot be added, choose unique one.", skill)
	}

	b.ExpertiseSkills = append(b.ExpertiseSkills, skill)

	return nil
}
