package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestSorcererExecuteSpellCastingAbility(t *testing.T) {
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
					{Name: types.AbilityCharisma, AbilityModifier: 2},
				},
				SpellSaveDC:    0,
				SpellAttackMod: 0,
			},
			expected: models.Character{
				Level:       4,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityCharisma, AbilityModifier: 2},
				},
				SpellSaveDC:    12,
				SpellAttackMod: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sorcerer := &Sorcerer{}
			sorcerer.executeSpellCastingAbility(tt.character)

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
