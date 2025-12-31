package manage

import (
	"fmt"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/tui/manage/class"
	"github.com/onioncall/dndgo/tui/manage/equipment"
	"github.com/onioncall/dndgo/tui/manage/info"
	"github.com/onioncall/dndgo/tui/manage/spells"
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

		if m.character != nil {
			innerWidth, availableHeight := m.getInnerDimensions()
			m.basicInfoTab = m.basicInfoTab.UpdateSize(innerWidth, availableHeight, *m.character)
			if m.character.SpellSaveDC > 0 {
				m.spellsTab = m.spellsTab.UpdateSize(innerWidth, availableHeight, *m.character)
			}
			m.equipmentTab = m.equipmentTab.UpdateSize(innerWidth, availableHeight, *m.character)
			m.classTab = m.classTab.UpdateSize(innerWidth, availableHeight, *m.character)
			m.notesTab = m.notesTab.UpdateSize(innerWidth, availableHeight, *m.character)
			m.helpTab = m.helpTab.UpdateSize(innerWidth, availableHeight, *m.character)
		}

		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			if m.character != nil {
				handlers.SaveCharacter(m.character)
			}
			return m, tea.Quit
		case "esc":
			if m.character != nil {
				handlers.SaveCharacter(m.character)
			}
			return m, func() tea.Msg { return tui.NavigateMsg{Page: tui.MenuPage} }
		case "tab":
			m.selectedTabIndex = (m.selectedTabIndex + 1) % len(m.tabs)
			// Skip spell tab if we're not rendering it
			if m.character.SpellSaveDC == 0 && m.selectedTabIndex == spellTab {
				m.selectedTabIndex++
			}

			return m, nil
		case "shift+tab":
			m.selectedTabIndex = (m.selectedTabIndex - 1 + len(m.tabs)) % len(m.tabs)
			// Skip spell tab if we're not rendering it
			if m.character.SpellSaveDC == 0 && m.selectedTabIndex == spellTab {
				m.selectedTabIndex--
			}

			return m, nil
		case "ctrl+s":
			// Clearing error as a part of this process
			m.err = nil
			m.cmdVisible = !m.cmdVisible
			m.cmdInput.Focus()

			if m.character != nil {
				innerWidth, availableHeight := m.getInnerDimensions()
				switch m.selectedTabIndex {
				case basicInfoTab:
					m.basicInfoTab = m.basicInfoTab.UpdateSize(innerWidth, availableHeight, *m.character)
				case spellTab:
					m.spellsTab = m.spellsTab.UpdateSize(innerWidth, availableHeight, *m.character)
				case equipmentTab:
					m.equipmentTab = m.equipmentTab.UpdateSize(innerWidth, availableHeight, *m.character)
				case classTab:
					m.classTab = m.classTab.UpdateSize(innerWidth, availableHeight, *m.character)
				case notesTab:
					m.notesTab = m.notesTab.UpdateSize(innerWidth, availableHeight, *m.character)
				case helpTab:
					m.helpTab = m.helpTab.UpdateSize(innerWidth, availableHeight, *m.character)
				}
			}

			return m, nil
		case "enter":
			if m.character == nil {
				return m, nil
			}

			if m.cmdVisible {
				searchInput := m.cmdInput.Value()
				m, m.selectedTabIndex, searchInput = m.executeUserCmd(searchInput, m.selectedTabIndex)
				m.cmdInput.SetValue("")
				m.cmdVisible = false
				handlers.SaveCharacter(m.character)
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
		if m.character != nil {
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
			case classTab:
				m.classTab, cmd = m.classTab.Update(msg)
			case notesTab:
				m.notesTab, cmd = m.notesTab.Update(msg)
			case helpTab:
				m.helpTab, cmd = m.helpTab.Update(msg)
			}
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
	case spellCmd:
		tab = spellTab
	case equipmentCmd:
		tab = equipmentTab
	case classCmd:
		tab = classTab
	case helpCmd:
		tab = helpTab
	case damageCmd:
		dmg, err := strconv.Atoi(inputAfterCmd)
		m.err = err
		m.character.DamageCharacter(int(dmg))
		m.basicInfoTab.HealthViewport.SetContent(info.GetHealthContent(*m.character))
	case recoverCmd:
		m.err = execRecoverCmd(inputAfterCmd, m.character)
		m.basicInfoTab.HealthViewport.SetContent(info.GetHealthContent(*m.character))
	case addTempCmd:
		temp, err := strconv.Atoi(inputAfterCmd)
		m.err = err
		m.character.AddTempHp(int(temp))
		m.basicInfoTab.HealthViewport.SetContent(info.GetHealthContent(*m.character))
	// TODO: Rename functionality will have to change with the support of multiple character files
	// case renameCmd:
	// 	if inputAfterCmd != "" {
	// 		m.character.RenameCharacter(inputAfterCmd)
	// 		m.basicInfoTab.BasicStatsViewport.SetContent(info.GetStatsContent(*m.character))
	// 	} else {
	// 		m.err = fmt.Errorf("name cannot be empty")
	// 	}
	case useSlotCmd:
		if m.character.SpellSaveDC == 0 {
			m.err = fmt.Errorf("Character cannot use spell commands")
			break
		}

		level, err := strconv.Atoi(inputAfterCmd)
		m.err = err
		m.character.UseSpellSlot(int(level))
		sWidth := m.spellsTab.SpellSlotsViewport.Width
		m.spellsTab.SpellSlotsViewport.SetContent(spells.GetSpellSlotContent(*m.character, sWidth))
	case recoverSlotCmd:
		if m.character.SpellSaveDC == 0 {
			m.err = fmt.Errorf("Character cannot use spell commands")
			break
		}

		level, err := strconv.Atoi(inputAfterCmd)
		m.err = err
		m.character.RecoverSpellSlots(int(level), 1)
		sWidth := m.spellsTab.SpellSlotsViewport.Width
		m.spellsTab.SpellSlotsViewport.SetContent(spells.GetSpellSlotContent(*m.character, sWidth))
	case addEquipmentCmd:
		m.err = execAddEquipmentCmd(inputAfterCmd, m.character)
		weWidth := m.equipmentTab.WornEquipmentViewport.Width
		m.equipmentTab.WornEquipmentViewport.SetContent(equipment.GetWornEquipmentContent(*m.character, weWidth))
	case equipCmd:
		m.err = execEquipCmd(inputAfterCmd, m.character)
		wpWidth := m.equipmentTab.WeaponsViewport.Width
		m.equipmentTab.WeaponsViewport.SetContent(equipment.GetWeaponsContent(*m.character, wpWidth))
	case unequipCmd:
		m.err = execUnequipCmd(inputAfterCmd, m.character)
		wpWidth := m.equipmentTab.WeaponsViewport.Width
		m.equipmentTab.WeaponsViewport.SetContent(equipment.GetWeaponsContent(*m.character, wpWidth))
	case addItemCmd:
		m.err = execModifyItemCmd(inputAfterCmd, true, m.character)
		bpWidth := m.equipmentTab.BackpackViewport.Width
		m.equipmentTab.BackpackViewport.SetContent(equipment.GetBackpackContent(*m.character, bpWidth))
	case removeItemCmd:
		m.err = execModifyItemCmd(inputAfterCmd, false, m.character)
		bpWidth := m.equipmentTab.BackpackViewport.Width
		m.equipmentTab.BackpackViewport.SetContent(equipment.GetBackpackContent(*m.character, bpWidth))
	case useClassTokenCmd:
		m.err = execUseClassTokenCmd(inputAfterCmd, m.character)
		m.classTab.DetailViewport.SetContent(class.GetClassDetails(*m.character))
	case recoverClassTokenCmd:
		m.err = execRecoverClassTokenCmd(inputAfterCmd, m.character)
		m.classTab.DetailViewport.SetContent(class.GetClassDetails(*m.character))
	default:
		m.err = fmt.Errorf("%s command not found", cmd)
	}

	return m, tab, newInput
}

