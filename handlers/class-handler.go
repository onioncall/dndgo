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
)

// We'll add more of these as needed
const (
	Bard      string = "bard"
	Barbarian string = "barbarian"
	Paladin   string = "paladin"
	Ranger    string = "ranger"
	Wizard    string = "wizard"
	Rogue     string = "rogue"
)

var ClassFileMap = map[string]string{
	Bard:      "bard-class.json",
	Barbarian: "barbarian-class.json",
	Ranger:    "default-ranger.json",
}

func LoadClass(classType string) (models.Class, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	// We aren't going to require a class file
	configPath := filepath.Join(homeDir, ".config/dndgo", "class.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil
	}

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to read class file: %w", err)
	}

	c, err := loadClassData(classType, fileData)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s class data: %w", classType, err)
	}

	return c, nil
}

func LoadClassTemplate(classType string) (models.IClass, error) {
	templateName := ClassFileMap[classType]
	if templateName == "" {
		return nil, fmt.Errorf("Unsupported class '%v'", classType)
	}

	fileData, err := defaultjsonconfigs.Content.ReadFile(templateName)
	if err != nil {
		return nil, fmt.Errorf("failed to read template class file: %w", err)
	}

	c, err := loadClassData(classType, fileData)
	if err != nil {
		return nil, fmt.Errorf("failed to load %s class data: %w", classType, err)
	}

	return c, nil
}

func SaveClassHandler(c models.Class) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "dndgo")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	filePath := filepath.Join(configDir, "class.json")

	characterJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling character to JSON: %w", err)
	}

	err = os.WriteFile(filePath, characterJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing character to file: %w", err)
	}

	fmt.Printf("Class json saved at: %s\n", filePath)
	return nil
}

func loadClassData(classType string, classData []byte) (models.Class, error) {
	var c models.Class
	var err error

	switch strings.ToLower(classType) {
	case Bard:
		c, err = class.LoadBard(classData)
	case Barbarian:
		c, err = class.LoadBarbarian(classData)
	case Ranger:
		c, err = class.LoadRanger(classData)
	default:
		return nil, fmt.Errorf("Unsupported class type '%v'", classType)
	}

	return c, err
}
