package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/character-management/types"
)

func TestRangerExecuteFightingStyle(t *testing.T) {
	tests := []struct {
		name      string
		character *models.Character
		ranger    *Ranger
		expected  models.Character
	}{
		{
			name: "Below level requirement",
			character: &models.Character{
				AC:    15,
				Level: 1,
				WornEquipment: types.WornEquipment{
					Armor: types.Armor{
						Name: "",
					},
				},
			},
			ranger: &Ranger{
				FightingStyle: types.FightingStyleDefense,
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
				WornEquipment: types.WornEquipment{
					Armor: types.Armor{
						Name: "",
					},
				},
			},
			ranger: &Ranger{
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
				WornEquipment: types.WornEquipment{
					Armor: types.Armor{
						Name: "",
					},
				},
			},
			ranger: &Ranger{
				FightingStyle: types.FightingStyleDefense,
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
				WornEquipment: types.WornEquipment{
					Armor: types.Armor{
						Name: "Leather Armor",
					},
				},
			},
			ranger: &Ranger{
				FightingStyle: types.FightingStyleDefense,
			},
			expected: models.Character{
				AC:    15,
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
