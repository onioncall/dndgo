package notes

import (
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/manage/notes/content"
	"github.com/onioncall/dndgo/tui/manage/notes/titles"
)

type paneEnum int

type NotesModel struct {
	ActivePane  paneEnum
	TitlePane   titles.NoteTitlesModel
	ContentPane content.NoteContentModel
	Initialized bool
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

func NewNotesModel() NotesModel {
	return NotesModel{}
}
func (m NotesModel) Init(character *models.Character) NotesModel {
	var selectedNote *models.Note
	if len(character.Notes) != 0 {
		selectedNote = &character.Notes[0]
	}

	m.TitlePane = titles.NewNoteTitlesModel(character.Notes)
	m.ContentPane = content.NewNoteContentModel(selectedNote)
	m.ActivePane = 0
	m.SetFocus(m.ActivePane)

	m.Initialized = true

	return m
}

func (m NotesModel) UpdateSize(innerWidth, availableHeight int) NotesModel {
	if !m.Initialized {
		return m
	}
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
	// Safety check: ensure the list is properly initialized with items
	if m.TitlePane.TitlesList.Items() == nil || len(m.TitlePane.TitlesList.Items()) == 0 {
		return ""
	}

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
