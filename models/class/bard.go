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
	preBuildMethods = append(preBuildMethods, b.JackOfAllTrades)
	preBuildMethods = append(preBuildMethods, b.Expertise)
	for _, m := range preBuildMethods {
		m(c)
	}
}

func (b *Bard) ExecutePreCalculateMethods(c *models.Character) {
}

func (b *Bard) TestPrint() {
	fmt.Println("success")
}

func (b *Bard) Expertise(c *models.Character) {
	if c.Level < 3 {
		return
	}

	for _, profToDouble := range b.SkillProficienciesToDouble {
		for i, cs := range c.Skills {
			if strings.ToLower(cs.Name) == strings.ToLower(profToDouble) {
				c.Skills[i].Bonus += c.Proficiency
			}
		}
	}
}

func (b *Bard) JackOfAllTrades(c *models.Character) {
	if (c.Level < 2) {
		return
	}

	for i, skill := range c.Skills {
		if !skill.Proficient {
			c.Skills[i].Bonus += int(math.Floor(float64(c.Proficiency / 2)))	
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
