package create

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateClassPage(msg tea.Msg) (Model, tea.Cmd) {
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
				m.err = m.addClass()
				if m.err != nil {
					return m, nil
				}

				for i := range m.inputs {
					m.inputs[i].SetValue("")
				}

				m.focused = 0
				m.inputs[classTypeInput].Focus()
				m.addButtonFocused = false
				m.viewportOffset = 0

				return m, nil
			} else if m.nextButtonFocused {
				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = abilitiesPage
				m.inputs = abilitiesInputs()
				m.populateSavedAbilitiesInputs()

				return m, nil
			} else if m.backButtonFocused {
				m.err = nil
				m.focused = 0
				m.backButtonFocused = false
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = basicInfoPage
				m.inputs = basicInfoInputs()
				m.populateBasicInfoInputs()

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

func (m *Model) addClass() error {
	levelValue := m.inputs[classLevelInput].Value()
	if levelValue == "" {
		levelValue = "0"
	}
	level, err := strconv.Atoi(levelValue)
	if err != nil {
		return fmt.Errorf("Invalid level, must be an integer")
	}

	classType := m.inputs[classTypeInput].Value()
	for _, existingClass := range m.character.ClassTypes {
		if strings.EqualFold(existingClass, classType) {
			return fmt.Errorf("Class has already been added for character")
		}
	}

	validClasses := []string{
		shared.ClassBarbarian,
		shared.ClassBard,
		shared.ClassCleric,
		shared.ClassDruid,
		shared.ClassFighter,
		shared.ClassMonk,
		shared.ClassPaladin,
		shared.ClassRanger,
		shared.ClassRogue,
		shared.ClassSorcerer,
		shared.ClassWarlock,
		shared.ClassWizard,
	}

	isValid := slices.Contains(validClasses, strings.ToLower(classType))

	if !isValid {
		return fmt.Errorf("Invalid class name")
	}

	m.character.ClassTypes = append(m.character.ClassTypes, classType)
	m.classMap[strings.ToLower(classType)] = level

	return nil
}
