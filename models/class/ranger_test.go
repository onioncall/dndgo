package class

import (
	"testing"

	"github.com/onioncall/dndgo/models"
	"github.com/onioncall/dndgo/types"
)

func TestRangerAppliedArchery(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	[]types.Weapon
		applied		bool
	}{
		{
			name: "No ranged weapon",
			character: &models.Character {
				Weapons: []types.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Dagger", Bonus: 2, Damage: "1d4", Range: "melee"},
				},
			},
			expected: []types.Weapon {
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				{Name: "Dagger", Bonus: 2, Damage: "1d4", Range: "melee"},
			},
			applied: false,
		},
		{
			name: "Range bonus applied",
			character: &models.Character {
				Weapons: []types.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged"},
				},
			},
			expected: []types.Weapon {
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
				WornEquipment: types.WornEquipment {
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
				WornEquipment: types.WornEquipment {
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
		expected	[]types.Weapon
		applied		bool
	}{
		{
			name: "No melee weapon",
			character: &models.Character {
				Weapons: []types.Weapon {
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged"},
				},
			},
			expected: []types.Weapon {
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged"},
			},
			applied: false,
		},
		{
			name: "Melee bonus applied",
			character: &models.Character {
				Weapons: []types.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				},
			},
			expected: []types.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Range: "melee"},
			},
			applied: true,
		},
		{
			name: "Multiple valid weapons, one bonus",
			character: &models.Character {
				Weapons: []types.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
				},
			},
			expected: []types.Weapon {
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
		expected	[]types.Weapon
		applied		bool
	}{
		{
			name: "No applicable weapons, bonus not applied",
			character: &models.Character {
				Abilities: []types.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon {
					{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged", Properties: []string {"two-handed"}},
				},
			},
			expected: []types.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged", Properties: []string {"two-handed"}},
			},
			applied: false,
		},
		{
			name: "One applicable weapon, bonus not applied",
			character: &models.Character {
				Abilities: []types.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Range: "ranged", Properties: []string {"two-handed"}},
				},
			},
			expected: []types.Weapon {
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Range: "melee", Properties: []string {"two-handed"}},
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee"},
			},
			applied: false,
		},
		{
			name: "Two applicable light weapons, bonus applied",
			character: &models.Character {
				Abilities: []types.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon {
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
				},
			},
			expected: []types.Weapon {
				{Name: "Club", Bonus: 4, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
				{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
			},
			applied: true,
		},
		{
			name: "Two applicable weapons, one light, both one handed, bonus applied",
			character: &models.Character {
				Abilities: []types.Abilities {
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon {
					{Name: "Rapier", Bonus: 2, Damage: "1d8", Range: "melee", Properties: []string {"finesse"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Range: "melee", Properties: []string {"light"}},
				},
			},
			expected: []types.Weapon {
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

func TestRanger_executeFightingStyle(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		ranger		*Ranger
		expected	models.Character	
	}{
		{
			name: "Not valid fighting style",
			character: &models.Character {
				AC: 15,
				Level: 3,
				WornEquipment: types.WornEquipment {
					Armour: "",
				},
			},
			ranger: &Ranger {
				FightingStyle: "the-worm",
			},
			expected: models.Character {
				AC: 15,
				Level: 3,
			},
		},
		{
			name: "Defense applied",
			character: &models.Character {
				AC: 15,
				Level: 3,
				WornEquipment: types.WornEquipment {
					Armour: "",
				},
			},
			ranger: &Ranger {
				FightingStyle: Defense,
			},
			expected: models.Character {
				AC: 16,
				Level: 3,
			},
		},
		{
			name: "Defense not applied (armor equiped)",
			character: &models.Character {
				AC: 15,
				Level: 3,
				WornEquipment: types.WornEquipment {
					Armour: "light-armor",
				},
			},
			ranger: &Ranger {
				FightingStyle: Defense,
			},
			expected: models.Character {
				AC: 15,
				Level: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ranger.executeFightingStyle(tt.character)
			result := tt.character

			if tt.expected.AC != result.AC {
				t.Errorf("AC- Expected: %d, Result: %d", tt.expected.AC, result.AC)
			}

			for i, e := range tt.expected.Weapons {
				if e.Bonus != result.Weapons[i].Bonus {
					t.Errorf("Weapon Bonus %s- Expected: %d, Result: %d", e.Name, e.Bonus, result.Weapons[i].Bonus)
				}
			}
		})
	}
}
