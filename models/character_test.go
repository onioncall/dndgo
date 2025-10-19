package models

import (
	"testing"

	"github.com/onioncall/dndgo/types"
)

func TestCharacterCalculateAbilitiesFromBase(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  []types.Abilities
	}{
		{
			name: "Ability mod round down",
			character: &Character{
				Level: 3,
				Abilities: []types.Abilities{
					{Name: "Strength", AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},
					{Name: "Dexterity", AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},
					{Name: "Constitution", AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", AbilityModifier: 2, Base: 14, SavingThrowsProficient: true},
				{Name: "Dexterity", AbilityModifier: 1, Base: 12, SavingThrowsProficient: false},
				{Name: "Constitution", AbilityModifier: 2, Base: 15, SavingThrowsProficient: true},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateAbilitiesFromBase()

			for i, e := range tt.expected {
				result := tt.character.Abilities[i].AbilityModifier
				if e.AbilityModifier != result {
					t.Errorf("Expected %d modifier, returned %d: %s", e.AbilityModifier, result, e.Name)
				}
			}
		})
	}
}

func TestCharacterCalculateSkillModifierFromBase(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  []types.Skill
	}{
		{
			name: "Multiple skills, different values",
			character: &Character{
				Skills: []types.Skill{
					{Name: "slight of hand", SkillModifier: 0, Proficient: false, Ability: "dexterity"},
					{Name: "persuasion", SkillModifier: 0, Proficient: false, Ability: "charisma"},
					{Name: "deception", SkillModifier: 0, Proficient: false, Ability: "charisma"},
				},
				Abilities: []types.Abilities{
					{Name: "Strength", AbilityModifier: 2, Base: 14, SavingThrowsProficient: true},
					{Name: "Dexterity", AbilityModifier: 1, Base: 12, SavingThrowsProficient: false},
					{Name: "Constitution", AbilityModifier: 2, Base: 15, SavingThrowsProficient: true},
					{Name: "Charisma", AbilityModifier: 0, Base: 10, SavingThrowsProficient: true},
				},
			},
			expected: []types.Skill{
				{Name: "slight of hand", SkillModifier: 1, Proficient: false, Ability: "dexterity"},
				{Name: "persuasion", SkillModifier: 0, Proficient: false, Ability: "charisma"},
				{Name: "deception", SkillModifier: 0, Proficient: false, Ability: "charisma"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateSkillModifierFromBase()

			if len(tt.character.Skills) != len(tt.expected) {
				t.Errorf("Skills Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Skills))
				return
			}

			for i, e := range tt.expected {
				result := tt.character.Skills[i].SkillModifier
				if e.SkillModifier != result {
					t.Errorf("Modifier %s- Expected: %d, returned: %d", e.Name, e.SkillModifier, result)
				}
			}
		})
	}
}

func TestCharacterCalculateProficiencyBonus(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  int
	}{
		{
			name: "Level 3 character",
			character: &Character{
				Level: 3,
			},
			expected: 2,
		},
		{
			name: "Level 8 character",
			character: &Character{
				Level: 8,
			},
			expected: 3,
		},
		{
			name: "Level 9 character",
			character: &Character{
				Level: 9,
			},
			expected: 4,
		},
		{
			name: "Level 13 character",
			character: &Character{
				Level: 13,
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateProficiencyBonusByLevel()

			result := tt.character.Proficiency
			if tt.expected != tt.character.Proficiency {
				t.Errorf("Proficiency Expected: %d, Returned: %d, Level %d", tt.expected, result, tt.character.Level)
			}
		})
	}
}

func TestCharacterCalculateAbilityScoreImprovement(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  []types.Abilities
	}{
		{
			name: "Level not high enough",
			character: &Character{
				Level: 3,
				Abilities: []types.Abilities{
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: "Strength", Bonus: 2},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, one ability increased by two",
			character: &Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, two abilities increased by one",
			character: &Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 1},
					{Ability: "Charisma", Bonus: 1},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 11, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 11, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, two abilities increased by two (failure)",
			character: &Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Charisma", Bonus: 2},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 8, one ability increased by two, and two abilities increased by one",
			character: &Character{
				Level: 8,
				Abilities: []types.Abilities{
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Charisma", Bonus: 1},
					{Ability: "Wisdom", Bonus: 1},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 11, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 11, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 20, one ability over maximum",
			character: &Character{
				Level: 20,
				Abilities: []types.Abilities{
					{Name: "Strength", Base: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Base: 12, SavingThrowsProficient: false},
					{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []types.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
				},
			},
			expected: []types.Abilities{
				{Name: "Strength", Base: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Base: 20, SavingThrowsProficient: false},
				{Name: "Constitution", Base: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Base: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Base: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Base: 10, SavingThrowsProficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateAbilityScoreImprovement()

			for i, e := range tt.expected {
				result := tt.character.Abilities[i]

				if e.Base != result.Base {
					t.Errorf("Ability Base %s- Expected: %d, Result: %d", e.Name, e.Base, result.Base)
				}
			}
		})
	}
}

func TestCharacterRecover(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  Character
	}{
		{
			name: "Recover Health, Spell Slots, Class Detail Slots",
			character: &Character{
				HPCurrent: 0,
				HPMax:     16,
				ClassName: "character",
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 4, Available: 1},
					{Level: 2, Slot: 2, Available: 0},
				},
			},
			expected: Character{
				HPCurrent: 16,
				HPMax:     16,
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 4, Available: 4},
					{Level: 2, Slot: 2, Available: 2},
				},
			},
		},
		{
			name: "Recover Health, Spell Slots, Multiple Class Detail Slots",
			character: &Character{
				HPCurrent: 0,
				HPMax:     16,
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 4, Available: 1},
					{Level: 2, Slot: 2, Available: 0},
				},
			},
			expected: Character{
				HPCurrent: 16,
				HPMax:     16,
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 4, Available: 4},
					{Level: 2, Slot: 2, Available: 2},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.Recover()

			if tt.character.HPCurrent != tt.expected.HPCurrent {
				t.Errorf("HPCurrent- Expected: %d, Result: %d", tt.character.HPCurrent, tt.character.HPCurrent)
			}

			// We should never mutate the max HP
			if tt.character.HPMax != tt.expected.HPMax {
				t.Errorf("HPMax- Expected: %d, Result: %d BAAAAD", tt.character.HPMax, tt.character.HPMax)
			}

			// TODO: Add implementation for class details in full recovery
			// for i, e := range tt.expected.ClassDetails.Slots {
			// 	result := tt.character.ClassDetails.Slots[i]
			//
			// 	if e.Available != result.Available {
			// 		t.Errorf("Class Detail Slot %s- Expected: %d, Result: %d", e.Name, e.Available, result.Available)
			// 	}
			// }

			for i, e := range tt.expected.SpellSlots {
				result := tt.character.SpellSlots[i]

				if e.Available != result.Available {
					t.Errorf("Spell Slot Level %d- Expected: %d, Result: %d", e.Level, e.Available, result.Available)
				}
			}
		})
	}
}

