package shared

type Ability struct {
	Name                   string `json:"name"`
	Base                   int    `json:"base"` // int 1-20, not to be mutated
	Adjusted               int    `json:"-"`    // int 1-20, value to mutate in code. Is not persisted
	AbilityModifier        int    `json:"-"`    // int between -10 and 10. Derived from Adjusted
	SavingThrowsProficient bool   `json:"saving-throws-proficient" clover:"saving-throws-proficient"`
}

type AbilityScoreImprovementItem struct {
	Ability string `json:"ability"`
	Bonus   int    `json:"bonus"`
}

const (
	AbilityStrength     string = "strength"
	AbilityDexterity    string = "dexterity"
	AbilityConstitution string = "constitution"
	AbilityIntelligence string = "intelligence"
	AbilityWisdom       string = "wisdom"
	AbilityCharisma     string = "charisma"
)
