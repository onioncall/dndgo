package manage

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m EquipmentModel) Init() tea.Cmd {
	return nil
}

func (m EquipmentModel) Update(msg tea.Msg) (EquipmentModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd
		m.wornEquipmentViewport, cmd = m.wornEquipmentViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.backpackViewport, cmd = m.backpackViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.weaponsViewport, cmd = m.weaponsViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
