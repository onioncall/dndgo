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

func (w *Warlock) PostCalculateSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(types.AbilityCharisma)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

func (w *Warlock) PostCalculateEldritchInvocations(c *models.Character) {
	if c.Level > 2 {
		return
	}

	switch w.Invocation {
	case InvocationArmorOfShadows:
	case InvocationFiendishVigor:
	case InvocationGiftOfEverLivingOnes:
	case InvocationAgonizingBlast:
	case InvocationLifedrinker:
	case InvocationImprovedPactWeapon:
	}
}

func applyArmorOfShadows(c *models.Character) bool {
	if c.WornEquipment.Armour != "" {
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
