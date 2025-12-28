package create

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) EquipmentPageView() string {
	if m.height == 0 {
		return ""
	}

	linesPerField := 2 // label + input
	headerLines := 4   // header + icons + padding
	footerLines := 4   // buttons + padding

	availableLines := m.height - headerLines - footerLines
	visibleFields := max(availableLines/linesPerField, 1)

	startIdx := m.viewportOffset
	endIdx := min(startIdx+visibleFields, len(m.inputs))

	labels := []string{
		"Head",
		"Amulet",
		"Cloak",
		"Hands/Arms",
		"Ring 1",
		"Ring 2",
		"Belt",
		"Boots",
		"Shield",
		"Armor",
		"Armor Proficient",
		"Armor Class",
		"Armor Type",
	}

	var formContent string

	// Top padding
	formContent += strings.Repeat("\n", 2)
	headerStyle := secondaryStyle.Width(41).Align(lipgloss.Center)
	formContent += headerStyle.Render("Worn Equipment:") + "\n\n"

	for i := startIdx; i < endIdx; i++ {
		formContent += fmt.Sprintf("%s\n%s\n",
			primaryStyle.Width(41).Render(labels[i]),
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
	formContent = getScrollIndicators(startIdx, endIdx, len(labels), visibleFields) + formContent + m.renderError()

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		formContent,
	)
}
