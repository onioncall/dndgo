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
		expected  WildShape
	}{
		{
			name: "Below level 20",
			character: &models.Character{
				Level: 16,
			},
			druid: &Druid{
				WildShape: WildShape{
					Available: 2,
					Maximum:   2,
				},
			},
			expected: WildShape{
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
				WildShape: WildShape{
					Available: 2,
					Maximum:   2,
				},
			},
			expected: WildShape{
				Available: 0,
				Maximum:   0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.druid.executeArchDruid(tt.character)
			result := tt.druid.WildShape

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
				Level:             3,
				ValidationEnabled: true,
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
				Level:             11,
				ValidationEnabled: true,
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
				ValidationEnabled: true,
				Level:             11,
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
				ValidationEnabled: false,
				Level:             11,
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
