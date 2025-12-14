package help

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	orange    = lipgloss.Color("#FFA500")
	lightBlue = lipgloss.Color("#5DC9E2")
	cream     = lipgloss.Color("#F9F6F0")
	darkGray  = lipgloss.Color("#767676")
)

func (m HelpModel) View(innerWidth, availableHeight int) string {
	helpVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lightBlue).
		Foreground(cream).
		Padding(1, 2).
		Width(innerWidth - 2).
		Height(availableHeight - 2)

	return helpVpStyle.Render(m.HelpViewport.View())
}
