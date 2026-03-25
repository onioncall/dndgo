package titles

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/onioncall/dndgo/character-management/models"
)

type NoteTitlesModel struct {
	TitlesList list.Model
	width      int
	height     int
}

type NoteTitleItem struct {
	NoteTitle   string
	NotePreview string
	paneWidth   int
}

func (i NoteTitleItem) FilterValue() string { return i.NoteTitle }
func (i NoteTitleItem) Title() string       { return i.NoteTitle }
func (i NoteTitleItem) Description() string {
	if len(i.NotePreview) == 0 {
		return "(empty)"
	}
	previewLen := int(float64(i.paneWidth) * 0.75)
	previewStr := i.NotePreview
	if len(i.NotePreview) > previewLen {
		previewStr = i.NotePreview[0:previewLen]
	}

	return fmt.Sprintf("%v...", previewStr)
}

func NewNoteTitlesModel(notes []models.Note) NoteTitlesModel {
	var m NoteTitlesModel
	m = m.SetTitlesList(notes)
	return m
}

func (m NoteTitlesModel) UpdateSize(width, height int) NoteTitlesModel {
	m.width = width
	m.height = height
	m.TitlesList.SetWidth(width)
	m.TitlesList.SetHeight(height)

	newItems := []list.Item{}
	for _, v := range m.TitlesList.Items() {
		item, _ := v.(NoteTitleItem)
		newItems = append(newItems, NoteTitleItem{
			NoteTitle:   item.NoteTitle,
			NotePreview: item.NotePreview,
			paneWidth:   width,
		})
	}
	m.TitlesList.SetItems(newItems)

	return m
}

func (m NoteTitlesModel) SetTitlesList(notes []models.Note) NoteTitlesModel {
	var items []list.Item
	for _, v := range notes {
		items = append(items, NoteTitleItem{NoteTitle: v.Title})
	}

	if m.TitlesList.Items() == nil {
		m.TitlesList = list.New(items, list.NewDefaultDelegate(), m.width, m.height)
	} else {
		m.TitlesList.SetItems(items)
	}
	return m
}
