package manage

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type BasicInfoModel struct {
	basicStatsViewport viewport.Model
	abilitiesViewport  viewport.Model
	healthViewport     viewport.Model
	skillsViewport     viewport.Model
}

func NewBasicInfoModel(character *models.Character) BasicInfoModel {
	statsVp := viewport.New(0, 0)
	statsContent := fmt.Sprintf(`Class: %s
Level: %d
Race: %s
Proficiency: +%d
Speed:  %d
Passive Perception: %d
Passive Insight: %d
AC: %d
Hit Dice: %s
	`, character.ClassName, character.Level, character.Race, character.Proficiency,
		character.Speed, character.PassivePerception, character.PassiveInsight,
		character.AC, character.HitDice)
	statsVp.SetContent(statsContent)

	// Content will be set in RefreshBasicInfo
	abilitiesVp := viewport.New(0, 0)

	healthStr := fmt.Sprintf("Current HP: %d | Max HP: %d | Temp HP: %d", character.HPCurrent, character.HPMax, character.HPTemp)
	healthVp := viewport.New(0, 0)
	healthVp.SetContent(healthStr)

	// Content will be set in InitializeContent
	skillsVp := viewport.New(0, 0)

	return BasicInfoModel{
		basicStatsViewport: statsVp,
		abilitiesViewport:  abilitiesVp,
		healthViewport:     healthVp,
		skillsViewport:     skillsVp,
	}
}

func (m *BasicInfoModel) SetAbilitiesContent(character *models.Character) {
	lineWidth := m.abilitiesViewport.Width - (abilitiesPadding * 2)
	abilitiesHeader := "Ability        -  Mod -  ST Mod"
	abilitiesStr := fmt.Sprintf("%s\n", abilitiesHeader)
	abilitiesStr += strings.Repeat("─", lineWidth) + "\n"
	for _, a := range character.Abilities {
		modStr := fmt.Sprintf("%d", a.AbilityModifier)
		if a.AbilityModifier >= 0 {
			modStr = fmt.Sprintf("+%d", a.AbilityModifier)
		}
		st := a.AbilityModifier
		if a.SavingThrowsProficient {
			st += character.Proficiency
		}
		stStr := fmt.Sprintf("%d", st)
		if st >= 0 {
			stStr = fmt.Sprintf("+%d", st)
		}

		abilityNameStr := fmt.Sprintf("%s%s", a.Name, strings.Repeat(" ", 13-utf8.RuneCountInString(a.Name)))
		abilityStr := fmt.Sprintf("%s  -  %s  -  %s", abilityNameStr, modStr, stStr)

		// Doing this so each line is the same width for centering purposes.
		// This is a space, but we are using the unicode so that lipgloss does not strip it out as a trailing space
		abilitiesStr += fmt.Sprintf("%s%s\n",
			abilityStr,
			strings.Repeat("\u00A0", utf8.RuneCountInString(abilitiesHeader)-utf8.RuneCountInString(abilityStr)))
	}

	m.abilitiesViewport.SetContent(abilitiesStr)
}

func (m *BasicInfoModel) SetSkillsContent(character *models.Character) {
	lineWidth := m.skillsViewport.Width - (skillsPadding * 2)
	skillsHeader := "Ability       -  Skills          -  Mod -  Proficient"
	skillsStr := fmt.Sprintf("%s\n", skillsHeader)
	skillsStr += strings.Repeat("─", lineWidth) + "\n"

	for _, s := range character.Skills {
		modStr := fmt.Sprintf("%d", s.SkillModifier)
		if s.SkillModifier >= 0 {
			modStr = fmt.Sprintf("+%d", s.SkillModifier)
		}
		profStr := " "
		if s.Proficient {
			profStr = " -  Proficient"
		}

		abilityStr := fmt.Sprintf("%s%s", s.Ability, strings.Repeat(" ", 13-utf8.RuneCountInString(s.Ability)))
		skillNameStr := fmt.Sprintf("%s%s", s.Name, strings.Repeat(" ", 15-utf8.RuneCountInString(s.Name)))
		skillStr := fmt.Sprintf("%s -  %s -  %s %s", abilityStr, skillNameStr, modStr, profStr)

		// Doing this so each line is the same width for centering purposes.
		// This is a space, but we are using the unicode so that lipgloss does not strip it out as a trailing space
		skillsStr += fmt.Sprintf("%s%s\n",
			skillStr,
			strings.Repeat("\u00A0", utf8.RuneCountInString(skillsHeader)-utf8.RuneCountInString(skillStr)))
	}

	m.skillsViewport.SetContent(skillsStr)
}

func (m *BasicInfoModel) SetHealthContent(character *models.Character) {
	healthStr := fmt.Sprintf("Current HP: %d | Max HP: %d | Temp HP: %d", character.HPCurrent, character.HPMax, character.HPTemp)
	m.healthViewport.SetContent(healthStr)
}

func (m *BasicInfoModel) UpdateSize(innerWidth, availableHeight int, character *models.Character) {
	// Column 1: 1/3 width, split vertically 50/50
	col1Width := innerWidth / 3
	boxHeight := availableHeight / 2

	basicStatsInnerWidth := col1Width - 2
	basicStatsInnerHeight := boxHeight - 2
	abilitiesInnerWidth := col1Width - 2
	abilitiesInnerHeight := boxHeight - 2

	m.basicStatsViewport.Width = basicStatsInnerWidth
	m.basicStatsViewport.Height = basicStatsInnerHeight
	m.abilitiesViewport.Width = abilitiesInnerWidth
	m.abilitiesViewport.Height = abilitiesInnerHeight

	// Column 2: 2/3 width, split 15/85
	col2Width := innerWidth * 2 / 3
	healthHeight := (availableHeight * 15) / 100
	skillsHeight := availableHeight - healthHeight

	healthInnerWidth := col2Width - 2
	healthInnerHeight := healthHeight - 2
	skillsInnerWidth := col2Width - 2
	skillsInnerHeight := skillsHeight - 2

	m.healthViewport.Width = healthInnerWidth
	m.healthViewport.Height = healthInnerHeight
	m.skillsViewport.Width = skillsInnerWidth
	m.skillsViewport.Height = skillsInnerHeight
}

func (m *BasicInfoModel) InitializeContent(character *models.Character) {
	m.SetHealthContent(character)
	m.SetAbilitiesContent(character)
	m.SetSkillsContent(character)
}
