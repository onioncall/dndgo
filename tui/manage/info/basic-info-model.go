package info

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type BasicInfoModel struct {
	BasicStatsViewport viewport.Model
	AbilitiesViewport  viewport.Model
	HealthViewport     viewport.Model
	SkillsViewport     viewport.Model
	contentSet         bool
}

func NewBasicInfoModel(character *models.Character) BasicInfoModel {
	statsVp := viewport.New(0, 0)
	healthVp := viewport.New(0, 0)
	abilitiesVp := viewport.New(0, 0)
	skillsVp := viewport.New(0, 0)

	return BasicInfoModel{
		BasicStatsViewport: statsVp,
		AbilitiesViewport:  abilitiesVp,
		HealthViewport:     healthVp,
		SkillsViewport:     skillsVp,
	}
}

func GetHealthContent(character *models.Character) string {
	healthContent := fmt.Sprintf("Current HP: %d | Max HP: %d | Temp HP: %d",
		character.HPCurrent, character.HPMax, character.HPTemp)

	return healthContent
}

func GetStatsContent(character *models.Character) string {
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

	return statsContent
}

func GetAbilitiesContent(character *models.Character, width int) string {
	width = width - (skillsPadding * 2) //padding on both sides
	abilitiesHeader := "Ability        -  Mod -  ST Mod"
	lineWidth := utf8.RuneCountInString(abilitiesHeader)
	abilitiesStr := fmt.Sprintf("%s\n", abilitiesHeader)
	abilitiesStr += fmt.Sprintf("%s\n", strings.Repeat("─", width))

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
			strings.Repeat("\u00A0", lineWidth-utf8.RuneCountInString(abilityStr)))
	}

	return abilitiesStr
}

func GetSkillsContent(character *models.Character, width int) string {
	width = width - (skillsPadding * 2) //padding on both sides
	skillsHeader := "Ability       -  Skills          -  Mod -  Proficient"
	lineWidth := utf8.RuneCountInString(skillsHeader)
	skillsStr := fmt.Sprintf("%s\n", skillsHeader)
	skillsStr += fmt.Sprintf("%s\n", strings.Repeat("─", width))

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
			strings.Repeat("\u00A0", lineWidth-utf8.RuneCountInString(skillStr)))
	}

	return skillsStr
}

func (m BasicInfoModel) UpdateSize(innerWidth, availableHeight int, character *models.Character) BasicInfoModel {
	// Column 1: 1/3 width, split vertically 50/50
	col1Width := innerWidth / 3
	boxHeight := availableHeight / 2

	basicStatsInnerWidth := col1Width - 2
	basicStatsInnerHeight := boxHeight - 2
	abilitiesInnerWidth := col1Width - 2
	abilitiesInnerHeight := boxHeight - 2

	m.BasicStatsViewport.Width = basicStatsInnerWidth
	m.BasicStatsViewport.Height = basicStatsInnerHeight
	m.AbilitiesViewport.Width = abilitiesInnerWidth
	m.AbilitiesViewport.Height = abilitiesInnerHeight

	// Column 2: 2/3 width, split 15/85
	col2Width := innerWidth * 2 / 3
	healthHeight := (availableHeight * 15) / 100
	skillsHeight := availableHeight - healthHeight

	healthInnerWidth := col2Width - 2
	healthInnerHeight := healthHeight - 2
	skillsInnerWidth := col2Width - 2
	skillsInnerHeight := skillsHeight - 2

	m.HealthViewport.Width = healthInnerWidth
	m.HealthViewport.Height = healthInnerHeight
	m.SkillsViewport.Width = skillsInnerWidth
	m.SkillsViewport.Height = skillsInnerHeight

	if !m.contentSet {
		statsContent := GetStatsContent(character)
		m.BasicStatsViewport.SetContent(statsContent)

		healthContent := GetHealthContent(character)
		m.HealthViewport.SetContent(healthContent)

		abilitiesContent := GetAbilitiesContent(character, m.AbilitiesViewport.Width)
		m.AbilitiesViewport.SetContent(abilitiesContent)

		skillsContent := GetSkillsContent(character, m.SkillsViewport.Width)
		m.SkillsViewport.SetContent(skillsContent)

		m.contentSet = true
	}

	return m
}
