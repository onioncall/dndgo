package handlers
import (
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
		if err != nil {
			return nil, fmt.Errorf("failed to load bard class: %w", err)
		}
	}

	return c, nil
}
