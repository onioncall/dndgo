package info

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m BasicInfoModel) Init() tea.Cmd {
	return nil
}

func (m BasicInfoModel) Update(msg tea.Msg) (BasicInfoModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd

		m.BasicStatsViewport, cmd = m.BasicStatsViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.AbilitiesViewport, cmd = m.AbilitiesViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.HealthViewport, cmd = m.HealthViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.SkillsViewport, cmd = m.SkillsViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
