package models

import (
	"testing"
)

func TestCharacter_calculateAttributesFromBase(t *testing.T) {
	tests := []struct {
		name 			string
		character 		*Character
		expectedStats 	[]Attribute
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
			expectedStats: []Attribute {
				{Name: "Strength", AbilityModifier: 2, Base: 14, SavingThrowsProficient: true},	
				{Name: "Dexterity", AbilityModifier: 1, Base: 12, SavingThrowsProficient: false},	
				{Name: "Constitution", AbilityModifier: 2, Base: 15, SavingThrowsProficient: true},	
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a deep copy to avoid modifying the original test data
			testChar := &Character {
				Level: test.character.Level,
				Attributes: make([]Attribute, len(test.character.Attributes)),
			}
			copy(testChar.Attributes, test.character.Attributes)

			testChar.calculateAttributesFromBase()

			for i, e := range test.expectedStats {
				testMod := testChar.Attributes[i].AbilityModifier
				if e.AbilityModifier != testMod {
					t.Errorf("Expected %d modifier, returned %d: %s", e.AbilityModifier, testMod, e.Name)
				}
			}
		})
	}
}

func TestCharacter_calculateSkillModifierFromBase(t *testing.T) {
	tests := []struct {
		name 			string
		character 		*Character
		expectedStats 	[]Skill
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
			expectedStats: []Skill{
				{Name: "slight of hand", SkillModifier: 1, Proficient: false, Attribute: "dexterity"},
				{Name: "persuasion", SkillModifier: 0, Proficient: false, Attribute: "charisma"},
				{Name: "deception", SkillModifier: 0, Proficient: false, Attribute: "charisma"},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Create a deep copy to avoid modifying the original test data
			testChar := &Character {
				Attributes: make([]Attribute, len(test.character.Attributes)),
				Skills: make([]Skill, len(test.character.Skills)),
			}
			copy(testChar.Attributes, test.character.Attributes)
			copy(testChar.Skills, test.character.Skills)

			testChar.calculateSkillModifierFromBase()
			
			for i, e := range test.expectedStats {
				testMod := testChar.Skills[i].SkillModifier
				if e.SkillModifier != testMod {
					t.Errorf("Expected %d modifier, returned %d: %s", e.SkillModifier, testMod, e.Name)
				}
			}
		})
	}
}


func TestCharacter_calculateProficiencyBonus(t *testing.T) {
	tests := []struct {
		name			string
		character		*Character
		expectedBonus 	int 
	}{
		{
			name: "Level 3 Character",
			character: &Character {
				Level: 3,
			},
			expectedBonus: 2,
		},
		{
			name: "Level 8 Character",
			character: &Character {
				Level: 8,
			},
			expectedBonus: 3,
		},
		{
			name: "Level 9 Character",
			character: &Character {
				Level: 9,
			},
			expectedBonus: 4,
		},
		{
			name: "Level 13 Character",
			character: &Character {
				Level: 13,
			},
			expectedBonus: 5,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.character.calculateProficiencyBonusByLevel()

			testMod :=  test.character.Proficiency
			if test.expectedBonus != test.character.Proficiency {
				t.Errorf("Expected Proficiency %d, returned %d: Level %d", test.expectedBonus, testMod, test.character.Level)
			}
		})
	}
}
