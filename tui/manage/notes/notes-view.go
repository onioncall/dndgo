package notes

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	brightBlue = lipgloss.Color("#7DF9FF")
	lightBlue  = lipgloss.Color("#5DC9E2")
	cream      = lipgloss.Color("#F9F6F0")
)

func (m NotesModel) View(innerWidth, availableHeight int) string {
	col1Width := (innerWidth * 1) / 3

	// Left side
	titlePaneStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, (padding * 2)).
		Width(col1Width - padding).
		Height(availableHeight - padding)

	// Right side
	col2Width := (innerWidth * 2) / 3
	contentPaneStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, padding).
		Width(col2Width - padding).
		Height(availableHeight - padding)

	switch m.ActivePane {
	case titlesPane:
		titlePaneStyle = showPaneAsFocused(titlePaneStyle)
	case contentPane:
		contentPaneStyle = showPaneAsFocused(contentPaneStyle)
	}

	titlePane := titlePaneStyle.Render(m.TitlePane.View())
	contentPane := contentPaneStyle.Render(m.ContentPane.View())

	return lipgloss.JoinHorizontal(lipgloss.Top, titlePane, contentPane)
}

func showPaneAsFocused(focusedVpStyle lipgloss.Style) lipgloss.Style {
	return focusedVpStyle.
		Border(lipgloss.ThickBorder()).
		BorderForeground(brightBlue)
}
