package spells

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	widthPadding int = 2
)

const (
	orange   = lipgloss.Color("#FFA500")
	darkGray = lipgloss.Color("#767676")
)

func (m SpellsModel) View(innerWidth, availableHeight int) string {
	// Column 1: 50% width, split 15/85 vertically
	col1Width := innerWidth / 2
	spellSaveDCHeight := (availableHeight * 15) / 100 // Back to 15%
	spellSlotsHeight := availableHeight - spellSaveDCHeight

	// Spell Save DC viewport (top of column 1)
	spellSaveDCStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(0, widthPadding).
		Width(col1Width - 2).
		Height(spellSaveDCHeight - 2).
		Align(lipgloss.Center)

	spellSaveDCVp := spellSaveDCStyle.Render(m.SpellSaveDCViewport.View())

	// Spell Slots viewport (bottom of column 1)
	spellSlotsStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(0, widthPadding).
		Width(col1Width - 2).
		Height(spellSlotsHeight - 2).
		Align(lipgloss.Center)

	spellSlotsVp := spellSlotsStyle.Render(m.SpellSlotsViewport.View())

	// Stack column 1 vertically
	column1 := lipgloss.JoinVertical(lipgloss.Left, spellSaveDCVp, spellSlotsVp)

	// Column 2: 50% width, full height
	col2Width := innerWidth / 2

	// Known Spells viewport
	knownSpellsStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(0, widthPadding).
		Width(col2Width - 2).
		Height(availableHeight - 2).
		Align(lipgloss.Center)

	knownSpellsVp := knownSpellsStyle.Render(m.KnownSpellsViewport.View())

	column2 := knownSpellsVp

	// Join columns horizontally
	return lipgloss.JoinHorizontal(lipgloss.Top, column1, column2)
}
