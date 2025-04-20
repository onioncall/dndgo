package handlers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/onioncall/dndgo/models"
)

const (
	create 	string = "create"
	update 	string = "update"
	add		string = "add"
	remove	string = "remove"
	backpack string = "backpack"
)

func HandleCharacter(query string) {
	splitQuery := strings.Split(query, " ")

	c, err := LoadCharacter()
	if err != nil {
		panic(err)
	}

	switch splitQuery[0] {
	case create: 
	case update:
		calculateCharacterStats(c)
		res := buildCharacter(c)
		saveCharacter(res, c.Path)
	case add:
		// addToCharacter(splitQuery[1:])
	}
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

	return &character, nil
}

// func addToCharacter(args []string, c *models.Character) {
// 	quantity := args[len(args)-1]
// 	switch args[0] {
// 		case backpack:
// 			c.AddItemToPack(args[1:len(args)-1],)
// 	}
// }

func calculateCharacterStats(c *models.Character) {
	calculateProficiencyBonusByLevel(c)	
	calculateProficienciesFromBase(c)
	calculateSkillBonusFromBase(c)
}

func calculateProficienciesFromBase(c *models.Character) {
	for i, prof := range c.Proficiencies {
		c.Proficiencies[i].Bonus = (prof.Base - 10) / 2
	}
}

func calculateSkillBonusFromBase(c *models.Character) {
	for i, skill := range c.Skills {
		// if this is too slow, I'll refactor this to use a map with the proficiency name as the key
		for _, prof := range c.Proficiencies {
			if skill.Proficiency == prof.Name {
				c.Skills[i].Bonus = prof.Bonus 
			}
		}
	}
}

func calculateProficiencyBonusByLevel(c *models.Character) {
	// Proficiency Bonus
	if c.Level <= 4 {
		c.Proficiency = 2	
	} else if c.Level > 4 && c.Level <= 8 {
		c.Proficiency = 3	
	} else if c.Level > 8 && c.Level <= 12 {
		c.Proficiency = 4	
	} else if c.Level > 12 && c.Level <= 16 {
		c.Proficiency = 5	
	} else if c.Level > 16 && c.Level <= 20 {
		c.Proficiency = 6	
	}
}

func buildCharacter(c *models.Character) string {
	var builder strings.Builder	
	nl := "\n"

	header := c.BuildHeader()
	for i := range header {
		builder.WriteString(header[i])	
	}
	builder.WriteString(nl)

	characterInfo := c.BuildCharacterInfo()
	for i := range characterInfo {
		builder.WriteString(characterInfo[i])	
	}
	builder.WriteString(nl)

	feats := c.BuildFeats()
	for i := range feats {
		builder.WriteString(feats[i])
	}
	builder.WriteString(nl)

	languages := c.BuildLanguages()
	for i := range languages {
		builder.WriteString(languages[i])
	}
	builder.WriteString(nl)

	generalStats := c.BuildGeneralStats()
	for i := range generalStats {
		builder.WriteString(generalStats[i])
	}
	builder.WriteString(nl)
	
	proficiencies := c.BuildProficiencies()
	for i := range proficiencies {
		builder.WriteString(proficiencies[i])
	}
	builder.WriteString(nl)

	skills := c.BuildSkills()
	for i := range skills {
		builder.WriteString(skills[i])
	}
	builder.WriteString(nl)

	spells := c.BuildSpells() 
	for i := range spells {
		builder.WriteString(spells[i]) 
	}
	builder.WriteString(nl)
	
	weapons := c.BuildWeapons()
	for i := range weapons {
		builder.WriteString(weapons[i]) 
	}
	builder.WriteString(nl)

	equipment := c.BuildEquipment()
	for i := range equipment {
		builder.WriteString(equipment[i]) 
	}
	builder.WriteString(nl)

	backpack := c.BuildBackpack()
	for i := range backpack {
		builder.WriteString(backpack[i]) 
	}
	builder.WriteString(nl)

    result := builder.String()
	return result
}

func saveCharacter(res string, path string) {
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
}

func ClearFile(filePath string) error {
    err := os.WriteFile(filePath, []byte{}, 0644)
    if err != nil {
        return fmt.Errorf("failed to clear file: %w", err)
    }
    
    return nil
}
