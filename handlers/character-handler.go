package handlers

import (
	"encoding/json"
"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/onioncall/dndgo/models"
)

const (
	create 	string = "create"
	update 	string = "update"
	add		string = "add"
	remove	string = "remove"
	backpack string = "backpack"
)

func HandleCharacter(c *models.Character) {
	if c.Class != nil {
		c.Class.ExecutePreCalculateMethods(c)
	}

	c.CalculateCharacterStats()

	if c.Class != nil {
		c.Class.ExecutePostCalculateMethods(c)
	}

	res := c.BuildCharacter()
	SaveCharacterMarkdown(res, c.Path)
}

func AddSpell(c *models.Character, spellQuery string) {
	r := SpellRequest {
		Name: spellQuery,
		PathType: SpellType,
	}

	s := r.GetSingle()
	caltrop := false

	if s.Level == 0 {
		caltrop = true
	}

	cs := models.CharacterSpell {
		IsCaltrop: caltrop,
		SlotLevel: s.Level,
		IsRitual: s.Ritual,
		Name: s.Name,
	}

	c.Spells = append(c.Spells, cs)

	sort.Slice(c.Spells, func(i, j int) bool {
		return c.Spells[i].SlotLevel < c.Spells[j].SlotLevel
	})
}

func SaveCharacterJson(c *models.Character) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("error getting home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "dndgo")
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	filePath := filepath.Join(configDir, "character.json")

	characterJSON, err := json.MarshalIndent(c, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshaling character to JSON: %w", err)
	}

	err = os.WriteFile(filePath, characterJSON, 0644)
	if err != nil {
		return fmt.Errorf("error writing character to file: %w", err)
	}

	fmt.Printf("Character json saved at: %s\n", filePath)
	return nil
}

func LoadCharacter() (*models.Character, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ".config/dndgo", "character.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("character file not found at %s: %w", configPath, err)
	}
	
	fileData, err := os.ReadFile(configPath)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to read character file: %w", err)
	}
	
	var character models.Character
	if err := json.Unmarshal(fileData, &character); err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to parse character data: %w", err)
	}

	class, err := LoadClass(character.ClassName)
	if err != nil {
		fmt.Println(err)
		return nil, fmt.Errorf("failed to load class file: %w", err)
	}
	
	character.Class = class

	return &character, nil
}

func SaveCharacterMarkdown(res string, path string) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(fmt.Sprintf("Error getting home directory: %v", err))
	}

	filePath := filepath.Join(homeDir, path, "character.md")

	dir := filepath.Dir(filePath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		panic(fmt.Sprintf("Error creating directories: %v", err))
	}

	err = ClearFile(filePath)
	if err != nil {
		panic(fmt.Sprintf("Error clearing file: %v", err))
	}

	err = os.WriteFile(filePath, []byte(res), 0644)
	if err != nil {
		panic(fmt.Sprintf("Error writing file: %v", err))
	}

	fmt.Printf("Character markdown saved at: %s\n", filePath)
}

func ClearFile(filePath string) error {
    err := os.WriteFile(filePath, []byte{}, 0644)
    if err != nil {
        return fmt.Errorf("failed to clear file: %w", err)
    }
    
    return nil
}
