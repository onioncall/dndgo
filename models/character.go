package models

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/logger"
)

type Character struct {
	Path              		string           	`json:"path"`
	Name              		string           	`json:"name"`
	Level             		int              	`json:"level"`
	ClassName         		string           	`json:"class-name"`
	Race              		string           	`json:"race"`
	Background        		string           	`json:"background"`
	Feats			  		[]GenericItem		`json:"feats"`
	Languages         		[]string         	`json:"languages"`
	Proficiency       		int
	PassivePerception 		int              	`json:"passive-perception"`
	PassiveInsight    		int              	`json:"passive-insight"`
	AC                		int              	`json:"ac"`
	SpellSaveDC       		int              	`json:"spell-save-dc"`
	HPCurrent		  		int			   		`json:"hp-current"`
	HPMax			  		int			   		`json:"hp-max"`
	Initiative        		int              	`json:"initiative"`
	Speed             		int              	`json:"speed"`
	HitDice           		string           	`json:"hit-dice"`
	Attributes     	  		[]Attribute			`json:"attributes"`
	Skills            		[]Skill          	`json:"skills"`
	Spells            		[]CharacterSpell 	`json:"spells"`
	SpellSlots        		[]SpellSlot       	`json:"spell-slots"`
	Weapons           		[]Weapon         	`json:"weapons"`
	BodyEquipment			BodyEquipment    	`json:"body-equipment"`
	Backpack          		[]BackpackItem   	`json:"backpack"`
	AbilityScoreImprovement []AbilityScoreImprovementItem `json:"ability-score-improvement"`
	Class			  		IClass				`json:"-"`	
}

type IClass interface {
	LoadMethods()
	ExecutePostCalculateMethods(c *Character)
	ExecutePreCalculateMethods(c *Character)
	PrintClassDetails(c *Character) []string
	UseClassSlots(string)
	RecoverClassSlots(string, int)
}

type GenericItem struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

type BackpackItem struct {
	Name	 string `json:"name"`
	Quantity int 	`json:"quantity"`
}

type Attribute struct {
	Name        		string `json:"name"`
	Base        		int    `json:"base"`
	Adjusted			int
	AbilityModifier		int
	SavingThrowsProficient  bool   `json:"saving-throws-proficient"`
}

type AbilityScoreImprovementItem struct {
	Ability string 	`json:"ability"`
	Bonus	int		`json:"bonus"`
}

type Skill struct {
	Attribute 		string `json:"attribute"`
	Name       		string `json:"name"`
	SkillModifier	int
	Proficient  	bool   `json:"proficient"`
}

type CharacterSpell struct {
	IsCaltrop bool   `json:"is-caltrop"`
	SlotLevel int    `json:"slot-level"`
	IsRitual    bool   `json:"ritual"`
	Name      string `json:"name"`
}

