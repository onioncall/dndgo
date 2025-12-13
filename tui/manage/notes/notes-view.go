package notes

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	darkGray  = lipgloss.Color("#767676")
)

func (m NotesModel) View(innerWidth, availableHeight int) string {
	col1Width := (innerWidth * 1) / 3

	// Column 1 Viewports
	titleVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Padding(0, 4).
		Width(col1Width - 2).
		Height(availableHeight - 2).
		Align(lipgloss.Center)

	titleVp := titleVpStyle.Render(m.TitleViewPort.View())

	col2Width := (innerWidth * 2) / 3

	// Column 2 Viewports
	notesVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Padding(0, 2).
		Width(col2Width - 2).
		Height(availableHeight - 2).
		Align(lipgloss.Center)

	notesVp := notesVpStyle.Render(m.NoteViewPort.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, titleVp, notesVp)
}
