package content

import "github.com/charmbracelet/lipgloss"

func (m NoteContentModel) View() string {
	hintStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	if m.IsEditing {
		hint := hintStyle.Render("tab to confirm • esc to cancel")
		view := m.ContentTextArea.View()
		if m.focused {
			view += "\n" + hint
		}
		return view
	} else {
		hint := hintStyle.Render("↑/k up • ↓/j down")
		view := m.ContentViewPort.View()
		if m.focused && m.ContentViewPort.TotalLineCount() > m.ContentViewPort.VisibleLineCount() {
			view += "\n" + hint
		}
		return view
	}
}
