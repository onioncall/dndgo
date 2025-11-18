package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
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
				Abilities: []shared.Abilities{
					{
						Name:            shared.AbilityStrength,
						AbilityModifier: 5,
					},
					{
						Name:            shared.AbilityDexterity,
						AbilityModifier: 3,
					},
					{
						Name:            shared.AbilityConstitution,
						AbilityModifier: 4,
					},
					{
						Name:            shared.AbilityIntelligence,
						AbilityModifier: 2,
					},
					{
						Name:            shared.AbilityWisdom,
						AbilityModifier: 0,
					},
					{
						Name:            shared.AbilityCharisma,
						AbilityModifier: -1,
					},
				},
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "Leather Armor",
					},
				},
			},
			expected: 0,
		},
		{
			name: "No armor, valid",
			character: &models.Character{
				AC: 0,
				Abilities: []shared.Abilities{
					{
						Name:            shared.AbilityStrength,
						AbilityModifier: 1,
					},
					{
						Name:            shared.AbilityDexterity,
						AbilityModifier: 3,
					},
					{
						Name:            shared.AbilityConstitution,
						AbilityModifier: 2,
					},
					{
						Name:            shared.AbilityIntelligence,
						AbilityModifier: 2,
					},
					{
						Name:            shared.AbilityWisdom,
						AbilityModifier: 3,
					},
					{
						Name:            shared.AbilityCharisma,
						AbilityModifier: 0,
					},
				},
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "",
					},
				},
			},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.executeUnarmoredDefense(tt.character)
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
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "Leather Armor",
					},
				},
			},
			expected: 0,
		},
		{
			name: "No armor, valid",
			character: &models.Character{
				Speed: 16,
				Level: 3,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "",
					},
				},
			},
			expected: 26,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}

			monk.executeUnarmoredMovement(tt.character)
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

			monk.executeMartialArts(tt.character)
			result := monk.MartialArts

			if tt.expected != result {
				t.Errorf("Martial Arts- Expected: %s, Result: %s", tt.expected, result)
			}
		})
	}
}

func TestMonkExecuteKiPoints(t *testing.T) {
	tests := []struct {
		name       string
		character  *models.Character
		monk       *Monk
		expected   shared.NamedToken
		expectedDC int
	}{
		{
			name: "Below level 2",
			character: &models.Character{
				Level:       1,
				Proficiency: 2,
				Abilities: []shared.Abilities{
					{Name: shared.AbilityDexterity, AbilityModifier: 4},
				},
			},
			monk: &Monk{
				ClassToken: shared.NamedToken{
					Name: "ki-points",
				},
			},
			expected: shared.NamedToken{
				Available: 0,
				Maximum:   0,
			},
		},
		{
			name: "Level 4",
			character: &models.Character{
				Level:       4,
				Proficiency: 3,
				Abilities: []shared.Abilities{
					{Name: shared.AbilityWisdom, AbilityModifier: 4},
				},
			},
			monk: &Monk{
				ClassToken: shared.NamedToken{
					Name: "ki-points",
				},
			},
			expected: shared.NamedToken{
				Available: 4,
				Maximum:   4,
				Name:      "ki-points",
			},
			expectedDC: 15,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tt.monk.executeKiPoints(tt.character)

			maxPointResult := tt.monk.ClassToken.Maximum
			spellSaveDCResult := tt.monk.KiSpellSaveDC
			expectedMax := tt.expected.Maximum
			expectedDC := tt.expectedDC

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
				Abilities: []shared.Abilities{
					{Name: shared.AbilityDexterity, AbilityModifier: 4},
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

			monk.executeDeflectMissles(tt.character)
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
		expected  []shared.Abilities
	}{
		{
			name: "Below level 14",
			character: &models.Character{
				Level: 10,
				Abilities: []shared.Abilities{
					{Name: shared.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
					{Name: shared.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: shared.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
					{Name: shared.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: shared.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: shared.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				},
			},
			expected: []shared.Abilities{
				{Name: shared.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
				{Name: shared.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				{Name: shared.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
				{Name: shared.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				{Name: shared.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				{Name: shared.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 14",
			character: &models.Character{
				Level: 14,
				Abilities: []shared.Abilities{
					{Name: shared.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
					{Name: shared.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: shared.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
					{Name: shared.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: shared.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: shared.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
				},
			},
			expected: []shared.Abilities{
				{Name: shared.AbilityStrength, AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
				{Name: shared.AbilityDexterity, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
				{Name: shared.AbilityConstitution, AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
				{Name: shared.AbilityWisdom, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
				{Name: shared.AbilityIntelligence, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
				{Name: shared.AbilityCharisma, AbilityModifier: 0, Base: 12, SavingThrowsProficient: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			monk := &Monk{}
			monk.executeDiamondSoul(tt.character)

			for i, e := range tt.expected {
				result := tt.character.Abilities[i].SavingThrowsProficient
				if e.SavingThrowsProficient != result {
					t.Errorf("Ability %s- Expected: %t, Result: %t", e.Name, e.SavingThrowsProficient, result)
				}
			}
		})
	}
}
