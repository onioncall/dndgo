package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

const (
	InvocationArmorOfShadows       = "armor of shadows"
	InvocationFiendishVigor        = "fiendish vigor"
	InvocationGiftOfEverLivingOnes = "gift of the ever-living ones"
	InvocationAgonizingBlast       = "agonizing blast"
	InvocationLifedrinker          = "lifedrinker"
	InvocationImprovedPactWeapon   = "improved pact weapon"
)

type Warlock struct {
	BaseClass
	OtherworldlyPatron string   `json:"otherworldly-patron"`
	Invocations        []string `json:"invocations"`
}

func LoadWarlock(data []byte) (*Warlock, error) {
	var warlock Warlock
	if err := json.Unmarshal(data, &warlock); err != nil {
		return &warlock, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &warlock, nil
}

func (w *Warlock) ExecutePostCalculateMethods(c *models.Character) {
	w.executeSpellCastingAbility(c)
	w.executeEldritchInvocations(c)
}

func (w *Warlock) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (w *Warlock) executeSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(shared.AbilityCharisma)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

// May implement more thoroughly in the future, but most of these invove game state that we can't mock
// in this app. Will look into in the future when I know more about how this class plays
func (w *Warlock) executeEldritchInvocations(c *models.Character) {
	// if c.Level > 2 {
	// 	return
	// }
	//
	// switch w.Invocation {
	// case InvocationArmorOfShadows:
	// 	applyArmorOfShadows(c)
	// }
}

func applyArmorOfShadows(c *models.Character) bool {
	if c.WornEquipment.Armor.Name != "" {
		return false
	}

	dexMod := c.GetMod(shared.AbilityDexterity)
	armorOfShadows := 13 + dexMod

	if !c.ValidationDisabled {
		if c.AC > armorOfShadows {
			logger.Info(fmt.Sprintf("Armor of Shadows AC (%d) lower than characters current AC (%d)",
				armorOfShadows, c.AC))
			return false
		}
	}

	c.AC = armorOfShadows
	return true
}

func (w *Warlock) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	if w.OtherworldlyPatron != "" && c.Level > 3 {
		s = append(s, fmt.Sprintf("Otherwordly Patron: *%s*\n\n", w.OtherworldlyPatron))
	}

	if len(w.Invocations) > 0 && c.Level > 3 {
		s = append(s, "Invocation:\n\n")
		for _, invocation := range w.Invocations {
			s = append(s, fmt.Sprintf("%s\n", invocation))
		}
	}

	if len(w.OtherFeatures) > 0 {
		for _, detail := range w.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			detailName := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, detailName)
			details := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, details)
		}
	}

	return s
}
