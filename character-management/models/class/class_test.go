package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestClassExecuteExpertiseShared(t *testing.T) {
	tests := []struct {
		name            string
		character       *models.Character
		expertiseSkills []string
		expected        []types.Skill
	}{
		{
			name: "Below level requirement",
			character: &models.Character{
				Level:       2,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "dexterity", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			expertiseSkills: []string{
				"persuasion",
				"deception",
			},
			expected: []types.Skill{
				{Name: "dexterity", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 4, Proficient: false},
				{Name: "deception", SkillModifier: 3, Proficient: false},
			},
		},
		{
			name: "Level 3 - two skill proficiencies doubled",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			expertiseSkills: []string{
				"persuasion",
				"deception",
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 3 - two skill proficiencies doubled, one removed",
			character: &models.Character{
				Level:       3,
				Proficiency: 2,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
				},
			},
			expertiseSkills: []string{
				"persuasion",
				"deception",
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 5, Proficient: false},
				{Name: "persuasion", SkillModifier: 6, Proficient: false},
				{Name: "deception", SkillModifier: 5, Proficient: false},
			},
		},
		{
			name: "Level 10, four skill proficiencies doubled",
			character: &models.Character{
				Level:       10,
				Proficiency: 4,
				Skills: []types.Skill{
					{Name: "nature", SkillModifier: 5, Proficient: false},
					{Name: "persuasion", SkillModifier: 4, Proficient: false},
					{Name: "deception", SkillModifier: 3, Proficient: false},
					{Name: "religion", SkillModifier: 2, Proficient: false},
					{Name: "survival", SkillModifier: 4, Proficient: false},
				},
			},
			expertiseSkills: []string{
				"persuasion",
				"deception",
				"nature",
				"religion",
			},
			expected: []types.Skill{
				{Name: "nature", SkillModifier: 9, Proficient: false},
				{Name: "persuasion", SkillModifier: 8, Proficient: false},
				{Name: "deception", SkillModifier: 7, Proficient: false},
				{Name: "religion", SkillModifier: 6, Proficient: false},
				{Name: "survival", SkillModifier: 4, Proficient: false},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executeExpertiseShared(tt.character, tt.expertiseSkills)

			if len(tt.character.Skills) != len(tt.expected) {
				t.Errorf("Skills Count- Expected: %d, Result: %d", len(tt.expected), len(tt.character.Skills))
				return
			}

			for i, e := range tt.expected {
				result := tt.character.Skills[i]

				if e.SkillModifier != result.SkillModifier {
					t.Errorf("Skill Modifier %s- Expected: %d, Result %d", e.Name, e.SkillModifier, result.SkillModifier)
				}
				if e.Proficient != result.Proficient {
					t.Errorf("Skill Proficient %s- Expected: %t, Result %t", e.Name, e.Proficient, result.Proficient)
				}
			}
		})
	}
}

func TestClassExecuteSpellCastingAbility(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  models.Character
	}{
		{
			name: "Abiltity mod +2, proficiency +2",
			character: &models.Character{
				Level:       4,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				SpellSaveDC:    0,
				SpellAttackMod: 0,
			},
			expected: models.Character{
				Level:       4,
				Proficiency: 2,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				SpellSaveDC:    12,
				SpellAttackMod: 4,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wizard := &Wizard{}
			wizard.executeSpellCastingAbility(tt.character)

			expectedDC := tt.expected.SpellSaveDC
			expectedAttackMod := tt.expected.SpellAttackMod
			resultDC := tt.character.SpellSaveDC
			resultAttackMod := tt.character.SpellAttackMod

			if expectedDC != resultDC {
				t.Errorf("Spell Save DC- Expected: %d, Result: %d", expectedDC, resultDC)
			}

			if expectedAttackMod != resultAttackMod {
				t.Errorf("Spell Attack Mod- Expected: %d, Result: %d", expectedAttackMod, resultAttackMod)
			}
		})
	}
}

func TestClassExecutePreparedSpells(t *testing.T) {
	tests := []struct {
		name           string
		character      *models.Character
		preparedSpells []string
		expected       models.Character
	}{
		{
			name: "No Prepared Spells",
			character: &models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			preparedSpells: []string{},
			expected: models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
		},
		{
			name: "One Prepared Spell",
			character: &models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: false},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
			preparedSpells: []string{
				"Some Spell",
			},
			expected: models.Character{
				Level: 4,
				Abilities: []types.Abilities{
					{Name: types.AbilityIntelligence, AbilityModifier: 2},
				},
				Spells: []types.CharacterSpell{
					{Name: "Some Spell", IsPrepared: true},
					{Name: "Different Spell", IsPrepared: false},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			executePreparedSpellsShared(tt.character, tt.preparedSpells)

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

func TestClassAppliedArchery(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  []types.Weapon
		applied   bool
	}{
		{
			name: "No ranged weapon",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
					{Name: "Dagger", Bonus: 2, Damage: "1d4", Ranged: false},
				},
			},
			expected: []types.Weapon{
				{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
				{Name: "Dagger", Bonus: 2, Damage: "1d4", Ranged: false},
			},
			applied: false,
		},
		{
			name: "Range bonus applied",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true},
				},
			},
			expected: []types.Weapon{
				{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
				{Name: "Longbow", Bonus: 4, Damage: "1d8", Ranged: true},
			},
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyArchery(tt.character).IsApplied
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

func TestClassAppliedDefense(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  int
		applied   bool
	}{
		{
			name: "Armor equiped, early return",
			character: &models.Character{
				AC: 15,
				WornEquipment: types.WornEquipment{
					Armor: types.Armor{
						Name: "Leather Armor",
					},
				},
			},
			expected: 15,
			applied:  false,
		},
		{
			name: "Armor not equiped, bonus added",
			character: &models.Character{
				AC: 15,
				WornEquipment: types.WornEquipment{
					Armor: types.Armor{
						Name: "",
					},
				},
			},
			expected: 16,
			applied:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyDefense(tt.character).IsApplied
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

func TestClassAppliedDueling(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  []types.Weapon
		applied   bool
	}{
		{
			name: "No melee weapon",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true},
				},
				PrimaryEquipped: "Longbow",
			},
			expected: []types.Weapon{
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true},
			},
			applied: false,
		},
		{
			name: "Melee bonus applied",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
				},
				PrimaryEquipped: "Club",
			},
			expected: []types.Weapon{
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Ranged: false},
			},
			applied: true,
		},
		{
			name: "Multiple valid weapons, one bonus",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
				},
				PrimaryEquipped: "Club",
			},
			expected: []types.Weapon{
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Ranged: false},
				{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
			},
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyDueling(tt.character).IsApplied
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

