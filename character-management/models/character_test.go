package models

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/shared"
)

func TestCharacterCalculateAbilitiesFromAdjusted(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  []shared.Ability
	}{
		{
			name: "Ability mod round down",
			character: &Character{
				Level: 3,
				Abilities: []shared.Ability{
					{Name: "Strength", AbilityModifier: 0, Adjusted: 14, SavingThrowsProficient: true},
					{Name: "Dexterity", AbilityModifier: 0, Adjusted: 12, SavingThrowsProficient: false},
					{Name: "Constitution", AbilityModifier: 0, Adjusted: 15, SavingThrowsProficient: true},
				},
			},
			expected: []shared.Ability{
				{Name: "Strength", AbilityModifier: 2, Adjusted: 14, SavingThrowsProficient: true},
				{Name: "Dexterity", AbilityModifier: 1, Adjusted: 12, SavingThrowsProficient: false},
				{Name: "Constitution", AbilityModifier: 2, Adjusted: 15, SavingThrowsProficient: true},
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

func TestCharacterCalculateSkillModifierFromAdjusted(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  []shared.Skill
	}{
		{
			name: "Multiple skills, different values",
			character: &Character{
				Skills: []shared.Skill{
					{Name: "slight of hand", SkillModifier: 0, Proficient: false, Ability: "dexterity"},
					{Name: "persuasion", SkillModifier: 0, Proficient: false, Ability: "charisma"},
					{Name: "deception", SkillModifier: 0, Proficient: false, Ability: "charisma"},
				},
				Abilities: []shared.Ability{
					{Name: "Strength", AbilityModifier: 2, Adjusted: 14, SavingThrowsProficient: true},
					{Name: "Dexterity", AbilityModifier: 1, Adjusted: 12, SavingThrowsProficient: false},
					{Name: "Constitution", AbilityModifier: 2, Adjusted: 15, SavingThrowsProficient: true},
					{Name: "Charisma", AbilityModifier: 0, Adjusted: 10, SavingThrowsProficient: true},
				},
			},
			expected: []shared.Skill{
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
		expected  []shared.Ability
	}{
		{
			name: "Level 4, one ability increased by two",
			character: &Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []shared.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
				},
			},
			expected: []shared.Ability{
				{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Adjusted: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, two abilities increased by one",
			character: &Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []shared.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 1},
					{Ability: "Charisma", Bonus: 1},
				},
			},
			expected: []shared.Ability{
				{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Adjusted: 11, SavingThrowsProficient: false},
				{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Adjusted: 11, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 4, two abilities increased by two",
			character: &Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []shared.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Charisma", Bonus: 2},
				},
			},
			expected: []shared.Ability{
				{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Adjusted: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Adjusted: 12, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 8, one ability increased by two, and two abilities increased by one",
			character: &Character{
				Level: 8,
				Abilities: []shared.Ability{
					{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []shared.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Charisma", Bonus: 1},
					{Ability: "Wisdom", Bonus: 1},
				},
			},
			expected: []shared.Ability{
				{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Adjusted: 12, SavingThrowsProficient: false},
				{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Adjusted: 11, SavingThrowsProficient: false},
				{Name: "Charisma", Adjusted: 11, SavingThrowsProficient: false},
			},
		},
		{
			name: "Level 20, one ability over maximum",
			character: &Character{
				Level: 20,
				Abilities: []shared.Ability{
					{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Dexterity", Adjusted: 12, SavingThrowsProficient: false},
					{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
					{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
				},
				AbilityScoreImprovement: []shared.AbilityScoreImprovementItem{
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
					{Ability: "Dexterity", Bonus: 2},
				},
			},
			expected: []shared.Ability{
				{Name: "Strength", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Dexterity", Adjusted: 20, SavingThrowsProficient: false},
				{Name: "Constitution", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Intelligence", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Wisdom", Adjusted: 10, SavingThrowsProficient: false},
				{Name: "Charisma", Adjusted: 10, SavingThrowsProficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateAbilityScoreImprovement()

			for i, e := range tt.expected {
				result := tt.character.Abilities[i]

				if e.Adjusted != result.Adjusted {
					t.Errorf("Ability Adjusted %s- Expected: %d, Result: %d", e.Name, e.Adjusted, result.Adjusted)
				}
			}
		})
	}
}

func TestCharacterCalculateAC(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  int
	}{
		{
			name: "Light armor, no shield",
			character: &Character{
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 2},
				},
				AC: 0,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Type:       shared.LightArmor,
						Name:       "Leather Armor",
						Class:      11,
						Proficient: true,
					},
				},
			},
			expected: 13,
		},
		{
			name: "Light armor, with shield",
			character: &Character{
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 2},
				},
				AC:              0,
				PrimaryEquipped: "sOmE SHieLd",
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Type:       shared.LightArmor,
						Name:       "Leather Armor",
						Class:      11,
						Proficient: true,
					},
					Shield: "Some Shield",
				},
			},
			expected: 15,
		},
		{
			name: "Medium armor, no shield",
			character: &Character{
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 4},
				},
				AC: 0,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Type:       shared.MediumArmor,
						Name:       "Scale Mail",
						Class:      14,
						Proficient: true,
					},
				},
			},
			expected: 16,
		},
		{
			name: "Heavy armor, no shield",
			character: &Character{
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 4},
				},
				AC: 0,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Type:       shared.HeavyArmor,
						Name:       "Chain Mail",
						Class:      16,
						Proficient: true,
					},
				},
			},
			expected: 16,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateAC()
			result := tt.character.AC

			if tt.expected != result {
				t.Errorf("AC- Expected: %d, Result: %d", tt.expected, result)
			}
		})
	}
}

