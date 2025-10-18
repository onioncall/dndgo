package types

type Abilities struct {
	Name        		string 		`json:"name"`
	Base        		int    		`json:"base"`
	Adjusted			int			`json:"-"`
	AbilityModifier		int			`json:"-"`
	SavingThrowsProficient  bool   	`json:"saving-throws-proficient"`
}

type AbilityScoreImprovementItem struct {
	Ability string 	`json:"ability"`
	Bonus	int		`json:"bonus"`
}

const (
	AbilityStrength			string = "strength"
	AbilityDexterity		string = "dexterity"
	AbilityConstitution		string = "constitution"
	AbilityIntelligence		string = "Intelligence"
	AbilityWisdom			string = "wisdom"
	AbilityCharisma			string = "charisma"
)
