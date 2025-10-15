package class

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
)

type Bard struct {
	SkillProficienciesToDouble 	[]string 						`json:"expertise"`
	College 					string 							`json:"college"`
	OtherFeatures 				[]models.ClassFeatures			`json:"other-features"`
	BardicInspiration			BardicInspiration				`json:"bardic-inspiration"`
}

type BardicInspiration struct {
	Available	int	`json:"available"`
	Slot		int	`json:"slot"`
}

func LoadBard(data []byte) (*Bard, error) {
	var bard Bard
	if err := json.Unmarshal(data, &bard); err != nil {
		errLog := fmt.Errorf("Failed to parse class data: %w", err)
		logger.HandleError(errLog, err)

		return nil, err
	}

	return &bard, nil
}

func (b *Bard) LoadMethods() {
}

func (b *Bard) ExecutePostCalculateMethods(c *models.Character) {
	models.PostCalculateMethods = append(models.PostCalculateMethods, b.executeJackOfAllTrades)
	models.PostCalculateMethods = append(models.PostCalculateMethods, b.executeExpertise)
	for _, m := range models.PostCalculateMethods {
		m(c)
	}
}

func (b *Bard) ExecutePreCalculateMethods(c *models.Character) {
	for _, m := range models.PreCalculateMethods {
		m(c)
	}
}

// At level 3, bards can pick two skills they are proficient in, and double the modifier. 
// They select two more at level 10
func (b *Bard) executeExpertise(c *models.Character) {
	if c.Level < 3 {
		return
	}
	
	if c.Level < 10 && len(b.SkillProficienciesToDouble) > 2 {
		// We'll allow the user to specify more, but only the first two get taken for it to be legal
		b.SkillProficienciesToDouble = b.SkillProficienciesToDouble[:2]
	} 

	if c.Level > 10 && len(b.SkillProficienciesToDouble) > 4 {
		// We'll allow the user to specify more, but only the first four get taken for it to be legal
		b.SkillProficienciesToDouble = b.SkillProficienciesToDouble[:2]
	}

	seen := make(map[string]bool)
	for _, profToDouble := range b.SkillProficienciesToDouble {
		if seen[profToDouble] == true {
			logger.HandleInfo("Bard Config Error - Expertise can not have dupliate proficiencies")
			return
		}
		seen[profToDouble] = true

		for i, cs := range c.Skills {
			if strings.ToLower(cs.Name) == strings.ToLower(profToDouble) {
				c.Skills[i].SkillModifier += c.Proficiency
			}
		}
	}
}

// At level 2, bards can add half their proficiency bonus (rounded down) to any ability check 
// that doesn't already use their proficiency bonus.
func (b *Bard) executeJackOfAllTrades(c *models.Character) {
	if (c.Level < 2) {
		return
	}

	for i, skill := range c.Skills {
		if !skill.Proficient {
			c.Skills[i].SkillModifier += int(math.Floor(float64(c.Proficiency / 2)))	
		} 
	}
}

func (b *Bard) PrintClassDetails(c *models.Character) []string {
	s := c.BuildClassDetailsHeader()

	if b.College != "" {
		collegeHeader := fmt.Sprintf("College: *%s*\n\n", b.College)
		s = append(s, collegeHeader)
	}

	if b.BardicInspiration.Available != 0 && b.BardicInspiration.Slot != 0 {
		bardicSlots := c.GetSlots(b.BardicInspiration.Available, b.BardicInspiration.Slot)
		biLine := fmt.Sprintf("**Bardic Inspiration**: %s\n\n", bardicSlots)
		s = append(s, biLine)
	}

	if len(b.SkillProficienciesToDouble) > 0 && c.Level >= 3 {
		expertiseHeader := fmt.Sprintf("Expertise\n")
		s = append(s, expertiseHeader)
		for _, exp := range b.SkillProficienciesToDouble {
			expLine := fmt.Sprintf("- %s\n", exp)
			s = append(s, expLine)
		}
		s = append(s, "\n")
	}

	if len(b.OtherFeatures) > 0 {
		for _, detail := range b.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			collegeDetailName := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, collegeDetailName)
			collegeDetail := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, collegeDetail)
		}
	}

	return s
}

// CLI

func (b *Bard) UseClassTokens(tokenName string) {
	// We only really need slot name for classes that have multiple slots
	// since bard only has bardic inspiration, we won't check the slot name value
	if b.BardicInspiration.Available <= 0 {
		logger.HandleInfo("No Bardic Inspiration tokens left")
		return
	} 

	b.BardicInspiration.Available--
}

func (b *Bard) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since bard only has bardic inspiration, we won't check the slot name value
	b.BardicInspiration.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || b.BardicInspiration.Available > b.BardicInspiration.Slot {
		b.BardicInspiration.Available = b.BardicInspiration.Slot
	}
}
