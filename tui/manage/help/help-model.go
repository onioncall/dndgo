package help

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type HelpModel struct {
	HelpViewport viewport.Model
}

func NewHelpModel() HelpModel {
	helpViewport := viewport.New(0, 0)
	helpContent := `Getting Started with dndgo

COMMANDS:
To modify your character, press Ctrl+S to open the command input.
Type a command and press Enter to execute it.

Available Commands:
  • damage <amount>        - Deal damage to your character
  • recover <amount>       - Heal your character (use "all" for full recovery)
  • temp <amount>          - Add temporary hit points
  • rename <name>          - Change your character's name
  • use-slot <level>       - Use a spell slot
  • recover-slot <level>   - Recover a spell slot
  • equip <weapon>         - Equip a weapon
  • unequip <slot>         - Unequip a weapon (primary/secondary)
  • add-item <name>/<qty>  - Add item to backpack
  • remove-item <name>/<qty> - Remove item from backpack

NAVIGATION:
  • Tab / Shift+Tab        - Switch between tabs
  • Esc                    - Return to main menu (saves character)
  • Ctrl+C                 - Quit (saves character)

TAB SHORTCUTS:
  • /b - Go to Basic Info
  • /s - Go to Spells
  • /e - Go to Equipment
  • /c - Go to Class

Press Ctrl+S to get started!`

	helpViewport.SetContent(helpContent)

	return HelpModel{
		HelpViewport: helpViewport,
	}
}

func (m HelpModel) UpdateSize(innerWidth, availableHeight int, character models.Character) HelpModel {
	helpInnerWidth := innerWidth - 4
	helpInnerHeight := availableHeight - 2

	m.HelpViewport.Width = helpInnerWidth
	m.HelpViewport.Height = helpInnerHeight

	return m
}
