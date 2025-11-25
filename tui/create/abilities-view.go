package create

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) AbilitiesPageView() string {
	if m.height == 0 {
		return ""
	}

	linesPerAbility := 3 // label + base + proficient
	headerLines := 4     // header + icons + padding
	footerLines := 4     // next button + back button + padding

	availableLines := m.height - headerLines - footerLines
	visibleFields := max(availableLines/linesPerAbility, 1)

	startAbilityIdx := m.viewportOffset / 2
	endAbilityIdx := min(startAbilityIdx+visibleFields, 6)

	labels := []string{
		"Strength		",
		"Dexterity		",
		"Constitution	",
		"Intelligence	",
		"Wisdom			",
		"Charisma		",
	}

	var formContent string

	// Top padding
	formContent += strings.Repeat("\n", 2)
	formContent += "Abilities:\n\n"

	// Looping over the start and end index instead of just the len(lables) for scroll functionality
	for i := startAbilityIdx; i < endAbilityIdx; i++ {
		baseIdx := i * 2
		proficientIdx := i*2 + 1

		formContent += fmt.Sprintf("%s\nBase:       %s\nProficient: %s\n",
			inputStyle.Width(41).Render(labels[i]),
			m.inputs[baseIdx].View(),
			m.inputs[proficientIdx].View(),
		)
	}

	nextText := "next"
	backText := "back"
	if m.nextButtonFocused {
		nextText = "[ next ]"
	} else if m.backButtonFocused {
		backText = "[ back ]"
	}

	formContent += "\n" + continueStyle.Render(nextText)
	formContent += "\n" + continueStyle.Render(backText)
	formContent = getScrollIndicators(startAbilityIdx, endAbilityIdx, len(labels), visibleFields) +
		formContent +
		m.renderError()

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		formContent,
	)
}

func (m *Model) updateViewportAbilities() {
	if m.height == 0 {
		return
	}
	linesPerAbility := 3 // label + base + proficient
	headerLines := 4     // header + icons + padding
	footerLines := 4     // buttons + padding
	availableLines := m.height - headerLines - footerLines
	visibleAbilities := max(availableLines/linesPerAbility, 1)

	maxAbilityOffset := max(0, 6-visibleAbilities)

	currentAbility := m.focused / 2
	if m.nextButtonFocused {
		m.viewportOffset = maxAbilityOffset * 2
		return
	}
	viewportAbility := m.viewportOffset / 2
	if currentAbility >= viewportAbility+visibleAbilities {
		m.viewportOffset = (currentAbility - visibleAbilities + 1) * 2
	}
	if currentAbility < viewportAbility {
		m.viewportOffset = currentAbility * 2
	}

	m.viewportOffset = min(m.viewportOffset, maxAbilityOffset*2)
}
