package manage

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	tui "github.com/onioncall/dndgo/tui/shared"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		innerWidth, availableHeight := m.getInnerDimensions()
		m.basicInfoTab = m.basicInfoTab.UpdateSize(innerWidth, availableHeight, m.character)
		m.spellsTab = m.spellsTab.UpdateSize(innerWidth, availableHeight, m.character)
		m.equipmentTab = m.equipmentTab.UpdateSize(innerWidth, availableHeight, m.character)

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			handlers.SaveCharacterJson(m.character)
			return m, tea.Quit
		case "esc":
			handlers.SaveCharacterJson(m.character)
			return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
		case "tab":
			m.selectedTabIndex = (m.selectedTabIndex + 1) % len(m.tabs)
			return m, nil
		case "shift+tab":
			m.selectedTabIndex = (m.selectedTabIndex - 1 + len(m.tabs)) % len(m.tabs)
			return m, nil
		case "ctrl+s":
			// Clearing error as a part of this process
			m.err = nil
			m.cmdVisible = !m.cmdVisible
			m.cmdInput.Focus()

			innerWidth, availableHeight := m.getInnerDimensions()
			switch m.selectedTabIndex {
			case basicInfoTab:
				m.basicInfoTab = m.basicInfoTab.UpdateSize(innerWidth, availableHeight, m.character)
			case spellTab:
				m.spellsTab = m.spellsTab.UpdateSize(innerWidth, availableHeight, m.character)
			case equipmentTab:
				m.equipmentTab = m.equipmentTab.UpdateSize(innerWidth, availableHeight, m.character)
			}

			return m, nil
		case "enter":
			if m.cmdVisible {
				searchInput := m.cmdInput.Value()
				m, m.selectedTabIndex, searchInput = m.executeUserCmd(searchInput, m.selectedTabIndex)
				m.cmdInput.SetValue("")
				m.cmdVisible = false
				handlers.HandleCharacter(m.character)
			}

			return m, nil
		}
	}

	var cmd tea.Cmd
	if m.cmdVisible {
		m.cmdInput, cmd = m.cmdInput.Update(msg)
		cmds = append(cmds, cmd)
		m.cmdInput.Focus()
	} else {
		m.cmdInput.Blur()
		switch m.selectedTabIndex {
		case basicInfoTab:
			m.basicInfoTab, cmd = m.basicInfoTab.Update(msg)
			cmds = append(cmds, cmd)
		case spellTab:
			m.spellsTab, cmd = m.spellsTab.Update(msg)
			cmds = append(cmds, cmd)
		case equipmentTab:
			m.equipmentTab, cmd = m.equipmentTab.Update(msg)
			cmds = append(cmds, cmd)
		}
	}

	return m, tea.Batch(cmds...)
}

func (m Model) executeUserCmd(cmdInput string, currentTab int) (Model, int, string) {
	cmd, inputAfterCmd, _ := strings.Cut(cmdInput, " ")
	tab := currentTab
	newInput := cmdInput
	switch strings.ToLower(cmd) {
	case basicInfoCmd:
		tab = basicInfoTab
		newInput = inputAfterCmd
	case spellCmd:
		tab = spellTab
		newInput = inputAfterCmd
	case equipmentCmd:
		tab = equipmentTab
		newInput = inputAfterCmd
	case classCmd:
		tab = classTab
		newInput = inputAfterCmd
	case damageCmd:
		dmg, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.DamageCharacter(int(dmg))
		m.basicInfoTab.healthViewport.SetContent(getHealthContent(m.character))
	case recoverCmd:
		health, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.HealCharacter(int(health))
		m.basicInfoTab.healthViewport.SetContent(getHealthContent(m.character))
	case addTempCmd:
		temp, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.AddTempHp(int(temp))
		m.basicInfoTab.healthViewport.SetContent(getHealthContent(m.character))
	case useSlotCmd:
		level, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.UseSpellSlot(int(level))
		m.spellsTab.spellSlotsViewport.SetContent(getSpellSlotContent(m.character))
	case recoverSlotCmd:
		level, err := strconv.ParseInt(inputAfterCmd, 10, 32)
		m.err = err
		m.character.RecoverSpellSlots(int(level), 1)
		m.spellsTab.spellSlotsViewport.SetContent(getSpellSlotContent(m.character))
	case addEquipmentCmd:
		m.err = execAddEquipmentCmd(inputAfterCmd, m.character)
		m.equipmentTab.wornEquipmentViewport.SetContent(getWornEquipmentContent(m.character))
	case equipCmd:
		m.err = execEquipCmd(inputAfterCmd, m.character)
		m.equipmentTab.weaponsViewport.SetContent(getWeaponsContent(m.character))
	case unequipCmd:
		m.err = execUnequipCmd(inputAfterCmd, m.character)
		m.equipmentTab.weaponsViewport.SetContent(getWeaponsContent(m.character))
	case addItemCmd:
		m.err = execModifyItemCmd(inputAfterCmd, true, m.character)
		m.equipmentTab.backpackViewport.SetContent(getBackpackContent(m.character))
	case removeItemCmd:
		m.err = execModifyItemCmd(inputAfterCmd, false, m.character)
		m.equipmentTab.backpackViewport.SetContent(getBackpackContent(m.character))
	default:
		m.err = fmt.Errorf("%s command not found", cmd)
	}

	return m, tab, newInput
}

