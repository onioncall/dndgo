package info

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	abilitiesPadding int = 2
	skillsPadding    int = 4
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
)

func (m BasicInfoModel) View(innerWidth, availableHeight int) string {
	col1Width := innerWidth / 3
	statsHeight := availableHeight / 2
	abilitiesHeight := availableHeight - statsHeight

	// Column 1 Viewports
	statsVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 4).
		Width(col1Width - 2).
		Height(statsHeight - 2)

	basicStatsVp := statsVpStyle.Render(m.BasicStatsViewport.View())

	abilitiesVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, abilitiesPadding).
		Width(col1Width - abilitiesPadding).
		Height(abilitiesHeight - 2).
		Align(lipgloss.Center)

	abilitiesVp := abilitiesVpStyle.Render(m.AbilitiesViewport.View())

	// Stack them vertically
	column1 := lipgloss.JoinVertical(lipgloss.Left, basicStatsVp, abilitiesVp)

	// Column 2: 2/3 horizontally, split 15/85 vertically
	col2Width := (innerWidth * 2) / 3
	healthHeight := (availableHeight * 15) / 100
	skillsHeight := availableHeight - healthHeight

	// Column 2 Viewports
	healthVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, 2).
		Width(col2Width - 2).
		Height(healthHeight - 2).
		Align(lipgloss.Center)

	healthVp := healthVpStyle.Render(m.HealthViewport.View())

	skillsVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(0, skillsPadding).
		Width(col2Width - 2).
		Height(skillsHeight - 2).
		Align(lipgloss.Center)

	skillsVp := skillsVpStyle.Render(m.SkillsViewport.View())

	column2 := lipgloss.JoinVertical(lipgloss.Left, healthVp, skillsVp)

	return lipgloss.JoinHorizontal(lipgloss.Top, column1, column2)
}
