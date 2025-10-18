package attributes

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
	Strength		string = "strength"
	Dexterity		string = "dexterity"
	Constitution	string = "constitution"
	Intelligence	string = "Intelligence"
	Wisdom			string = "wisdom"
	Charisma		string = "charisma"
)
