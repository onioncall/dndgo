package class

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/logger"
)

func executeExpertiseShared(c *models.Character, expertiseSkills []string) {
	if c.Level < 3 {
		return
	}

	if c.Level < 10 && len(expertiseSkills) > 2 {
		// We'll allow the user to specify more, but only the first two get taken for it to be legal
		expertiseSkills = expertiseSkills[:2]
	}

	if c.Level > 10 && len(expertiseSkills) > 4 {
		// We'll allow the user to specify more, but only the first four get taken for it to be legal
		expertiseSkills = expertiseSkills[:4]
	}

	seen := make(map[string]bool)
	for _, profToDouble := range expertiseSkills {
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

// If not wearing armor, Armor Class is boosted to 10 + modifiers outlined by implementing class
func executeUnarmoredDefenseShared(c *models.Character, abilities []string) {
	if c.WornEquipment.Armour != "" {
		return
	}

	c.AC = 0

	for _, charAbility := range c.Abilities {
		for _, classAbility := range abilities {
			if charAbility.Name == classAbility {
				c.AC += charAbility.AbilityModifier
			}
		}
	}

	c.AC += 10
}

func buildClassDetailsHeader() []string {
	s := make([]string, 0, 100)
	header := fmt.Sprintf("Class Details\n")
	spacer := fmt.Sprintf("---\n")
	s = append(s, header)
	s = append(s, spacer)

	return s
}
