package class

import (
	"encoding/json"
	"fmt"
	"math"
	"strings"

	"github.com/onioncall/dndgo/models"
)

type Bard struct {
	SkillProficienciesToDouble 	[]string 			`json:"expertise"`
	College 					College 			`json:"college"`
	OtherFeatures 				[]NameDetailPair	`json:"other-features"`
}

type College struct {
	Name	string 				`json:"name"`
	Details []NameDetailPair 	`json:"other-details"`
}

type NameDetailPair struct {
	Name 	string `json:"name"`
	Details string `json:"details"`
}

var preBuildMethods []func(c *models.Character)
var postBuildMethods []func(c *models.Character)

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
	preBuildMethods = append(preBuildMethods, b.jackOfAllTrades)
	preBuildMethods = append(preBuildMethods, b.expertise)
	for _, m := range preBuildMethods {
		m(c)
	}
}

func (b *Bard) ExecutePreCalculateMethods(c *models.Character) {
}

func (b *Bard) TestPrint() {
	fmt.Println("success")
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
				fmt.Println(cs.Name)
				fmt.Printf("Prof: %d\n", c.Proficiency)
				fmt.Printf("SkillMod: %d\n\n", c.Skills[i].SkillModifier)
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

func (b *Bard) PrintOtherFeatures() []string {
	if b.College.Name == "" {
		return nil
	}

	s := make([]string, 0, 100)	
	header := fmt.Sprintf("Sub-Class Details\n")
	spacer := fmt.Sprintf("---\n")
	s = append(s, header)
	s = append(s, spacer)

	expertiseHeader := fmt.Sprintf("Expertise\n")
	s = append(s, expertiseHeader)
	for _, exp := range b.SkillProficienciesToDouble {
		expLine := fmt.Sprintf("- %s\n", exp)
		s = append(s, expLine)
	}
	s = append(s, "\n")

	collegeHeader := fmt.Sprintf("*%s*\n\n", b.College.Name)
	s = append(s, collegeHeader)

	for _, collegeDetail := range b.College.Details {
		collegeDetailName := fmt.Sprintf("%s\n", collegeDetail.Name)
		s = append(s, collegeDetailName)
		collegeDetail := fmt.Sprintf("%s\n", collegeDetail.Details)
		s = append(s, collegeDetail)
	}
	s = append(s, "\n")
	
	// otherDetailHeader := fmt.Sprintf("*Other Class Details*\n\n")
	// s = append(s, otherDetailHeader)
	// for _, feature := range b.OtherFeatures {
	// 	otherDetailName := fmt.Sprintf("%s\n", feature.Name)
	// 	s = append(s, otherDetailName)
	// 	otherDetail := fmt.Sprintf("%s\n", feature.Details)
	// 	s = append(s, otherDetail)
	// }
	
	return s	
}
