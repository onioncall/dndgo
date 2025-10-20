package class

import (
	"strings"

	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
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
