package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestDruidExecuteArchDruid(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		druid     *Druid
		expected  types.NamedToken
	}{
		{
			name: "Below level 20",
			character: &models.Character{
				Level: 16,
			},
			druid: &Druid{
				ClassToken: types.NamedToken{
					Available: 2,
					Maximum:   2,
				},
			},
			expected: types.NamedToken{
				Available: 2,
				Maximum:   2,
			},
		},
		{
			name: "Over level 20",
			character: &models.Character{
				Level: 21,
			},
			druid: &Druid{
				ClassToken: types.NamedToken{
					Available: 2,
					Maximum:   2,
				},
			},
			expected: types.NamedToken{
				Available: 0,
				Maximum:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.druid.executeArchDruid(tt.character)
			result := tt.druid.ClassToken

			if tt.expected.Maximum != result.Maximum {
				t.Errorf("Wild Shape Max- Expected: %d, Result: %d", tt.expected.Maximum, result.Maximum)
			}

			if tt.expected.Available != result.Available {
				t.Errorf("Wild Shape Avl- Expected: %d, Result: %d", tt.expected.Available, result.Available)
			}
		})
	}
}

func TestDruidValidateCantripVersatility(t *testing.T) {
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
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: types.AbilityDexterity, Bonus: 2},
					{Ability: types.AbilityStrength, Bonus: 2},
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
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: types.AbilityDexterity, Bonus: 2},
					{Ability: types.AbilityStrength, Bonus: 2},
					{Ability: types.AbilityWisdom, Bonus: 1},
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
			druid := Druid{}
			result := druid.validateCantripVersatility(tt.character)

			if tt.expected != result {
				t.Errorf("Cantrip Versatility- Expected: %t, result: %t", tt.expected, result)
			}
		})
	}
}

func TestDruidExecuteSpellCastingAbility(t *testing.T) {
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
			druid := &Druid{}
			druid.executeSpellCastingAbility(tt.character)

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

func TestDruidExecutePreparedSpells(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		druid     Druid
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
			druid: Druid{
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
			tt.druid.executePreparedSpells(tt.character)

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
