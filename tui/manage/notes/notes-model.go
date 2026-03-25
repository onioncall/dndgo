package notes

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/manage/notes/content"
	"github.com/onioncall/dndgo/tui/manage/notes/titles"
)

type NotesModel struct {
	ActivePane  int
	TitlePane   titles.NoteTitlesModel
	ContentPane content.NoteContentModel
}

const (
	titlesPane int = iota
	contentPane
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

	return NotesModel{
		ActivePane:  0,
		TitlePane:   titles.NewNoteTitlesModel(character.Notes),
		ContentPane: content.NewNoteContentModel(selectedNote),
	}
}

func (m NotesModel) UpdateSize(innerWidth, availableHeight int, character models.Character) NotesModel {
	col1Width := (innerWidth * 1) / 3
	col2Width := (innerWidth * 2) / 3

	titleInnerWidth := col1Width - 2
	titleInnerHeight := availableHeight - 2
	noteInnerWidth := col2Width - 2
	noteInnerHeight := availableHeight - 2

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
