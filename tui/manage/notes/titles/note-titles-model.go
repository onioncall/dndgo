package titles

import (
	"github.com/charmbracelet/bubbles/list"
	"github.com/onioncall/dndgo/character-management/models"
)

type NoteTitlesModel struct {
	TitlesList list.Model
	width      int
	height     int
}

type NoteTitleItem struct {
	Title string
}

func (i NoteTitleItem) FilterValue() string { return i.Title }

func NewNoteTitlesModel(notes []models.Note) NoteTitlesModel {
	var m NoteTitlesModel
	m.SetTitlesList(notes)
	return m
}

func (m NoteTitlesModel) UpdateSize(width, height int) {
	m.width = width
	m.height = height
	m.TitlesList.SetWidth(width)
	m.TitlesList.SetHeight(height)
}

func (m NoteTitlesModel) SetTitlesList(notes []models.Note) {
	var items []list.Item
	for _, v := range notes {
		items = append(items, NoteTitleItem{Title: v.Title})
	}

	m.TitlesList = list.New(items, list.NewDefaultDelegate(), m.width, m.height)
}
