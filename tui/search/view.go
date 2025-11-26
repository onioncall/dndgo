package search

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	orange   = lipgloss.Color("#FFA500")
	darkGray = lipgloss.Color("#767676")
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
		BorderForeground(orange).
		Padding(0, 1)

	activeTab = tab.Border(activeTabBorder, true)

	tabGap = tab.
		BorderTop(false).
		BorderLeft(false).
		BorderRight(false)
)

func (m Model) View() string {
	if m.width == 0 || m.height == 0 {
		return ""
	}

	outerBorderMargin := 2
	containerStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(1, 2).
		Margin(outerBorderMargin).
		Width(m.width - (outerBorderMargin * 2) - 2).
		Height(m.height - (outerBorderMargin * 2) - 2)

	innerWidth := containerStyle.GetWidth() - containerStyle.GetHorizontalPadding()
	innerHeight := containerStyle.GetHeight() - containerStyle.GetVerticalPadding()

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
		fillerStyle := lipgloss.NewStyle().Foreground(orange)
		leftFiller := fillerStyle.Render(strings.Repeat("─", leftGap))
		rightFiller := fillerStyle.Render(strings.Repeat("─", rightGap))
		tabRow = lipgloss.JoinHorizontal(
			lipgloss.Bottom,
			leftFiller,
			tabRow,
			rightFiller,
		)
	}

	var searchRow string
	if m.searchVisible {
		searchStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(orange).
			Width(40)
		searchBox := searchStyle.Render(m.searchInput.View())
		searchRow = lipgloss.NewStyle().
			Width(innerWidth).
			Align(lipgloss.Center).
			Render(searchBox)
	}

	tabHeight := lipgloss.Height(tabRow)
	searchHeight := 0
	if m.searchVisible {
		searchHeight = lipgloss.Height(searchRow)
	}
	contentHeight := innerHeight - tabHeight - searchHeight

	contentBlock := lipgloss.NewStyle().
		Width(innerWidth).
		Height(contentHeight).
		Render(m.content)

	var sections []string
	sections = append(sections, tabRow)
	sections = append(sections, contentBlock)
	if m.searchVisible {
		sections = append(sections, searchRow)
	}

	inner := lipgloss.JoinVertical(lipgloss.Left, sections...)

	return containerStyle.Render(inner)
}
