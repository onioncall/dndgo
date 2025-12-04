package manage

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
)

type Model struct {
	width              int
	height             int
	selectedTabIndex   int
	tabs               []string
	cmdInput           textinput.Model
	cmdVisible         bool
	character          *models.Character
	contentInitialized bool
	err                error

	basicInfoTab BasicInfoModel
	spellsTab    SpellsModel
	equipmentTab EquipmentModel
}

// Tab constants
const (
	basicInfoTab = iota
	spellTab
	equipmentTab
	classTab
)

// Tab Commands
const (
	// Navigation
	basicInfoCmd = "/b"
	spellCmd     = "/s"
	equipmentCmd = "/e"
	classCmd     = "/c"

	// Basic Info
	damageCmd  = "damage"
	recoverCmd = "recover"
	addTempCmd = "temp"

	// Spell Slots
	useSlotCmd     = "use-slot"
	recoverSlotCmd = "recover-slot"

	// Equipment
	addEquipmentCmd    = "add-equipment"
	removeEquipmentCmd = "remove-equipment"
	equipCmd           = "equip"
	unequipCmd         = "unequip"
	addItemCmd         = "add-item"
	removeItemCmd      = "remove-item"
)

func NewModel() Model {
	character, err := handlers.LoadCharacter()
	if err != nil {
		panic("Failed to load character")
	}

	err = handlers.HandleCharacter(character)
	if err != nil {
		panic("Failed to handle character")
	}

	input := textinput.New()
	input.Focus()
	input.Placeholder = "Cmd..."
	input.Width = 38

	tabs := []string{"Basic Info", "Equipment", "Class", "Notes"}
	if character.SpellSaveDC > 0 {
		tabs = []string{"Basic Info", "Spells", "Equipment", "Class", "Notes"}
	}

	basicInfoTab := NewBasicInfoModel(character)
	spellsTab := NewSpellsModel(character)
	equipmentTab := NewEquipmentModel(character)

	return Model{
		width:            0,
		height:           0,
		selectedTabIndex: 0,
		tabs:             tabs,
		cmdInput:         input,
		cmdVisible:       false,
		basicInfoTab:     basicInfoTab,
		spellsTab:        spellsTab,
		equipmentTab:     equipmentTab,
		character:        character,
	}
}

func (m Model) getInnerDimensions() (width, height int) {
	outerBorderMargin := 2
	bottomBoxHeight := 0
	if m.cmdVisible || m.err != nil {
		bottomBoxHeight = 3
	}

	containerWidth := m.width - (outerBorderMargin * 2) - 2
	containerHeight := m.height - (outerBorderMargin * 2) - 2 - bottomBoxHeight

	containerPadding := 2
	innerWidth := containerWidth - containerPadding
	innerHeight := containerHeight

	tabHeight := 3
	availableHeight := innerHeight - tabHeight

	return innerWidth, availableHeight
}
