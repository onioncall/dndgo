package create

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) BackpackPageView() string {
	if m.height == 0 {
		return ""
	}
	var formContent string

	// Top padding
	formContent += strings.Repeat("\n", 2)
	headerStyle := secondaryStyle.Width(41).Align(lipgloss.Center)
	formContent += headerStyle.Render("Backpack Inventory:") + "\n\n"

	formContent += fmt.Sprintf("%s\n%s\n\n",
		primaryStyle.Width(41).Render("Item Name"),
		m.inputs[itemNameInput].View(),
	)
	formContent += fmt.Sprintf("%s\n%s\n",
		primaryStyle.Width(41).Render("Item Quantity"),
		m.inputs[itemQuantityInput].View(),
	)

	addItemText := "add "
	nextText := "save"
	menuText := "back"
	if m.addButtonFocused {
		addItemText = "[ add  ]"
	} else if m.nextButtonFocused {
		nextText = "[ save ]"
	} else if m.backButtonFocused {
		menuText = "[ back ]"
	}

	formContent += "\n" + secondaryStyle.Render(addItemText)
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
