package titles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/onioncall/dndgo/character-management/models"
)

type NoteTitlesModel struct {
	TitlesList list.Model
}

type NoteTitleItem struct {
	Title string
}

func (i NoteTitleItem) FilterValue() string { return i.Title }

func NewNoteTitlesModel(notes []models.Note) NoteTitlesModel {
	var items []list.Item
	for _, v := range notes {
		items = append(items, NoteTitleItem{Title: v.Title})
	}

	return NoteTitlesModel{
		TitlesList: list.New(items, list.NewDefaultDelegate(), 0, 0),
	}
}

func (m NoteTitlesModel) UpdateSize(width, height int) {
	m.TitlesList.SetWidth(width)
	m.TitlesList.SetHeight(height)
}
