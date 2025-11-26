package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Paladin struct {
	OtherFeatures        []models.ClassFeature `json:"other-features"`
	PreparedSpells       []string              `json:"prepared-spells"`
	OathSpells           []string              `json:"oath-spells"`
	ClassTokens          []shared.NamedToken   `json:"class-tokens"`
	FightingStyle        string                `json:"fighting-style"`
	FightingStyleFeature FightingStyleFeature  `json:"-"`
	SacredOath           string                `json:"sacred-oath"`
}

func LoadPaladin(data []byte) (*Paladin, error) {
	var paladin Paladin
	if err := json.Unmarshal(data, &paladin); err != nil {
		return &paladin, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &paladin, nil
}

func (p *Paladin) ExecutePostCalculateMethods(c *models.Character) {
	p.executeSpellCastingAbility(c)
	p.executePreparedSpells(c)
	p.executeClassTokens(c)
	p.executeFightingStyle(c)
	p.executeOathSpells(c)
}

func (p *Paladin) CalculateHitDice(level int) string {
	return fmt.Sprintf("%dd10", level)
}

func (s *Paladin) executeSpellCastingAbility(c *models.Character) {
	chrMod := c.GetMod(shared.AbilityCharisma)

	executeSpellSaveDC(c, chrMod)
	executeSpellAttackMod(c, chrMod)
}

func (p *Paladin) executeClassTokens(c *models.Character) {
	for i, token := range p.ClassTokens {
		if token.Name == "divine-sense" {
			p.ClassTokens[i].Maximum = 1 + c.GetMod(shared.AbilityCharisma)
		} else if token.Name == "lay-on-hands" {
			p.ClassTokens[i].Maximum = 5 * c.Level
		}
	}
}

// At level 2, Paladins adopt a fighting style as their specialty
// only one of these styles can be selected
func (p *Paladin) executeFightingStyle(c *models.Character) {
	if c.Level < 2 {
		return
	}

	invalidMsg := fmt.Sprintf("%s not one of the valid fighting styles, %s, %s, %s, %s",
		p.FightingStyle,
		shared.FightingStyleGreatWeaponFighting,
		shared.FightingStyleDefense,
		shared.FightingStyleDueling,
		shared.FightingStyleProtection)

	switch p.FightingStyle {
	case shared.FightingStyleGreatWeaponFighting:
		p.FightingStyleFeature = applyGreatWeaponFighting(c)
	case shared.FightingStyleDefense:
		p.FightingStyleFeature = applyDefense(c)
	case shared.FightingStyleDueling:
		p.FightingStyleFeature = applyDueling(c)
	case shared.FightingStyleProtection:
		p.FightingStyleFeature = applyProtection(c)
	default:
		logger.Info(invalidMsg)
	}
}

func (p *Paladin) executePreparedSpells(c *models.Character) {
	chrMod := c.GetMod(shared.AbilityCharisma)
	preparedSpellsMax := chrMod + (c.Level / 2)

	if !c.ValidationDisabled {
		if len(p.PreparedSpells) > preparedSpellsMax {
			logger.Info(fmt.Sprintf("%d exceeds the maximum amount of prepared spells (%d)",
				len(p.PreparedSpells), preparedSpellsMax))
		} else if len(p.PreparedSpells) < preparedSpellsMax {
			diff := preparedSpellsMax - len(p.PreparedSpells)
			logger.Info(fmt.Sprintf("You have %d prepared spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, p.PreparedSpells)
}

func (p *Paladin) executeOathSpells(c *models.Character) {
	oathSpellsMax := 0
	switch {
	case c.Level < 3:
		oathSpellsMax = 0
	case c.Level < 5:
		oathSpellsMax = 2
	case c.Level < 9:
		oathSpellsMax = 4
	case c.Level < 13:
		oathSpellsMax = 6
	case c.Level < 17:
		oathSpellsMax = 8
	case c.Level >= 17:
		oathSpellsMax = 10
	}

	if !c.ValidationDisabled {
		if len(p.OathSpells) > oathSpellsMax {
			logger.Info(fmt.Sprintf("%d exceeds the maximum amount of oath spells (%d)",
				len(p.OathSpells), oathSpellsMax))
		} else if len(p.OathSpells) < oathSpellsMax {
			diff := oathSpellsMax - len(p.OathSpells)
			logger.Info(fmt.Sprintf("You have %d oath spells not being used", diff))
		}
	}

	executePreparedSpellsShared(c, p.OathSpells)
}

func (p *Paladin) PrintClassDetails(c *models.Character) []string {
	s := buildClassDetailsHeader()

	for _, token := range p.ClassTokens {
		if token.Maximum == 0 || c.Level < token.Level {
			continue
		}

		switch token.Name {
		case "divine-sense":
			tokenSlots := c.GetSlots(token.Available, token.Maximum)
			s = append(s, fmt.Sprintf("*%s*: %s\n\n", "Divine Sense", tokenSlots))
		case "lay-on-hands":
			s = append(s, fmt.Sprintf("*Lay On Hands*: %d/%d\n\n", token.Available, token.Maximum))
		default:
			logger.Info(fmt.Sprintf("Invalid token name: %s", token.Name))
			continue
		}
	}

	if p.FightingStyleFeature.Name != "" && c.Level >= 2 {
		appliedText := "Requirements for fighting style not met."
		if p.FightingStyleFeature.IsApplied {
			appliedText = "Requirements for this fighting style are met, and any bonuses to armor or weapons have been applied to your character."
		}

		fightingStyleHeader := fmt.Sprintf("**Fighting Style**: *%s*\n", p.FightingStyleFeature.Name)
		fightingStyleDetail := fmt.Sprintf("%s\n%s\n\n", p.FightingStyleFeature.Details, appliedText)
		s = append(s, fightingStyleHeader)
		s = append(s, fightingStyleDetail)
	}

	if len(p.OtherFeatures) > 0 {
		for _, detail := range p.OtherFeatures {
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

// CLI

func (p *Paladin) UseClassTokens(tokenName string, quantity int) {
	token := getToken(tokenName, p.ClassTokens)

	if token == nil {
		logger.Info(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	if token.Available <= 0 {
		logger.Info(fmt.Sprintf("%s had no uses left", tokenName))
		return
	}

	token.Available -= quantity
}

func (p *Paladin) RecoverClassTokens(tokenName string, quantity int) {
	if tokenName == "all" {
		fullTokenRecovery(p.ClassTokens)
		return
	}

	token := getToken(tokenName, p.ClassTokens)

	if token == nil {
		logger.Info(fmt.Sprintf("Invalid token name: %s", tokenName))
		return
	}

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || token.Available > token.Maximum {
		token.Available = token.Maximum
	}
}

func (p *Paladin) GetTokens() []string {
	s := []string{}

	for _, token := range p.ClassTokens {
		s = append(s, token.Name)
	}

	return s
}