func TestCharacterCalculateWeaponBonus(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  []shared.Weapon
	}{
		{
			name: "Proficient finesse weapon",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 2},
					{Name: shared.AbilityStrength, AbilityModifier: 1},
				},
				Weapons: []shared.Weapon{
					{Name: "Rapier", Bonus: 0, Proficient: true, Properties: []string{shared.WeaponPropertyFinesse}},
				},
			},
			expected: []shared.Weapon{
				{Name: "Rapier", Bonus: 2, Proficient: true, Properties: []string{shared.WeaponPropertyFinesse}},
			},
		},
		{
			name: "Non-proficient finesse weapon",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 2},
					{Name: shared.AbilityStrength, AbilityModifier: 1},
				},
				Weapons: []shared.Weapon{
					{Name: "Rapier", Bonus: 0, Proficient: false, Properties: []string{shared.WeaponPropertyFinesse}},
				},
			},
			expected: []shared.Weapon{
				{Name: "Rapier", Bonus: 2, Proficient: false, Properties: []string{shared.WeaponPropertyFinesse}},
			},
		},
		{
			name: "Proficient finesse weapon, higher strength mod",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 2},
					{Name: shared.AbilityStrength, AbilityModifier: 3},
				},
				Weapons: []shared.Weapon{
					{Name: "Rapier", Bonus: 0, Proficient: true, Properties: []string{shared.WeaponPropertyFinesse}},
				},
			},
			expected: []shared.Weapon{
				{Name: "Rapier", Bonus: 3, Proficient: true, Properties: []string{shared.WeaponPropertyFinesse}},
			},
		},
		{
			name: "Proficient ranged weapon",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 2},
					{Name: shared.AbilityStrength, AbilityModifier: 3},
				},
				Weapons: []shared.Weapon{
					{Name: "Sling", Bonus: 0, Proficient: true, Ranged: true},
				},
			},
			expected: []shared.Weapon{
				{Name: "Sling", Bonus: 2, Proficient: true, Ranged: true},
			},
		},
		{
			name: "Proficient melee weapon",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityDexterity, AbilityModifier: 3},
					{Name: shared.AbilityStrength, AbilityModifier: 2},
				},
				Weapons: []shared.Weapon{
					{Name: "Club", Bonus: 0, Proficient: true, Ranged: false},
				},
			},
			expected: []shared.Weapon{
				{Name: "Club", Bonus: 2, Proficient: true, Ranged: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateWeaponBonus()

			for i, e := range tt.expected {
				result := tt.character.Weapons[i].Bonus
				if e.Bonus != result {
					t.Errorf("Weapon Bonus '%s'- Expected: %d, Result: %d", e.Name, e.Bonus, result)
				}
			}
		})
	}
}

