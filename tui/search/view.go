package search

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
		Foreground(orange).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
)

func (m *Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}
	outerBorderMargin := 2

	// Reserve space for search row outside the container
	searchHeight := 0
	if m.searchVisible {
		searchHeight = 3 // border top + content + border bottom
	}

	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(1, 2).
		MarginTop(outerBorderMargin).
		MarginLeft(outerBorderMargin).
		MarginRight(outerBorderMargin).
		Width(m.width - (outerBorderMargin * 2) - 2).
		Height(m.height - (outerBorderMargin * 2) - 2 - searchHeight)

	innerWidth := containerStyle.GetWidth() - containerStyle.GetHorizontalPadding()
	innerHeight := containerStyle.GetHeight() - containerStyle.GetVerticalPadding()

	// Search Tabs
	var rendered []string
	for i, t := range m.tabs {
		if i == m.selectedTabIndex {
			rendered = append(rendered, activeTab.Render(t))
		} else {
			rendered = append(rendered, tab.Render(t))
		}
	}

	tabRow := lipgloss.JoinHorizontal(lipgloss.Top, rendered...)
	tabsWidth := lipgloss.Width(tabRow)
	totalGap := innerWidth - tabsWidth
	leftGap := totalGap / 2
	rightGap := totalGap - leftGap

	if totalGap > 0 {
		fillerStyle := lipgloss.NewStyle().Foreground(lightBlue)
		leftFiller := fillerStyle.Render(strings.Repeat("─", leftGap))
		rightFiller := fillerStyle.Render(strings.Repeat("─", rightGap))
		tabRow = lipgloss.JoinHorizontal(
			lipgloss.Bottom,
			leftFiller,
			tabRow,
			rightFiller,
		)
	}

	tabHeight := lipgloss.Height(tabRow)
	contentHeight := innerHeight - tabHeight

	// Viewport
	m.viewport.Width = innerWidth
	m.viewport.Height = contentHeight

	contentBlock := m.viewport.View()

	var sections []string
	sections = append(sections, tabRow)
	sections = append(sections, contentBlock)
	inner := lipgloss.JoinVertical(lipgloss.Left, sections...)
	container := containerStyle.Render(inner)

	// Search Box
	if m.searchVisible {
		searchStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lightBlue).
			Foreground(cream).
			Width(40)
		searchBox := searchStyle.Render(m.searchInput.View())
		searchRow := lipgloss.NewStyle().
			Width(m.width).
			Align(lipgloss.Center).
			Render(searchBox)
		return lipgloss.JoinVertical(lipgloss.Left, container, searchRow)
	}

	return container
}
