package search

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type Model struct {
	width            int
	height           int
	err              error
	pageText         string
	content          string
	searchInput      textinput.Model
	selectedTabIndex int
	searchVisible    bool
	tabs             []string
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
	input.Placeholder = "Search..."
	input.Width = 38

	return Model{
		selectedTabIndex: 0,
		searchVisible:    true,
		searchInput:      input,
		tabs:             []string{"Spell", "Monster", "Equipment", "Feature"},
	}
}
