package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
)

type Rogue struct {
	models.BaseClass
	ExpertiseSkills []string `json:"expertise" clover:"expertise"`
	SneakAttack     string   `json:"-" clover:"-"`
}

func LoadRogue(data []byte) (*Rogue, error) {
	var ranger Rogue
	if err := json.Unmarshal(data, &ranger); err != nil {
		return &ranger, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &ranger, nil
}

func (r *Rogue) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (r *Rogue) ExecutePostCalculateMethods(c *models.Character) {
	r.executeExpertise(c)
	r.executeSneakAttack(c)
}

// At level 1, rogues can pick two skills they are proficient in, and double the modifier.
// They select two more at level 6
func (r *Rogue) executeExpertise(c *models.Character) {
	if c.Level < 6 && len(r.ExpertiseSkills) > 2 {
		// We'll allow the user to specify more, but only the first two get taken for it to be ExpertiseSkills
		r.ExpertiseSkills = r.ExpertiseSkills[:2]
	}

	if c.Level >= 6 && len(r.ExpertiseSkills) > 4 {
		// We'll allow the user to specify more, but only the first two get taken for it to be ExpertiseSkills
		r.ExpertiseSkills = r.ExpertiseSkills[:4]
	}
	executeExpertiseShared(c, r.ExpertiseSkills)
}

func (r *Rogue) executeSneakAttack(c *models.Character) {
	switch {
	case c.Level < 3:
		r.SneakAttack = "1d6"
	case c.Level < 5:
		r.SneakAttack = "2d6"
	case c.Level < 7:
		r.SneakAttack = "3d6"
	case c.Level < 9:
		r.SneakAttack = "4d6"
	case c.Level < 11:
		r.SneakAttack = "5d6"
	case c.Level < 13:
		r.SneakAttack = "6d6"
	case c.Level < 15:
		r.SneakAttack = "7d6"
	case c.Level < 17:
		r.SneakAttack = "8d6"
	case c.Level < 19:
		r.SneakAttack = "9d6"
	case c.Level >= 19:
		r.SneakAttack = "10d6"
	}
}

func (r *Rogue) ClassDetails(level int) string {
	var s string

	if len(r.ExpertiseSkills) > 0 {
		expertiseHeader := fmt.Sprintf("Expertise:\n")
		s += expertiseHeader

		for _, exp := range r.ExpertiseSkills {
			expLine := fmt.Sprintf("- %s\n", exp)
			s += expLine
		}

		s += "\n"
	}

	sneakAttackLine := fmt.Sprintf("*Sneak Attack*: %s\n\n", r.SneakAttack)
	s += sneakAttackLine

	return s
}
