package manage

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
		m.spellSaveDCViewport, cmd = m.spellSaveDCViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.spellSlotsViewport, cmd = m.spellSlotsViewport.Update(msg)
		cmds = append(cmds, cmd)

		m.knownSpellsViewport, cmd = m.knownSpellsViewport.Update(msg)
		cmds = append(cmds, cmd)
	}

	return m, tea.Batch(cmds...)
}
