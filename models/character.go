package models

import (
	"fmt"
	"strings"
)

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
	PassivePerception int              `json:"passive-perception"`
	PassiveInsight    int              `json:"passive-insight"`
	AC                int              `json:"ac"`
	SpellSaveDC       int              `json:"spell-save-dc"`
	HPCurrent		  int			   `json:"hp-current"`
	HPMax			  int			   `json:"hp-max"`
	Initiative        int              `json:"initiative"`
	Speed             int              `json:"speed"`
	HitDice           string           `json:"hit-dice"`
	Proficiencies     []ProficiencyStat`json:"proficiencies"`
	Skills            []Skill          `json:"skills"`
	Spells            []CharacterSpell `json:"spells"`
	SpellSlots        []SpellSlot       `json:"spell-slots"`
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
	IsRitual    bool   `json:"ritual"`
	Name      string `json:"name"`
}

type SpellSlot struct {
	Level		int	`json:"level"`
	Slot		int `json:"slot"`
	Available	int `json:"available"`
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

const (
	Head      string = "head"
	Amulet    string = "amulet"
	Cloak     string = "cloak"
	Armour    string = "armor"
	HandsArms string = "hands-arms"
	Ring      string = "ring"
	Ring2     string = "ring2"
	Belt      string = "belt"
	Boots     string = "boots"
)

// Load Character Details

func (c *Character) CalculateCharacterStats() {
	c.calculateProficiencyBonusByLevel()	
	c.calculateProficienciesFromBase()
	c.calculateSkillBonusFromBase()
}

func (c *Character) calculateProficienciesFromBase() {
	for i, prof := range c.Proficiencies {
		c.Proficiencies[i].Bonus = (prof.Base - 10) / 2
	}
}

func (c *Character) calculateSkillBonusFromBase() {
	for i, skill := range c.Skills {
		// if this is too slow, I'll refactor this to use a map with the proficiency name as the key
		for _, prof := range c.Proficiencies {
			if skill.Proficiency == prof.Name {
				c.Skills[i].Bonus = prof.Bonus 
			}
		}
	}
}

func (c *Character) calculateProficiencyBonusByLevel() {
	if c.Level <= 4 {
		c.Proficiency = 2	
	} else if c.Level > 4 && c.Level <= 8 {
		c.Proficiency = 3	
	} else if c.Level > 8 && c.Level <= 12 {
		c.Proficiency = 4	
	} else if c.Level > 12 && c.Level <= 16 {
		c.Proficiency = 5	
	} else if c.Level > 16 && c.Level <= 20 {
		c.Proficiency = 6	
	}
}

// Format Markdown

func (c *Character) BuildCharacter() string {
	var builder strings.Builder	
	nl := "\n"

	header := c.BuildHeader()
	for i := range header {
		builder.WriteString(header[i])	
	}
	builder.WriteString(nl)

	characterInfo := c.BuildCharacterInfo()
	for i := range characterInfo {
		builder.WriteString(characterInfo[i])	
	}
	builder.WriteString(nl)

	feats := c.BuildFeats()
	for i := range feats {
		builder.WriteString(feats[i])
	}
	builder.WriteString(nl)

	languages := c.BuildLanguages()
	for i := range languages {
		builder.WriteString(languages[i])
	}
	builder.WriteString(nl)

	generalStats := c.BuildGeneralStats()
	for i := range generalStats {
		builder.WriteString(generalStats[i])
	}
	builder.WriteString(nl)
	
	proficiencies := c.BuildProficiencies()
	for i := range proficiencies {
		builder.WriteString(proficiencies[i])
	}
	builder.WriteString(nl)

	skills := c.BuildSkills()
	for i := range skills {
		builder.WriteString(skills[i])
	}
	builder.WriteString(nl)

	spells := c.BuildSpells() 
	for i := range spells {
		builder.WriteString(spells[i]) 
	}
	builder.WriteString(nl)
	
	weapons := c.BuildWeapons()
	for i := range weapons {
		builder.WriteString(weapons[i]) 
	}
	builder.WriteString(nl)

	equipment := c.BuildEquipment()
	for i := range equipment {
		builder.WriteString(equipment[i]) 
	}
	builder.WriteString(nl)

	backpack := c.BuildBackpack()
	for i := range backpack {
		builder.WriteString(backpack[i]) 
	}
	builder.WriteString(nl)

    result := builder.String()
	return result
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
	nl := "\n"
	profBonusLine	:= fmt.Sprintf("Proficincy Bonus: +%d\n", c.Proficiency)
	passPerception	:= fmt.Sprintf("Passive Perception: %d\n", c.PassivePerception)
	passInsight		:= fmt.Sprintf("Passive Insight: %d\n", c.PassiveInsight)

	acLine 			:= fmt.Sprintf("AC: %d\n", c.AC)
	ssdcLine 		:= fmt.Sprintf("Spell Save DC: %d\n", c.SpellSaveDC)
	initiativeLine 	:= fmt.Sprintf("Initiative: %d\n", c.Initiative)
	speedLine 		:= fmt.Sprintf("Speed: %d\n", c.Speed)
	hpMaxLine 		:= fmt.Sprintf("HP: %d/%d\n", c.HPCurrent, c.HPMax)
	hitDiceLine 	:= fmt.Sprintf("Hit Dice: %s\n", c.HitDice)

	s := []string{
		profBonusLine,
		passPerception,
		passInsight,
		nl,
		acLine,
		ssdcLine,
		initiativeLine,
		speedLine,
		hpMaxLine,
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
	s := make([]string, 0, len(c.Spells) + 20)
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
		if spell.IsRitual {
			rString = "*"
		}

		spellRow := fmt.Sprintf("| %s | %d | %s | %s |\n", scString, spell.SlotLevel, rString, spell.Name)
		s = append(s, spellRow)
	}	
	s = append(s, nl)

	spellSlots 		:= fmt.Sprintf("- Spell Slots\n")
	s = append(s, spellSlots)

	for _, spellSlot := range c.SpellSlots {
		slotRow := fmt.Sprintf("	- Level %d: %d/%d\n", spellSlot.Level, spellSlot.Available, spellSlot.Slot)
		s = append(s, slotRow)
	}

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
	head 		:= fmt.Sprintf("	- Head: %s\n", c.BodyEquipment.Head)
	amulet 		:= fmt.Sprintf("	- Amulet: %s\n", c.BodyEquipment.Amulet)
	cloak 		:= fmt.Sprintf("	- Cloak: %s\n", c.BodyEquipment.Cloak)
	armor 		:= fmt.Sprintf("	- Armor: %s\n", c.BodyEquipment.Armour)
	hands 		:= fmt.Sprintf("	- Hands: %s\n", c.BodyEquipment.HandsArms)
	ring 		:= fmt.Sprintf("	- Ring: %s\n", c.BodyEquipment.Ring)
	ring2 		:= fmt.Sprintf("	- Ring: %s\n", c.BodyEquipment.Ring2)
	belt 		:= fmt.Sprintf("	- Belt: %s\n", c.BodyEquipment.Belt)
	boots 		:= fmt.Sprintf("	- Boots: %s\n", c.BodyEquipment.Boots)

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

// CLI Actions

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

func (c *Character) RemoveItemFromPack(item string, quantity int) {
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

func (c *Character) AddLanguage(language string) {
	c.Languages = append(c.Languages, language)
}

func (c *Character) AddEquipment(equipmentType string, equipmentName string) {
	equipmentName = strings.ToLower(equipmentName)
	switch equipmentType {
		case Head:
			c.BodyEquipment.Head = equipmentName
		case Amulet:
			c.BodyEquipment.Amulet = equipmentName
		case Cloak:
			c.BodyEquipment.Cloak = equipmentName
		case Armour:
			c.BodyEquipment.Armour = equipmentName
		case HandsArms:
			c.BodyEquipment.HandsArms = equipmentName
		case Ring:
			c.BodyEquipment.Ring = equipmentName
		case Ring2:
			c.BodyEquipment.Ring2 = equipmentName
		case Belt:
			c.BodyEquipment.Belt = equipmentName
		case Boots:
			c.BodyEquipment.Boots = equipmentName
		default:
			err := fmt.Sprintf("Invalid Equipment Type: %s", equipmentType)
			panic(err)
	}
}

func (c *Character) AddLevel() {
	c.Level += 1
}

func (c *Character) HealCharacter(hpInc int) {
	c.HPCurrent += hpInc

	if c.HPCurrent > c.HPMax {
		c.HPCurrent = c.HPMax
	}
}

func (c *Character) DamageCharacter(hpDecr int) {
	c.HPCurrent -= hpDecr

	if c.HPCurrent < 0 {
		c.HPCurrent = 0
	}
}

func (c *Character) UseSpellSlot(level int) {
	for i := range c.SpellSlots {
		if c.SpellSlots[i].Level == level {
			if c.SpellSlots[i].Available > 0 {
				c.SpellSlots[i].Available--
			}

			return
		}
	}

	panic("invalid level, must be 1-9") 
}

func (c *Character) RecoverSpellSlots(level int) {
	for i := range c.SpellSlots {
		if c.SpellSlots[i].Level == level {
			c.SpellSlots[i].Available = c.SpellSlots[i].Slot
		}
	}
}

func (c *Character) Recover() {
	c.HPCurrent = c.HPMax

	for i := range c.SpellSlots {
		c.SpellSlots[i].Available = c.SpellSlots[i].Slot
	}
}
