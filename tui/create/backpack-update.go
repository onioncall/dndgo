package create

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateBackpackPage(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.addButtonFocused {
				m.err = m.addBackpackItem()
				if m.err != nil {
					return m, nil
				}

				for i := range m.inputs {
					m.inputs[i].SetValue("")
				}

				m.focused = 0
				m.inputs[itemNameInput].Focus()
				m.addButtonFocused = false
				m.viewportOffset = 0

				return m, nil
			} else if m.nextButtonFocused {
				err := handlers.SaveCharacterJson(m.character)
				if err != nil {
					logger.HandleInfo("Failed to save character json")
				}
				class, err := handlers.LoadClassTemplate(strings.ToLower(m.character.ClassName))
				if err != nil {
					logger.HandleInfo("Failed to load class template")
				}
				err = handlers.SaveClassHandler(class)
				if err != nil {
					logger.HandleInfo("Failed to save class handler")
				}

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				for i := range m.inputs {
					m.inputs[i].SetValue("")
				}
				m.currentPage = dumpPage

				return m, nil
			} else if m.backButtonFocused {
				m.err = nil
				m.focused = 0
				m.backButtonFocused = false
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = wornEquipmentPage
				m.inputs = equipmentInputs()
				m.populateEquipmentInputs()

				return m, nil
			}

			m.nextInput(3)
		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
		case "shift+tab", "ctrl+k", "up":
			m.prevInput(3)
		case "tab", "ctrl+j", "down":
			m.nextInput(3)
		}

		m.addButtonFocused = m.focused == len(m.inputs)
		m.nextButtonFocused = m.focused == len(m.inputs)+1
		m.backButtonFocused = m.focused == len(m.inputs)+2

		for i := range m.inputs {
			m.inputs[i].Blur()
		}

		if m.focused < len(m.inputs) {
			m.inputs[m.focused].Focus()
		}

		m.updateViewportGeneric(2, 4, 5)
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) addBackpackItem() error {
	itemName := m.inputs[itemNameInput].Value()
	for _, item := range m.character.Backpack {
		if itemName == item.Name {
			return fmt.Errorf("Item already exists in backpack")
		}
	}

	itemQuantity, err := strconv.Atoi(m.inputs[itemQuantityInput].Value())
	if err != nil || itemQuantity < 1 {
		return fmt.Errorf("Invalid item quantity, must be a positive integer")
	}

	item := shared.BackpackItem{
		Name:     itemName,
		Quantity: itemQuantity,
	}

	m.character.Backpack = append(m.character.Backpack, item)

	return nil
}
