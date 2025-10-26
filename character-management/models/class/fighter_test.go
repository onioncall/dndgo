package class

import (
	"testing"

	"github.com/onioncall/dndgo/character-management/types"
)

func TestFighterUseClassTokens(t *testing.T) {
	tests := []struct {
		name      string
		fighter   *Fighter
		tokenName string
		expected  int
	}{
		{
			name: "One removed successfully, mixed case",
			fighter: &Fighter{
				Indomitable: types.NamedToken{
					Name:      "indomitable",
					Available: 1,
					Maximum:   1,
				},
				SecondWind: types.NamedToken{
					Name:      "second-wind",
					Available: 1,
					Maximum:   1,
				},
			},
			tokenName: "Indomitable",
			expected:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fighter.UseClassTokens(tt.tokenName)
			result := tt.fighter.Indomitable.Available

			if tt.expected != result {
				t.Errorf("Token- Expected: %d, Result: %d", tt.expected, result)
			}
		})
	}
}

func TestFighterRecoverClassTokens(t *testing.T) {
	tests := []struct {
		name      string
		fighter   *Fighter
		tokenName string
		quantity  int
		expected  int
	}{
		{
			name: "One recovered successfully, mixed case",
			fighter: &Fighter{
				Indomitable: types.NamedToken{
					Name:      "indomitable",
					Available: 1,
					Maximum:   3,
				},
				SecondWind: types.NamedToken{
					Name:      "second-wind",
					Available: 1,
					Maximum:   1,
				},
			},
			tokenName: "Indomitable",
			quantity:  1,
			expected:  1,
		},
		{
			name: "Multiple recovered successfully, mixed case",
			fighter: &Fighter{
				Indomitable: types.NamedToken{
					Name:      "indomitable",
					Available: 0,
					Maximum:   3,
				},
				SecondWind: types.NamedToken{
					Name:      "second-wind",
					Available: 1,
					Maximum:   1,
				},
			},
			tokenName: "Indomitable",
			quantity:  0,
			expected:  3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fighter.RecoverClassTokens(tt.tokenName, tt.quantity)
			result := tt.fighter.Indomitable.Available

			if tt.expected != result {
				t.Errorf("Token- Expected: %d, Result: %d", tt.expected, result)
			}
		})
	}
}
