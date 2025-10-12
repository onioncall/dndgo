package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/onioncall/dndgo/models"
	"github.com/onioncall/dndgo/models/class"
)

// We'll add more of these as needed
const(
	Bard string = "bard"
	Barbarian string = "barbarian"
	Paladin string = "paladin"
	Ranger string = "ranger"
	Wizard string = "wizard"
	Rogue string = "rogue"
)

func LoadClass(classType string) (models.IClass, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	// We arn't going to require a class file
	configPath := filepath.Join(homeDir, ".config/dndgo", "class.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil
	}
	
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to read character file: %w", err)
	}
	
	var c models.IClass
	switch strings.ToLower(classType) {
	case Bard: c, err = class.LoadBard(fileData)
	case Barbarian: c, err = class.LoadBarbarian(fileData)
	default: fmt.Println("BAD") 
	}

	if err != nil {
		return nil, fmt.Errorf("failed to load %s class data: %w", classType, err)
	}

	return c, nil
}

func SaveClassHandler(c *models.IClass) error {
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
