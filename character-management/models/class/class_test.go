package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestClassExecuteExpertiseShared(t *testing.T) {
	tests := []struct {
		name            string
		character       *models.Character
		expertiseSkills []string
		expected        []types.Skill
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
			expertiseSkills: []string{
				"persuasion",
				"deception",
			},
			expected: []types.Skill{
				{Name: "dexterity", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 4, Proficient: false},
				{Name: "deception", SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 3 - two skill proficiencies doubled",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			expertiseSkills: []string{
				"persuasion",
				"deception",
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 3 - two skill proficiencies doubled, one removed",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			expertiseSkills: []string{
				"persuasion",
				"deception",
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
			expertiseSkills: []string{
				"persuasion",
				"deception",
				"nature",
				"religion",
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
			executeExpertiseShared(tt.character, tt.expertiseSkills)

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

func TestClassExecuteSpellCastingAbility(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  models.Character
	}{
		{
			name: "Abiltity mod +2, proficiency +2",
			character: &models.Character{
				Level:       4,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				SpellSaveDC:    0,
				SpellAttackMod: 0,
			},
			expected: models.Character{
				Level:       4,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				SpellSaveDC:    12,
				SpellAttackMod: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wizard := &Wizard{}
			wizard.executeSpellCastingAbility(tt.character)

			expectedDC := tt.expected.SpellSaveDC
			expectedAttackMod := tt.expected.SpellAttackMod
			resultDC := tt.character.SpellSaveDC
			resultAttackMod := tt.character.SpellAttackMod

			if expectedDC != resultDC {
				t.Errorf("Spell Save DC- Expected: %d, Result: %d", expectedDC, resultDC)
			}

			if expectedAttackMod != resultAttackMod {
				t.Errorf("Spell Attack Mod- Expected: %d, Result: %d", expectedAttackMod, resultAttackMod)
			}
		})
	}
}

func TestClassExecutePreparedSpells(t *testing.T) {
	tests := []struct {
		name           string
		character      *models.Character
		preparedSpells []string
		expected       models.Character
	}{
		{
			name: "No Prepared Spells",
			character: &models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			preparedSpells: []string{},
			expected: models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
		},
		{
			name: "One Prepared Spell",
			character: &models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			preparedSpells: []string{
				"Some Spell",
			},
			expected: models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: true},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executePreparedSpellsShared(tt.character, tt.preparedSpells)

			for _, e := range tt.expected.Spells {
				for _, r := range tt.character.Spells {
					if e.Name == r.Name {
						if e.IsPrepared != r.IsPrepared {
							t.Errorf("Spell '%s' Is Prepared- Expected: %t, Result: %t", e.Name, e.IsPrepared, r.IsPrepared)
						}
					}
				}
			}
		})
	}
}
