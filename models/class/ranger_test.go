package class

import (
	"testing"

	"github.com/onioncall/dndgo/models"
	attr "github.com/onioncall/dndgo/types/attributes"
	eqmt "github.com/onioncall/dndgo/types/equipment"
)

func TestRangerAppliedArchery(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	[]eqmt.Weapon
		applied		bool
	}{
		{
			name: "No ranged weapon",
			character: &models.Character {
				Weapons: []eqmt.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Dagger", Bonus: 2, Damage: "1d4", Range: "melee"},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				{Name: "Dagger", Bonus: 2, Damage: "1d4", Range: "melee"},
			},
			applied: false,
		},
		{
			name: "Range bonus applied",
			character: &models.Character {
				Weapons: []eqmt.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged"},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				{Name: "Longbow", Bonus: 4, Damage: "1d8", Range: "ranged"},
			},
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyArchery(tt.character)
			result := tt.character.Weapons

			for i, e := range tt.expected {
				if e.Bonus != result[i].Bonus {
					t.Errorf("Weapon %s Bonus- Expected: %d, Result: %d", e.Name, e.Bonus, result[i].Bonus)
				}
			}

			if tt.applied != returned {
				t.Errorf("Not Applied Correctly- Expected: %t, Result: %t", tt.applied, returned)
			}
		})
	}
}

func TestRangerAppliedDefense(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	int
		applied		bool
	}{
		{
			name: "Armor equiped, early return",
			character: &models.Character {
				AC: 15,
				BodyEquipment: eqmt.BodyEquipment {
					Armour: "Light Armor",
				},
			},
			expected: 15,
			applied: false,
		},
		{
			name: "Armor not equiped, bonus added",
			character: &models.Character {
				AC: 15,
				BodyEquipment: eqmt.BodyEquipment {
					Armour: "",
				},
			},
			expected: 16,
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyDefense(tt.character)
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

func TestRangerAppliedDueling(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	[]eqmt.Weapon
		applied		bool
	}{
		{
			name: "No melee weapon",
			character: &models.Character {
				Weapons: []eqmt.Weapon {
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged"},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged"},
			},
			applied: false,
		},
		{
			name: "Melee bonus applied",
			character: &models.Character {
				Weapons: []eqmt.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Range: "melee"},
			},
			applied: true,
		},
		{
			name: "Multiple valid weapons, one bonus",
			character: &models.Character {
				Weapons: []eqmt.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Range: "melee"},
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
			},
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyDueling(tt.character)
			result := tt.character.Weapons

			for i, e := range tt.expected {
				if e.Bonus != result[i].Bonus {
					t.Errorf("Weapon %s Bonus- Expected: %d, Result: %d", e.Name, e.Bonus, result[i].Bonus)
				}
			}
			
			if tt.applied != returned {
				t.Errorf("Not Applied Correctly- Expected: %t, Result: %t", tt.applied, returned)
			}
		})
	}
}

func TestRangerAppliedTwoWeaponFighting(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	[]eqmt.Weapon
		applied		bool
	}{
		{
			name: "No applicable weapons, bonus not applied",
			character: &models.Character {
				Abilities: []attr.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []eqmt.Weapon {
					{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged", Properties: []string {"two-handed"}},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged", Properties: []string {"two-handed"}},
			},
			applied: false,
		},
		{
			name: "One applicable weapon, bonus not applied",
			character: &models.Character {
				Abilities: []attr.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []eqmt.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged", Properties: []string {"two-handed"}},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
			},
			applied: false,
		},
		{
			name: "Two applicable light weapons, bonus applied",
			character: &models.Character {
				Abilities: []attr.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []eqmt.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Club", Bonus: 4, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
			},
			applied: true,
		},
		{
			name: "Two applicable weapons, one light, both one handed, bonus applied",
			character: &models.Character {
				Abilities: []attr.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []eqmt.Weapon {
					{Name: "Rapier", Bonus: 2, Damage: "1d8", Range: "melee", Properties: []string {"finesse"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
				},
			},
			expected: []eqmt.Weapon {
				{Name: "Rapier", Bonus: 2, Damage: "1d8", Range: "melee", Properties: []string {"finesse"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
			},
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func (t *testing.T) {
			returned := applyTwoWeaponFighting(tt.character)
			result := tt.character.Weapons

			for i, e := range tt.expected {
				if e.Bonus != result[i].Bonus {
					t.Errorf("Weapon %s Bonus- Expected: %d, Result: %d", e.Name, e.Bonus, result[i].Bonus)
				}
			}

			if tt.applied != returned {
				t.Errorf("Not Applied Correctly- Expected: %t, Result: %t", tt.applied, returned)
			}
		})
	}
}
