package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/onioncall/dndgo/character-management/db"
	defaultjsonconfigs "github.com/onioncall/dndgo/character-management/default-json-configs"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
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
		c.HitDice = c.Class.CalculateHitDice(c.Level)

		if preCalculater, ok := c.Class.(models.PreCalculator); ok {
			preCalculater.ExecutePreCalculateMethods(c)
		}
	}

	c.CalculateCharacterStats()

	if c.Class != nil {
		if postCalculater, ok := c.Class.(models.PostCalculator); ok {
			postCalculater.ExecutePostCalculateMethods(c)
		}
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

	cs := shared.CharacterSpell{
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

func SaveCharacter(c *models.Character) error {
	return db.Repo.SyncCharacter(*c)
}

func LoadCharacter() (*models.Character, error) {
	character, err := db.Repo.GetCharacter()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve character from db:\n%w", err)
	}

	if character.ClassName != "" {
		class, err := LoadClass(character.ID, character.ClassName)
		if err != nil {
			return nil, fmt.Errorf("failed to load class file: %w", err)
		}

		character.Class = class
	}

	return character, nil
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
	character.Default = true

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

func ImportCharacterJson(characterJson []byte) error {
	var ch models.Character
	if err := json.Unmarshal(characterJson, ch); err != nil {
		return fmt.Errorf("Parsing error on character json content:\n%w", err)
	}

	if ch.ID == "" {
		if _, err := db.Repo.InsertCharacter(ch); err != nil {
			return fmt.Errorf("Failed to create character:\n%w", err)
		}
	} else {
		if err := db.Repo.SyncCharacter(ch); err != nil {
			return fmt.Errorf("Failed to update character:\n%w", err)
		}
	}

	return nil
}

func ExportCharacterJson(characterName string) ([]byte, error) {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate character with name '%v':\n%w", characterName, err)
	}

	data, err := json.Marshal(ch)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse existing character '%v':\n%w", characterName, err)
	}
	return data, nil
}
