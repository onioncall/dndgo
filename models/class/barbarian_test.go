package class

import (
	"os"
	"testing"

	"github.com/onioncall/dndgo/models"
)

func TestBarbarian_unarmoredDefense(t *testing.T) {
	tests := []struct {
		name		string
		character	*models.Character
		expected	int
	}{
		{
			name: "ArmorEquiped, EarlyReturn",
			character: &models.Character {
				AC: 0,
				Attributes: []models.Attribute {
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
				Attributes: []models.Attribute {
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

			barbarian.unarmoredDefense(tt.character)
			result := tt.character.AC

			if tt.expected != result {
				t.Errorf("AC- Expected: %d, Result: %d", tt.expected, result)		
			}
		})
	}
}

func TestBarbarian_primalKnowledge(t *testing.T) {
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
			tt.barbarian.primalKnowledge(tt.character)

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

func TestBarbarian_UseSlots(t *testing.T) {
	tests := []struct {
		name		string
		slotName 	string
		character 	*models.Character
		barbarian	*Barbarian
		expected	Rage
	}{
		{
			name: "One Use, Success",
			slotName: "rage",
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
			// Prevent from writing to terminal during tests
			original := os.Stdout
			os.Stdout, _ = os.Open(os.DevNull)
			defer func() { os.Stdout = original }()

			tt.barbarian.UseClassSlots(tt.slotName)	

			result := tt.barbarian.Rage.Available
			e := tt.expected.Available
			
			if e != result {
				t.Errorf("Rage- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}

func TestBarbarian_RecoverClassSlots(t *testing.T) {
	tests := []struct {
		name		string
		slotName 	string
		recover 	int
		character 	*models.Character
		barbarian	*Barbarian
		expected	Rage
	}{
		{
			name: "Recover By 1",
			slotName: "rage",
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
			slotName: "rage",
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
			tt.barbarian.RecoverClassSlots(tt.slotName, tt.recover)

			result := tt.barbarian.Rage.Available
			e := tt.expected.Available

			if e != result {
				t.Errorf("Rage- Expected: %d\nResult: %d", e, result)
			}
		})
	}
}
