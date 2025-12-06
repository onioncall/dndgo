package help

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	orange = lipgloss.Color("#FFA500")
)

func (m HelpModel) View(innerWidth, availableHeight int) string {
	helpVpStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(1, 2).
		Width(innerWidth - 2).
		Height(availableHeight - 2)

	return helpVpStyle.Render(m.HelpViewport.View())
}
