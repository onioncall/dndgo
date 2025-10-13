package class

import (
	"os"
	"testing"

	"github.com/onioncall/dndgo/models"
)

func TestBard_ExecuteExpertise(t *testing.T) {
	tests := []struct {
		name 			string
		character 		*models.Character
		bard			*Bard
		expected 		[]models.Skill
	}{
		{
			name: "Below Level Requirement",
			character: &models.Character {
				Level: 2,
				Proficiency: 2,
				Skills: []models.Skill {
					{Name: "dexterity", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard {
				SkillProficienciesToDouble: []string {
					"persuasion",
					"deception",
				},
			},
			expected: []models.Skill {
				{Name: "dexterity", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 4, Proficient: false},
				{Name: "deception", SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 3 - Two Skill Proficiencies Doubled",
			character: &models.Character {
				Level: 3,
				Proficiency: 2,
				Skills: []models.Skill {
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard {
				SkillProficienciesToDouble: []string {
					"persuasion",
					"deception",
				},
			},
			expected: []models.Skill {
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 3 - Two Skill Proficiencies Doubled, One Removed",
			character: &models.Character {
				Level: 3,
				Proficiency: 2,
				Skills: []models.Skill {
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard {
				SkillProficienciesToDouble: []string {
					"persuasion",
					"deception",
					"nature",
				},
			},
			expected: []models.Skill {
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 10 - Four Skill Proficiencies Doubled",
			character: &models.Character {
				Level: 10,
				Proficiency: 4,
				Skills: []models.Skill {
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
					{Name: "religion", SkillModifier: 2, Proficient: false},
					{Name: "survival", SkillModifier: 4, Proficient: false},
				},
			},
			bard: &Bard {
				SkillProficienciesToDouble: []string {
					"persuasion",
					"deception",
					"nature",
					"religion",
				},
			},
			expected: []models.Skill {
				{Name: "nature", SkillModifier: 9, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 7, Proficient: false},
				{Name: "religion", SkillModifier: 6, Proficient: false},
				{Name: "survival", SkillModifier: 4, Proficient: false},
			},
		},
	}

	for  _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.executeExpertise(tt.character)

			if len(tt.character.Skills) != len(tt.expected) {
				t.Errorf("Skills Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Skills))
				return
			}

			for i, e := range tt.expected {
				result := tt.character.Skills[i]

				if e.SkillModifier != result.SkillModifier {
					t.Errorf("Skill Modifier %s- Expected: %d, Result %d", e.Name, e.SkillModifier, result.SkillModifier)
				}
				if e.Proficient != result.Proficient {
					t.Errorf("Skill Proficient %s- Expected: %t, Result %t", e.Name, e.Proficient, result.Proficient)
				}
			}
		})
	}
}

func TestBard_ExecuteJackOfAllTrades(t *testing.T) {
	tests := []struct {
		name        string
		character   *models.Character
		bard		*Bard
		expected 	[]models.Skill
	}{
		{
			name: "Level 1 character - no bonus applied",
			character: &models.Character {
				Level:       1,
				Proficiency: 2,
				Skills: []models.Skill {
					{SkillModifier: 5, Proficient: false},
					{SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard{},
			expected: []models.Skill {
				{SkillModifier: 5, Proficient: false},
				{SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 2 character with non-proficient skills - bonus applied",
			character: &models.Character {
				Level:       2,
				Proficiency: 2,
				Skills: []models.Skill {
					{SkillModifier: 5, Proficient: false},
					{SkillModifier: 3, Proficient: false},
					{SkillModifier: 1, Proficient: false},
				},
			},
			bard: &Bard{},
			expected: []models.Skill {
				{SkillModifier: 6, Proficient: false},
				{SkillModifier: 4, Proficient: false},
				{SkillModifier: 2, Proficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.executeJackOfAllTrades(tt.character)
			result := tt.character

			if len(result.Skills) != len(tt.expected) {
				t.Errorf("Skills Count- Expected: %d, Result: %d", len(tt.expected), len(result.Skills))
				return
			}

			for i, e := range tt.expected {
				result := tt.expected[i]
				if e.SkillModifier != result.SkillModifier {
					t.Errorf("Skill Modifier %s- Expected: %d, Result %d", e.Name, e.SkillModifier, result.SkillModifier)
				}
				if e.Proficient != result.Proficient {
					t.Errorf("Skill Proficient %s- Expected: %t, Result %t", e.Name, e.Proficient, result.Proficient)
				}
			}
		})
	}
}

func TestBard_UseClassSlots(t *testing.T) {
	tests := []struct {
		name		string
		slotName 	string
		character 	*models.Character
		bard		*Bard
		expected	BardicInspiration
	}{
		{
			name: "One Use, Single Slot",
			slotName: "bardic inspiration",
			character: &models.Character{},
			bard: &Bard {
				BardicInspiration: BardicInspiration {
					Slot: 4,
					Available: 4,
				},
			},
			expected: BardicInspiration {
				Slot: 4,
				Available: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prevent from writing to terminal during tests
			original := os.Stdout
			os.Stdout, _ = os.Open(os.DevNull)
			defer func() { os.Stdout = original }()

			tt.bard.UseClassSlots(tt.slotName)	

			result := tt.bard.BardicInspiration.Available
			e := tt.expected.Available
			
			if e != result {
				t.Errorf("Bardic Inspiration- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}

func TestBard_RecoverClassSlots(t *testing.T) {
	tests := []struct {
		name		string
		slotName 	string
		recover 	int
		character 	*models.Character
		bard		*Bard
		expected	BardicInspiration
	}{
		{
			name: "Recover By 1",
			slotName: "bardic inspiration",
			recover: 1,
			bard: &Bard {
				BardicInspiration: BardicInspiration {
					Slot: 4,
					Available: 2,
				},
			},
			expected: BardicInspiration {
				Slot: 4,
				Available: 3,
			},
		},
		{
			name: "Full Recover",
			slotName: "bardic inspiration",
			recover: 0,
			bard: &Bard {
				BardicInspiration: BardicInspiration {
					Slot: 4,
					Available: 2,
				},
			},
			expected: BardicInspiration {
				Slot: 4,
				Available: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.RecoverClassSlots(tt.slotName, tt.recover)

			result := tt.bard.BardicInspiration.Available
			e := tt.expected.Available

			if e != result {
				t.Errorf("Bardic Inspiration- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}