func TestCharacterUseSpellSlot(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		level     int
		expected  []types.SpellSlot
	}{
		{
			name:  "Use Level 1 Slot",
			level: 1,
			character: &Character{
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 6, Available: 6},
					{Level: 2, Slot: 3, Available: 3},
				},
			},
			expected: []types.SpellSlot{
				{Level: 1, Slot: 6, Available: 5},
				{Level: 2, Slot: 3, Available: 3},
			},
		},
		{
			name:  "All Slots Used",
			level: 1,
			character: &Character{
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 6, Available: 0},
					{Level: 2, Slot: 3, Available: 3},
				},
			},
			expected: []types.SpellSlot{
				{Level: 1, Slot: 6, Available: 0},
				{Level: 2, Slot: 3, Available: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.UseSpellSlot(tt.level)

			for i, e := range tt.expected {
				result := tt.character.SpellSlots[i]

				if e != result {
					t.Errorf("Spell Slot Level %d- Expected: %d, Result: %d", e.Level, e.Available, result.Available)
				}
			}
		})
	}
}

func TestCharacterRecoverSpellSlots(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		level     int
		expected  []types.SpellSlot
	}{
		{
			name:  "Recover Level 1 Slot",
			level: 1,
			character: &Character{
				SpellSlots: []types.SpellSlot{
					{Level: 1, Slot: 6, Available: 3},
					{Level: 2, Slot: 3, Available: 3},
				},
			},
			expected: []types.SpellSlot{
				{Level: 1, Slot: 6, Available: 6},
				{Level: 2, Slot: 3, Available: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.RecoverSpellSlots(tt.level)

			for i, e := range tt.expected {
				result := tt.character.SpellSlots[i]

				if e != result {
					t.Errorf("Spell Slot Level %d- Expected: %d, Result: %d", e.Level, e.Available, result.Available)
				}
			}
		})
	}
}

func TestCharacterDamageCharacter(t *testing.T) {
	tests := []struct {
		name      string
		damage    int
		character *Character
		expected  Character
	}{
		{
			name:   "Some Damage",
			damage: 5,
			character: &Character{
				HPCurrent: 16,
				HPMax:     16,
			},
			expected: Character{
				HPCurrent: 11,
				HPMax:     16,
			},
		},
		{
			name:   "Damage Below Zero",
			damage: 16,
			character: &Character{
				HPCurrent: 11,
				HPMax:     16,
			},
			expected: Character{
				HPCurrent: 0,
				HPMax:     16,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.DamageCharacter(tt.damage)

			e := tt.expected
			result := tt.character
			if e.HPCurrent != result.HPCurrent {
				t.Errorf("HPCurrent- Expected: %d, Result: %d", e.HPCurrent, result.HPCurrent)
			}

			// We should never mutate the max HP
			if e.HPMax != result.HPMax {
				t.Errorf("HPMax- Expected: %d, Result: %d", e.HPMax, result.HPMax)
			}
		})
	}
}

