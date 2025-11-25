package create

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateEquipmentPage(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd = make([]tea.Cmd, len(m.inputs))

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			if m.nextButtonFocused {
				m.err = m.saveEquipment()
				if m.err != nil {
					return m, nil
				}

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = backpackPage
				m.inputs = backpackInputs()

				return m, nil
			} else if m.backButtonFocused {
				m.err = nil
				m.focused = 0
				m.backButtonFocused = false
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = weaponsPage
				m.inputs = weaponInputs()

				return m, nil
			}

			m.nextInput(2)

		case "ctrl+c":
			return m, tea.Quit
		case "esc":
			return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
		case "shift+tab", "ctrl+k", "up":
			m.prevInput(2)
		case "tab", "ctrl+j", "down":
			m.nextInput(2)
		}

		m.nextButtonFocused = m.focused == len(m.inputs)
		m.backButtonFocused = m.focused == len(m.inputs)+1

		for i := range m.inputs {
			m.inputs[i].Blur()
		}

		if m.focused < len(m.inputs) {
			m.inputs[m.focused].Focus()
		}

		m.updateViewportGeneric(2, 4, 4)

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) saveEquipment() error {
	m.character.WornEquipment.Head = m.inputs[headInput].Value()
	m.character.WornEquipment.Amulet = m.inputs[amuletInput].Value()
	m.character.WornEquipment.Cloak = m.inputs[cloakInput].Value()
	m.character.WornEquipment.HandsArms = m.inputs[handsArmsInput].Value()
	m.character.WornEquipment.Ring = m.inputs[ringInput].Value()
	m.character.WornEquipment.Ring2 = m.inputs[ring2Input].Value()
	m.character.WornEquipment.Belt = m.inputs[beltInput].Value()
	m.character.WornEquipment.Boots = m.inputs[bootsInput].Value()
	m.character.WornEquipment.Shield = m.inputs[shieldInput].Value()

	armorValue := m.inputs[armorInput].Value()
	if armorValue == "" {
		return nil
	}

	armorProficientValue := m.inputs[armorProficientInput].Value()
	if armorProficientValue == "" {
		armorProficientValue = "false"
	}
	armorProficient, err := strconv.ParseBool(armorProficientValue)
	if err != nil {
		return fmt.Errorf("Armor Proficient values must be booleans")
	}

	armorClass, err := strconv.Atoi(m.inputs[armorClassInput].Value())
	if err != nil {
		return fmt.Errorf("Invalid armor class, must be an integer")
	}

	armorType := m.inputs[armorTypeInput].Value()

	if strings.ToLower(armorType) != shared.LightArmor &&
		armorType != shared.MediumArmor &&
		armorType != shared.HeavyArmor {
		return fmt.Errorf("Invalid armor type, must be light, medium, or heavy")
	}

	armor := shared.Armor{
		Name:       armorValue,
		Proficient: armorProficient,
		Class:      armorClass,
		Type:       armorType,
	}

	m.character.WornEquipment.Armor = armor

	return nil
}

func (m *Model) populateEquipmentInputs() {
	m.inputs[headInput].SetValue(m.character.WornEquipment.Head)
	m.inputs[amuletInput].SetValue(m.character.WornEquipment.Amulet)
	m.inputs[cloakInput].SetValue(m.character.WornEquipment.Cloak)
	m.inputs[handsArmsInput].SetValue(m.character.WornEquipment.HandsArms)
	m.inputs[ringInput].SetValue(m.character.WornEquipment.Ring)
	m.inputs[ring2Input].SetValue(m.character.WornEquipment.Ring2)
	m.inputs[beltInput].SetValue(m.character.WornEquipment.Belt)
	m.inputs[bootsInput].SetValue(m.character.WornEquipment.Boots)
	m.inputs[shieldInput].SetValue(m.character.WornEquipment.Shield)

	armor := m.character.WornEquipment.Armor

	if armor.Name == "" {
		return
	}

	m.inputs[armorInput].SetValue(armor.Name)
	m.inputs[armorProficientInput].SetValue(strconv.FormatBool(armor.Proficient))
	m.inputs[armorClassInput].SetValue(strconv.FormatInt(int64(armor.Class), 10))
	m.inputs[armorTypeInput].SetValue(armor.Type)
}
