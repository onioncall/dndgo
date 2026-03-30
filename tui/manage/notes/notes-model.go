package notes

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/manage/notes/content"
	"github.com/onioncall/dndgo/tui/manage/notes/titles"
)

type paneEnum int

type NotesModel struct {
	ActivePane  paneEnum
	TitlePane   titles.NoteTitlesModel
	ContentPane content.NoteContentModel
	width       int
	height      int
}

const (
	titlesPane paneEnum = iota
	contentPane
)

const (
	border  int = 2
	padding int = 2
)

func NewNotesModel(character *models.Character) NotesModel {
	titleViewPort := viewport.New(0, 0)
	titleViewPort.SetContent("Note titles are under construction")

	noteViewPort := viewport.New(0, 0)
	noteViewPort.SetContent("Note contents are under construction")

	var selectedNote *models.Note
	if len(character.Notes) != 0 {
		selectedNote = &character.Notes[0]
	}

	model := NotesModel{
		ActivePane:  0,
		TitlePane:   titles.NewNoteTitlesModel(character.Notes),
		ContentPane: content.NewNoteContentModel(selectedNote),
	}
	model.SetFocus(model.ActivePane)

	return model
}

func (m NotesModel) UpdateSize(innerWidth, availableHeight int) NotesModel {
	m.width = innerWidth
	m.height = availableHeight

	col1Width := (innerWidth * 1) / 3
	col2Width := (innerWidth * 2) / 3

	titleInnerWidth := col1Width - padding
	titleInnerHeight := availableHeight - padding
	noteInnerWidth := col2Width - padding - (2 * border)
	noteInnerHeight := availableHeight - padding

	m.TitlePane = m.TitlePane.UpdateSize(titleInnerWidth, titleInnerHeight)
	m.ContentPane = m.ContentPane.UpdateSize(noteInnerWidth, noteInnerHeight)

	return m
}

func (m NotesModel) GetSelectedNoteTitle() string {
	if note, ok := m.TitlePane.TitlesList.SelectedItem().(titles.NoteTitleItem); ok {
		return note.NoteTitle
	}
	return ""
}

func (m NotesModel) UpdateNoteContent(content string) NotesModel {
	m.ContentPane = m.ContentPane.SetContent(content)
	return m
}

func (m NotesModel) SetNotesList(notes []models.Note) NotesModel {
	m.TitlePane = m.TitlePane.SetTitlesList(notes)
	return m
}

func (m NotesModel) IsEditing() bool {
	return m.ContentPane.IsEditing
}