func execModifyItemCmd(input string, isAdd bool, character *models.Character) error {
	splitInput := strings.Split(input, "/")
	itemName := input
	quantity := 1
	var err error

	if len(splitInput) == 2 {
		quantity, err = strconv.Atoi(splitInput[1])
		if err != nil {
			return fmt.Errorf("Invalid argument '%s', second (option argument must be an integer)", splitInput[1])
		}

		itemName = splitInput[0]
	} else if len(splitInput) != 1 {
		return fmt.Errorf("Invalid argument, (string, item name)/(optional int, quantity)")
	}

	if isAdd {
		character.AddItemToPack(itemName, quantity)
	} else {
		err = character.RemoveItemFromPack(itemName, quantity)
	}

	return err
}

func execUnequipCmd(input string, character *models.Character) error {
	// We're going to let the user optionally specify if they want to equip as primary or secondary.
	// If they don't specify, we'll equip the open slot. If no spot is open, we are going to equip primary
	isPrimary := false
	// handle secondary case first in instances where primary and secondary have the same name
	input = strings.ToLower(input)
	if input == "secondary" || input == strings.ToLower(character.SecondaryEquipped) {
		isPrimary = false
	} else if input == "primary" || input == strings.ToLower(character.PrimaryEquipped) {
		isPrimary = true
	} else {
		return fmt.Errorf("Invalid argument, (string, primary/secondary/weapon name)")
	}

	character.Unequip(isPrimary)

	return nil
}

func execEquipCmd(input string, character *models.Character) error {
	// We're going to let the user optionally specify if they want to equip as primary or secondary.
	// If they don't specify, we'll equip the open slot. If no spot is open, we are going to equip primary
	isPrimary := true
	splitInput := strings.Split(input, "/")
	weaponName := input // this can technically be weapon name or shield, but I don't want to put that everywhere

	if len(splitInput) == 1 {
		if character.PrimaryEquipped != "" && character.SecondaryEquipped == "" {
			isPrimary = false
		}
	} else if len(splitInput) == 2 {
		if strings.ToLower(splitInput[1]) == "primary" {
			isPrimary = true // Don't technically need to do this, but I think it's more clear
		} else if strings.ToLower(splitInput[1]) == "secondary" {
			isPrimary = false
		} else {
			return fmt.Errorf("Invalid argument '%s', (string, weapon/shield name)/(string, primary/secondary)",
				splitInput[1])
		}

		weaponName = splitInput[0]
	} else {
		return fmt.Errorf("Invalid argument, (string, weapon/shield name)/(string, primary/secondary)")
	}

	err := character.Equip(isPrimary, weaponName)
	return err
}

func execAddEquipmentCmd(input string, character *models.Character) error {
	splitInput := strings.Split(input, "/")
	if len(splitInput) < 2 {
		return fmt.Errorf("Too few arguments, (string, equipment type)/(string, equipment name)")
	} else if len(splitInput) > 2 {
		return fmt.Errorf("Too many arguments, (string, equipment type)/(string, equipment name)")
	}

	inputEquipmentType := splitInput[0]
	inputEquipmentName := splitInput[1]

	for _, wornEquipmentType := range shared.WornEquipmentTypes {
		if strings.ToLower(inputEquipmentType) == wornEquipmentType {
			character.AddEquipment(wornEquipmentType, inputEquipmentName)
			return nil
		}
	}

	return fmt.Errorf("Equipment type '%s', not found", inputEquipmentType)
}
