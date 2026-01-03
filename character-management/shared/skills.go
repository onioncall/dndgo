package shared

type Skill struct {
	Ability       string `json:"ability"`
	Name          string `json:"name"`
	SkillModifier int    `json:"-"`
	Proficient    bool   `json:"proficient"`
}

var Skills = []string{
	SkillAthletics,
	SkillAcrobatics,
	SkillSleightOfHand,
	SkillStealth,
	SkillArcana,
	SkillHistory,
	SkillInvestigation,
	SkillNature,
	SkillReligion,
	SkillAnimalHandling,
	SkillInsight,
	SkillMedicine,
	SkillPerception,
	SkillSurvival,
	SkillDeception,
	SkillIntimidation,
	SkillPerformance,
	SkillPersuasion,
}

const (
	// Strength-based skills
	SkillAthletics string = "athletics"

	// Dexterity-based skills
	SkillAcrobatics    string = "acrobatics"
	SkillSleightOfHand string = "sleight of hand"
	SkillStealth       string = "stealth"

	// Intelligence-based skills
	SkillArcana        string = "arcana"
	SkillHistory       string = "history"
	SkillInvestigation string = "investigation"
	SkillNature        string = "nature"
	SkillReligion      string = "religion"

	// Wisdom-based skills
	SkillAnimalHandling string = "animal handling"
	SkillInsight        string = "insight"
	SkillMedicine       string = "medicine"
	SkillPerception     string = "perception"
	SkillSurvival       string = "survival"

	// Charisma-based skills
	SkillDeception    string = "deception"
	SkillIntimidation string = "intimidation"
	SkillPerformance  string = "performance"
	SkillPersuasion   string = "persuasion"
)
