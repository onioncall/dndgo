package notes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/tui/manage/msgs"
)

func (m NotesModel) Init() tea.Cmd {
	return nil
}

func (m NotesModel) Update(msg tea.Msg) (NotesModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	var noteJustAdded bool

	switch msg := msg.(type) {
	case msgs.AddNoteMsg, msgs.EditNoteMsg:
		noteJustAdded = true
		m.ActivePane = contentPane

	case tea.KeyMsg:
		if m.ContentPane.IsEditing {
			// editing mode is for typing keys
			// so like do not handle key inputs
			break
		}

		switch msg.String() {
		case "h", "left":
			if m.ActivePane == contentPane {
				m.ActivePane = titlesPane
				// No updates propegate downwards during focus shift
				return m, nil
			}
		case "l", "right":
			if m.ActivePane == titlesPane {
				m.ActivePane = contentPane
				// No updates propegate downwards during focus shift
				return m, nil
			}
		}
	}

	// TitlePane only updates when in focus
	if m.ActivePane == titlesPane || noteJustAdded {
		m.TitlePane, cmd = m.TitlePane.Update(msg)
		if cmd != nil {
			cmds = append(cmds, cmd)
		}
	}

	// ContentPane always updates while notes tab is in view
	m.ContentPane, cmd = m.ContentPane.Update(msg)
	if cmd != nil {
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
