package menu

import (
	create "github.com/onioncall/dndgo/tui/create"
	"github.com/onioncall/dndgo/tui/manage"
	"github.com/onioncall/dndgo/tui/search"
)

type Model struct {
	width       int
	height      int
	selectedBtn int
	buttons     []string
	pageText    string
	currentPage string
	createPage  create.Model
	managePage  manage.Model
	searchPage  search.Model
}

func New() Model {
	return Model{
		selectedBtn: 1,
		buttons:     []string{"Create", "Manage", "Search"},
		pageText:    "dndgo",
		currentPage: "menu",
		createPage:  create.NewModel(),
		managePage:  manage.NewModel(),
		searchPage:  search.NewModel(),
	}
}