func TestCharacterHealCharacter(t *testing.T) {
	tests := []struct {
		name            string
		healthRecovered int
		character       *Character
		expected        Character
	}{
		{
			name:            "Some Recovery",
			healthRecovered: 4,
			character: &Character{
				HPCurrent: 11,
				HPMax:     16,
			},
			expected: Character{
				HPCurrent: 15,
				HPMax:     16,
			},
		},
		{
			name:            "Greater Than Full Recovery",
			healthRecovered: 16,
			character: &Character{
				HPCurrent: 11,
				HPMax:     16,
			},
			expected: Character{
				HPCurrent: 16,
				HPMax:     16,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.HealCharacter(tt.healthRecovered)

			e := tt.expected
			result := tt.character
			if e.HPCurrent != result.HPCurrent {
				t.Errorf("HPCurrent- Expected: %d, Result: %d", e.HPCurrent, result.HPCurrent)
			}

			// We should never mutate the max HP
			if e.HPMax != result.HPMax {
				t.Errorf("HPMax- Expected: %d, Result: %d", e.HPMax, result.HPMax)
			}
		})
	}
}

func TestCharacterAddEquipment(t *testing.T) {
	tests := []struct {
		name          string
		character     *Character
		equipmentType string
		equipmentName string
		expected      types.WornEquipment
	}{
		{
			name:          "Add Cloak",
			character:     &Character{},
			equipmentType: "cloak",
			equipmentName: "cloak of rad shit",
			expected: types.WornEquipment{
				Cloak: "cloak of rad shit",
			},
		},
		{
			name: "EquipmentType not valid",
			character: &Character{
				WornEquipment: types.WornEquipment{
					Cloak: "cloak of rad shit",
				},
			},
			equipmentType: "cloakwef",
			equipmentName: "cloak of cool shit",
			expected: types.WornEquipment{
				Cloak: "cloak of rad shit",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.AddEquipment(tt.equipmentType, tt.equipmentName)

			e := tt.expected.Cloak
			result := tt.character.WornEquipment.Cloak

			if e != result {
				t.Errorf("Cloak- Expected: %s. Result: %s", e, result)
			}
		})
	}
}

func TestCharacterRemoveItemFromBackpack(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		itemName  string
		quantity  int
		expected  []types.BackpackItem
	}{
		{
			name:     "Remove 1 Item",
			itemName: "soap",
			quantity: 5,
			character: &Character{
				Backpack: []types.BackpackItem{
					{Name: "soap", Quantity: 50},
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []types.BackpackItem{
				{Name: "soap", Quantity: 45},
				{Name: "gold", Quantity: 5},
			},
		},
		{
			name:     "Remove More Than Available Quantity",
			itemName: "soap",
			quantity: 51,
			character: &Character{
				Backpack: []types.BackpackItem{
					{Name: "soap", Quantity: 51},
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []types.BackpackItem{
				{Name: "soap", Quantity: 0},
				{Name: "gold", Quantity: 5},
			},
		},
		{
			name:     "Item Not In Backpack",
			itemName: "soapehrgerg",
			quantity: 50,
			character: &Character{
				Backpack: []types.BackpackItem{
					{Name: "soap", Quantity: 50},
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []types.BackpackItem{
				{Name: "soap", Quantity: 50},
				{Name: "gold", Quantity: 5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.RemoveItemFromPack(tt.itemName, tt.quantity)

			if len(tt.expected) != len(tt.character.Backpack) {
				t.Errorf("Item Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Backpack))
			}

			for i, e := range tt.expected {
				result := tt.character.Backpack[i]

				if e.Quantity != result.Quantity {
					t.Errorf("Item Quantity %s- Expected: %d, Result: %d", e.Name, e.Quantity, result.Quantity)
				}
			}
		})
	}
}

func TestCharacterAddItemToBackpack(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		itemName  string
		quantity  int
		expected  []types.BackpackItem
	}{
		{
			name:     "Add 1 New Item",
			itemName: "soap",
			quantity: 5,
			character: &Character{
				Backpack: []types.BackpackItem{
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []types.BackpackItem{
				{Name: "gold", Quantity: 5},
				{Name: "soap", Quantity: 5},
			},
		},
		{
			name:     "Add 1 Existing Item",
			itemName: "soap",
			quantity: 5,
			character: &Character{
				Backpack: []types.BackpackItem{
					{Name: "gold", Quantity: 5},
					{Name: "soap", Quantity: 5},
				},
			},
			expected: []types.BackpackItem{
				{Name: "gold", Quantity: 5},
				{Name: "soap", Quantity: 10},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.AddItemToPack(tt.itemName, tt.quantity)

			if len(tt.expected) != len(tt.character.Backpack) {
				t.Errorf("Item Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Backpack))
			}

			for i, e := range tt.expected {
				result := tt.character.Backpack[i]

				if e != result {
					t.Errorf("Item Quantity %s- Expected: %d, Result: %d", e.Name, e.Quantity, result.Quantity)
				}
			}
		})
	}
}
