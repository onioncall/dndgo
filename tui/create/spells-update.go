package create

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateSpellsPage(msg tea.Msg) (Model, tea.Cmd) {
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
				m.err = m.addSpell()
				if m.err != nil {
					return m, nil
				}

				for i := range m.inputs {
					m.inputs[i].SetValue("")
				}

				m.focused = 0
				m.inputs[spellNameInput].Focus()
				m.addButtonFocused = false
				m.viewportOffset = 0

				return m, nil
			} else if m.nextButtonFocused {
				m.configureSpellSlots()

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = weaponsPage
				m.inputs = weaponInputs()

				return m, nil
			} else if m.backButtonFocused {
				m.err = nil
				m.focused = 0
				m.backButtonFocused = false
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = skillsPage
				m.inputs = skillsInputs()
				m.populateSavedSkillInputs()

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

		m.updateViewportAbilities()
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) addSpell() error {
	spellName := m.inputs[spellNameInput].Value()
	for _, spell := range m.character.Spells {
		if spellName == spell.Name {
			return fmt.Errorf("Spell already exists in list of known spells")
		}
	}

	slotLevelValue := m.inputs[slotLevelInput].Value()
	if slotLevelValue == "" {
		slotLevelValue = "0"
	}
	slotLevel, err := strconv.Atoi(slotLevelValue)
	if err != nil {
		return fmt.Errorf("Invalid spell slot level, must be an integer")
	}

	isRitualValue := m.inputs[isRitualInput].Value()
	if isRitualValue == "" {
		isRitualValue = "false"
	}
	isRitual, err := strconv.ParseBool(isRitualValue)
	if err != nil {
		return fmt.Errorf("Is Ritual value must be boolean")
	}

	spell := shared.CharacterSpell{
		Name:      spellName,
		IsRitual:  isRitual,
		SlotLevel: slotLevel,
	}

	m.character.Spells = append(m.character.Spells, spell)

	return nil
}

func (m *Model) configureSpellSlots() {
	level := m.character.Level
	class := strings.ToLower(m.character.ClassName)

	// These classes don't support spells, no need to do any of this
	if class == shared.ClassBarbarian ||
		class == shared.ClassFighter ||
		class == shared.ClassMonk ||
		class == shared.ClassRogue {
		return
	}

	fullCasterProgression := map[int]map[int]int{
		1:  {1: 2},
		2:  {1: 3},
		3:  {1: 4, 2: 2},
		4:  {1: 4, 2: 3},
		5:  {1: 4, 2: 3, 3: 2},
		6:  {1: 4, 2: 3, 3: 3},
		7:  {1: 4, 2: 3, 3: 3, 4: 1},
		8:  {1: 4, 2: 3, 3: 3, 4: 2},
		9:  {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
		10: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
		11: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
		12: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1},
		13: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
		14: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1},
		15: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
		16: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1},
		17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2, 6: 1, 7: 1, 8: 1, 9: 1},
		18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 1, 7: 1, 8: 1, 9: 1},
		19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 1, 8: 1, 9: 1},
		20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 3, 6: 2, 7: 2, 8: 1, 9: 1},
	}

	halfCasterProgression := map[int]map[int]int{
		2:  {1: 2},
		3:  {1: 3},
		4:  {1: 3},
		5:  {1: 4, 2: 2},
		6:  {1: 4, 2: 2},
		7:  {1: 4, 2: 3},
		8:  {1: 4, 2: 3},
		9:  {1: 4, 2: 3, 3: 2},
		10: {1: 4, 2: 3, 3: 2},
		11: {1: 4, 2: 3, 3: 3},
		12: {1: 4, 2: 3, 3: 3},
		13: {1: 4, 2: 3, 3: 3, 4: 1},
		14: {1: 4, 2: 3, 3: 3, 4: 1},
		15: {1: 4, 2: 3, 3: 3, 4: 2},
		16: {1: 4, 2: 3, 3: 3, 4: 2},
		17: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
		18: {1: 4, 2: 3, 3: 3, 4: 3, 5: 1},
		19: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
		20: {1: 4, 2: 3, 3: 3, 4: 3, 5: 2},
	}

	warlockProgression := map[int]map[int]int{
		1:  {1: 1},
		2:  {1: 2},
		3:  {2: 2},
		4:  {2: 2},
		5:  {3: 2},
		6:  {3: 2},
		7:  {4: 2},
		8:  {4: 2},
		9:  {5: 2},
		10: {5: 2},
		11: {5: 3},
		12: {5: 3},
		13: {5: 3},
		14: {5: 3},
		15: {5: 3},
		16: {5: 3},
		17: {5: 4},
		18: {5: 4},
		19: {5: 4},
		20: {5: 4},
	}

	spellSlotTable := map[string]map[int]map[int]int{
		shared.ClassBard:     fullCasterProgression,
		shared.ClassCleric:   fullCasterProgression,
		shared.ClassDruid:    fullCasterProgression,
		shared.ClassSorcerer: fullCasterProgression,
		shared.ClassWizard:   fullCasterProgression,
		shared.ClassPaladin:  halfCasterProgression,
		shared.ClassRanger:   halfCasterProgression,
		shared.ClassWarlock:  warlockProgression,
	}

	slotMap := make(map[int]int)
	if classTable, ok := spellSlotTable[class]; ok {
		if levelSlots, ok := classTable[level]; ok {
			slotMap = levelSlots
		}
	}

	for i := 1; i <= 9; i++ {
		slot := shared.SpellSlot{
			Level:     i,
			Available: slotMap[i],
			Maximum:   slotMap[i],
		}

		m.character.SpellSlots = append(m.character.SpellSlots, slot)
	}
}
