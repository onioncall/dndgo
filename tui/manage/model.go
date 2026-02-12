package manage

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/onioncall/dndgo/character-management/handlers"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/logger"
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
	character          *models.Character
	contentInitialized bool
	currentClass       string
	err                error

	basicInfoTab info.BasicInfoModel
	spellsTab    spells.SpellsModel
	equipmentTab equipment.EquipmentModel
	classTab     class.ClassModel
	notesTab     notes.NotesModel
	helpTab      help.HelpModel

	keyBindings map[int]keyBinding
	visibleCmd  int

	commands       []string
	autoSuggestion string
}

type keyBinding struct {
	shortcut string
	cmdFunc  func(Model) Model
	input    *textinput.Model
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

const (
	paletteKeybinding = iota
	damageKeybinding
	recoverHpKeybinding
	longRestKeybinding
	useSpellKeybinding
	recoverSpellSlotKeybinding
	removeItemKeybinding
	addItemKeybinding
	useClassTokenKeybinding
	recoverClassTokenKeybinding
)

const cmdInactive = 99

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
	updateClassCmd  = "update-class"

	// Class
	useClassTokenCmd     = "use-token"
	recoverClassTokenCmd = "recover-token"
)

func NewModel() Model {
	character, err := handlers.LoadCharacter()
	if err != nil {
		logger.Info("Failed to load character")
	}

	if character != nil {
		err = handlers.HandleCharacter(character)
		if err != nil {
			logger.Info("Failed to handle character")
		}
	}

	defaultClass := ""
	if character != nil && len(character.ClassTypes) > 0 {
		defaultClass = character.ClassTypes[0]
	}

	paletteInput := textinput.New()
	paletteInput.Focus()
	paletteInput.Placeholder = "enter command..."
	paletteInput.Prompt = " cmd> "
	paletteInput.Width = 38

	damageInput := textinput.New()
	damageInput.Focus()
	damageInput.Placeholder = "hit points to reduce..."
	damageInput.Prompt = " damage> "
	damageInput.Width = 38

	recoverHpInput := textinput.New()
	recoverHpInput.Focus()
	recoverHpInput.Placeholder = "hit points to recover..."
	recoverHpInput.Prompt = " recover hp> "
	recoverHpInput.Width = 38

	longRestInput := textinput.New()
	longRestInput.Focus()
	longRestInput.Placeholder = "yes or no"
	longRestInput.Prompt = " long rest?> "
	longRestInput.Width = 38

	useSpellInput := textinput.New()
	useSpellInput.Focus()
	useSpellInput.Placeholder = "level of slot to use..."
	useSpellInput.Prompt = " cast spell> "
	useSpellInput.Width = 38

	useTokenInput := textinput.New()
	useTokenInput.Focus()
	useTokenInput.Placeholder = "class token to use..."
	useTokenInput.Prompt = " use token> "
	useTokenInput.Width = 38

	keyBindings := make(map[int]keyBinding)

	keyBindings[paletteKeybinding] = keyBinding{"ctrl+p", ExecPaletteKeyBinding, &paletteInput}
	keyBindings[damageKeybinding] = keyBinding{"ctrl+d", ExecDamageKeyBinding, &damageInput}
	keyBindings[recoverHpKeybinding] = keyBinding{"ctrl+r", ExecRecoverKeyBinding, &recoverHpInput}
	keyBindings[longRestKeybinding] = keyBinding{"ctrl+l", ExecLongRestKeyBinding, &longRestInput}
	keyBindings[useSpellKeybinding] = keyBinding{"ctrl+s", ExecUseSpellKeyBinding, &useSpellInput}
	keyBindings[useClassTokenKeybinding] = keyBinding{"ctrl+t", ExecUseClassTokenKeyBinding, &useTokenInput}

	// Currently can't get shift+char to work, so holding off on implementing the following until I do
	// keyBindings[recoverSpellSlotKeybinding] = "ctrl+S"
	// keyBindings[removeItemKeybinding] = "ctrl+i"
	// keyBindings[addItemKeybinding] = "ctrl+I"
	// keyBindings[recoverClassTokenKeybinding] = "ctrl+T"

	commands := []string{
		addEquipmentCmd,
		addItemCmd,
		damageCmd,
		equipCmd,
		recoverCmd,
		recoverSlotCmd,
		recoverClassTokenCmd,
		removeItemCmd,
		renameCmd,
		addTempCmd,
		unequipCmd,
		updateClassCmd,
		useSlotCmd,
		useClassTokenCmd,
	}

	tabs := []string{"Basic Info", "Spells", "Equipment", "Class", "Notes", "Help"}

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
		currentClass:     defaultClass,
		basicInfoTab:     basicInfoTab,
		spellsTab:        spellsTab,
		equipmentTab:     equipmentTab,
		classTab:         classTab,
		notesTab:         notesTab,
		helpTab:          helpTab,
		character:        character,
		keyBindings:      keyBindings,
		visibleCmd:       cmdInactive,
		commands:         commands,
		autoSuggestion:   "",
	}
}

func (m Model) getInnerDimensions() (width, height int) {
	outerBorderMargin := 2
	bottomBoxHeight := 0
	if m.visibleCmd != 99 || m.err != nil {
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
