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
	version     string
	currentPage string
	createPage  create.Model
	managePage  manage.Model
	searchPage  search.Model
}

func New(version string) Model {
	return Model{
		selectedBtn: 1,
		buttons:     []string{"Create", "Manage", "Search"},
		pageText:    "dndgo",
		version:     version,
		currentPage: "menu",
		createPage:  create.NewModel(),
		managePage:  manage.NewModel(),
		searchPage:  search.NewModel(),
	}
}
