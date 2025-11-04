package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/logger"
)

type Monk struct {
	MartialArts     string                 `json:"-"`
	KiPoints        Ki                     `json:"ki-points"`
	MosaicTradition string                 `json:"mosaic-tradition"`
	DeflectMissles  int                    `json:"-"`
	OtherFeatures   []models.ClassFeatures `json:"other-features"`
}

type Ki struct {
	KiSpellSaveDC int
	Maximum       int
	Available     int
}

func LoadMonk(data []byte) (*Monk, error) {
	var monk Monk
	if err := json.Unmarshal(data, &monk); err != nil {
		err = fmt.Errorf("Failed to parse class data: %w", err)
		panic(err)
	}

	return &monk, nil
}

func (m *Monk) ValidateMethods(c *models.Character) {
}

// func (m *Monk) ExecutePostCalculateMethods(c *models.Character) {
// 	m.PostCalculateUnarmoredDefense(c)
// 	m.PostCalculateMartialArts(c)
// 	m.PostCalculateUnarmoredMovement(c)
// 	m.PostCalculateDeflectMissles(c)
// 	m.PostCalculateKiPoints(c)
// }
//
// func (m *Monk) ExecutePreCalculateMethods(c *models.Character) {
// 	m.PreCalculateDiamondSoul(c)
// }

func (m *Monk) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd8", level)
}

// If not wearing armor, Armor Class is boosted to 10 + dex mod + wisdom mod
func (m *Monk) PostCalculateUnarmoredDefense(c *models.Character) {
	monkExpertiseAbilityModifiers := []string{
		types.AbilityDexterity,
		types.AbilityWisdom,
	}

	executeUnarmoredDefenseShared(c, monkExpertiseAbilityModifiers)
}

func (m *Monk) PostCalculateUnarmoredMovement(c *models.Character) {
	if c.WornEquipment.Armor.Name != "" || c.Level < 2 {
		return
	}

	c.Speed += 10
}

func (m *Monk) PostCalculateMartialArts(c *models.Character) {
	switch {
	case c.Level < 5:
		m.MartialArts = "1d4"
	case c.Level < 11:
		m.MartialArts = "1d6"
	case c.Level < 17:
		m.MartialArts = "1d8"
	case c.Level >= 17:
		m.MartialArts = "1d10"
	}
}

func (m *Monk) PostCalculateKiPoints(c *models.Character) {
	if c.Level < 2 {
		return
	}

	m.KiPoints.Maximum = c.Level
	m.KiPoints.Available = min(m.KiPoints.Available, m.KiPoints.Maximum)

	widomMod := 0
	for _, a := range c.Abilities {
		if a.Name == types.AbilityWisdom {
			widomMod = a.AbilityModifier
		}
	}

	m.KiPoints.KiSpellSaveDC = 8 + c.Proficiency + widomMod
}

func (m *Monk) PostCalculateDeflectMissles(c *models.Character) {
	if c.Level < 3 {
		return
	}

	m.DeflectMissles = (10 + c.Proficiency + c.Level) * -1
}

func (m *Monk) PreCalculateDiamondSoul(c *models.Character) {
	if c.Level < 14 {
		return
	}

	for i := range c.Abilities {
		c.Abilities[i].SavingThrowsProficient = true
	}
}

func (m *Monk) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	martialArts := fmt.Sprintf("*Martial Arts*: %s\n\n", m.MartialArts)
	s = append(s, martialArts)

	kiPoints := fmt.Sprintf("*Ki Points*: %d/%d\n\n", m.KiPoints.Available, m.KiPoints.Maximum)
	s = append(s, kiPoints)

	kiSpellSaveDC := fmt.Sprintf("*Ki Spell Save DC*: %d\n\n", m.KiPoints.KiSpellSaveDC)
	s = append(s, kiSpellSaveDC)

	if m.DeflectMissles > 0 {
		deflectMissles := fmt.Sprintf("*Deflect Missles Damage Reduction*: %d", m.DeflectMissles)
		s = append(s, deflectMissles)
	}

	if c.Level > 3 {
		mosaicTradition := fmt.Sprintf("*Mosaic Tradition*: %s\n\n", m.MosaicTradition)
		s = append(s, mosaicTradition)
	}

	if len(m.OtherFeatures) > 0 {
		for _, detail := range m.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			name := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, name)
			detail := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, detail)
		}
	}

	return s
}

// CLI

func (m *Monk) UseClassTokens(tokenName string) {
	// We only really need slot name for classes that have multiple slots
	// since monk only has ki points, we won't check the slot name value
	if m.KiPoints.Available <= 0 {
		logger.HandleInfo("No Ki Points Available")
		return
	}

	m.KiPoints.Available--
}

func (m *Monk) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since monk only has ki points, we won't check the slot name value
	m.KiPoints.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || m.KiPoints.Available > m.KiPoints.Maximum {
		m.KiPoints.Available = m.KiPoints.Maximum
	}
}
