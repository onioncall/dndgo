package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestWarlockAppliedArmorOfShadows(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  int
		applied   bool
	}{
		{
			name: "Armor equiped, early return",
			character: &models.Character{
				AC: 15,
				WornEquipment: types.WornEquipment{
					Armor: "Light Armor",
				},
			},
			expected: 15,
			applied:  false,
		},
		{
			name: "Armor not equiped, bonus added",
			character: &models.Character{
				AC: 15,
				Abilities: []types.Abilities{
					{Name: types.AbilityDexterity, AbilityModifier: 4},
				},
				WornEquipment: types.WornEquipment{
					Armor: "",
				},
			},
			expected: 17,
			applied:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyArmorOfShadows(tt.character)
			result := tt.character.AC

			if tt.expected != result {
				t.Errorf("AC- Expected: %d, Result: %d", tt.expected, result)
			}

			if tt.applied != returned {
				t.Errorf("Not Applied Correctly- Expected: %t, Result: %t", tt.applied, returned)
			}
		})
	}
}
