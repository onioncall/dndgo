package create

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) WeaponsPageView() string {
	if m.height == 0 {
		return ""
	}

	linesPerField := 2 // label/input in line for these
	headerLines := 4   // header + icons + padding
	footerLines := 5   // buttons + padding

	availableLines := m.height - headerLines - footerLines
	visibleFields := max(availableLines/linesPerField, 1)

	startIdx := m.viewportOffset
	endIdx := min(startIdx+visibleFields, len(m.inputs))

	var formContent string

	// Top padding
	formContent += strings.Repeat("\n", 2)
	headerStyle := secondaryStyle.Width(41).Align(lipgloss.Center)
	formContent += headerStyle.Render("Weapons:") + "\n\n"

	labels := []string{
		"Weapon Name",
		"Damage",
		"Proficient",
		"Ranged",
		"Normal Range",
		"Long Range",
		"Type",
		"Properties",
	}

	for i := startIdx; i < endIdx; i++ {
		formContent += fmt.Sprintf("%s\n%s\n",
			primaryStyle.Width(41).Render(labels[i]),
			m.inputs[i].View(),
		)
	}

	addSpellText := "add "
	nextText := "next"
	menuText := "back"
	if m.addButtonFocused {
		addSpellText = "[ add  ]"
	}
	if m.nextButtonFocused {
		nextText = "[ next ]"
	}
	if m.backButtonFocused {
		menuText = "[ back ]"
	}

	formContent += "\n" + secondaryStyle.Render(addSpellText)
	formContent += "\n" + secondaryStyle.Render(nextText)
	formContent += "\n" + secondaryStyle.Render(menuText)
	formContent = getScrollIndicators(startIdx, endIdx, len(labels), visibleFields) + formContent + m.renderError()

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		formContent,
	)
}
