package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestBardExecuteExpertise(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		bard      *Bard
		expected  []types.Skill
	}{
		{
			name: "Below level requirement",
			character: &models.Character{
				Level:       2,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "dexterity", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard{
				ExpertiseSkills: []string{
					"persuasion",
					"deception",
				},
			},
			expected: []types.Skill{
				{Name: "dexterity", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 4, Proficient: false},
				{Name: "deception", SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 3, two skill proficiencies doubled",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard{
				ExpertiseSkills: []string{
					"persuasion",
					"deception",
				},
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 3, two skill proficiencies doubled, one removed",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard{
				ExpertiseSkills: []string{
					"persuasion",
					"deception",
					"nature",
				},
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 10, four skill proficiencies doubled",
			character: &models.Character{
				Level:       10,
				Proficiency: 4,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
					{Name: "religion", SkillModifier: 2, Proficient: false},
					{Name: "survival", SkillModifier: 4, Proficient: false},
				},
			},
			bard: &Bard{
				ExpertiseSkills: []string{
					"persuasion",
					"deception",
					"nature",
					"religion",
				},
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 9, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 7, Proficient: false},
				{Name: "religion", SkillModifier: 6, Proficient: false},
				{Name: "survival", SkillModifier: 4, Proficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.PostCalculateExpertise(tt.character)

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

func TestBardExecuteJackOfAllTrades(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		bard      *Bard
		expected  []types.Skill
	}{
		{
			name: "Level 1 character - no bonus applied",
			character: &models.Character{
				Level:       1,
				Proficiency: 2,
				Skills: []types.Skill{
					{SkillModifier: 5, Proficient: false},
					{SkillModifier: 3, Proficient: false},
				},
			},
			bard: &Bard{},
			expected: []types.Skill{
				{SkillModifier: 5, Proficient: false},
				{SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 2 character with non-proficient skills - bonus applied",
			character: &models.Character{
				Level:       2,
				Proficiency: 2,
				Skills: []types.Skill{
					{SkillModifier: 5, Proficient: false},
					{SkillModifier: 3, Proficient: false},
					{SkillModifier: 1, Proficient: false},
				},
			},
			bard: &Bard{},
			expected: []types.Skill{
				{SkillModifier: 6, Proficient: false},
				{SkillModifier: 4, Proficient: false},
				{SkillModifier: 2, Proficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.PostCalculateJackOfAllTrades(tt.character)
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

func TestBardUseClassSlots(t *testing.T) {
	tests := []struct {
		name      string
		tokenName string
		character *models.Character
		bard      *Bard
		expected  BardicInspiration
	}{
		{
			name:      "One use, single slot",
			tokenName: "bardic inspiration",
			character: &models.Character{},
			bard: &Bard{
				BardicInspiration: BardicInspiration{
					Maximum:   4,
					Available: 4,
				},
			},
			expected: BardicInspiration{
				Maximum:   4,
				Available: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.UseClassTokens(tt.tokenName)

			result := tt.bard.BardicInspiration.Available
			e := tt.expected.Available

			if e != result {
				t.Errorf("Bardic Inspiration- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}

func TestBardRecoverClassSlots(t *testing.T) {
	tests := []struct {
		name      string
		tokenName string
		recover   int
		character *models.Character
		bard      *Bard
		expected  BardicInspiration
	}{
		{
			name:      "Recover by 1",
			tokenName: "bardic inspiration",
			recover:   1,
			bard: &Bard{
				BardicInspiration: BardicInspiration{
					Maximum:   4,
					Available: 2,
				},
			},
			expected: BardicInspiration{
				Maximum:   4,
				Available: 3,
			},
		},
		{
			name:      "Full recover",
			tokenName: "bardic inspiration",
			recover:   0,
			bard: &Bard{
				BardicInspiration: BardicInspiration{
					Maximum:   4,
					Available: 2,
				},
			},
			expected: BardicInspiration{
				Maximum:   4,
				Available: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.bard.RecoverClassTokens(tt.tokenName, tt.recover)

			result := tt.bard.BardicInspiration.Available
			e := tt.expected.Available

			if e != result {
				t.Errorf("Bardic Inspiration- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}
