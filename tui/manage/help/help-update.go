package help

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m HelpModel) Update(msg tea.Msg) (HelpModel, tea.Cmd) {
	var cmd tea.Cmd
	m.HelpViewport, cmd = m.HelpViewport.Update(msg)
	return m, cmd
}
