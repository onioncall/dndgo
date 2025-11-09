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
				ClassTokens: []types.NamedToken{
					{Name: "indomitable", Available: 1, Maximum: 1, Level: 1},
					{Name: "second-wind", Available: 1, Maximum: 1, Level: 1},
				},
			},
			tokenName: "Indomitable",
			expected:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fighter.UseClassTokens(tt.tokenName, 1)
			for _, token := range tt.fighter.ClassTokens {
				if tt.tokenName == token.Name {
					if tt.expected != token.Available {
						t.Errorf("Token- Expected: %d, Result: %d", tt.expected, token.Available)
					}
				}
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
				ClassTokens: []types.NamedToken{
					{Name: "indomitable", Available: 1, Maximum: 1, Level: 1},
					{Name: "second-wind", Available: 1, Maximum: 1, Level: 1},
				},
			},
			tokenName: "Indomitable",
			quantity:  1,
			expected:  1,
		},
		{
			name: "Multiple recovered successfully, mixed case",
			fighter: &Fighter{
				ClassTokens: []types.NamedToken{
					{Name: "indomitable", Available: 1, Maximum: 1, Level: 1},
					{Name: "second-wind", Available: 1, Maximum: 1, Level: 1},
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
			for _, token := range tt.fighter.ClassTokens {
				if tt.tokenName == token.Name {
					if tt.expected != token.Available {
						t.Errorf("Token- Expected: %d, Result: %d", tt.expected, token.Available)
					}
				}
			}
		})
	}
}
