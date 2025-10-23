package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	defaultjsonconfigs "github.com/onioncall/dndgo/default-json-configs"
	"github.com/onioncall/dndgo/models"
	"github.com/onioncall/dndgo/models/class"
	"github.com/onioncall/dndgo/types"
)

var ClassFileMap = map[string]string{
	types.ClassBarbarian: "barbarian.json",
	types.ClassBard:      "bard.json",
	types.ClassCleric:    "cleric.json",
	types.ClassDruid:     "druid.json",
	types.ClassFighter:   "fighter.json",
	types.ClassMonk:      "monk.json",
	types.ClassPaladin:   "paladin.json",
	types.ClassRanger:    "ranger.json",
	types.ClassRogue:     "rogue.json",
	types.ClassSorcerer:  "sorcerer.json",
	types.ClassWarlock:   "warlock.json",
	types.ClassWizard:    "wizard.json",
}

func LoadClass(classType string) (models.Class, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to get home directory: %w", err)
	}

	// We aren't going to require a class file
	configPath := filepath.Join(homeDir, ".config/dndgo", "class.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil
	}

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("Failed to read class file: %w", err)
	}

	c, err := loadClassData(classType, fileData)
	if err != nil {
		return nil, fmt.Errorf("Failed to load %s class data: %w", classType, err)
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

	c, err := loadClassData(classType, fileData)
	if err != nil {
		return nil, fmt.Errorf("Failed to load %s class data: %w", classType, err)
	}

	return c, nil
}

func SaveClassHandler(c models.Class) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error getting home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "dndgo")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("Error creating config directory: %w", err)
	}

	filePath := filepath.Join(configDir, "class.json")

	characterJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("Error marshaling character to JSON: %w", err)
	}

	err = os.WriteFile(filePath, characterJSON, 0644)
	if err != nil {
		return fmt.Errorf("Error writing character to file: %w", err)
	}

	return nil
}

func loadClassData(classType string, classData []byte) (models.Class, error) {
	var c models.Class
	var err error

	switch strings.ToLower(classType) {
	case types.ClassBarbarian:
		c, err = class.LoadBarbarian(classData)
	case types.ClassBard:
		c, err = class.LoadBard(classData)
	case types.ClassCleric:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	case types.ClassDruid:
		c, err = class.LoadDruid(classData)
	case types.ClassFighter:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	case types.ClassMonk:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	case types.ClassPaladin:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	case types.ClassRanger:
		c, err = class.LoadRanger(classData)
	case types.ClassRogue:
		c, err = class.LoadRogue(classData)
	case types.ClassSorcerer:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	case types.ClassWarlock:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	case types.ClassWizard:
		c = nil
		err = fmt.Errorf("%s not implemented yet", classType)
	default:
		return nil, fmt.Errorf("Unsupported class type '%s'", classType)
	}

	return c, err
}
