package titles

import "github.com/charmbracelet/lipgloss"

func (m NoteTitlesModel) View() string {
	hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	if m.NoNotes() {
		return hintStyle.Render("\n" + "No Notes yet!" + "\n\n" + "(use `add-note` command)")
	}

	// This was getting unset somehow
	m.TitlesList.SetShowHelp(false)

	view := m.TitlesList.View()
	if m.focused {
		view += "\n" + hintStyle.Render("↑/k up • ↓/j down")
	}
	return view
}
