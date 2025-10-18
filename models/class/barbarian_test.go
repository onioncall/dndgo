package class

import (
	"testing"

	"github.com/onioncall/dndgo/models"
	attr "github.com/onioncall/dndgo/types/attributes"
	eqmt "github.com/onioncall/dndgo/types/equipment"
)

func TestBarbarianExecuteUnarmoredDefense(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	int
	}{
		{
			name: "Armor equiped, early return",
			character: &models.Character {
				AC: 0,
				Abilities: []attr.Abilities {
					{
						Name: "Strength",
						AbilityModifier: 4,
					},
					{
						Name: "Dexterity",
						AbilityModifier: 3,
					},
					{
						Name: "Constitution",
						AbilityModifier: 2,
					},
					{
						Name: "Intelligence",
						AbilityModifier: 2,
					},
					{
						Name: "Wisdom",
						AbilityModifier: 2,
					},
					{
						Name: "Charisma",
						AbilityModifier: 2,
					},
				},
				BodyEquipment: eqmt.BodyEquipment {
					Armour: "Leather Armor",
				},
			},
			expected: 0,
		},
		{
			name: "No armor, valid",
			character: &models.Character {
				AC: 0,
				Abilities: []attr.Abilities {
					{
						Name: "Strength",
						AbilityModifier: 5,
					},
					{
						Name: "Dexterity",
						AbilityModifier: 3,
					},
					{
						Name: "Constitution",
						AbilityModifier: 4,
					},
					{
						Name: "Intelligence",
						AbilityModifier: 2,
					},
					{
						Name: "Wisdom",
						AbilityModifier: 0,
					},
					{
						Name: "Charisma",
						AbilityModifier: -1,
					},
				},
				BodyEquipment: eqmt.BodyEquipment {
					Armour: "",
				},
			},
			expected: 17,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			barbarian := &Barbarian{}

			barbarian.executeUnarmoredDefense(tt.character)
			result := tt.character.AC

			if tt.expected != result {
				t.Errorf("AC- Expected: %d, Result: %d", tt.expected, result)		
			}
		})
	}
}

func TestBarbarianExecutePrimalKnowledge(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		barbarian	Barbarian
		expected	[]attr.Skill
	}{
		{
			name: "Below level requirement",
			character: &models.Character {
				Level: 2,
				Proficiency: 2,
				Skills: []attr.Skill {
					{Name: "athletics", SkillModifier: 5, Proficient: false},
					{Name: "intimidation", SkillModifier: 4, Proficient: true},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			barbarian: Barbarian {
				PrimalKnowledge: []string {
					"athletics",
				},
			},
			expected: []attr.Skill {
				{Name: "athletics", SkillModifier: 5, Proficient: false},
				{Name: "intimidation", SkillModifier: 4, Proficient: true},
				{Name: "deception", SkillModifier: 3, Proficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.barbarian.executePrimalKnowledge(tt.character)

			result := tt.character.Skills 
			for i, e := range tt.expected {
				if e.Proficient != result[i].Proficient {
					t.Errorf("Skill Proficiency %s- Expected: %t , Result %t",
					e.Name, 
					e.Proficient, 
					result[i].Proficient)
				} 	
			}
		})
	}
}

func TestBarbarianExecutePrimalChampion(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	[]attr.Abilities
	}{
		{
			name: "Below level threshold",
			character: &models.Character {
				Level: 15,
				Abilities: []attr.Abilities {
					{Name: "Strength", Base: 16},
					{Name: "Constitution", Base: 16},
				},
			},
			expected: []attr.Abilities {
				{Name: "Strength", Base: 16},
				{Name: "Constitution", Base: 16},
			},
		},
		{
			name: "Meets level requirements, valid configuration",
			character: &models.Character {
				Level: 20,
				Abilities: []attr.Abilities {
					{Name: "Strength", Base: 17},
					{Name: "Constitution", Base: 17},
				},
			},
			expected: []attr.Abilities {
				{Name: "Strength", Base: 21},
				{Name: "Constitution", Base: 21},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			barbarian := &Barbarian{}
			
			barbarian.executePrimalChampion(tt.character)

			for i, e :=range  tt.expected {
				result := tt.character.Abilities[i]
				if e.Base != result.Base {
					t.Errorf("Ability %s- Expected: %d, Result: %d", e.Name, e.Base, result.Base)
				}
			}
		})
	}
}

func TestBarbarianUseSlots(t *testing.T) {
	tests := []struct {
		name		string
		tokenName 	string
		character 	*models.Character
		barbarian	*Barbarian
		expected	Rage
	}{
		{
			name: "One use, success",
			tokenName: "rage",
			character: &models.Character{},
			barbarian: &Barbarian {
				Rage: Rage {
					Slot: 4,
					Available: 4,
				},
			},
			expected: Rage {
				Slot: 4,
				Available: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.barbarian.UseClassTokens(tt.tokenName)	

			result := tt.barbarian.Rage.Available
			e := tt.expected.Available
			
			if e != result {
				t.Errorf("Rage- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}

func TestBarbarianRecoverClassSlots(t *testing.T) {
	tests := []struct {
		name		string
		tokenName 	string
		recover 	int
		character 	*models.Character
		barbarian	*Barbarian
		expected	Rage
	}{
		{
			name: "Recover by 1",
			tokenName: "rage",
			recover: 1,
			character: &models.Character{},
			barbarian: &Barbarian {
				Rage: Rage {
					Slot: 4,
					Available: 2,
				},
			},
			expected: Rage {
				Slot: 4,
				Available: 3,
			},
		},
		{
			name: "Full recover",
			tokenName: "rage",
			recover: 0,
			barbarian: &Barbarian {
				Rage: Rage {
					Slot: 4,
					Available: 2,
				},
			},
			expected: Rage {
				Slot: 4,
				Available: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.barbarian.RecoverClassTokens(tt.tokenName, tt.recover)

			result := tt.barbarian.Rage.Available
			e := tt.expected.Available

			if e != result {
				t.Errorf("Rage- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}
