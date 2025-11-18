package menu

import (
	tea "github.com/charmbracelet/bubbletea"

	"github.com/onioncall/dndgo/tui/shared"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.createPage, _ = m.createPage.Update(msg)
		m.managePage, _ = m.managePage.Update(msg)
		m.searchPage, _ = m.searchPage.Update(msg)
		return m, nil

	case shared.NavigateMsg:
		m.currentPage = msg.Page
		return m, nil
	}

	if m.currentPage != "menu" {
		var cmd tea.Cmd
		switch m.currentPage {
		case shared.InitPage:
			m.createPage, cmd = m.createPage.Update(msg)
		case shared.ManagePage:
			m.managePage, cmd = m.managePage.Update(msg)
		case shared.SearchPage:
			m.searchPage, cmd = m.searchPage.Update(msg)
		}
		return m, cmd
	}

	// Handle menu page input
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "q", "ctrl+c":
			return m, tea.Quit

		case "left", "shift+tab":
			m.selectedBtn--
			if m.selectedBtn < 0 {
				m.selectedBtn = len(m.buttons) - 1
			}
			return m, nil

		case "right", "tab":
			m.selectedBtn++
			if m.selectedBtn >= len(m.buttons) {
				m.selectedBtn = 0
			}
			return m, nil

		case "enter":
			switch m.selectedBtn {
			case 0:
				m.currentPage = shared.InitPage
			case 1:
				m.currentPage = shared.ManagePage
			case 2:
				m.currentPage = shared.SearchPage
			}
			return m, nil
		}
	}

	return m, nil
}
