package create

import (
	"fmt"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) UpdateSkillsPage(msg tea.Msg) (Model, tea.Cmd) {
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
				m.err = m.saveSkills()
				if m.err != nil {
					return m, nil
				}

				m.err = nil
				m.focused = 0
				m.nextButtonFocused = false
				m.viewportOffset = 0

				if hasSpellClass(m.character.ClassTypes) {
					m.currentPage = spellsPage
					m.inputs = spellInputs()
				} else {
					m.currentPage = weaponsPage
					m.inputs = weaponInputs()
				}

				return m, nil
			} else if m.backButtonFocused {
				m.err = nil
				m.focused = 0
				m.backButtonFocused = false
				m.nextButtonFocused = false
				m.viewportOffset = 0

				m.currentPage = abilitiesPage
				m.inputs = abilitiesInputs()
				m.populateSavedAbilitiesInputs()

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

		m.updateViewportGeneric(1, 4, 4)

	case errMsg:
		m.err = msg
		return m, nil
	}

	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return m, tea.Batch(cmds...)
}

func (m *Model) saveSkills() error {
	athleticsProfValue := m.inputs[athleticsInput].Value()
	if athleticsProfValue == "" {
		athleticsProfValue = "false"
	}
	athleticsProf, err := strconv.ParseBool(athleticsProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	athleticsSkill := shared.Skill{
		Name:       skillToAbility[athleticsInput].name,
		Ability:    skillToAbility[athleticsInput].ability,
		Proficient: athleticsProf,
	}

	acrobaticsProfValue := m.inputs[acrobaticsInput].Value()
	if acrobaticsProfValue == "" {
		acrobaticsProfValue = "false"
	}
	acrobaticsProf, err := strconv.ParseBool(acrobaticsProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	acrobaticsSkill := shared.Skill{
		Name:       skillToAbility[acrobaticsInput].name,
		Ability:    skillToAbility[acrobaticsInput].ability,
		Proficient: acrobaticsProf,
	}

	sleightOfHandProfValue := m.inputs[sleightOfHandInput].Value()
	if sleightOfHandProfValue == "" {
		sleightOfHandProfValue = "false"
	}
	sleightOfHandProf, err := strconv.ParseBool(sleightOfHandProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	sleightOfHandSkill := shared.Skill{
		Name:       skillToAbility[sleightOfHandInput].name,
		Ability:    skillToAbility[sleightOfHandInput].ability,
		Proficient: sleightOfHandProf,
	}

	stealthProfValue := m.inputs[stealthInput].Value()
	if stealthProfValue == "" {
		stealthProfValue = "false"
	}
	stealthProf, err := strconv.ParseBool(stealthProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	stealthSkill := shared.Skill{
		Name:       skillToAbility[stealthInput].name,
		Ability:    skillToAbility[stealthInput].ability,
		Proficient: stealthProf,
	}

	arcanaProfValue := m.inputs[arcanaInput].Value()
	if arcanaProfValue == "" {
		arcanaProfValue = "false"
	}
	arcanaProf, err := strconv.ParseBool(arcanaProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	arcanaSkill := shared.Skill{
		Name:       skillToAbility[arcanaInput].name,
		Ability:    skillToAbility[arcanaInput].ability,
		Proficient: arcanaProf,
	}

	historyProfValue := m.inputs[historyInput].Value()
	if historyProfValue == "" {
		historyProfValue = "false"
	}
	historyProf, err := strconv.ParseBool(historyProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	historySkill := shared.Skill{
		Name:       skillToAbility[historyInput].name,
		Ability:    skillToAbility[historyInput].ability,
		Proficient: historyProf,
	}

	investigationProfValue := m.inputs[investigationInput].Value()
	if investigationProfValue == "" {
		investigationProfValue = "false"
	}
	investigationProf, err := strconv.ParseBool(investigationProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	investigationSkill := shared.Skill{
		Name:       skillToAbility[investigationInput].name,
		Ability:    skillToAbility[investigationInput].ability,
		Proficient: investigationProf,
	}

	natureProfValue := m.inputs[natureInput].Value()
	if natureProfValue == "" {
		natureProfValue = "false"
	}
	natureProf, err := strconv.ParseBool(natureProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	natureSkill := shared.Skill{
		Name:       skillToAbility[natureInput].name,
		Ability:    skillToAbility[natureInput].ability,
		Proficient: natureProf,
	}

	religionProfValue := m.inputs[religionInput].Value()
	if religionProfValue == "" {
		religionProfValue = "false"
	}
	religionProf, err := strconv.ParseBool(religionProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	religionSkill := shared.Skill{
		Name:       skillToAbility[religionInput].name,
		Ability:    skillToAbility[religionInput].ability,
		Proficient: religionProf,
	}

	animalHandlingProfValue := m.inputs[animalHandlingInput].Value()
	if animalHandlingProfValue == "" {
		animalHandlingProfValue = "false"
	}
	animalHandlingProf, err := strconv.ParseBool(animalHandlingProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	animalHandlingSkill := shared.Skill{
		Name:       skillToAbility[animalHandlingInput].name,
		Ability:    skillToAbility[animalHandlingInput].ability,
		Proficient: animalHandlingProf,
	}

	insightProfValue := m.inputs[insightInput].Value()
	if insightProfValue == "" {
		insightProfValue = "false"
	}
	insightProf, err := strconv.ParseBool(insightProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	insightSkill := shared.Skill{
		Name:       skillToAbility[insightInput].name,
		Ability:    skillToAbility[insightInput].ability,
		Proficient: insightProf,
	}

	medicineProfValue := m.inputs[medicineInput].Value()
	if medicineProfValue == "" {
		medicineProfValue = "false"
	}
	medicineProf, err := strconv.ParseBool(medicineProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	medicineSkill := shared.Skill{
		Name:       skillToAbility[medicineInput].name,
		Ability:    skillToAbility[medicineInput].ability,
		Proficient: medicineProf,
	}

	perceptionProfValue := m.inputs[perceptionInput].Value()
	if perceptionProfValue == "" {
		perceptionProfValue = "false"
	}
	perceptionProf, err := strconv.ParseBool(perceptionProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	perceptionSkill := shared.Skill{
		Name:       skillToAbility[perceptionInput].name,
		Ability:    skillToAbility[perceptionInput].ability,
		Proficient: perceptionProf,
	}

	survivalProfValue := m.inputs[survivalInput].Value()
	if survivalProfValue == "" {
		survivalProfValue = "false"
	}
	survivalProf, err := strconv.ParseBool(survivalProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	survivalSkill := shared.Skill{
		Name:       skillToAbility[survivalInput].name,
		Ability:    skillToAbility[survivalInput].ability,
		Proficient: survivalProf,
	}

	deceptionProfValue := m.inputs[deceptionInput].Value()
	if deceptionProfValue == "" {
		deceptionProfValue = "false"
	}
	deceptionProf, err := strconv.ParseBool(deceptionProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	deceptionSkill := shared.Skill{
		Name:       skillToAbility[deceptionInput].name,
		Ability:    skillToAbility[deceptionInput].ability,
		Proficient: deceptionProf,
	}

	intimidationProfValue := m.inputs[intimidationInput].Value()
	if intimidationProfValue == "" {
		intimidationProfValue = "false"
	}
	intimidationProf, err := strconv.ParseBool(intimidationProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	intimidationSkill := shared.Skill{
		Name:       skillToAbility[intimidationInput].name,
		Ability:    skillToAbility[intimidationInput].ability,
		Proficient: intimidationProf,
	}

	performanceProfValue := m.inputs[performanceInput].Value()
	if performanceProfValue == "" {
		performanceProfValue = "false"
	}
	performanceProf, err := strconv.ParseBool(performanceProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	performanceSkill := shared.Skill{
		Name:       skillToAbility[performanceInput].name,
		Ability:    skillToAbility[performanceInput].ability,
		Proficient: performanceProf,
	}

	persuasionProfValue := m.inputs[persuasionInput].Value()
	if persuasionProfValue == "" {
		persuasionProfValue = "false"
	}
	persuasionProf, err := strconv.ParseBool(persuasionProfValue)
	if err != nil {
		return fmt.Errorf("Skill Proficient values must be booleans")
	}
	persuasionSkill := shared.Skill{
		Name:       skillToAbility[persuasionInput].name,
		Ability:    skillToAbility[persuasionInput].ability,
		Proficient: persuasionProf,
	}

	m.character.Skills = []shared.Skill{
		athleticsSkill,
		acrobaticsSkill,
		sleightOfHandSkill,
		stealthSkill,
		arcanaSkill,
		historySkill,
		investigationSkill,
		natureSkill,
		religionSkill,
		animalHandlingSkill,
		insightSkill,
		medicineSkill,
		perceptionSkill,
		survivalSkill,
		deceptionSkill,
		intimidationSkill,
		performanceSkill,
		persuasionSkill,
	}

	return nil
}

func (m *Model) populateSavedSkillInputs() {
	if len(m.character.Skills) <= 0 {
		return
	}

	currentSkillsMap := make(map[string]shared.Skill)
	for _, s := range m.character.Skills {
		currentSkillsMap[s.Name] = s
	}

	for i := range skillToAbility {
		skillName := skillToAbility[i].name
		if skill, ok := currentSkillsMap[skillName]; ok {
			m.inputs[i].SetValue(strconv.FormatBool(skill.Proficient))
		}
	}
}
