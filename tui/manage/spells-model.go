package manage

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type SpellsModel struct {
	spellSaveDCViewport viewport.Model
	spellSlotsViewport  viewport.Model
	knownSpellsViewport viewport.Model
}

func NewSpellsModel(character *models.Character) SpellsModel {
	spellSaveDCVp := viewport.New(0, 0)
	dcStr := fmt.Sprintf("Spell Save DC: %d", character.SpellSaveDC)
	spellSaveDCVp.SetContent(dcStr)

	// These work differently from basic info, since basic info has known sizes
	// and spells are different from character to character
	spellSlotsVp := viewport.New(0, 0)
	slotContent := getSpellSlotContent(character)
	spellSlotsVp.SetContent(slotContent)

	knownSpellsVp := viewport.New(0, 0)
	spellContent := getKnownSpellContent(character)
	knownSpellsVp.SetContent(spellContent)

	return SpellsModel{
		spellSaveDCViewport: spellSaveDCVp,
		spellSlotsViewport:  spellSlotsVp,
		knownSpellsViewport: knownSpellsVp,
	}
}

func getKnownSpellContent(character *models.Character) string {
	maxSpellNameWidth := 0

	for _, s := range character.Spells {
		maxSpellNameWidth = max(utf8.RuneCountInString(s.Name), maxSpellNameWidth)
	}

	knownSpellsHeader := fmt.Sprintf("Level  - Name%s - Ritual - Prepared",
		strings.Repeat(" ", maxSpellNameWidth-4))
	knownSpellsContent := fmt.Sprintf("%s\n", knownSpellsHeader)
	knownSpellsContent += strings.Repeat("─", utf8.RuneCountInString(knownSpellsHeader)) + "\n"

	for _, s := range character.Spells {
		ritualStr := strings.Repeat("\u00A0", 6)
		if s.IsRitual {
			ritualStr = "Ritual"
		}
		preparedStr := strings.Repeat("\u00A0", 8)
		if s.IsRitual {
			preparedStr = "Prepared"
		}

		nameLen := utf8.RuneCountInString(s.Name)
		knownSpellStr := fmt.Sprintf("lvl: %d - %s%s - %s - %s",
			s.SlotLevel, s.Name, strings.Repeat(" ", maxSpellNameWidth-nameLen), ritualStr, preparedStr)
		knownSpellsContent += fmt.Sprintf("%s\n\n", knownSpellStr)
	}

	return knownSpellsContent
}

func getSpellSlotContent(character *models.Character) string {
	slotHeader := "Spell Slots"
	lineWidth := utf8.RuneCountInString(slotHeader)
	slotLineContent := ""

	// We're setting this a little backwards since the header isn't the longest field, and we don't know how
	// long the longest line is without going through it. But we're doing this to get the length of the line under
	// the header
	for _, s := range character.SpellSlots {
		slots := character.GetSlots(s.Available, s.Maximum)
		level := strconv.FormatInt(int64(s.Level), 10)
		slotLine := fmt.Sprintf("lvl: %s - %s", level, slots)
		slotLineWidth := utf8.RuneCountInString(slotLine)
		lineWidth = max(slotLineWidth, lineWidth)
		slotLineContent += fmt.Sprintf("%s%s\n\n", slotLine, strings.Repeat("\u00A0", lineWidth-slotLineWidth))
	}

	slotContent := fmt.Sprintf("%s\n\n", slotHeader)
	slotContent += slotLineContent

	return slotContent
}

func (m SpellsModel) UpdateSize(innerWidth, availableHeight int, character *models.Character) SpellsModel {
	// Column 1: 1/2 width, split vertically 15/85
	col1Width := innerWidth / 2
	spellSaveDCHeight := (availableHeight * 15) / 100
	spellSlotsHeight := availableHeight - spellSaveDCHeight

	spellSaveDCInnerWidth := col1Width - 2
	spellSaveDCInnerHeight := spellSaveDCHeight - 2
	spellSlotsInnerWidth := col1Width - 2
	spellSlotsInnerHeight := spellSlotsHeight - 2

	m.spellSaveDCViewport.Height = spellSaveDCInnerHeight
	m.spellSaveDCViewport.Width = spellSaveDCInnerWidth
	m.spellSlotsViewport.Height = spellSlotsInnerHeight
	m.spellSlotsViewport.Width = spellSlotsInnerWidth

	col2Width := innerWidth / 2

	knownSpellsInnerHeight := availableHeight - 2
	knownSpellsInnerWidth := col2Width - 2

	m.knownSpellsViewport.Height = knownSpellsInnerHeight
	m.knownSpellsViewport.Width = knownSpellsInnerWidth

	return m
}
