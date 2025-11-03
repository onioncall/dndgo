package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestClericValidateCantripVersatility(t *testing.T) {
	tests := []struct {
		name      string // description of this test case
		character *models.Character
		expected  bool
	}{
		{
			name: "Below level 4, valid",
			character: &models.Character{
				Level:              3,
				ValidationDisabled: false,
				Spells: []types.CharacterSpell{
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{},
			},
			expected: true,
		},
		{
			name: "Below level 12, valid",
			character: &models.Character{
				Level:              11,
				ValidationDisabled: false,
				Spells: []types.CharacterSpell{
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: types.AbilityDexterity, Bonus: 2},
					{Ability: types.AbilityStrength, Bonus: 2},
					{Ability: types.AbilityStrength, Bonus: 1},
				},
			},
			expected: true,
		},
		{
			name: "Below level 12, invalid",
			character: &models.Character{
				ValidationDisabled: false,
				Level:              11,
				Spells: []types.CharacterSpell{
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: types.AbilityDexterity, Bonus: 2},
					{Ability: types.AbilityStrength, Bonus: 2},
					{Ability: types.AbilityWisdom, Bonus: 2},
				},
			},
			expected: false,
		},
		{
			name: "Below level 12, validation disabled",
			character: &models.Character{
				ValidationDisabled: true,
				Level:              11,
				Spells: []types.CharacterSpell{
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: true},
					{SlotLevel: 0, IsRitual: false},
					{SlotLevel: 0, IsRitual: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: types.AbilityDexterity, Bonus: 2},
					{Ability: types.AbilityStrength, Bonus: 2},
					{Ability: types.AbilityWisdom, Bonus: 1},
				},
			},
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleric := Cleric{}
			result := cleric.validateCantripVersatility(tt.character)

			if tt.expected != result {
				t.Errorf("Cantrip Versatility- Expected: %t, result: %t", tt.expected, result)
			}
		})
	}
}

func TestClericExecuteSpellCastingAbility(t *testing.T) {
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
					{Name: types.AbilityWisdom, AbilityModifier: 2},
				},
				SpellSaveDC:    0,
				SpellAttackMod: 0,
			},
			expected: models.Character{
				Level:       4,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityWisdom, AbilityModifier: 2},
				},
				SpellSaveDC:    12,
				SpellAttackMod: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cleric := &Cleric{}
			cleric.PostCalculateSpellCastingAbility(tt.character)

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

func TestClericExecutePreparedSpells(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		cleric    Cleric
		expected  models.Character
	}{
		{
			name: "One Prepared Spell",
			character: &models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityWisdom, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			cleric: Cleric{
				PreparedSpells: []string{
					"Some Spell",
				},
			},
			expected: models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityWisdom, AbilityModifier: 2},
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
			tt.cleric.PostCalculatePreparedSpells(tt.character)

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
