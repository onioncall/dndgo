package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	defaultjsonconfigs "github.com/onioncall/dndgo/character-management/default-json-configs"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
	"github.com/onioncall/dndgo/search/handlers"
)

const (
	create   string = "create"
	update   string = "update"
	add      string = "add"
	remove   string = "remove"
	backpack string = "backpack"
)

func HandleCharacter(c *models.Character) error {
	if c.Class != nil {
		c.Class.ExecutePreCalculateMethods(c)
	}

	c.CalculateCharacterStats()

	if c.Class != nil {
		c.Class.ExecutePostCalculateMethods(c)
	}

	res := c.BuildCharacter()
	err := SaveCharacterMarkdown(res, c.Path)
	if err != nil {
		return fmt.Errorf("Failed to save character markdown, Path: %s\nError: %s", c.Path, err)
	}

	return nil
}

func AddSpell(c *models.Character, spellQuery string) error {
	r := handlers.SpellRequest{
		Name:     spellQuery,
		PathType: handlers.SpellType,
	}

	s, err := r.GetSingle()
	if err != nil {
		return fmt.Errorf("Failed To get spell (%s) to add: %w", spellQuery, err)
	}

	cs := types.CharacterSpell{
		SlotLevel: s.Level,
		IsRitual:  s.Ritual,
		Name:      s.Name,
	}

	c.Spells = append(c.Spells, cs)

	// sorting spells by level
	sort.Slice(c.Spells, func(i, j int) bool {
		return c.Spells[i].SlotLevel < c.Spells[j].SlotLevel
	})

	return nil
}

func GetConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("error getting home directory: %w", err)
	}

	configDir := filepath.Join(homeDir, ".config", "dndgo")
	return configDir, nil
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

	return nil
}

func LoadCharacter() (*models.Character, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get home directory: %w", err)
	}

	configPath := filepath.Join(homeDir, ".config", "dndgo", "character.json")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, fmt.Errorf("character file not found at %s: %w", configPath, err)
	}

	fileData, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read character file: %w", err)
	}

	var character models.Character
	if err := json.Unmarshal(fileData, &character); err != nil {
		return nil, fmt.Errorf("failed to parse character data: %w", err)
	}

	if character.ClassName != "" {
		class, err := LoadClass(character.ClassName)
		if err != nil {
			return nil, fmt.Errorf("failed to load class file: %w", err)
		}

		character.Class = class
	}

	return &character, nil
}

func LoadCharacterTemplate(characterName string, className string) (*models.Character, error) {
	filePath := "character.json"
	fileData, err := defaultjsonconfigs.Content.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read template character file: %w", err)
	}

	var character models.Character
	if err := json.Unmarshal(fileData, &character); err != nil {
		return nil, fmt.Errorf("failed to parse character data: %w", err)
	}
	character.Name = characterName
	character.ClassName = className

	return &character, nil
}

func SaveCharacterMarkdown(res string, path string) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("Error getting home directory: %w", err)
	}

	// If path is empty, we're going to default to the config path
	if path == "" {
		path = filepath.Join(".config", "dndgo")
	}

	path = filepath.Join(homeDir, path, "character.md")

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("Error creating directories: %w", err)
	}

	err = ClearFile(path)
	if err != nil {
		return fmt.Errorf("Error clearing file: %w", err)
	}

	err = os.WriteFile(path, []byte(res), 0644)
	if err != nil {
		return fmt.Errorf("Error writing file: %w", err)
	}

	return nil
}

func ClearFile(filePath string) error {
	err := os.WriteFile(filePath, []byte{}, 0644)
	if err != nil {
		return fmt.Errorf("failed to clear file: %w", err)
	}

	return nil
}
