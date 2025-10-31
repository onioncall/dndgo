package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestRogueExecuteExpertise(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		rogue     *Rogue
		expected  []types.Skill
	}{
		{
			name: "Level 1, two skill proficiencies doubled",
			character: &models.Character{
				Level:       1,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "dexterity", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			rogue: &Rogue{
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
			name: "Level 6, four skill proficiencies doubled",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			rogue: &Rogue{
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
			rogue: &Rogue{
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
			name: "Level 10, four skill proficiencies doubled, one removed",
			character: &models.Character{
				Level:       10,
				Proficiency: 4,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
					{Name: "religion", SkillModifier: 2, Proficient: false},
					{Name: "survival", SkillModifier: 4, Proficient: false},
					{Name: "perception", SkillModifier: 6, Proficient: false},
				},
			},
			rogue: &Rogue{
				ExpertiseSkills: []string{
					"persuasion",
					"deception",
					"nature",
					"religion",
					"perception",
				},
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 9, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 7, Proficient: false},
				{Name: "religion", SkillModifier: 6, Proficient: false},
				{Name: "survival", SkillModifier: 4, Proficient: false},
				{Name: "perception", SkillModifier: 6, Proficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.rogue.PostCalculateExpertise(tt.character)

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

func TestRogueExecuteSneakAttack(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  string
	}{
		{
			name: "Level 1, 1d6",
			character: &models.Character{
				Level: 1,
			},
			expected: "1d6",
		},
		{
			name: "Level 3, 2d6",
			character: &models.Character{
				Level: 1,
			},
			expected: "1d6",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rogue := &Rogue{}
			rogue.PostCalculateSneakAttack(tt.character)

			result := rogue.SneakAttack
			if tt.expected != result {
				t.Errorf("Sneak Attack- Expected: %s, Result: %s", tt.expected, result)
			}
		})
	}
}
