package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/models"
)

type Barbarian struct {
	Path			string					`json:"path"`
	OtherFeatures	[]models.ClassFeatures	`json:"other-features"`
	Rage			Rage					`json:"rage"`
	PrimalKnowledge	[]string				`json:"primal-knowledge"`
}

type Rage struct {
	Available	int	`json:"available"`
	Slot		int	`json:"slot"`
	Damage		int `json:"damage"`
}

func LoadBarbarian(data []byte) (*Barbarian, error) {
	var barbarian Barbarian
	if err := json.Unmarshal(data, &barbarian); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to parse character data: %w", err)
	}

	return &barbarian, nil
}

func (b *Barbarian) LoadMethods() {
	fmt.Println("TestPrint")
}

func (b *Barbarian) ExecutePostCalculateMethods(c *models.Character) {
	models.PostCalculateMethods = append(models.PostCalculateMethods, b.executeUnarmoredDefense)
	for _, m := range models.PostCalculateMethods {
		m(c)
	}
}

func (b *Barbarian) ExecutePreCalculateMethods(c *models.Character) {
	models.PreCalculateMethods = append(models.PreCalculateMethods, b.executePrimalChampion)
	for _, m := range models.PreCalculateMethods {
		m(c)
	}
}

// At level 3, You gain proficiency in one skill of your choice from the list of skills
// available to barbarians at 1st level. 
func (b *Barbarian) executePrimalKnowledge(c *models.Character) {
	if c.Level < 3 {
		return
	}

	avaliableSkills := []string {
		"Animal Handling",
		"Athletics",
		"Intimidation",
		"Nature",
		"Perception",
		"Survival",
	}
	

	// TODO: Refactor this to make it more performant, the three loops are going to be the least efficient way to handle this
	for _, pk := range b.PrimalKnowledge {
		// Find matching skill for each skill in primal primalKnowledge()
		for _, s := range c.Skills {
			// ensure that we only run this for skills available to barbarian at first level
			for i, as := range avaliableSkills {
				if as == pk && pk == s.Name {
					c.Skills[i].Proficient = true
					break
				} else if i == len(avaliableSkills) {
					fmt.Printf("Primal Knowledge Skill %s was not one of the six available skills at level 1: %s\n",
					pk, 
					strings.Join(avaliableSkills, ", "))
				}
			}
		}
	}
}

// If not wearing armor, Armor Class is boosted to 10 + dex mod + constitution mod
func (b *Barbarian) executeUnarmoredDefense(c *models.Character) {
	if c.BodyEquipment.Armour != "" {
		return
	}

	c.AC = 0
	
	for _, a := range c.Attributes {
		if strings.ToLower(a.Name) != "dexterity" && strings.ToLower(a.Name) != "constitution" {
			continue
		}

		c.AC += a.AbilityModifier
	}

	c.AC += 10
}

func (b *Barbarian) PrintClassDetails(c *models.Character) []string {
	s := c.BuildClassDetailsHeader()

	if b.Rage.Available != 0 && b.Rage.Slot != 0 {
		rageSlots := c.GetSlots(b.Rage.Available, b.Rage.Slot)
		rageLine := fmt.Sprintf("**Rage**: %s - Damage: +%d\n\n", rageSlots, b.Rage.Damage)
		s = append(s, rageLine)
	}

	if b.Path != "" {
		pathHeader := fmt.Sprintf("Primal Path: *%s*\n\n", b.Path)
		s = append(s, pathHeader)
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

// At level 20, your Strength and Constitution scores increase by 4. Your maximum for those scores is now 24. 
func (b *Barbarian) executePrimalChampion(c *models.Character) {
	if c.Level < 20 {
		return
	}

	for i, a := range c.Attributes {
		if strings.ToLower(a.Name) != "strength" && strings.ToLower(a.Name) != "constitution" {
			continue
		}

		c.Attributes[i].Base += 4
	}
}

// CLI

func (b *Barbarian) UseClassSlots(slotName string) {
	// We only really need slot name for classes that have multiple slots
	// since barbarian only has rage, we won't check the slot name value
	if b.Rage.Available <= 0 {
		fmt.Println("Class slot had no uses left")
		return
	} 

	b.Rage.Available--
}

func (b *Barbarian) RecoverClassSlots(slotName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since barbarian only has rage, we won't check the slot name value
	b.Rage.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || b.Rage.Available > b.Rage.Slot {
		b.Rage.Available = b.Rage.Slot
	}
}
