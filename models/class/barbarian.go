package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
)

type Barbarian struct {
	Path            string                 `json:"path"`
	OtherFeatures   []models.ClassFeatures `json:"other-features"`
	Rage            Rage                   `json:"rage"`
	PrimalKnowledge []string               `json:"primal-knowledge"`
}

type Rage struct {
	Available int `json:"available"`
	Maximum   int `json:"maximum"`
	Damage    int `json:"damage"`
}

func LoadBarbarian(data []byte) (*Barbarian, error) {
	var barbarian Barbarian
	if err := json.Unmarshal(data, &barbarian); err != nil {
		err = fmt.Errorf("Failed to parse class data: %w", err)
		panic(err)
	}

	return &barbarian, nil
}

func (b *Barbarian) ValidateMethods(c *models.Character) {
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

	avaliableSkills := []string{
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
					info := fmt.Sprintf("Primal Knowledge Skill %s was not one of the six available skills at level 1: %s\n",
						pk,
						strings.Join(avaliableSkills, ", "))

					logger.HandleInfo(info)
				}
			}
		}
	}
}

// If not wearing armor, Armor Class is boosted to 10 + dex mod + constitution mod
func (b *Barbarian) executeUnarmoredDefense(c *models.Character) {
	if c.WornEquipment.Armour != "" {
		return
	}

	c.AC = 0

	for _, a := range c.Abilities {
		if strings.ToLower(a.Name) != "dexterity" && strings.ToLower(a.Name) != "constitution" {
			continue
		}

		c.AC += a.AbilityModifier
	}

	c.AC += 10
}

func (b *Barbarian) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if b.Rage.Available != 0 && b.Rage.Maximum != 0 {
		rageSlots := c.GetSlots(b.Rage.Available, b.Rage.Maximum)
		rageLine := fmt.Sprintf("**Rage**: %s - Damage: +%d\n\n", rageSlots, b.Rage.Damage)
		s = append(s, rageLine)
	}

	if b.Path != "" && c.Level > 3 {
		pathHeader := fmt.Sprintf("Primal Path: *%s*\n\n", b.Path)
		s = append(s, pathHeader)
	}

	if len(b.OtherFeatures) > 0 {
		for _, detail := range b.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			detailHeader := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, detailHeader)
			detail := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, detail)
		}
	}

	return s
}

// At level 20, your Strength and Constitution scores increase by 4. Your maximum for those scores is now 24.
func (b *Barbarian) executePrimalChampion(c *models.Character) {
	if c.Level < 20 {
		return
	}

	for i, a := range c.Abilities {
		if strings.ToLower(a.Name) != "strength" && strings.ToLower(a.Name) != "constitution" {
			continue
		}

		c.Abilities[i].Base += 4
	}
}

// CLI

func (b *Barbarian) UseClassTokens(tokenName string) {
	// We only really need token name for classes that have multiple tokens
	// since barbarian only has rage, we won't check the token name value
	if b.Rage.Available <= 0 {
		logger.HandleInfo("Rage had no uses left")
		return
	}

	b.Rage.Available--
}

func (b *Barbarian) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need token name for classes that have multiple tokens
	// since barbarian only has rage, we won't check the token name value
	b.Rage.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || b.Rage.Available > b.Rage.Maximum {
		b.Rage.Available = b.Rage.Maximum
	}
}
