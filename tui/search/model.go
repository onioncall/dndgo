package search

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
)

type Model struct {
	width            int
	height           int
	err              error
	pageText         string
	searchInput      textinput.Model
	selectedTabIndex int
	searchVisible    bool
	tabs             []string
	viewport         viewport.Model
}

const (
	spellTab = iota
	monsterTab
	equipmentTab
	featureTab
)

// Tab Commands
const (
	spellCmd     = "/s"
	monsterCmd   = "/m"
	equipmentCmd = "/e"
	featureCmd   = "/f"
)

func NewModel() Model {
	input := textinput.New()
	input.Focus()
	input.Placeholder = "Search..."
	input.Width = 38

	return Model{
		selectedTabIndex: 0,
		searchVisible:    true,
		searchInput:      input,
		viewport:         viewport.New(0, 0),
		tabs:             []string{"Spell", "Monster", "Equipment", "Feature"},
	}
}
