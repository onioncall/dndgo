package content

import (
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type NoteContentModel struct {
	ContentViewPort viewport.Model
	ContentTextArea textarea.Model
	IsEditing       bool
	focused         bool
	text            string
	width           int
	height          int
}

const hintHeight int = 1

func NewNoteContentModel(note *models.Note) NoteContentModel {
	vp := viewport.New(0, 0)
	ta := textarea.New()

	if note != nil {
		vp.SetContent(note.Content)
		ta.SetValue(note.Content)
	}

	return NoteContentModel{
		ContentViewPort: viewport.New(0, 0),
		ContentTextArea: textarea.New(),
		IsEditing:       false,
		text:            note.Content,
	}
}

func (m NoteContentModel) UpdateSize(width, height int) NoteContentModel {
	m.width = width
	m.height = height
	m.ContentViewPort.Width = width
	m.ContentTextArea.SetWidth(width)
	if m.focused {
		m.ContentViewPort.Height = height - hintHeight - 1
		m.ContentTextArea.SetHeight(height - hintHeight)
	} else {
		m.ContentViewPort.Height = height
		m.ContentTextArea.SetHeight(height)
	}

	return m
}

func (m NoteContentModel) SetContent(content string) NoteContentModel {
	m.ContentViewPort.SetContent(content)
	m.ContentTextArea.SetValue(content)
	m.text = content
	return m
}

func (m NoteContentModel) GetContent() string {
	return m.text
}

func (m *NoteContentModel) SetFocused(focused bool) {
	m.focused = focused
}
