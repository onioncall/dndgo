package create

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateWeaponsPage(msg tea.Msg) (Model, tea.Cmd) {
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
				m.err = m.addWeapon()
				if m.err != nil {
					return m, nil
				}

				for i := range m.inputs {
					m.inputs[i].SetValue("")
				}

				m.focused = 0
				m.inputs[weaponNameInput].Focus()
				m.addButtonFocused = false
				m.viewportOffset = 0

				return m, nil
			} else if m.nextButtonFocused {
				if len(m.character.Weapons) > 0 {
					m.character.PrimaryEquipped = m.character.Weapons[0].Name
				}
				if len(m.character.Weapons) > 1 {
					m.character.PrimaryEquipped = m.character.Weapons[1].Name
				}

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = wornEquipmentPage
				m.inputs = equipmentInputs()
				m.populateEquipmentInputs()

				return m, nil
			} else if m.backButtonFocused {
				m.err = nil
				m.focused = 0
				m.backButtonFocused = false
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = spellsPage
				m.inputs = spellInputs()

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

func (m *Model) addWeapon() error {
	weaponName := m.inputs[weaponNameInput].Value()
	for _, weapon := range m.character.Weapons {
		if weaponName == weapon.Name {
			return fmt.Errorf("Weapon already exists in weapon list")
		}
	}

	proficientValue := m.inputs[proficientWeaponInput].Value()
	if proficientValue == "" {
		proficientValue = "false"
	}
	proficient, err := strconv.ParseBool(proficientValue)
	if err != nil {
		return fmt.Errorf("Weapon proficient value must be boolean")
	}

	rangedValue := m.inputs[rangedInput].Value()
	if rangedValue == "" {
		rangedValue = "false"
	}
	ranged, err := strconv.ParseBool(rangedValue)
	if err != nil {
		return fmt.Errorf("Weapon proficient value must be boolean")
	}

	normalRange, err := strconv.Atoi(m.inputs[normalRangeInput].Value())
	if err != nil {
		return fmt.Errorf("Invalid normal range, must be an integer")
	}

	longRange, err := strconv.Atoi(m.inputs[longRangeInput].Value())
	if err != nil {
		return fmt.Errorf("Invalid long range, must be an integer")
	}

	weapon := shared.Weapon{
		Name:       weaponName,
		Proficient: proficient,
		Ranged:     ranged,
		Range: shared.WeaponRange{
			NormalRange: normalRange,
			LongRange:   longRange,
		},
		Type:       m.inputs[typeInput].Value(),
		Properties: strings.Split(m.inputs[propertiesInput].Value(), ", "),
	}

	m.character.Weapons = append(m.character.Weapons, weapon)

	return nil
}
