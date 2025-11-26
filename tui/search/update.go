package search

import (
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
		return m, nil

	case tea.KeyMsg:
		result := ""
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			m.content = ""
			return m, func() tea.Msg { return shared.NavigateMsg{Page: shared.MenuPage} }
		case "tab":
			m.selectedTabIndex = (m.selectedTabIndex + 1) % len(m.tabs)
		case "shift+tab":
			m.selectedTabIndex = (m.selectedTabIndex - 1 + len(m.tabs)) % len(m.tabs)
		case "ctrl+s":
			m.searchVisible = !m.searchVisible
		case "enter":
			searchInput := m.searchInput.Value()
			cmd, searchInputAfterCmd, _ := strings.Cut(searchInput, " ")
			switch strings.ToLower(cmd) {
			case spellCmd:
				m.selectedTabIndex = spellTab
				searchInput = searchInputAfterCmd
			case monsterCmd:
				m.selectedTabIndex = monsterTab
				searchInput = searchInputAfterCmd
			case equipmentCmd:
				m.selectedTabIndex = equipmentTab
				searchInput = searchInputAfterCmd
			case featureCmd:
				m.selectedTabIndex = featureTab
				searchInput = searchInputAfterCmd
			}

			switch m.selectedTabIndex {
			case spellTab:
				result, m.err = handlers.HandleSpellRequest(searchInput, m.width-10) // subtracting for border and padding
			case monsterTab:
			case equipmentTab:
			case featureTab:
			}

			m.content = result
			m.searchVisible = false

			m.searchInput.SetValue("")
		}
	}

	if m.searchVisible {
		var cmd tea.Cmd
		m.searchInput, cmd = m.searchInput.Update(msg)
		cmds = append(cmds, cmd)
		m.searchInput.Focus()
	} else {
		m.searchInput.Blur()
	}

	return m, tea.Batch(cmds...)
}
