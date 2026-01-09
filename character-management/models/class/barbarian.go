package class

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Barbarian struct {
	models.BaseClass
	ClassToken      shared.NamedToken `json:"class-token" clover:"class-token"`
	RageDamage      int               `json:"-" clover:"-"`
	PrimalKnowledge []string          `json:"primal-knowledge" clover:"primal-knowledge"`
}

const rageToken string = "rage"

func LoadBarbarian(data []byte) (*Barbarian, error) {
	var barbarian Barbarian
	if err := json.Unmarshal(data, &barbarian); err != nil {
		err = fmt.Errorf("Failed to parse class data: %w", err)
		panic(err)
	}

	return &barbarian, nil
}

func (b *Barbarian) ExecutePostCalculateMethods(c *models.Character) {
	b.executeRage(c)
	b.executeUnarmoredDefense(c)
	b.executePrimalKnowledge(c)
}

func (b *Barbarian) ExecutePreCalculateMethods(c *models.Character) {
	b.executePrimalChampion(c)
}

func (b *Barbarian) CalculateHitDice() string {
	return fmt.Sprintf("%dd12", b.Level)
}

func (b *Barbarian) executeRage(c *models.Character) {
	if b.ClassToken.Name == "" {
		return
	} else if b.ClassToken.Name != rageToken {
		logger.Info("Invalid Class Token Name")
		return
	}

	switch {
	case c.Level < 3:
		b.ClassToken.Maximum = 2
	case c.Level < 6:
		b.ClassToken.Maximum = 3
	case c.Level < 12:
		b.ClassToken.Maximum = 4
	case c.Level < 17:
		b.ClassToken.Maximum = 5
	case c.Level < 20:
		b.ClassToken.Maximum = 6
	case c.Level >= 20:
		b.ClassToken.Maximum = 0 // unlimited
	}

	// Unfortunately these don't line up and putting them in the same switch is gross
	switch {
	case c.Level < 9:
		b.RageDamage = 2
	case c.Level < 16:
		b.RageDamage = 3
	case c.Level > 12:
		b.RageDamage = 4
	}
}

// At level 3, You gain proficiency in one skill of your choice from the list of skills
// available to barbarians at 1st level.
func (b *Barbarian) executePrimalKnowledge(c *models.Character) {
	if b.Level < 3 {
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

					logger.Info(info)
				}
			}
		}
	}
}

// If not wearing armor, Armor Class is boosted to 10 + dex mod + constitution mod
func (b *Barbarian) executeUnarmoredDefense(c *models.Character) {
	barbarianExpertiseAbilityModifiers := []string{
		shared.AbilityDexterity,
		shared.AbilityConstitution,
	}

	executeUnarmoredDefenseShared(c, barbarianExpertiseAbilityModifiers)
}

func (b *Barbarian) ClassDetails() string {
	var s string

	rageSlots := formatTokens(b.ClassToken, rageToken, b.Level)
	rageLine := fmt.Sprintf("**Rage**: %s - Damage: +%d\n", rageSlots, b.RageDamage)
	s += rageLine

	return s
}

// At level 20, your Strength and Constitution scores increase by 4. Your maximum for those scores is now 24.
func (b *Barbarian) executePrimalChampion(c *models.Character) {
	if b.Level < 20 {
		return
	}

	for i, a := range c.Abilities {
		if strings.ToLower(a.Name) != "strength" && strings.ToLower(a.Name) != "constitution" {
			continue
		}

		c.Abilities[i].Adjusted += 4
	}
}

// CLI

func (b *Barbarian) UseClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != rageToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, b.ClassType))
		return
	}

	if b.ClassToken.Available <= 0 {
		logger.Info("Rage had no uses left")
		return
	}

	b.ClassToken.Available -= quantity
}

func (b *Barbarian) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != rageToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, b.ClassType))
		return
	}

	b.ClassToken.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || b.ClassToken.Available > b.ClassToken.Maximum {
		b.ClassToken.Available = b.ClassToken.Maximum
	}
}

func (b *Barbarian) GetTokens() []string {
	return []string{
		"rage",
	}
}
