package create

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateBasicInfoPage(msg tea.Msg) (Model, tea.Cmd) {
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
				m.err = m.saveBasicInfo()
				if m.err != nil {
					return m, nil
				}

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = classPage
				m.inputs = classInputs()

				return m, nil
			} else if m.backButtonFocused {
				return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
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

func (m *Model) saveBasicInfo() error {
	nameValue := m.inputs[nameInput].Value()
	for _, existing := range m.existingNames {
		if strings.EqualFold(existing, nameValue) {
			return fmt.Errorf("A character with this name already exists on this machine, please choose another")
		}
	}

	hpValue := m.inputs[hpInput].Value()
	if hpValue == "" {
		hpValue = "0"
	}
	hp, err := strconv.Atoi(hpValue)
	if err != nil {
		return fmt.Errorf("Invalid hp, must be an integer")
	}

	speedValue := m.inputs[speedInput].Value()
	if speedValue == "" {
		speedValue = "0"
	}
	speed, err := strconv.Atoi(speedValue)
	if err != nil {
		return fmt.Errorf("Invalid speed, must be an integer")
	}

	m.character.Name = m.inputs[nameInput].Value()
	m.character.Race = m.inputs[raceInput].Value()
	m.character.Background = m.inputs[backgroundInput].Value()
	m.character.Languages = strings.Split(m.inputs[languagesInput].Value(), ", ")
	m.character.HPCurrent = hp
	m.character.HPMax = hp
	m.character.Speed = speed

	return nil
}

func (m *Model) populateBasicInfoInputs() {
	if m.character.Name == "" {
		// This is going to be a hack mostly for development. If the character is not named,
		// we're not going to worry about populating the basic info inputs
		return
	}

	hpStr := strconv.Itoa(m.character.HPMax)
	speedStr := strconv.Itoa(m.character.Speed)

	m.inputs[nameInput].SetValue(m.character.Name)
	m.inputs[raceInput].SetValue(m.character.Race)
	m.inputs[backgroundInput].SetValue(m.character.Background)
	m.inputs[languagesInput].SetValue(strings.Join(m.character.Languages, ", "))
	m.inputs[hpInput].SetValue(hpStr)
	m.inputs[speedInput].SetValue(speedStr)
}
