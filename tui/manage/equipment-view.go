package manage

import (
	"github.com/charmbracelet/lipgloss"
)

func (m EquipmentModel) View(innerWidth, availableHeight int) string {
	// Row 1: 50/50 horizontal split, taking 50% of height
	row1Height := availableHeight / 2
	row2Height := availableHeight - row1Height

	// Worn Equipment viewport (left side of row 1)
	wornWidth := innerWidth / 2
	wornEquipmentStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(0, 2).
		Width(wornWidth - 2).
		Height(row2Height - 2).
		Align(lipgloss.Center)

	wornEquipmentVp := wornEquipmentStyle.Render(m.wornEquipmentViewport.View())

	// Backpack viewport (right side of row 1)
	backpackWidth := innerWidth - wornWidth
	backpackStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(0, 2).
		Width(backpackWidth - 2).
		Height(row2Height - 2).
		Align(lipgloss.Center)

	backpackVp := backpackStyle.Render(m.backpackViewport.View())

	// Join row 1 horizontally
	row1 := lipgloss.JoinHorizontal(lipgloss.Top, wornEquipmentVp, backpackVp)

	// Row 2: Weapons viewport (full width)
	weaponsStyle := lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(orange).
		Padding(0, 2).
		Width(innerWidth - 2).
		Height(row2Height - 2).
		Align(lipgloss.Center)

	weaponsVp := weaponsStyle.Render(m.weaponsViewport.View())

	// Join rows vertically
	return lipgloss.JoinVertical(lipgloss.Left, row1, weaponsVp)
}