func execRecoverClassTokenCmd(input string, character *models.Character) error {
	splitInput := strings.Split(input, "/")
	tokenName := input
	quantity := 0 // By default for recover, we assume a full recover unless a quantity is specified
	var err error

	if len(splitInput) == 2 {
		quantity, err = strconv.Atoi(splitInput[1])
		if err != nil {
			return fmt.Errorf("Invalid argument '%s', second (option argument must be an integer)", splitInput[1])
		}

		tokenName = splitInput[0]
	} else if len(splitInput) != 1 {
		return fmt.Errorf("Invalid argument, (string, token name)/(optional int, quantity)")
	}

	character.RecoverClassTokens(tokenName, quantity)
	return err
}

func execUseClassTokenCmd(input string, character *models.Character) error {
	splitInput := strings.Split(input, "/")
	tokenName := input
	quantity := 1 // By default for use token, we assume one use unless a quantity is specified
	var err error

	if len(splitInput) == 2 {
		quantity, err = strconv.Atoi(splitInput[1])
		if err != nil {
			return fmt.Errorf("Invalid argument '%s', second (option argument must be an integer)", splitInput[1])
		}

		tokenName = splitInput[0]
	} else if len(splitInput) != 1 {
		return fmt.Errorf("Invalid argument, (string, token name)/(optional int, quantity)")
	}

	character.UseClassTokens(tokenName, quantity)
	return err
}

func execRecoverCmd(input string, character *models.Character) error {
	if input == "all" {
		character.Recover()
	} else {
		health, err := strconv.Atoi(input)
		if err != nil {
			return err
		}

		character.HealCharacter(int(health))
	}

	return nil
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
