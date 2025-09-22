package class

import (
	"testing"

	"github.com/onioncall/dndgo/models"
)

func TestBard_expertise(t *testing.T) {
	tests := []struct {
		name 			string
		character 		*models.Character
		bard			*Bard
		expectedSKills 	[]models.Skill
	}{
		{
			name: "Below Level Requirement",
			character: &models.Character {
				Level: 2,
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
			expectedSKills: []models.Skill {
				{Name: "dexterity", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 4, Proficient: false},
				{Name: "deception", SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 3 - Two Skill Proficiencies Doubled",
			character: &models.Character {
				Level: 3,
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
			expectedSKills: []models.Skill {
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 6, Proficient: false},
			},
		},
		{
			name: "Level 3 - Two Skill Proficiencies Doubled, One Removed",
			character: &models.Character {
				Level: 3,
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
			expectedSKills: []models.Skill {
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 6, Proficient: false},
			},
		},
		{
			name: "Level 10 - Four Skill Proficiencies Doubled",
			character: &models.Character {
				Level: 10,
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
			expectedSKills: []models.Skill {
				{Name: "nature", SkillModifier: 10, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 6, Proficient: false},
				{Name: "religion", SkillModifier: 4, Proficient: false},
				{Name: "survival", SkillModifier: 4, Proficient: false},
			},
		},
	}

	for  _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			// Create a deep copy to avoid modifying the original test data
			testChar := &models.Character{
				Level:       test.character.Level,
				Proficiency: test.character.Proficiency,
				Skills:      make([]models.Skill, len(test.character.Skills)),
			}
			copy(testChar.Skills, test.character.Skills)

			test.bard.expertise(testChar)

			for i, skill := range testChar.Skills {
				expected := test.expectedSKills[i]

				if skill.SkillModifier != expected.SkillModifier {
					t.Errorf("Skill[%d].SkillModifier = %d, expected %d", i, skill.SkillModifier, expected.SkillModifier)
				}
				if skill.Proficient != expected.Proficient {
					t.Errorf("Skill[%d].Proficient = %v, expected %v", i, skill.Proficient, expected.Proficient)
				}
			}
		})
	}
}

func TestBard_jackOfAllTrades(t *testing.T) {
	tests := []struct {
		name           string
		character      *models.Character
		expectedSkills []models.Skill
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
			expectedSkills: []models.Skill {
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
			expectedSkills: []models.Skill {
				{SkillModifier: 6, Proficient: false},
				{SkillModifier: 4, Proficient: false},
				{SkillModifier: 2, Proficient: false},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			bard := &Bard{}
			
			// Create a deep copy to avoid modifying the original test data
			testChar := &models.Character {
				Level:       test.character.Level,
				Proficiency: test.character.Proficiency,
				Skills:      make([]models.Skill, len(test.character.Skills)),
			}
			copy(testChar.Skills, test.character.Skills)

			bard.jackOfAllTrades(testChar)

			if len(testChar.Skills) != len(test.expectedSkills) {
				t.Errorf("Expected %d skills, got %d", len(test.expectedSkills), len(testChar.Skills))
				return
			}

			for i, skill := range testChar.Skills {
				expected := test.expectedSkills[i]
				if skill.SkillModifier != expected.SkillModifier {
					t.Errorf("Skill[%d].SkillModifier = %d, expected %d", i, skill.SkillModifier, expected.SkillModifier)
				}
				if skill.Proficient != expected.Proficient {
					t.Errorf("Skill[%d].Proficient = %v, expected %v", i, skill.Proficient, expected.Proficient)
				}
			}
		})
	}
}
