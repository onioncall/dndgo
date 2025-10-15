package class

import (
	"testing"

	"github.com/onioncall/dndgo/models"
)

func TestBarbarianExecuteUnarmoredDefense(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	int
	}{
		{
			name: "ArmorEquiped, EarlyReturn",
			character: &models.Character {
				AC: 0,
				Abilities: []models.Abilities {
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
				BodyEquipment: models.BodyEquipment {
					Armour: "Leather Armor",
				},
			},
			expected: 0,
		},
		{
			name: "No Armor, Valid",
			character: &models.Character {
				AC: 0,
				Abilities: []models.Abilities {
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
				BodyEquipment: models.BodyEquipment {
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
		expected	[]models.Skill
	}{
		{
			name: "Below Level Requirement",
			character: &models.Character {
				Level: 2,
				Proficiency: 2,
				Skills: []models.Skill {
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
			expected: []models.Skill {
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
		expected	[]models.Abilities
	}{
		{
			name: "Below Level Threshold",
			character: &models.Character {
				Level: 15,
				Abilities: []models.Abilities {
					{Name: "Strength", Base: 16},
					{Name: "Constitution", Base: 16},
				},
			},
			expected: []models.Abilities {
				{Name: "Strength", Base: 16},
				{Name: "Constitution", Base: 16},
			},
		},
		{
			name: "Meets Level Requirements, Valid Configuration",
			character: &models.Character {
				Level: 20,
				Abilities: []models.Abilities {
					{Name: "Strength", Base: 17},
					{Name: "Constitution", Base: 17},
				},
			},
			expected: []models.Abilities {
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
			name: "One Use, Success",
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
			name: "Recover By 1",
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
			name: "Full Recover",
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
