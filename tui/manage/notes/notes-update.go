package notes

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/tui/manage/msgs"
)

func (m NotesModel) Init() tea.Cmd {
	return nil
}

func (m NotesModel) SetFocus(pane paneEnum) NotesModel {
	m.ActivePane = pane
	if pane == titlesPane {
		m.TitlePane.SetFocused(true)
		m.ContentPane.SetFocused(false)
	}
	if pane == contentPane {
		m.TitlePane.SetFocused(false)
		m.ContentPane.SetFocused(true)
	}
	return m.UpdateSize(m.width, m.height)
}

func (m NotesModel) Update(msg tea.Msg) (NotesModel, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd
	var noteJustAdded bool

	switch msg := msg.(type) {
	case msgs.AddNoteMsg, msgs.EditNoteMsg:
		noteJustAdded = true
		m = m.SetFocus(contentPane)

	case msgs.NoteUpdatedMsg:
		m = m.SetFocus(titlesPane)

	case tea.KeyMsg:
		if m.ContentPane.IsEditing {
			// editing mode is for typing keys
			// so like do not handle key inputs
			break
		}

		switch msg.String() {
		case "h", "left":
			if m.ActivePane == contentPane {
				m = m.SetFocus(titlesPane)
			}
		case "l", "right":
			if m.ActivePane == titlesPane {
				m = m.SetFocus(contentPane)
			}
		}
	}

	// TitlePane only updates when in focus
	// Or if we just added a new note and are immediately opening for editing
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
