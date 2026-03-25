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
	text            string
}

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
	m.ContentViewPort.Width = width
	m.ContentViewPort.Height = height
	m.ContentTextArea.SetWidth(width)
	m.ContentTextArea.SetHeight(height)
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
