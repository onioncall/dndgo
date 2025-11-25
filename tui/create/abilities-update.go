package create

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateAbilitiesPage(msg tea.Msg) (Model, tea.Cmd) {
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
				m.err = m.saveAbilities()
				if m.err != nil {
					return m, nil
				}

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = skillsPage
				m.inputs = skillsInputs()
				m.populateSavedSkillInputs()

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

		if !m.nextButtonFocused && !m.backButtonFocused {
			m.inputs[m.focused].Focus()
		}

		m.updateViewportAbilities()
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) saveAbilities() error {
	strBaseValue := m.inputs[strengthBaseInput].Value()
	if strBaseValue == "" {
		strBaseValue = "0"
	}
	strBase, err := strconv.Atoi(strBaseValue)
	if err != nil {
		return fmt.Errorf("Base values must be integers")
	}
	strStpValue := m.inputs[strengthProficientInput].Value()
	if strStpValue == "" {
		strStpValue = "false"
	}
	strStp, err := strconv.ParseBool(strStpValue)
	if err != nil {
		return fmt.Errorf("Savings Throws Proficient values must be booleans")
	}
	strength := shared.Ability{
		Name:                   shared.AbilityStrength,
		Base:                   strBase,
		SavingThrowsProficient: strStp,
	}

	dexBaseValue := m.inputs[dexterityBaseInput].Value()
	if dexBaseValue == "" {
		dexBaseValue = "0"
	}
	dexBase, err := strconv.Atoi(dexBaseValue)
	if err != nil {
		return fmt.Errorf("Base values must be integers")
	}
	dexStpValue := m.inputs[dexterityProficientInput].Value()
	if dexStpValue == "" {
		dexStpValue = "false"
	}
	dexStp, err := strconv.ParseBool(dexStpValue)
	if err != nil {
		return fmt.Errorf("Savings Throws Proficient values must be booleans")
	}
	dexterity := shared.Ability{
		Name:                   shared.AbilityDexterity,
		Base:                   dexBase,
		SavingThrowsProficient: dexStp,
	}

	conBaseValue := m.inputs[constitutionBaseInput].Value()
	if conBaseValue == "" {
		conBaseValue = "0"
	}
	conBase, err := strconv.Atoi(conBaseValue)
	if err != nil {
		return fmt.Errorf("Base values must be integers")
	}
	conStpValue := m.inputs[constitutionProficientInput].Value()
	if conStpValue == "" {
		conStpValue = "false"
	}
	conStp, err := strconv.ParseBool(conStpValue)
	if err != nil {
		return fmt.Errorf("Savings Throws Proficient values must be booleans")
	}
	constitution := shared.Ability{
		Name:                   shared.AbilityConstitution,
		Base:                   conBase,
		SavingThrowsProficient: conStp,
	}

	intBaseValue := m.inputs[intelligenceBaseInput].Value()
	if intBaseValue == "" {
		intBaseValue = "0"
	}
	intBase, err := strconv.Atoi(intBaseValue)
	if err != nil {
		return fmt.Errorf("Base values must be integers")
	}
	intStpValue := m.inputs[intelligenceProficientInput].Value()
	if intStpValue == "" {
		intStpValue = "false"
	}
	intStp, err := strconv.ParseBool(intStpValue)
	if err != nil {
		return fmt.Errorf("Savings Throws Proficient values must be booleans")
	}
	intelligence := shared.Ability{
		Name:                   shared.AbilityIntelligence,
		Base:                   intBase,
		SavingThrowsProficient: intStp,
	}

	wisBaseValue := m.inputs[wisdomBaseInput].Value()
	if wisBaseValue == "" {
		wisBaseValue = "0"
	}
	wisBase, err := strconv.Atoi(wisBaseValue)
	if err != nil {
		return fmt.Errorf("Base values must be integers")
	}
	wisStpValue := m.inputs[wisdomProficientInput].Value()
	if wisStpValue == "" {
		wisStpValue = "false"
	}
	wisStp, err := strconv.ParseBool(wisStpValue)
	if err != nil {
		return fmt.Errorf("Savings Throws Proficient values must be booleans")
	}
	wisdom := shared.Ability{
		Name:                   shared.AbilityWisdom,
		Base:                   wisBase,
		SavingThrowsProficient: wisStp,
	}

	chaBaseValue := m.inputs[charismaBaseInput].Value()
	if chaBaseValue == "" {
		chaBaseValue = "0"
	}
	chaBase, err := strconv.Atoi(chaBaseValue)
	if err != nil {
		return fmt.Errorf("Base values must be integers")
	}
	chaStpValue := m.inputs[charismaProficientInput].Value()
	if chaStpValue == "" {
		chaStpValue = "false"
	}
	chaStp, err := strconv.ParseBool(chaStpValue)
	if err != nil {
		return fmt.Errorf("Savings Throws Proficient values must be booleans")
	}
	charisma := shared.Ability{
		Name:                   shared.AbilityCharisma,
		Base:                   chaBase,
		SavingThrowsProficient: chaStp,
	}

	m.character.Abilities = []shared.Ability{
		strength,
		dexterity,
		constitution,
		intelligence,
		wisdom,
		charisma,
	}
	return nil
}

func (m *Model) populateSavedAbilitiesInputs() {
	if len(m.character.Abilities) <= 0 {
		return
	}

	// this represents the abilities that have already been saved to the character.
	currentAbilitiesMap := make(map[string]shared.Ability)
	for _, a := range m.character.Abilities {
		currentAbilitiesMap[a.Name] = a
	}

	for i := range m.inputs {
		// Because there are two inputs per ability, this is the "index" for the actual corresponding ability
		abilityIndex := i / 2
		abilityName := strings.ToLower(abilityInputsMap[abilityIndex])
		if ability, ok := currentAbilitiesMap[abilityName]; ok {
			if i%2 == 0 {
				m.inputs[i].SetValue(strconv.Itoa(ability.Base))
			} else {
				m.inputs[i].SetValue(strconv.FormatBool(ability.SavingThrowsProficient))
			}
		}
	}
}
