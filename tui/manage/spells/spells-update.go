package spells

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m SpellsModel) Init() tea.Cmd {
	return nil
}

func (m SpellsModel) Update(msg tea.Msg) (SpellsModel, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		var cmd tea.Cmd
		m.SpellSaveDCViewport, cmd = m.SpellSaveDCViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.SpellSlotsViewport, cmd = m.SpellSlotsViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.KnownSpellsViewport, cmd = m.KnownSpellsViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
