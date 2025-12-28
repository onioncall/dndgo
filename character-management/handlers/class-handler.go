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

func LoadClass(characterId string, className string) (models.Class, error) {
	// Filthy trick, we can use loadClassDataFromType an empty instance of the correct class
	c, err := loadClassDataFromType(className, make([]byte, 0))
	if err != nil {
		return nil, fmt.Errorf("Failed to load class type:\n%w", err)
	}

	err = db.Repo.GetClass(characterId, c)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve class from db:\n%w", err)
	}

	return c, nil
}

func LoadClassTemplate(classType string) (models.Class, error) {
	templateName := ClassFileMap[classType]
	if templateName == "" {
		return nil, fmt.Errorf("Unsupported class '%s'", classType)
	}

	fileData, err := defaultjsonconfigs.Content.ReadFile(templateName)
	if err != nil {
		return nil, fmt.Errorf("Failed to read template class file: %w", err)
	}

	c, err := loadClassDataFromType(classType, fileData)
	if err != nil {
		return nil, fmt.Errorf("Failed to load %s class data: %w", classType, err)
	}

	return c, nil
}

func SaveClassHandler(c models.Class) error {
	return db.Repo.SyncClass(c)
}

func loadClassData(classData []byte) (models.Class, error) {
	var baseClass models.BaseClass
	if err := json.Unmarshal(classData, baseClass); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal class data to base struct:\n%w", err)
	}

	return loadClassDataFromType(baseClass.GetClassName(), classData)
}

func loadClassDataFromType(classType string, classData []byte) (models.Class, error) {
	var c models.Class
	var err error

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

	return c, err
}

func ImportClassJson(classJson []byte) error {
	c, err := loadClassData(classJson)
	if err != nil {
		return err
	}

	if c.GetCharacterId() == "" {
		return fmt.Errorf("No CharacterID found. CharacterID is required.")
	}

	ec, err := loadClassDataFromType(c.GetClassName(), make([]byte, 0))
	if err != nil {
		return fmt.Errorf("Failed to load existing class data:\n%w", err)
	}
	err = db.Repo.GetClass(c.GetCharacterId(), ec)
	if err != nil {
		return fmt.Errorf("Failed to retrieve existing class data:\n%w", err)
	}

	if ec.GetClassName() == "" {
		if err := db.Repo.InsertClass(c); err != nil {
			return fmt.Errorf("Failed to create class:\n%w", err)
		}
	} else {

		if ec.GetClassName() != c.GetClassName() {
			return fmt.Errorf("Class switching not supported (%v -> %v)", c.GetClassName(), ec.GetClassName())
		}

		if err := db.Repo.SyncClass(c); err != nil {
			return fmt.Errorf("Failed to update class:\n%w", err)
		}
	}

	return nil
}

func ExportClassJson(characterName string) ([]byte, error) {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate character with name '%v':\n%w", characterName, err)
	}

	c, err := LoadClass(ch.ID, ch.ClassName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate class for character '%v':\n%w", characterName, err)
	}

	j, err := json.Marshal(c)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal class for character '%v':\n%w", characterName, err)
	}

	return j, nil
}
