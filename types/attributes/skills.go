package attributes

type Skill struct {
	Ability 		string 	`json:"ability"`
	Name       		string 	`json:"name"`
	SkillModifier	int		`json:"-"`
	Proficient  	bool   	`json:"proficient"`
}

const (
    // Strength
    Athletics string = "athletics"
    
    // Dexterity
    Acrobatics     string = "acrobatics"
    SleightOfHand  string = "sleight of hand"
    Stealth        string = "stealth"
    
    // Intelligence
    Arcana         string = "arcana"
    History        string = "history"
    Investigation  string = "investigation"
    Nature         string = "nature"
    Religion       string = "religion"
    
    // Wisdom
    AnimalHandling string = "animal handling"
    Insight        string = "insight"
    Medicine       string = "medicine"
    Perception     string = "perception"
    Survival       string = "survival"
    
    // Charisma
    Deception      string = "deception"
    Intimidation   string = "intimidation"
    Performance    string = "performance"
    Persuasion     string = "persuasion"
)
