package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/shared"
)

func TestPaladinExecuteFightingStyle(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		paladin   *Paladin
		expected  models.Character
	}{
		{
			name: "Below level requirement",
			character: &models.Character{
				AC:    15,
				Level: 1,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "",
					},
				},
			},
			paladin: &Paladin{
				FightingStyle: shared.FightingStyleDefense,
			},
			expected: models.Character{
				AC:    15,
				Level: 3,
			},
		},
		{
			name: "Not valid fighting style",
			character: &models.Character{
				AC:    15,
				Level: 3,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "",
					},
				},
			},
			paladin: &Paladin{
				FightingStyle: "the-worm",
			},
			expected: models.Character{
				AC:    15,
				Level: 3,
			},
		},
		{
			name: "Defense applied",
			character: &models.Character{
				AC:    15,
				Level: 3,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "",
					},
				},
			},
			paladin: &Paladin{
				FightingStyle: shared.FightingStyleDefense,
			},
			expected: models.Character{
				AC:    16,
				Level: 3,
			},
		},
		{
			name: "Defense not applied (armor equiped)",
			character: &models.Character{
				AC:    15,
				Level: 3,
				WornEquipment: shared.WornEquipment{
					Armor: shared.Armor{
						Name: "Leather Armor",
					},
				},
			},
			paladin: &Paladin{
				FightingStyle: shared.FightingStyleDefense,
			},
			expected: models.Character{
				AC:    15,
				Level: 3,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.paladin.executeFightingStyle(tt.character)
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

func TestPaladinExecutePreparedSpells(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		Paladin   Paladin
		expected  models.Character
	}{
		{
			name: "One Prepared Spell",
			character: &models.Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: shared.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []shared.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			Paladin: Paladin{
				PreparedSpells: []string{
					"Some Spell",
				},
			},
			expected: models.Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: shared.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []shared.CharacterSpell{
					{Name: "Some Spell", IsPrepared: true},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Paladin.executePreparedSpells(tt.character)

			for _, e := range tt.expected.Spells {
				for _, r := range tt.character.Spells {
					if e.Name == r.Name {
						if e.IsPrepared != r.IsPrepared {
							t.Errorf("Spell '%s' Is Prepared- Expected: %t, Result: %t", e.Name, e.IsPrepared, r.IsPrepared)
						}
					}
				}
			}
		})
	}
}

func TestPaladinExecuteOathSpells(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		Paladin   Paladin
		expected  models.Character
	}{
		{
			name: "One Oath Spell",
			character: &models.Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: shared.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []shared.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			Paladin: Paladin{
				OathSpells: []string{
					"Some Spell",
				},
			},
			expected: models.Character{
				Level: 4,
				Abilities: []shared.Ability{
					{Name: shared.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []shared.CharacterSpell{
					{Name: "Some Spell", IsPrepared: true},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.Paladin.executeOathSpells(tt.character)

			for _, e := range tt.expected.Spells {
				for _, r := range tt.character.Spells {
					if e.Name == r.Name {
						if e.IsPrepared != r.IsPrepared {
							t.Errorf("Oath Spell '%s'- Expected: %t, Result: %t", e.Name, e.IsPrepared, r.IsPrepared)
						}
					}
				}
			}
		})
	}
}
