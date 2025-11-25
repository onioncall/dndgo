package create

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	tui "github.com/onioncall/dndgo/tui/shared"
)

var dumpLength int

func (m Model) DumpView() string {
	if m.height == 0 {
		return ""
	}

	// Build all content
	var allLines []string

	allLines = append(allLines, "Basic Info")
	allLines = append(allLines, fmt.Sprintf("Name: %v", m.character.Name))
	allLines = append(allLines, fmt.Sprintf("Level: %v", m.character.Level))
	allLines = append(allLines, fmt.Sprintf("Class: %v", m.character.ClassName))
	allLines = append(allLines, fmt.Sprintf("Race: %v", m.character.Race))
	allLines = append(allLines, fmt.Sprintf("Background: %v", m.character.Background))
	allLines = append(allLines, fmt.Sprintf("Languages: %v", strings.Join(m.character.Languages, ", ")))
	allLines = append(allLines, fmt.Sprintf("HP Current: %v", m.character.HPCurrent))
	allLines = append(allLines, fmt.Sprintf("HP Max: %v", m.character.HPMax))
	allLines = append(allLines, fmt.Sprintf("Speed: %v", m.character.Speed))
	allLines = append(allLines, "")

	allLines = append(allLines, "Abilities")
	for _, a := range m.character.Abilities {
		allLines = append(allLines, fmt.Sprintf("%s: %d | %v", a.Name, a.Base, a.SavingThrowsProficient))
	}
	allLines = append(allLines, "")

	allLines = append(allLines, "Skills")
	for _, s := range m.character.Skills {
		allLines = append(allLines, fmt.Sprintf("%s: %v", s.Name, s.Proficient))
	}
	allLines = append(allLines, "")

	allLines = append(allLines, "Spells")
	for _, s := range m.character.Spells {
		allLines = append(allLines, fmt.Sprintf("%s: %v | %d", s.Name, s.IsRitual, s.SlotLevel))
	}
	allLines = append(allLines, "")

	allLines = append(allLines, "Spell Slots")
	for _, s := range m.character.SpellSlots {
		allLines = append(allLines, fmt.Sprintf("%d: %d | %d", s.Level, s.Available, s.Maximum))
	}
	allLines = append(allLines, "")

	allLines = append(allLines, "Weapons")
	for _, w := range m.character.Weapons {
		allLines = append(allLines, fmt.Sprintf("%s: %s | %v | %v", w.Name, w.Damage, w.Proficient, w.Range))
	}
	allLines = append(allLines, "")

	allLines = append(allLines, "Worn Equipment")
	allLines = append(allLines, fmt.Sprintf("Head: %v", m.character.WornEquipment.Head))
	allLines = append(allLines, fmt.Sprintf("Amulet: %v", m.character.WornEquipment.Amulet))
	allLines = append(allLines, fmt.Sprintf("Cloak: %v", m.character.WornEquipment.Cloak))
	allLines = append(allLines, fmt.Sprintf("Hands/Arms: %v", m.character.WornEquipment.HandsArms))
	allLines = append(allLines, fmt.Sprintf("Ring: %v", m.character.WornEquipment.Ring))
	allLines = append(allLines, fmt.Sprintf("Ring2: %v", m.character.WornEquipment.Ring2))
	allLines = append(allLines, fmt.Sprintf("Belt: %v", m.character.WornEquipment.Belt))
	allLines = append(allLines, fmt.Sprintf("Boots: %v", m.character.WornEquipment.Boots))
	allLines = append(allLines, fmt.Sprintf("Shield: %v", m.character.WornEquipment.Shield))
	allLines = append(allLines, fmt.Sprintf("Armor: %v", m.character.WornEquipment.Armor.Name))
	allLines = append(allLines, "")

	allLines = append(allLines, "Backpack Items")
	for _, b := range m.character.Backpack {
		allLines = append(allLines, fmt.Sprintf("%s: %d", b.Name, b.Quantity))
	}
	allLines = append(allLines, "")

	dumpLength = len(allLines)
	// Calculate visible lines
	availableLines := m.height - 2 // Leave room for instructions
	startIdx := m.viewportOffset
	endIdx := min(startIdx+availableLines, len(allLines))

	// Build visible content
	var formContent string
	for i := startIdx; i < endIdx; i++ {
		formContent += allLines[i] + "\n"
	}

	formContent += "\n(Use arrow keys to scroll, ESC to exit)"

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		formContent,
	)
}

func (m Model) DumpUpdate(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
		case "up", "k":
			if m.viewportOffset > 0 {
				m.viewportOffset--
			}
		case "down", "j":
			// Calculate total lines for boundary check
			totalLines := dumpLength
			availableLines := m.height - 2
			if m.viewportOffset < totalLines-availableLines {
				m.viewportOffset++
			}
		}
	}
	return m, nil
}
