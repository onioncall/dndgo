package notes

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type NotesModel struct {
	TitleViewPort viewport.Model
	NoteViewPort  viewport.Model
}

func NewNotesModel() NotesModel {
	titleViewPort := viewport.New(0, 0)
	titleViewPort.SetContent("Note titles are")

	noteViewPort := viewport.New(0, 0)
	noteViewPort.SetContent("Note contents are under construction")

	return NotesModel{
		TitleViewPort: titleViewPort,
		NoteViewPort:  noteViewPort,
	}
}

func (m NotesModel) UpdateSize(innerWidth, availableHeight int, character models.Character) NotesModel {
	col1Width := (innerWidth * 1) / 3
	col2Width := (innerWidth * 2) / 3

	titleInnerWidth := col1Width - 2
	titleInnerHeight := availableHeight - 2
	noteInnerWidth := col2Width - 2
	noteInnerHeight := availableHeight - 2

	m.TitleViewPort.Width = titleInnerWidth
	m.TitleViewPort.Height = titleInnerHeight
	m.NoteViewPort.Width = noteInnerWidth
	m.NoteViewPort.Height = noteInnerHeight

	return m
}
