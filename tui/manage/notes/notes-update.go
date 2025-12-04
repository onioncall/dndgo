package notes

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m NotesModel) Init() tea.Cmd {
	return nil
}

func (m NotesModel) Update(msg tea.Msg) (NotesModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd

		m.TitleViewPort, cmd = m.TitleViewPort.Update(msg)
		cmds = append(cmds, cmd)

		m.NoteViewPort, cmd = m.NoteViewPort.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
