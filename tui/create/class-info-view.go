package create

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) ClassPageView() string {
	if m.height == 0 {
		return ""
	}

	var formContent string

	// Top padding
	formContent += strings.Repeat("\n", 2)
	headerStyle := secondaryStyle.Width(41).Align(lipgloss.Center)
	formContent += headerStyle.Render("Class:") + "\n\n"

	formContent += fmt.Sprintf("%s\n%s\n\n",
		primaryStyle.Width(41).Render("Class Type"),
		m.inputs[classTypeInput].View(),
	)
	formContent += fmt.Sprintf("%s\n%s\n",
		primaryStyle.Width(41).Render("Class Level"),
		m.inputs[classLevelInput].View(),
	)

	addClassText := "add "
	nextText := "next"
	menuText := "back"
	if m.addButtonFocused {
		addClassText = "[ add  ]"
	} else if m.nextButtonFocused {
		nextText = "[ next ]"
	} else if m.backButtonFocused {
		menuText = "[ back ]"
	}

	formContent += "\n" + secondaryStyle.Render(addClassText)
	formContent += "\n" + secondaryStyle.Render(nextText)
	formContent += "\n" + secondaryStyle.Render(menuText)
	formContent = formContent + m.renderError()

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Top,
		formContent,
	)
}
