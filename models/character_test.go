package models

import (
	"testing"
)

func TestCharacter_calculateAttributesFromBase(t *testing.T) {
	tests := []struct {
		name 			string
		character 		*Character
		expected	 	[]Attribute
	}{
		{
			name: "Ability Mod Round Down",
			character: &Character {
				Level: 3,
				Attributes: []Attribute {
					{Name: "Strength", AbilityModifier: 0, Base: 14, SavingThrowsProficient: true},	
					{Name: "Dexterity", AbilityModifier: 0, Base: 12, SavingThrowsProficient: false},	
					{Name: "Constitution", AbilityModifier: 0, Base: 15, SavingThrowsProficient: true},	
				},
			},
			expected: []Attribute {
				{Name: "Strength", AbilityModifier: 2, Base: 14, SavingThrowsProficient: true},	
				{Name: "Dexterity", AbilityModifier: 1, Base: 12, SavingThrowsProficient: false},	
				{Name: "Constitution", AbilityModifier: 2, Base: 15, SavingThrowsProficient: true},	
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateAttributesFromBase()

			for i, e := range tt.expected {
				result := tt.character.Attributes[i].AbilityModifier
				if e.AbilityModifier != result {
					t.Errorf("Expected %d modifier, returned %d: %s", e.AbilityModifier, result, e.Name)
				}
			}
		})
	}
}

func TestCharacter_calculateSkillModifierFromBase(t *testing.T) {
	tests := []struct {
		name 		string
		character 	*Character
		expected 	[]Skill
	}{
		{
			name: "Multiple Skills, different values",
			character: &Character {
				Skills: []Skill {
					{Name: "slight of hand", SkillModifier: 0, Proficient: false, Attribute: "dexterity"},
					{Name: "persuasion", SkillModifier: 0, Proficient: false, Attribute: "charisma"},
					{Name: "deception", SkillModifier: 0, Proficient: false, Attribute: "charisma"},
				},
				Attributes: []Attribute {
					{Name: "Strength", AbilityModifier: 2, Base: 14, SavingThrowsProficient: true},	
					{Name: "Dexterity", AbilityModifier: 1, Base: 12, SavingThrowsProficient: false},	
					{Name: "Constitution", AbilityModifier: 2, Base: 15, SavingThrowsProficient: true},	
					{Name: "Charisma", AbilityModifier: 0, Base: 10, SavingThrowsProficient: true},	
				},
			},
			expected: []Skill {
				{Name: "slight of hand", SkillModifier: 1, Proficient: false, Attribute: "dexterity"},
				{Name: "persuasion", SkillModifier: 0, Proficient: false, Attribute: "charisma"},
				{Name: "deception", SkillModifier: 0, Proficient: false, Attribute: "charisma"},
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
					t.Errorf("Modifier %s- Expected: %d, returned: %d", e.Name,  e.SkillModifier, result)
				}
			}
		})
	}
}


func TestCharacter_calculateProficiencyBonus(t *testing.T) {
	tests := []struct {
		name		string
		character	*Character
		expected 	int 
	}{
		{
			name: "Level 3 Character",
			character: &Character {
				Level: 3,
			},
			expected: 2,
		},
		{
			name: "Level 8 Character",
			character: &Character {
				Level: 8,
			},
			expected: 3,
		},
		{
			name: "Level 9 Character",
			character: &Character {
				Level: 9,
			},
			expected: 4,
		},
		{
			name: "Level 13 Character",
			character: &Character {
				Level: 13,
			},
			expected: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.calculateProficiencyBonusByLevel()

			result :=  tt.character.Proficiency
			if tt.expected != tt.character.Proficiency {
				t.Errorf("Proficiency Expected: %d, Returned: %d, Level %d", tt.expected, result, tt.character.Level)
			}
		})
	}
}

