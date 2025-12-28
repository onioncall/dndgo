package create

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) SkillsPageView() string {
	if m.height == 0 {
		return ""
	}

	linesPerField := 1 // label/input in line for these
	headerLines := 4   // header + icons + padding
	footerLines := 4   // buttons + padding

	availableLines := m.height - headerLines - footerLines
	visibleFields := max(availableLines/linesPerField, 1)
	startIdx := m.viewportOffset
	endIdx := min(startIdx+visibleFields, len(m.inputs))

	var formContent string

	// Top padding
	formContent += strings.Repeat("\n", 2)
	headerStyle := secondaryStyle.Width(70).Align(lipgloss.Center)
	formContent += headerStyle.Render("Skill Proficiencies:") + "\n\n"

	for i := startIdx; i < endIdx; i++ {
		formContent += fmt.Sprintf("%s %s %s\n",
			primaryStyle.Width(30).Render(skillToAbility[i].ability),
			primaryStyle.Width(30).Render(skillToAbility[i].name),
			m.inputs[i].View(),
		)
	}

	nextText := "next"
	menuText := "back"
	if m.nextButtonFocused {
		nextText = "[ next ]"
	} else if m.backButtonFocused {
		menuText = "[ back ]"
	}

	formContent += "\n" + secondaryStyle.Render(nextText)
	formContent += "\n" + secondaryStyle.Render(menuText)
	formContent = getScrollIndicators(startIdx, endIdx, len(skillToAbility), visibleFields) + formContent + m.renderError()

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		formContent,
	)
}