type ClassFeatures struct {
	Name 	string 	`json:"name"`
	Level	int		`json:"level"`
	Details string 	`json:"details"`
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

var PreCalculateMethods []func(c *Character)
var PostCalculateMethods []func(c *Character)

// Load Character Details

func (c *Character) CalculateCharacterStats() {
	c.calculateProficiencyBonusByLevel()	
	c.calculateAttributesFromBase()
	c.calculateAbilityScoreImprovement()
	c.calculateSkillModifierFromBase()
}

func (c *Character) calculateAttributesFromBase() {
	for i, a := range c.Attributes {
		c.Attributes[i].AbilityModifier = (a.Base - 10) / 2
	}
}

func (c *Character) calculateSkillModifierFromBase() {
	for i, skill := range c.Skills {
		// if this is too slow, I'll refactor this to use a map with the proficiency name as the key
		for _, a := range c.Attributes {
			if strings.ToLower(skill.Attribute) == strings.ToLower(a.Name) {
				c.Skills[i].SkillModifier = a.AbilityModifier 
			}
		}
	}
}

// At level 4, select one ability and add 2 to that score,
// or select two abilities and add 1 to each score (max of 20).
// They get to do this again at levels 8, 12, 16, and 19
func (c *Character) calculateAbilityScoreImprovement() {
	maxBonus := 0
	switch {
	case c.Level < 4:  maxBonus = 0
	case c.Level < 8:  maxBonus = 2
	case c.Level < 12: maxBonus = 4
	case c.Level < 16: maxBonus = 6
	case c.Level < 19: maxBonus = 8
	case c.Level >= 19: maxBonus = 10
	}

	if maxBonus == 0 {
		return // don't qualify yet
	}

	bonusSum := 0
	for _, item := range c.AbilityScoreImprovement {
		bonusSum += item.Bonus
	}

	if bonusSum > maxBonus {
		fmt.Printf("Ability Score Bonus (%d) exceeds available for level (%d)\n", bonusSum, maxBonus)
		return
	}
	
	for _, item := range c.AbilityScoreImprovement {
		for i := range c.Attributes {
			if strings.ToLower(c.Attributes[i].Name) == strings.ToLower(item.Ability) {
				c.Attributes[i].Base += item.Bonus
				c.Attributes[i].Base = min(20, c.Attributes[i].Base)
				break
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
	
	proficiencies := c.BuildAttributes()
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

	abilityScoreImprovement := c.BuildAbilityScoreImprovement()
	for i := range abilityScoreImprovement {
		builder.WriteString(abilityScoreImprovement[i]) 
	}
	builder.WriteString(nl)

	if c.Class != nil {
		otherClassFeatures := c.Class.PrintClassDetails(c)
		for i := range otherClassFeatures {
			builder.WriteString(otherClassFeatures[i]) 
		}
		builder.WriteString(nl)
	}

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
	classLine 		:= fmt.Sprintf("Class: %s\n", c.ClassName)
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
	featLine := fmt.Sprintf("- Feats:\n")
	s = append(s, featLine)

	for _, feat := range c.Feats {
		featRow := fmt.Sprintf("	- %s: %s\n", feat.Name, feat.Desc)
		s = append(s, featRow)
	}
	s = append(s, "---")

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
	proficiency		:= fmt.Sprintf("Proficiency: +%d\n", c.Proficiency)
	passPerception	:= fmt.Sprintf("Passive Perception: %d\n", c.PassivePerception)
	passInsight		:= fmt.Sprintf("Passive Insight: %d\n", c.PassiveInsight)

	acLine 			:= fmt.Sprintf("AC: %d\n", c.AC)
	ssdcLine 		:= fmt.Sprintf("Spell Save DC: %d\n", c.SpellSaveDC)
	initiativeLine 	:= fmt.Sprintf("Initiative: %d\n", c.Initiative)
	speedLine 		:= fmt.Sprintf("Speed: %d\n", c.Speed)
	hpMaxLine 		:= fmt.Sprintf("HP: %d/%d\n", c.HPCurrent, c.HPMax)
	hitDiceLine 	:= fmt.Sprintf("Hit Dice: %s\n", c.HitDice)

	s := []string{
		proficiency,
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

func (c *Character) BuildAttributes() []string {
	s := make([]string, 0, len(c.Attributes) + 3)
	profHeader := fmt.Sprintf("*Attributes*\n\n")
	s = append(s, profHeader)

	profTopRow := fmt.Sprintf("| Proficiency  | Base  | Modifier | Saving Throws |\n") 
	profSpacer := fmt.Sprintf("| --- | --- | --- | --- |\n")
	s = append(s, profTopRow)
	s = append(s, profSpacer)

	for _, attr := range c.Attributes {
		abMod := attr.AbilityModifier
		if attr.SavingThrowsProficient {
			abMod += c.Proficiency	
		}

		abBaseString := ""
		if attr.AbilityModifier > 0 {
			abBaseString = "+"
		}

		abModString := ""
		if abMod > 0 {
			abModString = "+"
		}

		abModString = fmt.Sprintf("%s%d", abModString, abMod)
		abBaseString = fmt.Sprintf("%s%d", abBaseString, attr.AbilityModifier)

		profRow := fmt.Sprintf("| %s | %d | %s | %s |\n", attr.Name, attr.Base, abBaseString, abModString)
		s = append(s, profRow)
	}

	return s
}

func (c *Character) BuildSkills() []string {
	s := make([]string, 0, len(c.Skills) + 10)
	skillHeader		:= fmt.Sprintf("*Skills*\n\n")
	s = append(s, skillHeader)

	skillTopRow 	:= fmt.Sprintf("| Skill | Attribute | Modifier |\n") 
	skillSpacer		:= fmt.Sprintf("| --- | --- | --- |\n")
	s = append(s, skillTopRow)
	s = append(s, skillSpacer)

	for _, skill := range c.Skills {
		if skill.Proficient {
			skill.SkillModifier += c.Proficiency
		}

		skillModifierString := ""
		if skill.SkillModifier > 0 {
			skillModifierString = "+"
		}

		skillModifierString = fmt.Sprintf("%s%d", skillModifierString, skill.SkillModifier)
		skillRow := fmt.Sprintf("| %s | %s | %s |\n", skill.Name, skill.Attribute, skillModifierString)
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
		fullCircle := strings.Repeat("● ", spellSlot.Available)
		hollowCircle := strings.Repeat("○ ", (spellSlot.Slot - spellSlot.Available))
		slotRow := fmt.Sprintf("	- Level %d: %s%s\n", spellSlot.Level, fullCircle, hollowCircle)
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

func (c *Character) BuildClassDetailsHeader() []string {
	s := make([]string, 0, 100)	
	header := fmt.Sprintf("Class Details\n")
	spacer := fmt.Sprintf("---\n")
	s = append(s, header)
	s = append(s, spacer)

	return s
}

func (c *Character) BuildAbilityScoreImprovement() []string {
	s := make([]string, 0, len(c.AbilityScoreImprovement))
	if len(c.AbilityScoreImprovement) > 0 && c.Level >= 4 {
		abilityMap := make(map[string]int)

		// Build map
		for _, ability := range c.AbilityScoreImprovement {
			_, exists := abilityMap[ability.Ability]
			if exists {
				abilityMap[ability.Ability] += ability.Bonus
			}
		}

		abilityScoreImprovementHeader := fmt.Sprintf("Ability Score Improvement\n")
		s = append(s, abilityScoreImprovementHeader)

		// TODO: Make this more sophistocated so we don't need to loop through this twice
		for _, ability := range c.AbilityScoreImprovement {
			expLine := fmt.Sprintf("- %s: +%d\n", ability.Ability, ability.Bonus)
			s = append(s, expLine)
		}
		s = append(s, "\n")
	}

	return s
}

func (c *Character) GetSlots(available int, max int) string {
	fullCircle := strings.Repeat("● ", available)
	hollowCircle := strings.Repeat("○ ", (max - available))

	return fmt.Sprintf("%s%s", fullCircle, hollowCircle)
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
				info := fmt.Sprintf("Quantity to remove (%d) greater than quantity in pack (%d)", quantity, packItem.Quantity)
				logger.HandleInfo(info)
			}

			c.Backpack[i].Quantity -= quantity
			return
		}
	}

	msg := fmt.Sprintf("Item %s not found in pack", item)
	fmt.Println(msg)
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
			info := fmt.Sprintf("Invalid Equipment Type: %s", equipmentType)
			logger.HandleInfo(info)
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
	if c.HPCurrent <= 0 {
		logger.HandleInfo("Character had no health left")
		return
	}

	c.HPCurrent -= hpDecr

	// reset to zero if the decremented amount is greater than remaining health
	if c.HPCurrent < 0 {
		c.HPCurrent = 0
	}
}

func (c *Character) UseSpellSlot(level int) {
	for i := range c.SpellSlots {
		if c.SpellSlots[i].Level == level {
			if c.SpellSlots[i].Available <= 0 {
				info := fmt.Sprintf("Spell Slot Level %d: already at zero", level)
				logger.HandleInfo(info)

				return
			}

			c.SpellSlots[i].Available--
			return
		}
	}

	logger.HandleInfo("Invalid level, must be 1-9") 
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

	if c.Class != nil {
		c.Class.RecoverClassSlots("", 0)
	}
}

func (c *Character) UseClassSlots(name string) {
	c.Class.UseClassSlots(name)
}

func (c *Character) RecoverClassSlots(name string, quantity int) {
	c.Class.RecoverClassSlots(name, quantity)
}
