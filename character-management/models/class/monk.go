package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Monk struct {
	models.BaseClass
	MartialArts    string            `json:"-" clover:"-"`
	ClassToken     shared.NamedToken `json:"class-token" clover:"class-token"`
	KiSpellSaveDC  int               `json:"-" clover:"-"`
	DeflectMissles int               `json:"-" clover:"-"`
}

const kiPointsToken string = "ki-points"

func LoadMonk(data []byte) (*Monk, error) {
	var monk Monk
	if err := json.Unmarshal(data, &monk); err != nil {
		err = fmt.Errorf("Failed to parse class data: %w", err)
		panic(err)
	}

	return &monk, nil
}

func (m *Monk) ExecutePostCalculateMethods(c *models.Character) {
	m.executeUnarmoredDefense(c)
	m.executeMartialArts(c)
	m.executeUnarmoredMovement(c)
	m.executeDeflectMissles(c)
	m.executeKiPoints(c)
}

func (m *Monk) ExecutePreCalculateMethods(c *models.Character) {
	m.executeDiamondSoul(c)
}

func (m *Monk) CalculateHitDice() string {
	return fmt.Sprintf("%dd8", m.Level)
}

func (m *Monk) executeUnarmoredDefense(c *models.Character) {
	monkExpertiseAbilityModifiers := []string{
		shared.AbilityDexterity,
		shared.AbilityWisdom,
	}

	executeUnarmoredDefenseShared(c, monkExpertiseAbilityModifiers)
}

func (m *Monk) executeUnarmoredMovement(c *models.Character) {
	if c.WornEquipment.Armor.Name != "" || m.Level < 2 {
		return
	}

	c.Speed += 10
}

func (m *Monk) executeMartialArts(c *models.Character) {
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

func (m *Monk) executeKiPoints(c *models.Character) {
	if c.Level < 2 || m.ClassToken.Name == "" {
		return
	} else if m.ClassToken.Name != kiPointsToken {
		logger.Info("Invalid Class Token Name")
		return
	}

	m.ClassToken.Maximum = c.Level
	m.ClassToken.Available = min(m.ClassToken.Available, m.ClassToken.Maximum)

	wisMod := c.GetMod(shared.AbilityWisdom)

	m.KiSpellSaveDC = 8 + c.Proficiency + wisMod
}

func (m *Monk) executeDeflectMissles(c *models.Character) {
	if c.Level < 3 {
		return
	}

	m.DeflectMissles = (10 + c.Proficiency + c.Level) * -1
}

func (m *Monk) executeDiamondSoul(c *models.Character) {
	if c.Level < 14 {
		return
	}

	for i := range c.Abilities {
		c.Abilities[i].SavingThrowsProficient = true
	}
}

func (m *Monk) ClassDetails() string {
	var s string

	martialArts := fmt.Sprintf("*Martial Arts*: %s\n\n", m.MartialArts)
	s += martialArts

	if m.ClassToken.Maximum != 0 && m.ClassToken.Name == kiPointsToken {
		s += fmt.Sprintf("*Ki Points*: %d/%d\n\n", m.ClassToken.Available, m.ClassToken.Maximum)
		s += fmt.Sprintf("*Ki Spell Save DC*: %d\n\n", m.KiSpellSaveDC)
	}

	if m.DeflectMissles > 0 {
		deflectMissles := fmt.Sprintf("*Deflect Missles Damage Reduction*: %d", m.DeflectMissles)
		s += deflectMissles
	}

	return s
}

// CLI

func (m *Monk) UseClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != kiPointsToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, m.ClassType))
		return
	}

	if m.ClassToken.Available <= 0 {
		logger.Info("No Ki Points Available")
		return
	}

	m.ClassToken.Available -= quantity
}

func (m *Monk) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName != "" && tokenName != kiPointsToken {
		logger.Info(fmt.Sprintf("Invalid token name '%s' for class '%s'", tokenName, m.ClassType))
		return
	}

	m.ClassToken.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || m.ClassToken.Available > m.ClassToken.Maximum {
		m.ClassToken.Available = m.ClassToken.Maximum
	}
}

func (m *Monk) GetTokens() []string {
	return []string{
		kiPointsToken,
	}
}