func TestCharacterCalculatePassiveStats(t *testing.T) {
	tests := []struct {
		name      string
		character *Character
		expected  Character
	}{
		{
			name: "Proficient perception, non-proficient insight",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityWisdom, AbilityModifier: 4},
				},
				Skills: []shared.Skill{
					{Name: shared.SkillPerception, SkillModifier: 2, Proficient: true},
					{Name: shared.SkillInsight, SkillModifier: 3, Proficient: false},
				},
			},
			expected: Character{
				PassivePerception: 16,
				PassiveInsight:    14,
			},
		},
		{
			name: "Proficient perception, non-proficient insight",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityWisdom, AbilityModifier: 4},
				},
				Skills: []shared.Skill{
					{Name: shared.SkillPerception, SkillModifier: 2, Proficient: false},
					{Name: shared.SkillInsight, SkillModifier: 3, Proficient: true},
				},
			},
			expected: Character{
				PassivePerception: 14,
				PassiveInsight:    16,
			},
		},
		{
			name: "Non-proficient perception, non-proficient insight",
			character: &Character{
				Proficiency: 2,
				Abilities: []shared.Ability{
					{Name: shared.AbilityWisdom, AbilityModifier: 4},
				},
				Skills: []shared.Skill{
					{Name: shared.SkillPerception, SkillModifier: 2, Proficient: false},
					{Name: shared.SkillInsight, SkillModifier: 3, Proficient: false},
				},
			},
			expected: Character{
				PassivePerception: 14,
				PassiveInsight:    14,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculatePassiveStats()

			resultPerception := tt.character.PassivePerception
			resultInsight := tt.character.PassiveInsight
			expectedPerception := tt.expected.PassivePerception
			expectedInsight := tt.expected.PassiveInsight

			if expectedPerception != resultPerception {
				t.Errorf("Passive Perception- Expected: %d, Result: %d", expectedPerception, resultPerception)
			}
			if expectedInsight != resultInsight {
				t.Errorf("Passive resultInsight- Expected: %d, Result: %d", expectedInsight, resultInsight)
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
				HPCurrent:  0,
				HPMax:      16,
				ClassTypes: []string{shared.ClassBard},
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 4, Available: 1},
					{Level: 2, Maximum: 2, Available: 0},
				},
			},
			expected: Character{
				HPCurrent: 16,
				HPMax:     16,
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 4, Available: 4},
					{Level: 2, Maximum: 2, Available: 2},
				},
			},
		},
		{
			name: "Recover Health, Spell Slots, Multiple Class Detail Slots",
			character: &Character{
				HPCurrent: 0,
				HPMax:     16,
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 4, Available: 1},
					{Level: 2, Maximum: 2, Available: 0},
				},
			},
			expected: Character{
				HPCurrent: 16,
				HPMax:     16,
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 4, Available: 4},
					{Level: 2, Maximum: 2, Available: 2},
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
		expected  []shared.SpellSlot
	}{
		{
			name:  "Use Level 1 Slot",
			level: 1,
			character: &Character{
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 6, Available: 6},
					{Level: 2, Maximum: 3, Available: 3},
				},
			},
			expected: []shared.SpellSlot{
				{Level: 1, Maximum: 6, Available: 5},
				{Level: 2, Maximum: 3, Available: 3},
			},
		},
		{
			name:  "All Slots Used",
			level: 1,
			character: &Character{
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 6, Available: 0},
					{Level: 2, Maximum: 3, Available: 3},
				},
			},
			expected: []shared.SpellSlot{
				{Level: 1, Maximum: 6, Available: 0},
				{Level: 2, Maximum: 3, Available: 3},
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
		name            string
		character       *Character
		level           int
		recoverQuantity int
		expected        []shared.SpellSlot
	}{
		{
			name:  "Recover Level 1 Slot",
			level: 1,
			character: &Character{
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 6, Available: 3},
					{Level: 2, Maximum: 3, Available: 3},
				},
			},
			expected: []shared.SpellSlot{
				{Level: 1, Maximum: 6, Available: 6},
				{Level: 2, Maximum: 3, Available: 3},
			},
			recoverQuantity: 0, // full recover
		},
		{
			name:  "Recover Level 1 Slot by 1",
			level: 1,
			character: &Character{
				SpellSlots: []shared.SpellSlot{
					{Level: 1, Maximum: 6, Available: 3},
					{Level: 2, Maximum: 3, Available: 3},
				},
			},
			expected: []shared.SpellSlot{
				{Level: 1, Maximum: 6, Available: 4},
				{Level: 2, Maximum: 3, Available: 3},
			},
			recoverQuantity: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.RecoverSpellSlots(tt.level, tt.recoverQuantity)

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
		{
			name:   "Temp HP, with remainder",
			damage: 4,
			character: &Character{
				HPCurrent: 11,
				HPMax:     16,
				HPTemp:    5,
			},
			expected: Character{
				HPCurrent: 11,
				HPMax:     16,
				HPTemp:    1,
			},
		},
		{
			name:   "Temp HP, with damage left over",
			damage: 7,
			character: &Character{
				HPCurrent: 11,
				HPMax:     16,
				HPTemp:    5,
			},
			expected: Character{
				HPCurrent: 9,
				HPMax:     16,
				HPTemp:    0,
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

			if e.HPTemp != result.HPTemp {
				t.Errorf("HPTemp- Expected: %d, Result: %d", e.HPTemp, result.HPTemp)
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

// TODO: Rename functionality will have to change with the support of multiple character files
// func TestCharacterRenameCharacter(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		newName      string
// 		character    *Character
// 		expectedName string
// 	}{
// 		{
// 			name:    "Simple Name Change",
// 			newName: "Aragorn",
// 			character: &Character{
// 				Name: "Nim",
// 			},
// 			expectedName: "Aragorn",
// 		},
// 		{
// 			name:    "Name With Spaces",
// 			newName: "Gandalf the Grey",
// 			character: &Character{
// 				Name: "Nim",
// 			},
// 			expectedName: "Gandalf the Grey",
// 		},
// 		{
// 			name:    "Empty String",
// 			newName: "",
// 			character: &Character{
// 				Name: "Nim",
// 			},
// 			expectedName: "",
// 		},
// 		{
// 			name:    "Special Characters",
// 			newName: "Drizzt Do'Urden",
// 			character: &Character{
// 				Name: "Nim",
// 			},
// 			expectedName: "Drizzt Do'Urden",
// 		},
// 		{
// 			name:    "Unicode Characters",
// 			newName: "Søren Kierkegaard",
// 			character: &Character{
// 				Name: "Nim",
// 			},
// 			expectedName: "Søren Kierkegaard",
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tt.character.RenameCharacter(tt.newName)
//
// 			if tt.character.Name != tt.expectedName {
// 				t.Errorf("Expected name: %q, Result: %q", tt.expectedName, tt.character.Name)
// 			}
// 		})
// 	}
// }

func TestCharacterAddEquipment(t *testing.T) {
	tests := []struct {
		name          string
		character     *Character
		equipmentType string
		equipmentName string
		expected      shared.WornEquipment
	}{
		{
			name:          "Add Cloak",
			character:     &Character{},
			equipmentType: "cloak",
			equipmentName: "cloak of rad shit",
			expected: shared.WornEquipment{
				Cloak: "cloak of rad shit",
			},
		},
		{
			name: "EquipmentType not valid",
			character: &Character{
				WornEquipment: shared.WornEquipment{
					Cloak: "cloak of rad shit",
				},
			},
			equipmentType: "cloakwef",
			equipmentName: "cloak of cool shit",
			expected: shared.WornEquipment{
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
		expected  []shared.BackpackItem
	}{
		{
			name:     "Remove 1 Item",
			itemName: "soap",
			quantity: 5,
			character: &Character{
				Backpack: []shared.BackpackItem{
					{Name: "soap", Quantity: 50},
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []shared.BackpackItem{
				{Name: "soap", Quantity: 45},
				{Name: "gold", Quantity: 5},
			},
		},
		{
			name:     "Remove More Than Available Quantity",
			itemName: "soap",
			quantity: 51,
			character: &Character{
				Backpack: []shared.BackpackItem{
					{Name: "soap", Quantity: 51},
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []shared.BackpackItem{
				{Name: "soap", Quantity: 0},
				{Name: "gold", Quantity: 5},
			},
		},
		{
			name:     "Item Not In Backpack",
			itemName: "soapehrgerg",
			quantity: 50,
			character: &Character{
				Backpack: []shared.BackpackItem{
					{Name: "soap", Quantity: 50},
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []shared.BackpackItem{
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
		expected  []shared.BackpackItem
	}{
		{
			name:     "Add 1 New Item",
			itemName: "soap",
			quantity: 5,
			character: &Character{
				Backpack: []shared.BackpackItem{
					{Name: "gold", Quantity: 5},
				},
			},
			expected: []shared.BackpackItem{
				{Name: "gold", Quantity: 5},
				{Name: "soap", Quantity: 5},
			},
		},
		{
			name:     "Add 1 Existing Item",
			itemName: "soap",
			quantity: 5,
			character: &Character{
				Backpack: []shared.BackpackItem{
					{Name: "gold", Quantity: 5},
					{Name: "soap", Quantity: 5},
				},
			},
			expected: []shared.BackpackItem{
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

func TestCharacterEquip(t *testing.T) {
	tests := []struct {
		name       string
		character  *Character
		itemName   string
		isPrimary  bool
		ePrimary   string
		eSecondary string
	}{
		{
			name: "No items equipped, equip primary weapon",
			character: &Character{
				PrimaryEquipped:   "",
				SecondaryEquipped: "",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Dagger"},
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Rapier",
			isPrimary:  true,
			ePrimary:   "Rapier",
			eSecondary: "",
		},
		{
			name: "No items equipped, equip primary shield",
			character: &Character{
				PrimaryEquipped:   "",
				SecondaryEquipped: "",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Dagger"},
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Some Shield",
			isPrimary:  true,
			ePrimary:   "Some Shield",
			eSecondary: "",
		},
		{
			name: "No items equipped, equip secondary weapon",
			character: &Character{
				PrimaryEquipped:   "",
				SecondaryEquipped: "",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "dagger"}, // testing casing
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Dagger",
			isPrimary:  false,
			ePrimary:   "",
			eSecondary: "Dagger",
		},
		{
			name: "Primary equipped, equip secondary weapon",
			character: &Character{
				PrimaryEquipped:   "Rapier",
				SecondaryEquipped: "",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Dagger"},
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Dagger",
			isPrimary:  false,
			ePrimary:   "Rapier",
			eSecondary: "Dagger",
		},
		{
			name: "Both equipped, equip secondary weapon as primary",
			character: &Character{
				PrimaryEquipped:   "Rapier",
				SecondaryEquipped: "Club",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Dagger"},
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Club",
			isPrimary:  true,
			ePrimary:   "Club",
			eSecondary: "",
		},
		{
			name: "Equipping already equipped primary as secondary",
			character: &Character{
				PrimaryEquipped:   "Rapier",
				SecondaryEquipped: "",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Dagger"},
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Rapier",
			isPrimary:  false,
			ePrimary:   "",
			eSecondary: "Rapier",
		},
		{
			name: "Equipping weapon we have two of as primary and secondary",
			character: &Character{
				PrimaryEquipped:   "Rapier",
				SecondaryEquipped: "Club",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Club"},
					{Name: "Club"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Club",
			isPrimary:  true,
			ePrimary:   "Club",
			eSecondary: "Club",
		},
		{
			name: "Equipment not found",
			character: &Character{
				PrimaryEquipped:   "",
				SecondaryEquipped: "",
				Weapons: []shared.Weapon{
					{Name: "Rapier"},
					{Name: "Club"},
					{Name: "Dagger"},
				},
				WornEquipment: shared.WornEquipment{
					Shield: "Some Shield",
				},
			},
			itemName:   "Longbow",
			isPrimary:  true,
			ePrimary:   "",
			eSecondary: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.Equip(tt.isPrimary, tt.itemName)

			rPrimary := tt.character.PrimaryEquipped
			rSecondary := tt.character.SecondaryEquipped

			if tt.ePrimary != rPrimary {
				t.Errorf("Primary- Expected: %s, Result: %s", tt.ePrimary, rPrimary)
			}
			if tt.eSecondary != rSecondary {
				t.Errorf("Secondary- Expected: %s, Result: %s", tt.eSecondary, rSecondary)
			}
		})
	}
}
