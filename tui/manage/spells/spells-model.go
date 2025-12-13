package spells

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/shared"
)

type SpellsModel struct {
	SpellSaveDCViewport viewport.Model
	SpellSlotsViewport  viewport.Model
	KnownSpellsViewport viewport.Model
	contentSet          bool
}

func NewSpellsModel() SpellsModel {
	spellSaveDCVp := viewport.New(0, 0)
	spellSlotsVp := viewport.New(0, 0)
	knownSpellsVp := viewport.New(0, 0)

	return SpellsModel{
		SpellSaveDCViewport: spellSaveDCVp,
		SpellSlotsViewport:  spellSlotsVp,
		KnownSpellsViewport: knownSpellsVp,
	}
}

func GetKnownSpellContent(character models.Character, width int) string {
	width = width - (widthPadding * 2) //padding on both sides
	longestSpellNameWidth := 0
	maxSpellNameWidth := width - 29 // based on width of viewport and characters in header

	spellNames := make(map[string]string)
	for _, s := range character.Spells {
		spellNameLen := utf8.RuneCountInString(s.Name)
		spellNames[s.Name] = shared.TruncateString(s.Name, maxSpellNameWidth)

		// if the spell name is truncated, we want to use the max length instead
		newSpellLen := min(spellNameLen, maxSpellNameWidth)
		longestSpellNameWidth = max(newSpellLen, longestSpellNameWidth)
	}

	knownSpellsHeader := fmt.Sprintf("Level  - Name%s - Ritual - Prepared",
		strings.Repeat(" ", longestSpellNameWidth-4))

	knownSpellsContent := fmt.Sprintf("%s\n", knownSpellsHeader)
	knownSpellsContent += fmt.Sprintf("%s\n", strings.Repeat("─", width))

	for _, s := range character.Spells {
		ritualStr := strings.Repeat("\u00A0", 6)
		if s.IsRitual {
			ritualStr = "Ritual"
		}
		preparedStr := strings.Repeat("\u00A0", 8)
		if s.IsRitual {
			preparedStr = "Prepared"
		}

		nameLen := utf8.RuneCountInString(spellNames[s.Name])
		knownSpellStr := fmt.Sprintf("lvl: %d - %s%s - %s - %s",
			s.SlotLevel, spellNames[s.Name], strings.Repeat(" ", longestSpellNameWidth-nameLen), ritualStr, preparedStr)
		knownSpellsContent += fmt.Sprintf("%s\n\n", knownSpellStr)
	}

	return knownSpellsContent
}

func GetSpellSlotContent(character models.Character, width int) string {
	width = width - (widthPadding * 2) //padding on both sides
	slotHeader := "Spell Slots"
	slotContent := fmt.Sprintf("%s\n", slotHeader)
	slotContent += fmt.Sprintf("%s\n", strings.Repeat("─", width))

	maxLineWidth := utf8.RuneCountInString(slotHeader)

	var slotLines []string
	for _, s := range character.SpellSlots {
		slots := models.GetSlots(s.Available, s.Maximum)
		level := strconv.FormatInt(int64(s.Level), 10)
		slotLine := fmt.Sprintf("lvl: %s - %s", level, slots)
		lineLength := utf8.RuneCountInString(slotLine)
		maxLineWidth = max(lineLength, maxLineWidth)
		slotLines = append(slotLines, slotLine)
	}

	for _, line := range slotLines {
		length := utf8.RuneCountInString(line)
		slotContent += fmt.Sprintf("%s%s\n\n", line, strings.Repeat("\u00A0", maxLineWidth-length))
	}

	return slotContent
}

func (m SpellsModel) UpdateSize(innerWidth, availableHeight int, character models.Character) SpellsModel {
	// Column 1: 1/2 width, split vertically 15/85
	col1Width := innerWidth / 2
	spellSaveDCHeight := (availableHeight * 15) / 100
	spellSlotsHeight := availableHeight - spellSaveDCHeight

	spellSaveDCInnerWidth := col1Width - 2
	spellSaveDCInnerHeight := spellSaveDCHeight - 2
	spellSlotsInnerWidth := col1Width - 2
	spellSlotsInnerHeight := spellSlotsHeight - 2

	m.SpellSaveDCViewport.Height = spellSaveDCInnerHeight
	m.SpellSaveDCViewport.Width = spellSaveDCInnerWidth
	m.SpellSlotsViewport.Height = spellSlotsInnerHeight
	m.SpellSlotsViewport.Width = spellSlotsInnerWidth

	col2Width := innerWidth / 2

	knownSpellsInnerHeight := availableHeight - 2
	knownSpellsInnerWidth := col2Width - 2

	m.KnownSpellsViewport.Height = knownSpellsInnerHeight
	m.KnownSpellsViewport.Width = knownSpellsInnerWidth

	if !m.contentSet {
		dcStr := fmt.Sprintf("Spell Save DC: %d", character.SpellSaveDC)
		m.SpellSaveDCViewport.SetContent(dcStr)

		spellSlotsContent := GetSpellSlotContent(character, m.SpellSaveDCViewport.Width)
		m.SpellSlotsViewport.SetContent(spellSlotsContent)

		knownSpellsContent := GetKnownSpellContent(character, m.KnownSpellsViewport.Width)
		m.KnownSpellsViewport.SetContent(knownSpellsContent)
	}

	return m
}
