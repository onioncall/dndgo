package class

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	darkGray  = lipgloss.Color("#767676")
)

func (m ClassModel) View(innerWidth, availableHeight int) string {
	col1Width := innerWidth / 2
	col1Height := availableHeight / 2

	// Column 1 Viewports
	tokenVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Padding(0, 4).
		Width(col1Width - 2).
		Height(col1Height - 2).
		Align(lipgloss.Center)

	tokenVp := tokenVpStyle.Render(m.TokenViewPort.View())

	detailVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Padding(0, 2).
		Width(col1Width - 2).
		Height(col1Height - 2).
		Align(lipgloss.Center)

	detialVp := detailVpStyle.Render(m.DetailViewPort.View())

	// Stack them vertically
	column1 := lipgloss.JoinVertical(lipgloss.Left, tokenVp, detialVp)

	col2Width := innerWidth / 2

	// Column 2 Viewports
	otherFeaturesVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Padding(0, 2).
		Width(col2Width - 2).
		Height(availableHeight - 2)

	otherFeaturesVp := otherFeaturesVpStyle.Render(m.OtherFeaturesViewPort.View())

	column2 := otherFeaturesVp

	return lipgloss.JoinHorizontal(lipgloss.Top, column1, column2)
}
