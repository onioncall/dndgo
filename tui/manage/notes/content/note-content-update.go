package content

import tea "github.com/charmbracelet/bubbletea"

type NoteUpdatedMsg struct{}

func (m NoteContentModel) Update(msg tea.Msg) (NoteContentModel, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "shift+enter":
			if m.IsEditing {
				m.IsEditing = false
				cmds = append(cmds, func() tea.Msg { return NoteUpdatedMsg{} })
			}
		}
	}

	m.ContentTextArea, cmd = m.ContentTextArea.Update(msg)
	cmds = append(cmds, cmd)
	m.ContentViewPort, cmd = m.ContentViewPort.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