func TestClassAppliedTwoWeaponFighting(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		expected  []types.Weapon
		applied   bool
	}{
		{
			name: "No applicable weapons, bonus not applied",
			character: &models.Character{
				Abilities: []types.Abilities{
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon{
					{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true, Properties: []string{"two-handed"}},
				},
				PrimaryEquipped: "Greataxe",
			},
			expected: []types.Weapon{
				{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true, Properties: []string{"two-handed"}},
			},
			applied: false,
		},
		{
			name: "One applicable weapon, bonus not applied",
			character: &models.Character{
				Abilities: []types.Abilities{
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon{
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true, Properties: []string{"two-handed"}},
				},
				PrimaryEquipped: "Longbow",
			},
			expected: []types.Weapon{
				{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
				{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true, Properties: []string{"two-handed"}},
			},
			applied: false,
		},
		{
			name: "Two applicable light weapons, bonus applied",
			character: &models.Character{
				Abilities: []types.Abilities{
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon{
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
				},
				PrimaryEquipped:   "Club",
				SecondaryEquipped: "Club",
			},
			expected: []types.Weapon{
				{Name: "Club", Bonus: 4, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
				{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
			},
			applied: true,
		},
		{
			name: "Two applicable weapons, one light, both one handed, bonus applied",
			character: &models.Character{
				Abilities: []types.Abilities{
					{Name: "Dexterity", Base: 14, AbilityModifier: 2},
				},
				Weapons: []types.Weapon{
					{Name: "Rapier", Bonus: 2, Damage: "1d8", Ranged: false, Properties: []string{"finesse"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
				},
				PrimaryEquipped:   "Rapier",
				SecondaryEquipped: "Club",
			},
			expected: []types.Weapon{
				{Name: "Rapier", Bonus: 2, Damage: "1d8", Ranged: false, Properties: []string{"finesse"}},
				{Name: "Club", Bonus: 4, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
			},
			applied: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyTwoWeaponFighting(tt.character).IsApplied
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

func TestClassAppliedGreatWeaponFighting(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		applied   bool
	}{
		{
			name: "Two handed equipped, bonus applied",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Greataxe", Bonus: 2, Damage: "1d12", Ranged: false, Properties: []string{"two-handed"}},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true, Properties: []string{"two-handed"}},
				},
				PrimaryEquipped: "Greataxe",
			},
			applied: true,
		},
		{
			name: "Two handed secondary equipped, bonus applied",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false},
					{Name: "Longbow", Bonus: 2, Damage: "1d8", Ranged: true, Properties: []string{"two-handed"}},
				},
				SecondaryEquipped: "Longbow",
			},
			applied: true,
		},
		{
			name: "No applicable weapons, bonus not applied",
			character: &models.Character{
				Weapons: []types.Weapon{
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
					{Name: "Club", Bonus: 2, Damage: "1d4", Ranged: false, Properties: []string{"light"}},
				},
				PrimaryEquipped:   "Club",
				SecondaryEquipped: "Club",
			},
			applied: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyGreatWeaponFighting(tt.character).IsApplied

			if tt.applied != returned {
				t.Errorf("Not Applied Correctly- Expected: %t, Result: %t", tt.applied, returned)
			}
		})
	}
}

func TestClassAppliedProtection(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		applied   bool
	}{
		{
			name: "Shield equipped, bonus applied",
			character: &models.Character{
				PrimaryEquipped: "Some Shield",
				WornEquipment: types.WornEquipment{
					Shield: "Some Shield",
				},
			},
			applied: true,
		},
		{
			name: "Shield not equipped, bonus not applied",
			character: &models.Character{
				PrimaryEquipped:   "",
				SecondaryEquipped: "",
				WornEquipment: types.WornEquipment{
					Shield: "Some Shield",
				},
			},
			applied: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			returned := applyProtection(tt.character).IsApplied

			if tt.applied != returned {
				t.Errorf("Not Applied Correctly- Expected: %t, Result: %t", tt.applied, returned)
			}
		})
	}
}
