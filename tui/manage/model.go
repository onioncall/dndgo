package manage

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/manage/class"
	"github.com/onioncall/dndgo/tui/manage/equipment"
	"github.com/onioncall/dndgo/tui/manage/help"
	"github.com/onioncall/dndgo/tui/manage/info"
	"github.com/onioncall/dndgo/tui/manage/notes"
	"github.com/onioncall/dndgo/tui/manage/spells"
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

	basicInfoTab info.BasicInfoModel
	spellsTab    spells.SpellsModel
	equipmentTab equipment.EquipmentModel
	classTab     class.ClassModel
	notesTab     notes.NotesModel
	helpTab      help.HelpModel
}

// Tab constants
const (
	basicInfoTab = iota
	spellTab
	equipmentTab
	classTab
	notesTab
	helpTab
)

// Tab Commands
const (
	// Navigation
	basicInfoCmd = "/b"
	spellCmd     = "/s"
	equipmentCmd = "/e"
	classCmd     = "/c"
	helpCmd      = "/h"

	// Basic Info
	damageCmd  = "damage"
	recoverCmd = "recover"
	addTempCmd = "temp"
	renameCmd  = "rename"

	// Spell Slots
	useSlotCmd     = "use-slot"
	recoverSlotCmd = "recover-slot"

	// Equipment
	addEquipmentCmd = "add-equipment"
	equipCmd        = "equip"
	unequipCmd      = "unequip"
	addItemCmd      = "add-item"
	removeItemCmd   = "remove-item"

	// Class
	useClassTokenCmd     = "use-token"
	recoverClassTokenCmd = "recover-token"
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

	tabs := []string{"Basic Info", "Equipment", "Class", "Notes", "Help"}
	if character.SpellSaveDC > 0 {
		tabs = []string{"Basic Info", "Spells", "Equipment", "Class", "Notes", "Help"}
	}

	basicInfoTab := info.NewBasicInfoModel()
	spellsTab := spells.NewSpellsModel()
	equipmentTab := equipment.NewEquipmentModel()
	classTab := class.NewClassModel()
	notesTab := notes.NewNotesModel()
	helpTab := help.NewHelpModel()

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
		classTab:         classTab,
		notesTab:         notesTab,
		helpTab:          helpTab,
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
