package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
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
	OtherworldlyPatron string                 `json:"otherworldly-patron"`
	Invocation         string                 `json:"invocation"`
	OtherFeatures      []models.ClassFeatures `json:"other-features"`
}

func LoadWarlock(data []byte) (*Warlock, error) {
	var warlock Warlock
	if err := json.Unmarshal(data, &warlock); err != nil {
		return &warlock, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &warlock, nil
}

func (w *Warlock) ValidateMethods(c *models.Character) {
}

func (w *Warlock) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

func (w *Warlock) PostCalculateSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(types.AbilityCharisma)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

// May implement more thoroughly in the future, but most of these invove game state that we can't mock
// in this app. Will look into in the future when I know more about how this class plays
func (w *Warlock) PostCalculateEldritchInvocations(c *models.Character) {
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

	dexMod := c.GetMod(types.AbilityDexterity)
	armorOfShadows := 13 + dexMod

	if !c.ValidationDisabled {
		if c.AC > armorOfShadows {
			logger.HandleInfo(fmt.Sprintf("Armor of Shadows AC (%d) lower than characters current AC (%d)",
				armorOfShadows, c.AC))
			return false
		}
	}

	c.AC = armorOfShadows
	return true
}
