package notes

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
)

func (m NotesModel) View(innerWidth, availableHeight int) string {
	col1Width := (innerWidth * 1) / 3

	// Column 1 Viewports
	titleVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 4).
		Width(col1Width - 2).
		Height(availableHeight - 2).
		Align(lipgloss.Center)

	col2Width := (innerWidth * 2) / 3

	// Column 2 Viewports
	notesVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 2).
		Width(col2Width - 2).
		Height(availableHeight - 2).
		Align(lipgloss.Center)

	switch m.ViewPorts[m.ActiveViewPortIdx] {
	case "title":
		titleVpStyle = showViewportAsFocused(titleVpStyle)
	case "notes":
		notesVpStyle = showViewportAsFocused(notesVpStyle)
	default:
		panic(fmt.Sprintf("no view port set, %v, %v", m.ActiveViewPortIdx, m.ViewPorts[m.ActiveViewPortIdx]))
	}

	titleVp := titleVpStyle.Render(m.TitleViewPort.View())
	notesVp := notesVpStyle.Render(m.NoteViewPort.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, titleVp, notesVp)
}

func showViewportAsFocused(focusedVpStyle lipgloss.Style) lipgloss.Style {
	return focusedVpStyle.
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("#7DF9FF"))
}
