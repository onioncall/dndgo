package search

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/search/handlers"
	"github.com/onioncall/dndgo/tui/shared"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.updateViewportSize()
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.viewport.SetContent("")
			return m, func() tea.Msg { return shared.NavigateMsg{Page: shared.MenuPage} }
		case "tab":
			m.selectedTabIndex = (m.selectedTabIndex + 1) % len(m.tabs)
		case "shift+tab":
			m.selectedTabIndex = (m.selectedTabIndex - 1 + len(m.tabs)) % len(m.tabs)
		case "ctrl+s":
			m.searchVisible = !m.searchVisible
			m.updateViewportSize()
		case "enter":
			searchInput := m.searchInput.Value()
			m.selectedTabIndex, searchInput = searchUserCmd(searchInput)

			var result string
			var err error

			if searchInput == "" {
				result = "not found"
			} else {
				result, err = searchUserInput(searchInput, m.selectedTabIndex, m.viewport.Width)
			}

			m.searchInput.SetValue("")
			if result != "" {
				m.searchVisible = false
			}

			m.err = err
			m.viewport.SetContent(result)
			m.viewport.GotoTop()
			m.updateViewportSize()
		}
	}

	var cmd tea.Cmd
	if m.searchVisible {
		m.searchInput, cmd = m.searchInput.Update(msg)
		cmds = append(cmds, cmd)
		m.searchInput.Focus()
	} else {
		m.searchInput.Blur()
		m.viewport, cmd = m.viewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}

func searchUserCmd(searchInput string) (int, string) {
	cmd, searchInputAfterCmd, _ := strings.Cut(searchInput, " ")

	var tab int
	var newSearchInput string

	switch strings.ToLower(cmd) {
	case spellCmd:
		tab = spellTab
		newSearchInput = searchInputAfterCmd
	case monsterCmd:
		tab = monsterTab
		newSearchInput = searchInputAfterCmd
	case equipmentCmd:
		tab = equipmentTab
		newSearchInput = searchInputAfterCmd
	case featureCmd:
		tab = featureTab
		newSearchInput = searchInputAfterCmd
	}

	return tab, newSearchInput
}

func searchUserInput(input string, currentIndex int, width int) (string, error) {
	var result string
	var err error

	if input == "" {
		return result, nil
	}

	switch currentIndex {
	case spellTab:
		if input == "list" {
			result, err = handlers.HandleSpellListRequest()
		}
		result, err = handlers.HandleSpellRequest(input, width)
	case monsterTab:
		if input == "list" {
			result, err = handlers.HandleMonsterListRequest()
		}
		result, err = handlers.HandleMonsterRequest(input, width)
	case equipmentTab:
		if input == "list" {
			result, err = handlers.HandleEquipmentListRequest()
		}
		result, err = handlers.HandleEquipmentRequest(input, width)
	case featureTab:
		if input == "list" {
			result, err = handlers.HandleFeatureListRequest()
		}
		result, err = handlers.HandleFeatureRequest(input, width)
	}

	if result == "" {
		result = fmt.Sprintf("%s not found", input)
	}

	return result, err
}

func (m *Model) updateViewportSize() {
	outerBorderMargin := 2
	searchHeight := 0
	if m.searchVisible {
		searchHeight = 3
	}

	containerWidth := m.width - (outerBorderMargin * 2) - 2
	containerHeight := m.height - (outerBorderMargin * 2) - 2 - searchHeight

	innerWidth := containerWidth - 4       // padding (2 each side)
	innerHeight := containerHeight - 2 - 2 // border + padding

	tabHeight := 3

	m.viewport.Width = innerWidth
	m.viewport.Height = innerHeight - tabHeight
}
