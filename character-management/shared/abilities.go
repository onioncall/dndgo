package shared

type Ability struct {
	Name                   string `json:"name" clover:"name"`
	Base                   int    `json:"base" clover:"base"` // int 1-20, not to be mutated
	Adjusted               int    `json:"-" clover:"-"`       // int 1-20, value to mutate in code. Is not persisted
	AbilityModifier        int    `json:"-" clover:"-"`       // int between -10 and 10. Derived from Adjusted
	SavingThrowsProficient bool   `json:"saving-throws-proficient" clover:"saving-throws-proficient"`
}

type AbilityScoreImprovementItem struct {
	Ability string `json:"ability" clover:"ability"`
	Bonus   int    `json:"bonus" clover:"bonus"`
}

const (
	AbilityStrength     string = "strength"
	AbilityDexterity    string = "dexterity"
	AbilityConstitution string = "constitution"
	AbilityIntelligence string = "intelligence"
	AbilityWisdom       string = "wisdom"
	AbilityCharisma     string = "charisma"
)
