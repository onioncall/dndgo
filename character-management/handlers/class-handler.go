package handlers

import (
	"encoding/json"
	"fmt"
	"slices"
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

// Consider refactor to use an actual custom error
var noClassNameError = fmt.Errorf("No class name provided in class data")

func LoadClass(characterId string, classType string) (models.Class, error) {
	c, err := newClassInst(classType)
	if err != nil {
		return nil, fmt.Errorf("Failed to load class type: %w", err)
	}

	err = db.Repo.GetClass(characterId, c, classType)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve class from db: %w", err)
	}

	return c, nil
}

func LoadClassesByCharacterId(characterId string, classTypes []string) ([]models.Class, error) {
	classInstances := []models.Class{}
	classTypesLower := make([]string, len(classTypes)) // doing this to avoid mutating casing of original slice
	for _, classType := range classTypes {
		classTypesLower = append(classTypesLower, strings.ToLower(classType))
		c, err := newClassInst(classType)
		if err != nil {
			return nil, fmt.Errorf("Failed to load class type: %w", err)
		}

		classInstances = append(classInstances, c)
	}

	err := db.Repo.GetClassesByCharacterId(characterId, classInstances)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve class from db: %w", err)
	}

	for _, class := range classInstances {
		if !slices.Contains(classTypesLower, strings.ToLower(class.GetClassType())) {
			return nil, fmt.Errorf("Class '%s' does not match defined classes by character.", class.GetClassType())
		}
	}

	return classInstances, nil
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

	class, err := loadClassInstFromType(baseClass.GetClassType(), classData)
	if err != nil {
		return nil, fmt.Errorf("Failed to load class instance from type:\n%w", err)
	}

	return class, nil
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

	c.SetClassType(classType)

	return c, err
}

func ImportClassJson(classJson []byte, characterName string, classType string) error {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return fmt.Errorf("Failed to retrieve character with name '%v': %w", characterName, err)
	}

	c, err := loadClassInst(classJson)
	if err != nil {
		return fmt.Errorf("Failed to load class instance from json:\n%w", err)
	}

	if c.GetCharacterId() != "" && c.GetCharacterId() != ch.ID {
		return fmt.Errorf("Character ID (%s) for character name '%s' does not match character ID in class (%s).",
			ch.ID, characterName, c.GetCharacterId())
	}

	classes, err := LoadClassesByCharacterId(ch.ID, ch.ClassTypes)
	if err != nil {
		return fmt.Errorf("Failed to load class for character '%v': %w", characterName, err)
	}

	var ec models.Class
	for _, class := range classes {
		if strings.EqualFold(class.GetClassType(), classType) || len(classes) == 1 {
			ec = class
		}
	}

	if ec.GetCharacterId() == "" {
		if err := db.Repo.InsertClass(c); err != nil {
			return fmt.Errorf("Failed to create class: %w", err)
		}
	} else {
		if ec.GetClassType() != c.GetClassType() {
			return fmt.Errorf("Class switching not supported (%v -> %v)", c.GetClassType(), ec.GetClassType())
		}

		if err := db.Repo.SyncClass(c); err != nil {
			return fmt.Errorf("Failed to update class: %w", err)
		}
	}

	return nil
}

func ExportClassJson(characterName string, classType string) ([]byte, error) {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate character with name '%v': %w", characterName, err)
	}

	classes, err := LoadClassesByCharacterId(ch.ID, ch.ClassTypes)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate class for character '%v': %w", characterName, err)
	}

	var classToExport models.Class
	for _, class := range classes {
		if strings.EqualFold(class.GetClassType(), classType) || len(classes) == 1 {
			classToExport = class
		}
	}

	j, err := json.MarshalIndent(classToExport, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal class for character '%v': %w", characterName, err)
	}

	return j, nil
}
