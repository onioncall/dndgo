package manage

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/tui/shared"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			return m, func() tea.Msg { return shared.NavigateMsg{Page: shared.MenuPage} }
		}
	}

	return m, nil
}
