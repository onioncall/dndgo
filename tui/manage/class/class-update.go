package class

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m ClassModel) Init() tea.Cmd {
	return nil
}

func (m ClassModel) Update(msg tea.Msg) (ClassModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd

		m.SubClassViewport, cmd = m.SubClassViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.DetailViewport, cmd = m.DetailViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.OtherFeaturesViewport, cmd = m.OtherFeaturesViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
