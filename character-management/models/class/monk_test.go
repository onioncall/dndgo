package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestMonkExecuteUnarmoredDefense(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  int
	}{
		{
			name: "Armor equiped, early return",
			character: &models.Character{
				AC: 0,
				Abilities: []types.Abilities{
					{
						Name:            types.AbilityStrength,
						AbilityModifier: 5,
					},
					{
						Name:            types.AbilityDexterity,
						AbilityModifier: 3,
					},
					{
						Name:            types.AbilityConstitution,
						AbilityModifier: 4,
					},
					{
						Name:            types.AbilityIntelligence,
						AbilityModifier: 2,
					},
					{
						Name:            types.AbilityWisdom,
						AbilityModifier: 0,
					},
					{
						Name:            types.AbilityCharisma,
						AbilityModifier: -1,
					},
				},
				WornEquipment: types.WornEquipment{
					Armour: "Leather Armor",
				},
			},
			expected: 0,
		},
		{
			name: "No armor, valid",
			character: &models.Character{
				AC: 0,
				Abilities: []types.Abilities{
					{
						Name:            types.AbilityStrength,
						AbilityModifier: 1,
					},
					{
						Name:            types.AbilityDexterity,
						AbilityModifier: 3,
					},
					{
						Name:            types.AbilityConstitution,
						AbilityModifier: 2,
					},
					{
						Name:            types.AbilityIntelligence,
						AbilityModifier: 2,
					},
					{
						Name:            types.AbilityWisdom,
						AbilityModifier: 3,
					},
					{
						Name:            types.AbilityCharisma,
						AbilityModifier: 0,
					},
				},
				WornEquipment: types.WornEquipment{
					Armour: "",
				},
			},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.PostCalculateUnarmoredDefense(tt.character)
			result := tt.character.AC

			if tt.expected != result {
				t.Errorf("AC- Expected: %d, Result: %d", tt.expected, result)
			}
		})
	}
}

func TestMonkExecuteUnarmoredMovement(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  int
	}{
		{
			name: "Armor equiped, early return",
			character: &models.Character{
				Speed: 0,
				Level: 3,
				WornEquipment: types.WornEquipment{
					Armour: "Leather Armor",
				},
			},
			expected: 0,
		},
		{
			name: "No armor, valid",
			character: &models.Character{
				Speed: 16,
				Level: 3,
				WornEquipment: types.WornEquipment{
					Armour: "",
				},
			},
			expected: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.PostCalculateUnarmoredMovement(tt.character)
			result := tt.character.Speed

			if tt.expected != result {
				t.Errorf("Speed- Expected: %d, Result: %d", tt.expected, result)
			}
		})
	}
}

func TestMonkExecuteMartialArts(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  string
	}{
		{
			name: "Below level 5",
			character: &models.Character{
				Level: 4,
			},
			expected: "1d4",
		},
		{
			name: "Below level 17",
			character: &models.Character{
				Level: 15,
			},
			expected: "1d8",
		},
		{
			name: "Above level 17",
			character: &models.Character{
				Level: 20,
			},
			expected: "1d10",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.PostCalculateMartialArts(tt.character)
			result := monk.MartialArts

			if tt.expected != result {
				t.Errorf("Martial Arts- Expected: %s, Result: %s", tt.expected, result)
			}
		})
	}
}

func TestMonkExecuteKiPoints(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  Ki
	}{
		{
			name: "Below level 2",
			character: &models.Character{
				Level:       1,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityDexterity, AbilityModifier: 4},
				},
			},
			expected: Ki{
				Available:     0,
				Maximum:       0,
				KiSpellSaveDC: 0,
			},
		},
		{
			name: "Level 4",
			character: &models.Character{
				Level:       4,
				Proficiency: 3,
				Abilities: []types.Abilities{
					{Name: types.AbilityWisdom, AbilityModifier: 4},
				},
			},
			expected: Ki{
				Available:     4,
				Maximum:       4,
				KiSpellSaveDC: 15,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.PostCalculateKiPoints(tt.character)

			maxPointResult := monk.KiPoints.Maximum
			spellSaveDCResult := monk.KiPoints.KiSpellSaveDC
			expectedMax := tt.expected.Maximum
			expectedDC := tt.expected.KiSpellSaveDC

			if expectedMax != maxPointResult {
				t.Errorf("Ki Point Max- Expected: %d, Result: %d", expectedMax, maxPointResult)
			}

			if expectedDC != spellSaveDCResult {
				t.Errorf("Ki Spell Save DC- Expected: %d, Result: %d", expectedDC, spellSaveDCResult)
			}
		})
	}
}

func TestMonkExecuteDeflectMissles(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  int
	}{
		{
			name: "Below level 3",
			character: &models.Character{
				Level:       2,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityDexterity, AbilityModifier: 4},
				},
			},
			expected: 0,
		},
		{
			name: "Level 4",
			character: &models.Character{
				Level:       4,
				Proficiency: 3,
			},
			expected: -17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.PostCalculateDeflectMissles(tt.character)
			result := monk.DeflectMissles

			if tt.expected != result {
				t.Errorf("Deflect Missles- Expected: %d, Result: %d", tt.expected, result)
			}
		})
	}
}

func TestMonkExecuteDiamond(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  []types.Abilities
	}{
		{
			name: "Below level 14",
			character: &models.Character{
				Level: 10,
				Abilities: []types.Abilities{
					{Name: types.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
					{Name: types.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: types.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
					{Name: types.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: types.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: types.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				},
			},
			expected: []types.Abilities{
				{Name: types.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
				{Name: types.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				{Name: types.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
				{Name: types.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				{Name: types.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				{Name: types.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 14",
			character: &models.Character{
				Level: 14,
				Abilities: []types.Abilities{
					{Name: types.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
					{Name: types.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: types.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
					{Name: types.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: types.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: types.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				},
			},
			expected: []types.Abilities{
				{Name: types.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
				{Name: types.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
				{Name: types.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
				{Name: types.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
				{Name: types.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
				{Name: types.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}
			monk.PreCalculateDiamondSoul(tt.character)

			for i, e := range tt.expected {
				result := tt.character.Abilities[i].SavingThrowsProficient
				if e.SavingThrowsProficient != result {
					t.Errorf("Ability %s- Expected: %t, Result: %t", e.Name, e.SavingThrowsProficient, result)
				}
			}
		})
	}
}
