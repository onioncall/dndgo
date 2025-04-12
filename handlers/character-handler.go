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
	create string = "create"
	update string = "update"
)

func HandleCharacter(query string) {
	switch query {
	case create: 
	case update:
		c, err := LoadCharacter()
		if err != nil {
			panic(err)
		}
		calculateCharacterStats(c)
		res := buildCharacter(c)
		saveCharacter(res, c.Path)
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


func calculateCharacterStats(c *models.Character) {
	calculateProficiencyBonusByLevel(c)	
	calculateProficienciesFromBase(c)
}

func calculateProficienciesFromBase(c *models.Character) {
	for i, prof := range c.Proficiencies {
		c.Proficiencies[i].Adjusted = prof.Base

		if prof.Proficient {
			c.Proficiencies[i].Adjusted += c.Proficiency
		}

		bonus := (c.Proficiencies[i].Adjusted - 10) / 2
		sign := ""
		if bonus > 0 {
			sign = "+"
		}

		c.Proficiencies[i].Bonus = fmt.Sprintf("%s%d", sign, bonus) 
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

	// Header
	header := fmt.Sprintf("#DnD Character\n")
	builder.WriteString(header)
	builder.WriteString(nl)

	nameLine := fmt.Sprintf("**Name: %s**\n", c.Name)
	builder.WriteString(nameLine)
	builder.WriteString(nl)

	// Character Info
	levelLine 		:= fmt.Sprintf("Level: %d\n", c.Level)
	classLine 		:= fmt.Sprintf("Class: %s\n", c.Class)
	raceLine 		:= fmt.Sprintf("Race: %s\n", c.Race)
	backgroundLine 	:= fmt.Sprintf("Background: %s\n", c.Background)
	builder.WriteString(levelLine)
	builder.WriteString(classLine)
	builder.WriteString(raceLine)
	builder.WriteString(backgroundLine)
	builder.WriteString(nl)

	// Feats
	featsLine			:= fmt.Sprintf("- Feats:\n")
	builder.WriteString(featsLine)
	for _, feat := range c.Feats {
		featRow := fmt.Sprintf(" - %s: %s\n", feat.Name, feat.Desc)
		builder.WriteString(featRow)
	}
	builder.WriteString(nl)

	// Languages
	languagesLine		:= fmt.Sprintf("- Languages:\n")
	builder.WriteString(languagesLine)
	for _, lang := range c.Languages {
		languageRow := fmt.Sprintf(" - %s\n", lang)
		builder.WriteString(languageRow)
	}

	builder.WriteString(nl)

	profBonusLine	:= fmt.Sprintf("Proficincy Bonus: +%d\n", c.Proficiency)
	passReception	:= fmt.Sprintf("Passive Reception: %d\n", c.PassiveReception)
	passInsight		:= fmt.Sprintf("Passive Insight: %d\n", c.PassiveInsight)
	builder.WriteString(profBonusLine)
	builder.WriteString(passReception)
	builder.WriteString(passInsight)
	builder.WriteString(nl)

	acLine 			:= fmt.Sprintf("AC: %d\n", c.AC)
	initiativeLine 	:= fmt.Sprintf("Initiative: %d\n", c.Initiative)
	speedLine 		:= fmt.Sprintf("Speed: %d\n", c.Speed)
	// hpMaxLine 		:= fmt.Sprintf("HP Max: %d", c.HPMax)
	hitDiceLine 	:= fmt.Sprintf("Hit Dice: %s\n", c.HitDice)
	builder.WriteString(acLine)
	builder.WriteString(initiativeLine)
	builder.WriteString(speedLine)
	// builder.WriteString(hpMaxLine)
	builder.WriteString(hitDiceLine)
	builder.WriteString(nl)

	// Proficiencies
	profHeader		:= fmt.Sprintf("*Proficiencies*\n")
	builder.WriteString(profHeader)
	builder.WriteString(nl)


	profTopRow 		:= fmt.Sprintf("| Proficiency  | Base  | Bonus | Saving Throws | ST Prof |\n") 
	profSpacer		:= fmt.Sprintf("| --- | --- | --- | --- | --- |\n")
	builder.WriteString(profTopRow)
	builder.WriteString(profSpacer)
	for _, prof := range c.Proficiencies {
		stProf := " "
		if prof.Proficient {
			stProf = "*"
		}

		profRow := fmt.Sprintf("| %s | %d | %s | %s |\n", prof.Name, prof.Adjusted, prof.Bonus, stProf)
		builder.WriteString(profRow)
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
