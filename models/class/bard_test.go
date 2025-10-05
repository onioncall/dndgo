package class

import (
	"os"
	"testing"

	"github.com/onioncall/dndgo/models"
)

func TestBard_expertise(t *testing.T) {
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
			tt.bard.expertise(tt.character)

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

func TestBard_jackOfAllTrades(t *testing.T) {
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
			tt.bard.jackOfAllTrades(tt.character)
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

func TestBard_abilityScoreImprovement(t *testing.T) {
	tests := []struct {
		name 		string
		bard		*Bard
		character 	*models.Character
		expected	[]models.Attribute
	}{
		{
			name: "Level not high enough",
			bard: &Bard {
				AbilityScoreImprovement: []AbilityScoreImprovementItem {
					{Ability: "Strength", Bonus: 2},
				},
			},
			character: &models.Character {
				Level: 3,
				Attributes: []models.Attribute {
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
			},
			expected: []models.Attribute {
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, one ability increased by two",
			bard: &Bard {
				AbilityScoreImprovement: []AbilityScoreImprovementItem {
					{Ability: "Dexterity", Bonus: 2},
				},
			},
			character: &models.Character {
				Level: 4,
				Attributes: []models.Attribute {
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
			},
			expected: []models.Attribute {
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, two abilities increased by one",
			bard: &Bard {
				AbilityScoreImprovement: []AbilityScoreImprovementItem {
					{Ability: "Dexterity", Bonus: 1},
					{Ability: "Charisma", Bonus: 1},
				},
			},
			character: &models.Character {
				Level: 4,
				Attributes: []models.Attribute {
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
			},
			expected: []models.Attribute {
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 11, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 11, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, two abilities increased by two (failure)",
			bard: &Bard {
				AbilityScoreImprovement: []AbilityScoreImprovementItem {
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Charisma", Bonus: 2},
				},
			},
			character: &models.Character {
				Level: 4,
				Attributes: []models.Attribute {
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
			},
			expected: []models.Attribute {
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 8, one ability increased by two, and two abilities increased by one",
			bard: &Bard {
				AbilityScoreImprovement: []AbilityScoreImprovementItem {
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Charisma", Bonus: 1},
					{Ability: "Wisdom", Bonus: 1},
				},
			},
			character: &models.Character {
				Level: 8,
				Attributes: []models.Attribute {
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
			},
			expected: []models.Attribute {
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 11, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 11, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 20, one ability over maximum",
			bard: &Bard {
				AbilityScoreImprovement: []AbilityScoreImprovementItem {
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
				},
			},
			character: &models.Character {
				Level: 20,
				Attributes: []models.Attribute {
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 12, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
			},
			expected: []models.Attribute {
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 20, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prevent from writing to terminal during tests
			original := os.Stdout
			os.Stdout, _ = os.Open(os.DevNull)
			defer func() { os.Stdout = original }()

			tt.bard.abilityScoreImprovement(tt.character)

			for i, e := range tt.expected {
				result := tt.character.Attributes[i]
				
				if e.Base != result.Base {
					t.Errorf("Attribute Base %s- Expected: %d, Result: %d", e.Name, e.Base, result.Base)
				}
			}
		})
	}
}
