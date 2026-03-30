package titles

import "github.com/charmbracelet/lipgloss"

func (m NoteTitlesModel) View() string {
	view := m.TitlesList.View()
	if m.focused {
		view += "\n" + lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render("↑/k up • ↓/j down")
	}
	return view
}
