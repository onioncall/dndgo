package notes

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m NotesModel) Init() tea.Cmd {
	return nil
}

func (m NotesModel) Update(msg tea.Msg) (NotesModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "h", "left":
			m.ActivePaneIdx = (m.ActivePaneIdx - 1 + len(m.Panes)) % len(m.Panes)
		case "l", "right":
			m.ActivePaneIdx = (m.ActivePaneIdx + 1) % len(m.Panes)
		}
	}
	m.TitlePane, cmd = m.TitlePane.Update(msg)
	cmds = append(cmds, cmd)

	m.ContentPane, cmd = m.ContentPane.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
