package manage

import (
	"testing"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/manage/info"
)

func TestExecuteUserCmd_Rename(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		initialName  string
		expectedName string
		expectError  bool
	}{
		{
			name:         "valid simple name",
			input:        "rename Aragorn",
			initialName:  "Nim",
			expectedName: "Aragorn",
			expectError:  false,
		},
		{
			name:         "valid name with spaces",
			input:        "rename Gandalf the Grey",
			initialName:  "Nim",
			expectedName: "Gandalf the Grey",
			expectError:  false,
		},
		{
			name:         "empty name",
			input:        "rename ",
			initialName:  "Nim",
			expectedName: "Nim", // Should not change
			expectError:  true,
		},
		{
			name:         "rename with only whitespace",
			input:        "rename",
			initialName:  "Nim",
			expectedName: "Nim", // Should not change
			expectError:  true,
		},
		{
			name:         "special characters and apostrophes",
			input:        "rename Drizzt Do'Urden",
			initialName:  "Nim",
			expectedName: "Drizzt Do'Urden",
			expectError:  false,
		},
		{
			name:         "unicode characters",
			input:        "rename Søren Kierkegaard",
			initialName:  "Nim",
			expectedName: "Søren Kierkegaard",
			expectError:  false,
		},
		{
			name:         "very long name",
			input:        "rename Taumatawhakatangihangakoauauotamateaturipukakapikimaungahoronukupokaiwhenuakitanatahu",
			initialName:  "Nim",
			expectedName: "Taumatawhakatangihangakoauauotamateaturipukakapikimaungahoronukupokaiwhenuakitanatahu",
			expectError:  false,
		},
		{
			name:         "name with numbers",
			input:        "rename Agent 47",
			initialName:  "Nim",
			expectedName: "Agent 47",
			expectError:  false,
		},
		{
			name:         "case insensitive command",
			input:        "RENAME Legolas",
			initialName:  "Nim",
			expectedName: "Legolas",
			expectError:  false,
		},
		{
			name:         "mixed case command",
			input:        "ReNaMe Gimli",
			initialName:  "Nim",
			expectedName: "Gimli",
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup model with test character
			character := &models.Character{
				Name:      tt.initialName,
				ClassName: "Bard",
				Level:     3,
				Race:      "Human",
			}

			// Initialize basic info tab with viewports
			basicInfoTab := info.BasicInfoModel{
				BasicStatsViewport: viewport.New(80, 20),
				AbilitiesViewport:  viewport.New(80, 20),
				HealthViewport:     viewport.New(80, 20),
				SkillsViewport:     viewport.New(80, 20),
			}

			m := Model{
				character:    character,
				basicInfoTab: basicInfoTab,
			}

			// Execute command
			m, _, _ = m.executeUserCmd(tt.input, 0)

			// Check name was updated correctly
			if m.character.Name != tt.expectedName {
				t.Errorf("expected name %q, got %q", tt.expectedName, m.character.Name)
			}

			// Check error state
			if tt.expectError && m.err == nil {
				t.Error("expected error but got none")
			}
			if !tt.expectError && m.err != nil {
				t.Errorf("unexpected error: %v", m.err)
			}
		})
	}
}
