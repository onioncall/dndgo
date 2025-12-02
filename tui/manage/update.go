package manage

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/handlers"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		innerWidth, availableHeight := m.getInnerDimensions()

		// Only update the active tab size
		switch m.selectedTabIndex {
		case basicInfoTab:
			m.basicInfoTab.UpdateSize(innerWidth, availableHeight, m.character)
		}

		// Initiallize all content on load
		if !m.contentInitialized {
			m.basicInfoTab.InitializeContent(m.character)
		}

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			handlers.SaveCharacterJson(m.character)
			return m, tea.Quit
		case "esc":
			handlers.SaveCharacterJson(m.character)
			return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
		case "tab":
			m.selectedTabIndex = (m.selectedTabIndex + 1) % len(m.tabs)
			return m, nil
		case "shift+tab":
			m.selectedTabIndex = (m.selectedTabIndex - 1 + len(m.tabs)) % len(m.tabs)
			return m, nil
		case "ctrl+s":
			// Clearing error as a part of this process
			m.err = nil
			m.cmdVisible = !m.cmdVisible
			m.cmdInput.Focus()
			// Update tab sizes when cmd visibility changes
			innerWidth, availableHeight := m.getInnerDimensions()
			m.basicInfoTab.UpdateSize(innerWidth, availableHeight, m.character)

			return m, nil
		case "enter":
			if m.cmdVisible {
				searchInput := m.cmdInput.Value()
				m, m.selectedTabIndex, searchInput = m.executeUserCmd(searchInput, m.selectedTabIndex)
				m.cmdInput.SetValue("")
				m.cmdVisible = false
				handlers.HandleCharacter(m.character)
			}

			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.cmdVisible {
		m.cmdInput, cmd = m.cmdInput.Update(msg)
		cmds = append(cmds, cmd)
		m.cmdInput.Focus()
	} else {
		m.cmdInput.Blur()
		switch m.selectedTabIndex {
		case basicInfoTab:
			cmd = m.basicInfoTab.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) executeUserCmd(cmdInput string, currentTab int) (Model, int, string) {
	cmd, inputAfterCmd, _ := strings.Cut(cmdInput, " ")
	tab := currentTab
	newInput := cmdInput
	switch strings.ToLower(cmd) {
	case basicInfoCmd:
		tab = basicInfoTab
		newInput = inputAfterCmd
	case spellCmd:
		tab = spellTab
		newInput = inputAfterCmd
	case equipmentCmd:
		tab = equipmentTab
		newInput = inputAfterCmd
	case classCmd:
		tab = classTab
		newInput = inputAfterCmd
	case damageCmd:
		dmg, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.DamageCharacter(int(dmg))
		m.basicInfoTab.SetHealthContent(m.character)
	case recoverCmd:
		health, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.HealCharacter(int(health))
		m.basicInfoTab.SetHealthContent(m.character)
	case addTempCmd:
		temp, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.AddTempHp(int(temp))
		m.basicInfoTab.SetHealthContent(m.character)
	default:
		m.err = fmt.Errorf("%s command not found", cmd)
	}

	return m, tab, newInput
}