func TestCharacter_UseClassSlots(t *testing.T) {
	tests := []struct {
		name		string
		slotName 	string
		character 	*Character
		expected	[]ClassSlot
	}{
		{
			name: "One Use, Single Slot",
			slotName: "bardic inspiration",
			character: &Character {
				ClassDetails: ClassDetails {
					Slots: []ClassSlot {
						{Name: "bardic inspiration", Slot: 4, Available: 4},
					},
				},
			},
			expected: []ClassSlot {
				{Name: "bardic inspiration", Slot: 4, Available: 3},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.UseClassSlots(tt.slotName)	

			if len(tt.character.ClassDetails.Slots) != len(tt.expected) {
				t.Errorf("Skills Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Skills))
				return
			}

			for i, e := range tt.expected {
				result := tt.character.ClassDetails.Slots[i]
				if e != result {
					t.Errorf("Expected ClassSlot: %v\nReturned: %v", e, result)
				}
			}
		})
	}
}

func TestCharacter_RecoverClassDetailSlots(t *testing.T) {
	tests := []struct {
		name		string
		slotName 	string
		character 	*Character
		expected	[]ClassSlot
	}{
		{
			name: "Recover Single Slot",
			slotName: "bardic inspiration",
			character: &Character {
				ClassDetails: ClassDetails {
					Slots: []ClassSlot {
						{Name: "bardic inspiration", Slot: 4, Available: 0},
					},
				},
			},
			expected: []ClassSlot {
				{Name: "bardic inspiration", Slot: 4, Available: 4},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.character.RecoverClassDetailSlots(tt.slotName)

			if len(tt.character.ClassDetails.Slots) != len(tt.expected) {
				t.Errorf("Skills Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Skills))
				return
			}

			for i, e := range tt.expected {
				result := tt.character.ClassDetails.Slots[i]
				if e != result {
					t.Errorf("Expected ClassSlot: %v\nReturned: %v", e, result)
				}
			}
		})
	}
}

func TestCharacter_Recover(t *testing.T) {
	tests := []struct {
		name		string
		character	*Character
		expected	Character
	}{
		{
			name: "Recover Health, Spell Slots, Class Detail Slots",
			character: &Character{
				HPCurrent: 0,
				HPMax: 16,
				ClassDetails: ClassDetails {
					Slots: []ClassSlot {
						{Name: "bardic inspiration", Slot: 4, Available: 0},
					},
				},
				SpellSlots: []SpellSlot {
					{Level: 1, Slot: 4, Available: 1},
					{Level: 2, Slot: 2, Available: 0},
				},
			},
			expected: Character{
				HPCurrent: 16,
				HPMax: 16,
				SpellSlots: []SpellSlot {
					{Level: 1, Slot: 4, Available: 4},
					{Level: 2, Slot: 2, Available: 2},
				},
				ClassDetails: ClassDetails {
					Slots: []ClassSlot {
						{Name: "bardic inspiration", Slot: 4, Available: 4},
					},
				},
			},
		},
		{
			name: "Recover Health, Spell Slots, Multiple Class Detail Slots",
			character: &Character{
				HPCurrent: 0,
				HPMax: 16,
				ClassDetails: ClassDetails {
					Slots: []ClassSlot {
						{Name: "bardic inspiration", Slot: 4, Available: 0},
						{Name: "some other charge or token", Slot: 3, Available: 1},
					},
				},
				SpellSlots: []SpellSlot {
					{Level: 1, Slot: 4, Available: 1},
					{Level: 2, Slot: 2, Available: 0},
				},
			},
			expected: Character{
				HPCurrent: 16,
				HPMax: 16,
				SpellSlots: []SpellSlot {
					{Level: 1, Slot: 4, Available: 4},
					{Level: 2, Slot: 2, Available: 2},
				},
				ClassDetails: ClassDetails {
					Slots: []ClassSlot {
						{Name: "bardic inspiration", Slot: 4, Available: 4},
						{Name: "some other charge or token", Slot: 3, Available: 3},
					},
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

			for i, e := range tt.expected.ClassDetails.Slots {
				result := tt.character.ClassDetails.Slots[i]

				if e.Available != result.Available {
					t.Errorf("Class Detail Slot %s- Expected: %d, Result: %d", e.Name, e.Available, result.Available)
				}
			}

			for i, e := range tt.expected.SpellSlots {
				result := tt.character.SpellSlots[i]

				if e.Available != result.Available {
					t.Errorf("Spell Slot Level %d- Expected: %d, Result: %d", e.Level, e.Available, result.Available)
				}
			}
		})
	}
}

func TestCharacter_UseSpellSlot(t *testing.T) {
	tests := []struct {
		name		string
		character	*Character
		level		int
		expected	[]SpellSlot
	}{
		{
			name: "Use Level 1 Slot",
			level: 1,
			character: &Character {
				SpellSlots: []SpellSlot {
					{Level: 1, Slot: 6, Available: 6},
					{Level: 2, Slot: 3, Available: 3},
				},
			},
			expected: []SpellSlot {
				{Level: 1, Slot: 6, Available: 5},
				{Level: 2, Slot: 3, Available: 3},
			},
		},
		{
			name: "All Slots Used",
			level: 1,
			character: &Character {
				SpellSlots: []SpellSlot {
					{Level: 1, Slot: 6, Available: 0},
					{Level: 2, Slot: 3, Available: 3},
				},
			},
			expected: []SpellSlot {
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
