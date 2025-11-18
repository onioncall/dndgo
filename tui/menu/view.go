package menu

import (
	"fmt"
	"github.com/onioncall/dndgo/tui/shared"
	"strings"
)

func (m Model) View() string {
	if m.currentPage == "menu" {
		return m.renderMenu()
	}

	switch m.currentPage {
	case shared.InitPage:
		return m.createPage.View()
	case shared.ManagePage:
		return m.managePage.View()
	case shared.SearchPage:
		return m.searchPage.View()
	default:
		return "Unknown page"
	}
}

func (m Model) renderMenu() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	var content strings.Builder

	// Center the text vertically
	textLines := 3 // pageText + empty line + buttons
	topPadding := (m.height - textLines) / 2

	for range topPadding {
		content.WriteString("\n")
	}

	// Center and write page text
	leftPadding := max((m.width-len(m.pageText))/2, 0)
	content.WriteString(strings.Repeat(" ", leftPadding))
	content.WriteString(m.pageText)
	content.WriteString("\n\n")

	// Build buttons line
	var buttonsLine strings.Builder
	for i, btn := range m.buttons {
		if i == m.selectedBtn {
			buttonsLine.WriteString(fmt.Sprintf("[ %s ]", btn))
		} else {
			buttonsLine.WriteString(fmt.Sprintf("  %s  ", btn))
		}
		if i < len(m.buttons)-1 {
			buttonsLine.WriteString("  ")
		}
	}

	// Center buttons
	btnLeftPadding := max((m.width-buttonsLine.Len())/2, 0)
	content.WriteString(strings.Repeat(" ", btnLeftPadding))
	content.WriteString(buttonsLine.String())

	return content.String()
}
