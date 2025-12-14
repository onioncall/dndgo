package class

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
)

func (m ClassModel) View(innerWidth, availableHeight int) string {
	col1Width := innerWidth / 2
	subClassHeight := (availableHeight * 15) / 100
	detailsHeight := availableHeight - subClassHeight

	// Column 1 Viewports
	subClassVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 4).
		Width(col1Width - 2).
		Height(subClassHeight - 2).
		Align(lipgloss.Center)

	subClassVp := subClassVpStyle.Render(m.SubClassViewPort.View())

	detailVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 4).
		Width(col1Width - 2).
		Height(detailsHeight - 2)

	detialVp := detailVpStyle.Render(m.DetailViewport.View())

	// Stack them vertically
	column1 := lipgloss.JoinVertical(lipgloss.Left, subClassVp, detialVp)

	col2Width := innerWidth / 2

	// Column 2 Viewports
	otherFeaturesVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 2).
		Width(col2Width - 2).
		Height(availableHeight - 2)

	otherFeaturesVp := otherFeaturesVpStyle.Render(m.OtherFeaturesViewPort.View())

	column2 := otherFeaturesVp

	return lipgloss.JoinHorizontal(lipgloss.Top, column1, column2)
}
