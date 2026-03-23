package titles

import tea "github.com/charmbracelet/bubbletea"

type NoteSelectedMsg struct{}

func (m NoteTitlesModel) Update(msg tea.Msg) (NoteTitlesModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds = append(cmds, func() tea.Msg { return NoteSelectedMsg{} })
		}
	}

	m.TitlesList, cmd = m.TitlesList.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
