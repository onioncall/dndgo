package manage

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m BasicInfoModel) Init() tea.Cmd {
	return nil
}

func (m BasicInfoModel) Update(msg tea.Msg) tea.Cmd {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd
		m.basicStatsViewport, cmd = m.basicStatsViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.abilitiesViewport, cmd = m.abilitiesViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.healthViewport, cmd = m.healthViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.skillsViewport, cmd = m.skillsViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return tea.Batch(cmds...)
}
