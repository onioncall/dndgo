package models

type Character struct {
	Path              string           `json:"path"`
	Name              string           `json:"name"`
	Level             int              `json:"level"`
	Class             string           `json:"class"`
	Race              string           `json:"race"`
	Background        string           `json:"background"`
	Feats             []Feat           `json:"feats"`
	Languages         []string         `json:"languages"`
	Proficiency       int              `json:"proficiency"`
	PassiveReception  int              `json:"passive-reception"`
	PassiveInsight    int              `json:"passive-insight"`
	AC                int              `json:"ac"`
	// HPMax			  int			   `json:"hp-max"`
	Initiative        int              `json:"initiative"`
	Speed             int              `json:"speed"`
	HitDice           string           `json:"hit-dice"`
	Proficiencies     []ProficiencyStat`json:"proficiencies"`
	Skills            []Skill          `json:"skills"`
	Spells            []CharacterSpell `json:"spells"`
	SpellSlots        SpellSlots       `json:"spell-slots"`
	Weapons           []Weapon         `json:"weapons"`
	BodyEquipment     BodyEquipment    `json:"body-equipment"`
	Backpack          []string         `json:"backpack"`
}

type Feat struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type ProficiencyStat struct {
	Name        string `json:"name"`
	Base        int    `json:"base"`
	Adjusted	int
	Bonus		string
	Proficient  bool   `json:"proficient"`
}

type Skill struct {
	Proficiency string `json:"proficiency"`
	Skill       string `json:"skill"`
	Proficient  bool   `json:"proficient"`
}

type CharacterSpell struct {
	IsCaltrop bool   `json:"is-caltrop"`
	SlotLevel int    `json:"slot-level"`
	Ritual    bool   `json:"ritual"`
	Name      string `json:"name"`
}

type SpellSlots struct {
	Level1 int `json:"level1"`
	Level2 int `json:"level2"`
	Level3 int `json:"level3"`
	Level4 int `json:"level4"`
	Level5 int `json:"level5"`
	Level6 int `json:"level6"`
	Level7 int `json:"level7"`
	Level8 int `json:"level8"`
	Level9 int `json:"level9"`
}

type Weapon struct {
	Name   string `json:"name"`
	Bonus  int    `json:"bonus"`
	Damage string `json:"damage"`
	Type   string `json:"type"`
}

type BodyEquipment struct {
	Head      string `json:"head"`
	Amulet    string `json:"amulet"`
	Cloak     string `json:"cloak"`
	Armour    string `json:"armour"`
	HandsArms string `json:"hands-arms"`
	Ring      string `json:"ring"`
	Ring2     string `json:"ring2"`
	Belt      string `json:"belt"`
	Boots     string `json:"boots"`
}
