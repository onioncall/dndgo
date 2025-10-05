package class

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/onioncall/dndgo/models"
)

type Bard struct {
	SkillProficienciesToDouble 	[]string 						`json:"expertise"`
	AbilityScoreImprovement		[]AbilityScoreImprovementItem	`json:"ability-score-improvement"`
	College 					string 							`json:"college"`
	OtherDetails 				[]ClassDetails					`json:"other-Details"`
}

type ClassDetails struct {
	Name 	string 	`json:"name"`
	Level	int		`json:"level"`
	Details string 	`json:"details"`
}

type AbilityScoreImprovementItem struct {
	Ability string 	`json:"ability"`
	Bonus	int		`json:"bonus"`
}

var preCalculateMethods []func(c *models.Character)
var postCalculateMethods []func(c *models.Character)

func LoadBard(data []byte) (*Bard, error) {
	var bard Bard
	if err := json.Unmarshal(data, &bard); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to parse character data: %w", err)
	}

	return &bard, nil
}

func (b *Bard) LoadMethods() {
}

func (b *Bard) ExecutePostCalculateMethods(c *models.Character) {
	postCalculateMethods = append(postCalculateMethods, b.jackOfAllTrades)
	postCalculateMethods = append(postCalculateMethods, b.expertise)
	for _, m := range postCalculateMethods {
		m(c)
	}
}

func (b *Bard) ExecutePreCalculateMethods(c *models.Character) {
	preCalculateMethods = append(preCalculateMethods, b.abilityScoreImprovement)
	for _, m := range preCalculateMethods {
		m(c)
	}
}

// At level 3, bards can pick two skills they are proficient in, and double the modifier. 
// They select two more at level 10
func (b *Bard) expertise(c *models.Character) {
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
			panic("Bard Config Error - Expertise can not have dupliate proficiencies")
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
func (b *Bard) jackOfAllTrades(c *models.Character) {
	if (c.Level < 2) {
		return
	}

	for i, skill := range c.Skills {
		if !skill.Proficient {
			c.Skills[i].SkillModifier += int(math.Floor(float64(c.Proficiency / 2)))	
		} 
	}
}

// At level 4, Bards may select one ability and add 2 to that score,
// or select two abilities and add 1 to each score (max of 20).
// They get to do this again at levels 8, 12, 16, and 19
func (b *Bard) abilityScoreImprovement(c *models.Character) {
	maxBonus := 0
	switch {
	case c.Level < 4:  maxBonus = 0
	case c.Level < 8:  maxBonus = 2
	case c.Level < 12: maxBonus = 4
	case c.Level < 16: maxBonus = 6
	case c.Level < 19: maxBonus = 8
	case c.Level >= 19: maxBonus = 10
	}

	if maxBonus == 0 {
		return // don't qualify yet
	}
	
	bonusSum := 0
	for _, item := range b.AbilityScoreImprovement {
		bonusSum += item.Bonus
	}

	if bonusSum > maxBonus {
		fmt.Printf("Ability Score Bonus (%d) exceeds available for level (%d)\n", bonusSum, maxBonus)
		return
	}
	
	for _, item := range b.AbilityScoreImprovement {
		for i := range c.Attributes {
			if strings.ToLower(c.Attributes[i].Name) == strings.ToLower(item.Ability) {
				c.Attributes[i].Base += item.Bonus
				c.Attributes[i].Base = min(20, c.Attributes[i].Base)
				break
			}
		}
	}
}

func (b *Bard) PrintOtherFeatures(c *models.Character) []string {
	s := make([]string, 0, 100)	
	header := fmt.Sprintf("Class Details\n")
	spacer := fmt.Sprintf("---\n")
	s = append(s, header)
	s = append(s, spacer)

	if len(b.SkillProficienciesToDouble) > 0 && c.Level >= 3 {
		expertiseHeader := fmt.Sprintf("Expertise\n")
		s = append(s, expertiseHeader)
		for _, exp := range b.SkillProficienciesToDouble {
			expLine := fmt.Sprintf("- %s\n", exp)
			s = append(s, expLine)
		}
		s = append(s, "\n")
	}

	if len(b.AbilityScoreImprovement) > 0 && c.Level >= 4 {
		abilityMap := make(map[string]int)

		// Build map
		for _, ability := range b.AbilityScoreImprovement {
			_, exists := abilityMap[ability.Ability]
			if exists {
				abilityMap[ability.Ability] += ability.Bonus
			}
		}

		abilityScoreImprovementHeader := fmt.Sprintf("Ability Score Improvement\n")
		s = append(s, abilityScoreImprovementHeader)

		// TODO: Make this more sophistocated so we don't need to loop through this twice
		for _, ability := range b.AbilityScoreImprovement {
			expLine := fmt.Sprintf("- %s: +%d\n", ability.Ability, ability.Bonus)
			s = append(s, expLine)
		}
		s = append(s, "\n")
	}

	if b.College != "" {
		collegeHeader := fmt.Sprintf("*%s*\n\n", b.College)
		s = append(s, collegeHeader)
	}

	if len(b.OtherDetails) > 0 {
		for _, detail := range b.OtherDetails {
			if detail.Level > c.Level {
				continue
			}

			collegeDetailName := fmt.Sprintf("---\n%s\n", detail.Name)
			s = append(s, collegeDetailName)
			collegeDetail := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, collegeDetail)
		}
		s = append(s, "\n")
	}

	return s
}
