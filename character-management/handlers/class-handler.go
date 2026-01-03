package handlers

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/db"
	defaultjsonconfigs "github.com/onioncall/dndgo/character-management/default-json-configs"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/models/class"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

var ClassFileMap = map[string]string{
	shared.ClassBarbarian: "barbarian.json",
	shared.ClassBard:      "bard.json",
	shared.ClassCleric:    "cleric.json",
	shared.ClassDruid:     "druid.json",
	shared.ClassFighter:   "fighter.json",
	shared.ClassMonk:      "monk.json",
	shared.ClassPaladin:   "paladin.json",
	shared.ClassRanger:    "ranger.json",
	shared.ClassRogue:     "rogue.json",
	shared.ClassSorcerer:  "sorcerer.json",
	shared.ClassWarlock:   "warlock.json",
	shared.ClassWizard:    "wizard.json",
}

// Consider refactor to use an actual custom error
var noClassNameError = fmt.Errorf("No class name provided in class data")

func LoadClass(characterId string, className string) (models.Class, error) {
	c, err := newClassInst(className)
	if err != nil {
		return nil, fmt.Errorf("Failed to load class type: %w", err)
	}

	err = db.Repo.GetClass(characterId, c)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve class from db: %w", err)
	}

	return c, nil
}

func LoadClassTemplate(classType string) (models.Class, error) {
	templateName := ClassFileMap[strings.ToLower(classType)]
	if templateName == "" {
		return nil, fmt.Errorf("Unsupported class '%s'", classType)
	}

	fileData, err := defaultjsonconfigs.Content.ReadFile(templateName)
	if err != nil {
		return nil, fmt.Errorf("Failed to read template class file: %w", err)
	}

	c, err := loadClassInstFromType(classType, fileData)
	if err != nil {
		return nil, fmt.Errorf("Failed to load %s class data: %w", classType, err)
	}

	return c, nil
}

func SaveClass(c models.Class) error {
	return db.Repo.SyncClass(c)
}

func loadClassInst(classData []byte) (models.Class, error) {
	var baseClass models.BaseClass
	if err := json.Unmarshal(classData, &baseClass); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal class data to base struct: %w", err)
	}

	return loadClassInstFromType(baseClass.GetClassName(), classData)
}

func newClassInst(classType string) (models.Class, error) {
	return loadClassInstFromType(classType, []byte("{}"))
}

func loadClassInstFromType(classType string, classData []byte) (models.Class, error) {
	var c models.Class
	var err error

	if string(classType) == "" {
		return nil, noClassNameError
	}

	switch strings.ToLower(classType) {
	case shared.ClassBarbarian:
		c, err = class.LoadBarbarian(classData)
	case shared.ClassBard:
		c, err = class.LoadBard(classData)
	case shared.ClassCleric:
		c, err = class.LoadCleric(classData)
	case shared.ClassDruid:
		c, err = class.LoadDruid(classData)
	case shared.ClassFighter:
		c, err = class.LoadFighter(classData)
	case shared.ClassMonk:
		c, err = class.LoadMonk(classData)
	case shared.ClassPaladin:
		c, err = class.LoadPaladin(classData)
	case shared.ClassRanger:
		c, err = class.LoadRanger(classData)
	case shared.ClassRogue:
		c, err = class.LoadRogue(classData)
	case shared.ClassSorcerer:
		c, err = class.LoadSorcerer(classData)
	case shared.ClassWarlock:
		c, err = class.LoadWarlock(classData)
	case shared.ClassWizard:
		c, err = class.LoadWizard(classData)
	default:
		return nil, fmt.Errorf("Unsupported class type '%s'", classType)
	}

	c.SetClassName(classType)

	return c, err
}

func ImportClassJson(classJson []byte, characterName string) error {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return fmt.Errorf("Failed to retrieve character with name '%v': %w", characterName, err)
	}

	c, err := loadClassInst(classJson)
	if err != nil {
		// If json is missing class-name, use class name from character as fallback.
		// When multiclassing is introduced, this will no longer be an option.
		if err == noClassNameError {
			c, err = loadClassInstFromType(ch.ClassName, classJson)
			if err != nil {
				return fmt.Errorf("failed to load class: %w", err)
			}
		} else {
			return err
		}
	}

	if c.GetCharacterId() != "" && c.GetCharacterId() != ch.ID {
		logger.Warnf("character-id provided in json '%v' does not match ID of character '%v', this value will not be included in the import", c.GetCharacterId(), characterName)
	}
	c.SetCharacterId(ch.ID)

	ec, err := newClassInst(c.GetClassName())
	if err != nil {
		return fmt.Errorf("Failed to load existing class data: %w", err)
	}
	err = db.Repo.GetClass(c.GetCharacterId(), ec)
	if err != nil {
		return fmt.Errorf("Failed to retrieve existing class data: %w", err)
	}

	if ec.GetCharacterId() == "" {
		if err := db.Repo.InsertClass(c); err != nil {
			return fmt.Errorf("Failed to create class: %w", err)
		}
	} else {

		if ec.GetClassName() != c.GetClassName() {
			return fmt.Errorf("Class switching not supported (%v -> %v)", c.GetClassName(), ec.GetClassName())
		}

		if err := db.Repo.SyncClass(c); err != nil {
			return fmt.Errorf("Failed to update class: %w", err)
		}
	}

	return nil
}

func ExportClassJson(characterName string) ([]byte, error) {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate character with name '%v': %w", characterName, err)
	}

	c, err := LoadClass(ch.ID, ch.ClassName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate class for character '%v': %w", characterName, err)
	}

	j, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal class for character '%v': %w", characterName, err)
	}

	return j, nil
}
