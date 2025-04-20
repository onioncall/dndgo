package models

import "fmt"

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
	Backpack          []BackpackItem   `json:"backpack"`
}

type Feat struct {
	Name string `json:"name"`
	Desc string `json:"desc"`
}

type BackpackItem struct {
	Name	 string `json:"name"`
	Quantity int 	`json:"quantity"`
}

type ProficiencyStat struct {
	Name        string `json:"name"`
	Base        int    `json:"base"`
	Adjusted	int
	Bonus		int
	SavingThrowsProficient  bool   `json:"saving-throws-proficient"`
}

type Skill struct {
	Proficiency string `json:"proficiency"`
	Name       	string `json:"name"`
	Bonus		int
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
	Name   		string `json:"name"`
	Bonus  		int    `json:"bonus"`
	Damage 		string `json:"damage"`
	Type   		string `json:"type"`
	Properties 	string `json:"properties"`
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

func (c *Character) BuildHeader() []string {
	header 		:= fmt.Sprintf("# DnD Character\n\n")
	nameLine 	:= fmt.Sprintf("**Name: %s**\n", c.Name)

	s := []string{header, nameLine}
	return s
}

func (c *Character) BuildCharacterInfo() []string {
	levelLine 		:= fmt.Sprintf("Level: %d\n", c.Level)
	classLine 		:= fmt.Sprintf("Class: %s\n", c.Class)
	raceLine 		:= fmt.Sprintf("Race: %s\n", c.Race)
	backgroundLine 	:= fmt.Sprintf("Background: %s\n", c.Background)

	s := []string {
		levelLine,
		classLine,
		raceLine,
		backgroundLine,
	}

	return s
}

func (c *Character) BuildFeats() []string {
	s := make([]string, 0, len(c.Feats) + 1)
	featsLine := fmt.Sprintf("- Feats:\n")
	s = append(s, featsLine)

	for _, feat := range c.Feats {
		featRow := fmt.Sprintf("	- %s: %s\n", feat.Name, feat.Desc)
		s = append(s, featRow)
	}

	return s
}

func (c *Character) BuildLanguages() []string {
	s := make([]string, 0, len(c.Languages) + 1)
	languagesLine := fmt.Sprintf("- Languages:\n")
	s = append(s, languagesLine)

	for _, lang := range c.Languages {
		languageRow := fmt.Sprintf("	- %s\n", lang)
		s = append(s, languageRow)
	}

	return s
}

func (c *Character) BuildGeneralStats() []string {
	nl := "nl"
	profBonusLine	:= fmt.Sprintf("Proficincy Bonus: +%d\n", c.Proficiency)
	passReception	:= fmt.Sprintf("Passive Reception: %d\n", c.PassiveReception)
	passInsight		:= fmt.Sprintf("Passive Insight: %d\n", c.PassiveInsight)

	acLine 			:= fmt.Sprintf("AC: %d\n", c.AC)
	initiativeLine 	:= fmt.Sprintf("Initiative: %d\n", c.Initiative)
	speedLine 		:= fmt.Sprintf("Speed: %d\n", c.Speed)
	// hpMaxLine 		:= fmt.Sprintf("HP Max: %d", c.HPMax)
	hitDiceLine 	:= fmt.Sprintf("Hit Dice: %s\n", c.HitDice)

	s := []string{
		profBonusLine,
		passReception,
		passInsight,
		nl,
		acLine,
		initiativeLine,
		speedLine,
		// hpMaxLine,
		hitDiceLine,
	}

	return s
}

func (c *Character) BuildProficiencies() []string {
	s := make([]string, 0, len(c.Proficiencies) + 3)
	profHeader := fmt.Sprintf("*Proficiencies*\n\n")
	s = append(s, profHeader)

	profTopRow := fmt.Sprintf("| Proficiency  | Base  | Bonus | Saving Throws |\n") 
	profSpacer := fmt.Sprintf("| --- | --- | --- | --- |\n")
	s = append(s, profTopRow)
	s = append(s, profSpacer)

	for _, prof := range c.Proficiencies {
		stBonus := prof.Bonus
		if prof.SavingThrowsProficient {
			stBonus += c.Proficiency	
		}

		profBonusString := ""
		if prof.Bonus > 0 {
			profBonusString = "+"
		}

		stBonusString := ""
		if stBonus > 0 {
			stBonusString = "+"
		}

		stBonusString = fmt.Sprintf("%s%d", stBonusString, stBonus)
		profBonusString = fmt.Sprintf("%s%d", profBonusString, prof.Bonus)

		profRow := fmt.Sprintf("| %s | %d | %s | %s |\n", prof.Name, prof.Bonus, profBonusString, stBonusString)
		s = append(s, profRow)
	}

	return s
}

func (c *Character) BuildSkills() []string {
	s := make([]string, 0, len(c.Skills) + 10)
	skillHeader		:= fmt.Sprintf("*Skills*\n\n")
	s = append(s, skillHeader)

	skillTopRow 	:= fmt.Sprintf("| Skill | Proficiency | Bonus |\n") 
	skillSpacer		:= fmt.Sprintf("| --- | --- | --- |\n")
	s = append(s, skillTopRow)
	s = append(s, skillSpacer)

	for _, skill := range c.Skills {
		if skill.Proficient {
			skill.Bonus += c.Proficiency
		}

		skillBonusString := ""
		if skill.Bonus > 0 {
			skillBonusString = "+"
		}

		skillBonusString = fmt.Sprintf("%s%d", skillBonusString, skill.Bonus)
		skillRow := fmt.Sprintf("| %s | %s | %s |\n", skill.Name, skill.Proficiency, skillBonusString)
		s = append(s, skillRow)
	}

	return s
}

func (c *Character) BuildSpells() []string {
	s := make([]string, 0, len(c.Spells) + 10)
	nl := "\n"
	spellHeader		:= fmt.Sprintf("*Spells*\n\n")
	s = append(s, spellHeader)

	spellTopRow 	:= fmt.Sprintf("| C/S | Slot Level | Ritual | Spell |\n") 
	spellSpacer		:= fmt.Sprintf("| --- | --- | --- | --- |\n")
	s = append(s, spellTopRow)
	s = append(s, spellSpacer)

	for _, spell := range c.Spells {
		scString := "S"
		if spell.IsCaltrop {
			scString = "C"
		}

		rString := " "
		if spell.Ritual {
			rString = "*"
		}

		spellRow := fmt.Sprintf("| %s | %d | %s | %s |\n", scString, spell.SlotLevel, rString, spell.Name)
		s = append(s, spellRow)
	}	
	s = append(s, nl)

	spellSlots := fmt.Sprintf("- Spell Slots\n")
	spellSlot1 := fmt.Sprintf("		- Level 1: %d\n", c.SpellSlots.Level1)
	spellSlot2 := fmt.Sprintf("		- Level 2: %d\n", c.SpellSlots.Level2)
	spellSlot3 := fmt.Sprintf("		- Level 3: %d\n", c.SpellSlots.Level3)
	spellSlot4 := fmt.Sprintf("		- Level 4: %d\n", c.SpellSlots.Level4)
	spellSlot5 := fmt.Sprintf("		- Level 5: %d\n", c.SpellSlots.Level5)
	spellSlot6 := fmt.Sprintf("		- Level 6: %d\n", c.SpellSlots.Level6)
	spellSlot7 := fmt.Sprintf("		- Level 7: %d\n", c.SpellSlots.Level7)
	spellSlot8 := fmt.Sprintf("		- Level 8: %d\n", c.SpellSlots.Level8)
	spellSlot9 := fmt.Sprintf("		- Level 9: %d\n", c.SpellSlots.Level9)
	s = append(s, spellSlots)
	s = append(s, spellSlot1)
	s = append(s, spellSlot2)
	s = append(s, spellSlot3)
	s = append(s, spellSlot4)
	s = append(s, spellSlot5)
	s = append(s, spellSlot6)
	s = append(s, spellSlot7)
	s = append(s, spellSlot8)
	s = append(s, spellSlot9)

	return s
}

func (c *Character) BuildWeapons() []string {
	s := make([]string, 0, len(c.Weapons) + 10)
	weaponsHeader := fmt.Sprintf("*Weapons*\n\n")
	s = append(s, weaponsHeader)

	weaponTopRow := fmt.Sprintf("| Weapon | Bonus | Damage | Type | Properties |\n") 
	weaponSpacer := fmt.Sprintf("| --- | --- | --- | --- | --- |\n")
	s = append(s, weaponTopRow)
	s = append(s, weaponSpacer)

	for _, weapon := range c.Weapons {
		wBonusString := ""
		if weapon.Bonus > 0 {
			wBonusString = "+"
		}

		wBonusString = fmt.Sprintf("%s%d", wBonusString, weapon.Bonus)
		weaponRow := fmt.Sprintf(
			"| %s | %s | %s | %s | %s |\n", 
			weapon.Name, 
			wBonusString, 
			weapon.Damage, 
			weapon.Type, 
			weapon.Properties)
			s = append(s, weaponRow)
	}

	return s
}

func (c *Character) BuildEquipment() []string {
	equipmentHeader := fmt.Sprintf("*Equipment*\n\n")

	bodyEquipment 	:= fmt.Sprintf("- Body Equipment\n")
	head 		:= fmt.Sprintf(" - Head: %s\n", c.BodyEquipment.Head)
	amulet 		:= fmt.Sprintf(" - Amulet: %s\n", c.BodyEquipment.Amulet)
	cloak 		:= fmt.Sprintf(" - Cloak: %s\n", c.BodyEquipment.Cloak)
	armor 		:= fmt.Sprintf(" - Armor: %s\n", c.BodyEquipment.Armour)
	hands 		:= fmt.Sprintf(" - Hands: %s\n", c.BodyEquipment.HandsArms)
	ring 		:= fmt.Sprintf(" - Ring: %s\n", c.BodyEquipment.Ring)
	ring2 		:= fmt.Sprintf(" - Ring: %s\n", c.BodyEquipment.Ring2)
	belt 		:= fmt.Sprintf(" - Belt: %s\n", c.BodyEquipment.Belt)
	boots 		:= fmt.Sprintf(" - Boots: %s\n", c.BodyEquipment.Boots)

	s := []string {
		equipmentHeader,
		bodyEquipment,
		head,
		amulet,
		cloak,
		armor,
		hands,
		ring,
		ring2,
		belt,
		boots,
	}

	return s
}

func (c *Character) BuildBackpack() []string {
	s := make([]string, 0, len(c.Weapons) + 10)
	packHeader := fmt.Sprintf("*Backpack*\n\n")
	s = append(s, packHeader)

	itemTopRow := fmt.Sprintf("| Item | Quantity |\n") 
	itemSpacer := fmt.Sprintf("| --- | --- |\n")
	s = append(s, itemTopRow)
	s = append(s, itemSpacer)

	for _, item := range c.Backpack {
		itemRow := fmt.Sprintf("| %s | %d |\n", item.Name, item.Quantity)
		s = append(s, itemRow)
	}

	return s
}

func (c *Character) AddItemToPack(item string, quantity int) {
	for i, packItem := range c.Backpack {
		if packItem.Name == item {
			c.Backpack[i].Quantity += quantity
			return
		}
	}

	newItem := BackpackItem {
		Name: item,
		Quantity: quantity,
	}

	c.Backpack = append(c.Backpack, newItem)
}

func (c *Character) RemoveItemToPack(item string, quantity int) {
	for i, packItem := range c.Backpack {
		if packItem.Name == item {
			if packItem.Quantity < quantity {
				err := fmt.Sprintf("Quantity to remove (%d) greater than quantity in pack (%d)", quantity, packItem.Quantity)
				panic(err)
			}

			c.Backpack[i].Quantity -= quantity
			return
		}
	}
}
