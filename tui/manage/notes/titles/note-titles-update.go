package titles

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/tui/manage/msgs"
)

func (m NoteTitlesModel) Update(msg tea.Msg) (NoteTitlesModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			cmds = append(cmds, func() tea.Msg { return msgs.EditNoteMsg{} })
		case "j", "k", "down", "up":
			cmds = append(cmds, func() tea.Msg { return msgs.NoteSelectedMsg{} })
		}
	case msgs.AddNoteMsg:
		m.TitlesList.Select(len(m.TitlesList.Items()) - 1)
	}

	m.TitlesList, cmd = m.TitlesList.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}
