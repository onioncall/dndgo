package manage

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
)

var (
	activeTabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      " ",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}
	tabBorder = lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┴",
		BottomRight: "┴",
	}
	tab = lipgloss.NewStyle().
		Border(tabBorder, true).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true).
			Foreground(orange).
			Bold(true)

	inactiveTab = lipgloss.NewStyle().
			Border(tabBorder, true).
			BorderForeground(lightBlue).
			Foreground(darkGray).
			Padding(0, 1)
)

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	if m.character == nil {
		return m.renderNoCharacter()
	}

	outerBorderMargin := 2
	bottomBoxHeight := 0
	if m.cmdVisible || m.err != nil {
		bottomBoxHeight = 3
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 1).
		MarginTop(outerBorderMargin).
		MarginLeft(outerBorderMargin).
		MarginRight(outerBorderMargin).
		Width(m.width - (outerBorderMargin * 2) - 2).
		Height(m.height - (outerBorderMargin * 2) - 2 - bottomBoxHeight)

	innerWidth := containerStyle.GetWidth() - containerStyle.GetHorizontalPadding()
	innerHeight := containerStyle.GetHeight() - containerStyle.GetVerticalPadding()

	// Render tab row with character name
	tabRow := m.renderTabRow(innerWidth)
	tabHeight := lipgloss.Height(tabRow)

	availableHeight := innerHeight - tabHeight
	availableWidth := innerWidth

	innerContent := m.renderActiveTabContent(availableWidth, availableHeight)

	// Build content
	var sections []string
	sections = append(sections, tabRow)
	sections = append(sections, innerContent)

	inner := lipgloss.JoinVertical(lipgloss.Left, sections...)
	container := containerStyle.Render(inner)

	// Show error box if there's an error, otherwise show cmd box if visible
	if m.err != nil {
		errorBox := m.renderErrorBox()
		return lipgloss.JoinVertical(lipgloss.Left, container, errorBox)
	}

	if m.cmdVisible {
		cmdBox := m.renderCmdBox()
		return lipgloss.JoinVertical(lipgloss.Left, container, cmdBox)
	}

	return container
}

func (m Model) renderActiveTabContent(width, height int) string {
	switch m.selectedTabIndex {
	case basicInfoTab:
		return m.basicInfoTab.View(width, height)
	case spellTab:
		return m.spellsTab.View(width, height)
	case equipmentTab:
		return m.equipmentTab.View(width, height)
	case classTab:
		return m.classTab.View(width, height)
	case notesTab:
		return m.notesTab.View(width, height)
	case helpTab:
		return m.helpTab.View(width, height)
	default:
		return m.renderPlaceholderContent(width, height)
	}
}

func (m Model) renderTabRow(innerWidth int) string {
	headerBorder := lipgloss.Border{
		Top:         "─",
		Bottom:      "─",
		Left:        "│",
		Right:       "│",
		TopLeft:     "╭",
		TopRight:    "╮",
		BottomLeft:  "┘",
		BottomRight: "└",
	}

	// Render character name
	headerText := lipgloss.NewStyle().
		Border(headerBorder, true).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 1).
		Bold(true).
		Render(m.character.Name)

	// Build tabs
	var rendered []string
	for i, t := range m.tabs {
		if i == m.selectedTabIndex {
			rendered = append(rendered, activeTab.Render(t))
		} else if m.character.SpellSaveDC == 0 && i == spellTab {
			rendered = append(rendered, inactiveTab.Render(t))
		} else {
			rendered = append(rendered, tab.Render(t))
		}
	}
	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, rendered...)
	tabsWidth := lipgloss.Width(tabRow)

	// Add header text to the left of tabs
	headerTextWidth := lipgloss.Width(headerText)
	totalGap := innerWidth - tabsWidth - headerTextWidth
	leftGap := totalGap / 2
	rightGap := totalGap - leftGap

	if totalGap > 0 {
		fillerStyle := lipgloss.NewStyle().Foreground(lightBlue)
		leftFiller := fillerStyle.Render(strings.Repeat("─", leftGap))
		rightFiller := fillerStyle.Render(strings.Repeat("─", rightGap))
		return lipgloss.JoinHorizontal(
			lipgloss.Bottom,
			headerText,
			leftFiller,
			tabRow,
			rightFiller,
		)
	}

	return lipgloss.JoinHorizontal(
		lipgloss.Bottom,
		headerText,
		tabRow,
	)
}

func (m Model) renderCmdBox() string {
	searchStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Width(40)
	searchBox := searchStyle.Render(m.cmdInput.View())
	return lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render(searchBox)
}

// Placeholder until we implement tab models
func (m Model) renderPlaceholderContent(width, height int) string {
	placeholder := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Align(lipgloss.Center, lipgloss.Center).
		Render("This tab is under construction")
	return placeholder
}

func (m Model) renderErrorBox() string {
	errorStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FF0000")).
		Padding(0, 1).
		Foreground(lipgloss.Color("#FF0000")).
		Width(40)
	errorBox := errorStyle.Render(m.err.Error())

	return lipgloss.NewStyle().
		Width(m.width).
		Align(lipgloss.Center).
		Render(errorBox)
}

func (m Model) renderNoCharacter() string {
	noCharacterStyle := lipgloss.NewStyle().
		Padding(0, 1).
		Border(lipgloss.RoundedBorder()).
		Align(lipgloss.Center, lipgloss.Center).
		BorderForeground(lightBlue).
		Foreground(cream)

	noCharacterBox := noCharacterStyle.Render("No character loaded, consider creating a character!\n\nPress ctrl+c to exit or 'esc' to return to menu.")

	return lipgloss.NewStyle().
		Width(m.width).
		Height(m.height).
		Align(lipgloss.Center, lipgloss.Center).
		Render(noCharacterBox)
}
