package menu

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/tui/shared"
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

	textLines := 3 // pageText + empty line + buttons
	topPadding := (m.height - textLines) / 2

	for range topPadding {
		content.WriteString("\n")
	}

	// For centering horizontally
	leftPadding := max((m.width-len(m.pageText))/2, 0)
	content.WriteString(strings.Repeat(" ", leftPadding))
	content.WriteString(m.pageText)
	content.WriteString("\n\n")

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

	btnLeftPadding := max((m.width-buttonsLine.Len())/2, 0)
	content.WriteString(strings.Repeat(" ", btnLeftPadding))
	content.WriteString(buttonsLine.String())

	// Calculate remaining space to push version to bottom
	currentLines := topPadding + textLines
	versionPadding := m.height - currentLines - 3 // 3 lines from bottom (with padding)
	for range versionPadding {
		content.WriteString("\n")
	}

	// Add version info at bottom
	content.WriteString("\n")
	versionText := fmt.Sprintf("v%s", m.version)
	versionLeftPadding := max((m.width-len(versionText))/2, 0)
	content.WriteString(strings.Repeat(" ", versionLeftPadding))
	content.WriteString(versionText)

	return content.String()
}
