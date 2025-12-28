package equipment

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
		m.WornEquipmentViewport, cmd = m.WornEquipmentViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.BackpackViewport, cmd = m.BackpackViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.WeaponsViewport, cmd = m.WeaponsViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
