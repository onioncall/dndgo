package menu

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/onioncall/dndgo/tui/shared"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
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

	textStyle := lipgloss.NewStyle().Foreground(cream)
	selectedBtnStyle := lipgloss.NewStyle().Foreground(cream).Bold(true)
	unselectedBtnStyle := lipgloss.NewStyle().Foreground(cream)

	var content strings.Builder
	textLines := 3 // pageText + empty line + buttons
	topPadding := (m.height - textLines) / 2
	for range topPadding {
		content.WriteString("\n")
	}

	leftPadding := max((m.width-len(m.pageText))/2, 0)
	content.WriteString(strings.Repeat(" ", leftPadding))
	content.WriteString(textStyle.Render(m.pageText))
	content.WriteString("\n\n")

	var visualButtons []string
	for i, btn := range m.buttons {
		if i == m.selectedBtn {
			visualButtons = append(visualButtons, fmt.Sprintf("[ %s ]", btn))
		} else {
			visualButtons = append(visualButtons, fmt.Sprintf("  %s  ", btn))
		}
	}
	
	visualWidth := 0
	for i, vBtn := range visualButtons {
		visualWidth += len(vBtn)
		if i < len(visualButtons)-1 {
			visualWidth += 2 // spacing
		}
	}

	var buttonsLine strings.Builder
	for i, btn := range m.buttons {
		if i == m.selectedBtn {
			buttonsLine.WriteString(selectedBtnStyle.Render(fmt.Sprintf("[ %s ]", btn)))
		} else {
			buttonsLine.WriteString(unselectedBtnStyle.Render(fmt.Sprintf("  %s  ", btn)))
		}
		if i < len(m.buttons)-1 {
			buttonsLine.WriteString("  ")
		}
	}

	btnLeftPadding := max((m.width-visualWidth)/2, 0)
	content.WriteString(strings.Repeat(" ", btnLeftPadding))
	content.WriteString(buttonsLine.String())

	return content.String()
}
