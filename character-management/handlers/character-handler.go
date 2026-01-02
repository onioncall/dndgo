package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/onioncall/dndgo/character-management/db"
	defaultjsonconfigs "github.com/onioncall/dndgo/character-management/default-json-configs"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/search/handlers"
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
		return fmt.Errorf("Failed to save character markdown, Path: %s Error: %s", c.Path, err)
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

func CreateCharacter(c *models.Character) error {
	isUnique, err := IsUniqueCharacterName(c.Name)
	if err != nil {
		return fmt.Errorf("Failed to create character, unable to determine name uniqueness")
	} else if !isUnique {
		return fmt.Errorf("Character name '%s' is not unique", c.Name)
	}

	ex, err := db.Repo.GetCharacter()
	if err != nil {
		return fmt.Errorf("Failed to check for existing default character during creation: %w", err)
	}

	if ex == nil {
		c.Default = true
	}

	cid, err := db.Repo.InsertCharacter(*c)
	if err != nil {
		return fmt.Errorf("Failed to insert new character: %w", err)
	}

	if c.Class == nil && c.ClassName != "" {
		c.Class, err = LoadClassTemplate(c.ClassName)
	}

	if c.Class != nil {
		c.Class.SetCharacterId(cid)
		c.Class.SetClassName(c.ClassName)

		db.Repo.InsertClass(c.Class)
	}

	return nil
}

// Removes default from other character(s) and sets default for character with provided name
func SetDefaultCharacter(name string) error {
	// Getting character first because we don't want to clear the existing default until we know
	// that this one exists
	character, err := db.Repo.GetCharacterByName(name)
	if err != nil {
		return fmt.Errorf("Failed to get character with name '%s':\n%w", name, err)
	}

	defaultCharacters, err := db.Repo.GetDefaultCharacters()
	if err != nil {
		return fmt.Errorf("Failed to get default characters:\n%w", err)
	}

	for _, dc := range defaultCharacters {
		// if we only have one character and it's already the default, there's nothing left to do.
		if character.Name == dc.Name && len(defaultCharacters) == 1 {
			return nil
		}

		dc.Default = false
		err = db.Repo.SyncCharacter(dc)
		// We have to decide between a situation where we end up with multiple default characters here by returning early
		// or having no default characters by logging this error and moving on. Open to changing my mind about this.
		if err != nil {
			return fmt.Errorf("Failed to remove default from character '%s':\n%w", dc.Name, err)
		}
	}

	character.Default = true
	err = db.Repo.SyncCharacter(*character)
	if err != nil {
		return fmt.Errorf("Failed to update character '%s', to default:\n%w", name, err)
	}

	return nil
}

func SaveCharacter(c *models.Character) error {
	return db.Repo.SyncCharacter(*c)
}

func DeleteCharacter(name string) error {
	character, err := db.Repo.GetCharacterByName(name)
	if err != nil {
		return fmt.Errorf("Failed to find character to delete with name '%s':\n%w", name, err)
	}

	err = db.Repo.DeleteClassesByCharacterId(character.ID)
	if err != nil {
		return fmt.Errorf("Failed to delete class for character '%s':\n%w", name, err)
	}

	err = db.Repo.DeleteCharacter(character.ID)
	if err != nil {
		return fmt.Errorf("Failed to delete character '%s' (class files were successfully deleted):\n%w", name, err)
	}

	return nil
}

func LoadCharacter() (*models.Character, error) {
	character, err := db.Repo.GetCharacter()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve character from db: %w", err)
	}

	if character != nil && character.ClassName != "" {
		class, err := LoadClass(character.ID, character.ClassName)
		if err != nil {
			return nil, fmt.Errorf("failed to load class file: %w", err)
		}

		character.Class = class
	}

	return character, nil
}

func GetCharacterNames() ([]string, error) {
	return db.Repo.GetCharacterNames()
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
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return fmt.Errorf("Error creating directories: %w", err)
	}

	err = ClearFile(path)
	if err != nil {
		return fmt.Errorf("Error clearing file: %w", err)
	}

	err = os.WriteFile(path, []byte(res), 0o644)
	if err != nil {
		return fmt.Errorf("Error writing file: %w", err)
	}

	return nil
}

func ClearFile(filePath string) error {
	err := os.WriteFile(filePath, []byte{}, 0o644)
	if err != nil {
		return fmt.Errorf("failed to clear file: %w", err)
	}

	return nil
}

func ImportCharacterJson(characterJson []byte) error {
	var ch models.Character
	err := json.Unmarshal(characterJson, &ch)
	if err != nil {
		return fmt.Errorf("Parsing error on character json content: %w", err)
	}

	isUnique, err := IsUniqueCharacterName(ch.Name)
	if err != nil {
		return fmt.Errorf("Failed to create character, unable to determine name uniqueness")
	} else if !isUnique {
		return fmt.Errorf("Character name '%s' is not unique", ch.Name)
	}

	var existing *models.Character
	if ch.ID != "" {
		existing, err = db.Repo.GetCharacterById(ch.ID)
		if err != nil {
			return fmt.Errorf("Failed to check for existing character with specified ID in db: %w", err)
		}
	}

	if ch.ID == "" || existing == nil {
		defaultc, err := db.Repo.GetCharacter()
		if err != nil {
			return fmt.Errorf("Failed to check for existing 'default' character in db: %w", err)
		}
		if defaultc == nil {
			ch.Default = true
		}

		if _, err := db.Repo.InsertCharacter(ch); err != nil {
			return fmt.Errorf("Failed to create character: %w", err)
		}
	} else {
		if err := db.Repo.SyncCharacter(ch); err != nil {
			return fmt.Errorf("Failed to update character: %w", err)
		}
	}

	return nil
}

func ExportCharacterJson(characterName string) ([]byte, error) {
	ch, err := db.Repo.GetCharacterByName(characterName)
	if err != nil {
		return nil, fmt.Errorf("Failed to locate character with name '%v': %w", characterName, err)
	}

	data, err := json.MarshalIndent(ch, "", "    ")
	if err != nil {
		return nil, fmt.Errorf("Failed to parse existing character '%v': %w", characterName, err)
	}
	return data, nil
}

func IsUniqueCharacterName(name string) (bool, error) {
	names, err := GetCharacterNames()
	if err != nil {
		return false, fmt.Errorf("Failed to get list of existing character names:\n%w", err)
	}

	for _, en := range names {
		if strings.EqualFold(name, en) {
			return false, nil
		}
	}

	return true, nil
}
