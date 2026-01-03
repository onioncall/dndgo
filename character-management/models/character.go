package models

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/character-management/shared"
	"github.com/onioncall/dndgo/logger"
)

type Character struct {
	ID                      string                               `json:"id" clover:"id"`
	Default                 bool                                 `json:"default" clover:"default"`
	Path                    string                               `json:"path" clover:"path"`
	ValidationDisabled      bool                                 `json:"validation-disabled" clover:"validation-disabled"`
	Name                    string                               `json:"name" clover:"name"`
	Level                   int                                  `json:"level" clover:"level"`
	ClassName               string                               `json:"class-name" clover:"class-name"`
	Race                    string                               `json:"race" clover:"race"`
	Background              string                               `json:"background" clover:"background"`
	Feats                   []GenericItem                        `json:"feats" clover:"feats"`
	Languages               []string                             `json:"languages" clover:"languages"`
	Proficiency             int                                  `json:"-" clover:"-"`
	PassivePerception       int                                  `json:"-" clover:"-"`
	PassiveInsight          int                                  `json:"-" clover:"-"`
	AC                      int                                  `json:"-" clover:"-"`
	SpellSaveDC             int                                  `json:"-" clover:"-"`
	SpellAttackMod          int                                  `json:"-" clover:"-"`
	HPCurrent               int                                  `json:"hp-current" clover:"hp-current"`
	HPMax                   int                                  `json:"hp-max" clover:"hp-max"`
	HPTemp                  int                                  `json:"hp-temp" clover:"hp-temp"`
	Speed                   int                                  `json:"speed" clover:"speed"`
	HitDice                 string                               `json:"-" clover:"-"`
	Abilities               []shared.Ability                     `json:"abilities" clover:"abilities"`
	Skills                  []shared.Skill                       `json:"skills" clover:"skills"`
	Spells                  []shared.CharacterSpell              `json:"spells" clover:"spells"`
	SpellSlots              []shared.SpellSlot                   `json:"spell-slots" clover:"spell-slots"`
	Weapons                 []shared.Weapon                      `json:"weapons" clover:"weapons"`
	PrimaryEquipped         string                               `json:"primary-equipped" clover:"primary-equipped"`
	SecondaryEquipped       string                               `json:"secondary-equipped" clover:"secondary-equipped"`
	WornEquipment           shared.WornEquipment                 `json:"worn-equipment" clover:"worn-equipment"`
	Backpack                []shared.BackpackItem                `json:"backpack" clover:"backpack"`
	AbilityScoreImprovement []shared.AbilityScoreImprovementItem `json:"ability-score-improvement" clover:"ability-score-improvement"`
	Class                   Class                                `json:"-" clover:"-"`
}

type GenericItem struct {
	Name string `json:"name"`
	Desc string `json:"description"`
}

var (
	PreCalculateMethods  []func(c *Character)
	PostCalculateMethods []func(c *Character)
)

// Load Character Details

func (c *Character) CalculateCharacterStats() {
	c.calculateProficiencyBonusByLevel()
	c.calculateAbilityScoreImprovement()
	c.calculateAbilitiesFromBase()
	c.calculateSkillModifierFromBase()
	c.calculateAC()
	c.calculatePassiveStats()
	c.calculateWeaponBonus()
}

func (c *Character) calculateAbilitiesFromBase() {
	for i, a := range c.Abilities {
		c.Abilities[i].AbilityModifier = (a.Base - 10) / 2
	}
}

func (c *Character) calculateSkillModifierFromBase() {
	for i, skill := range c.Skills {
		// if this is too slow, I'll refactor this to use a map with the proficiency name as the key
		for _, a := range c.Abilities {
			if strings.EqualFold(skill.Ability, a.Name) {
				c.Skills[i].SkillModifier = a.AbilityModifier

				if skill.Proficient {
					c.Skills[i].SkillModifier += c.Proficiency
				}
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
	case c.Level < 4:
		maxBonus = 0
	case c.Level < 8:
		maxBonus = 2
	case c.Level < 12:
		maxBonus = 4
	case c.Level < 16:
		maxBonus = 6
	case c.Level < 19:
		maxBonus = 8
	case c.Level >= 19:
		maxBonus = 10
	}

	bonusSum := 0
	for _, item := range c.AbilityScoreImprovement {
		bonusSum += item.Bonus
	}

	if bonusSum > maxBonus {
		info := fmt.Sprintf("Ability Score Bonus (%d) exceeds available for level (%d)\n", bonusSum, maxBonus)
		logger.Info(info)
	}

	for _, item := range c.AbilityScoreImprovement {
		for i := range c.Abilities {
			if strings.EqualFold(c.Abilities[i].Name, item.Ability) {
				c.Abilities[i].Base += item.Bonus
				c.Abilities[i].Base = min(20, c.Abilities[i].Base)
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

func (c *Character) calculateAC() {
	c.AC = c.WornEquipment.Armor.Class

	switch c.WornEquipment.Armor.Type {
	case shared.LightArmor:
		c.AC += c.GetMod(shared.AbilityDexterity)
	case shared.MediumArmor:
		c.AC += min(c.GetMod(shared.AbilityDexterity), 2)
	case shared.HeavyArmor:
		// this just uses the armor class which we've already accounted for
	case "":
		c.AC += 10 + c.GetMod(shared.AbilityDexterity)
	}

	if c.WornEquipment.Shield != "" {
		if strings.EqualFold(c.PrimaryEquipped, c.WornEquipment.Shield) ||
			strings.EqualFold(c.SecondaryEquipped, c.WornEquipment.Shield) {
			c.AC += 2
		}
	}
}

func (c *Character) calculatePassiveStats() {
	wisMod := c.GetMod(shared.AbilityWisdom)
	c.PassivePerception = 10 + wisMod
	c.PassiveInsight = 10 + wisMod

	for _, skill := range c.Skills {
		if strings.ToLower(skill.Name) == shared.SkillPerception {
			if skill.Proficient {
				c.PassivePerception += c.Proficiency
			}
		} else if strings.ToLower(skill.Name) == shared.SkillInsight {
			if skill.Proficient {
				c.PassiveInsight += c.Proficiency
			}
		}
	}
}

func (c *Character) calculateWeaponBonus() {
	for i, weapon := range c.Weapons {
		c.Weapons[i].Bonus = 0
		dexMod := c.GetMod(shared.AbilityDexterity)
		strMod := c.GetMod(shared.AbilityStrength)
		modApplied := false

		for _, prop := range weapon.Properties {
			// We'll prioritize the weapon properties that have a choice between dex and str mods
			prop = strings.ToLower(prop)
			if prop == shared.WeaponPropertyFinesse || prop == shared.WeaponPropertyThrown {
				c.Weapons[i].Bonus += max(dexMod, strMod)
				modApplied = true
			}
		}

		// If we haven't applied our mod from properties, we'll apply it based on range
		if !modApplied {
			if weapon.Ranged {
				c.Weapons[i].Bonus += dexMod
			} else {
				c.Weapons[i].Bonus += strMod
			}
		}

		if weapon.Proficient {
			c.Weapons[i].Bonus += c.Proficiency
		}

		// Since custom weapons are sometimes a thing, we'll allow the user to specify a custom bonus
		c.Weapons[i].Bonus += weapon.CustomBonus
	}
}

func (c *Character) GetMod(abilityName string) int {
	for _, ability := range c.Abilities {
		if strings.EqualFold(ability.Name, abilityName) {
			return ability.AbilityModifier
		}
	}

	return 0
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

	proficiencies := c.BuildAbilities()
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
		builder.WriteString("Class Details\n")
		builder.WriteString("---\n")

		subClass := c.Class.GetSubClass(c.Level)
		if subClass != "" {
			builder.WriteString(fmt.Sprintf("Sub-Class: %s\n\n", subClass))
		}

		details := c.Class.ClassDetails(c.Level)
		if details != "" {
			builder.WriteString(details + "\n")
		}

		classFeatures := c.Class.GetClassFeatures(c.Level)
		if classFeatures != "" {
			builder.WriteString(classFeatures + "\n")
		}
		builder.WriteString(nl)
	}

	result := builder.String()
	return result
}

func (c *Character) BuildHeader() []string {
	header := "# DnD Character\n\n"
	nameLine := fmt.Sprintf("**Name: %s**\n", c.Name)

	s := []string{header, nameLine}
	return s
}

func (c *Character) BuildCharacterInfo() []string {
	levelLine := fmt.Sprintf("Level: %d\n", c.Level)
	classLine := fmt.Sprintf("Class: %s\n", c.ClassName)
	raceLine := fmt.Sprintf("Race: %s\n", c.Race)
	backgroundLine := fmt.Sprintf("Background: %s\n", c.Background)

	s := []string{
		levelLine,
		classLine,
		raceLine,
		backgroundLine,
	}

	return s
}

func (c *Character) BuildFeats() []string {
	s := make([]string, 0, len(c.Feats)+1)

	if len(c.Feats) < 1 || c.Feats[0].Name == "" {
		return s
	}

	featLine := "- Feats:\n"
	s = append(s, featLine)

	for _, feat := range c.Feats {
		featRow := fmt.Sprintf("	- %s: %s\n", feat.Name, feat.Desc)
		s = append(s, featRow)
	}
	s = append(s, "---")

	return s
}

func (c *Character) BuildLanguages() []string {
	s := make([]string, 0, len(c.Languages)+1)
	languagesLine := "- Languages:\n"
	s = append(s, languagesLine)

	for _, lang := range c.Languages {
		languageRow := fmt.Sprintf("	- %s\n", lang)
		s = append(s, languageRow)
	}

	return s
}

func (c *Character) BuildGeneralStats() []string {
	nl := "\n"
	proficiency := fmt.Sprintf("Proficiency: +%d\n", c.Proficiency)
	passPerception := fmt.Sprintf("Passive Perception: %d\n", c.PassivePerception)
	passInsight := fmt.Sprintf("Passive Insight: %d\n", c.PassiveInsight)

	acLine := fmt.Sprintf("AC: %d\n", c.AC)
	ssdcLine := fmt.Sprintf("Spell Save DC: %d\n", c.SpellSaveDC)
	speedLine := fmt.Sprintf("Speed: %d\n", c.Speed)
	hpLine := fmt.Sprintf("HP: %d/%d", c.HPCurrent, c.HPMax)

	if c.HPTemp > 0 {
		hpLine += fmt.Sprintf(" +%d", c.HPTemp)
	}

	hitDiceLine := fmt.Sprintf("Hit Dice: %s\n", c.HitDice)

	s := []string{
		proficiency,
		passPerception,
		passInsight,
		nl,
		acLine,
		ssdcLine,
		speedLine,
		hpLine,
		nl,
		hitDiceLine,
	}

	return s
}

func (c *Character) BuildAbilities() []string {
	s := make([]string, 0, len(c.Abilities)+3)
	profHeader := "*Abilities*\n\n"
	s = append(s, profHeader)

	profTopRow := "| Proficiency  | Base  | Modifier | Saving Throws |\n"
	profSpacer := "| --- | --- | --- | --- |\n"
	s = append(s, profTopRow)
	s = append(s, profSpacer)

	for _, types := range c.Abilities {
		abMod := types.AbilityModifier
		if types.SavingThrowsProficient {
			abMod += c.Proficiency
		}

		abBaseString := ""
		if types.AbilityModifier > 0 {
			abBaseString = "+"
		}

		abModString := ""
		if abMod > 0 {
			abModString = "+"
		}

		abModString = fmt.Sprintf("%s%d", abModString, abMod)
		abBaseString = fmt.Sprintf("%s%d", abBaseString, types.AbilityModifier)

		profRow := fmt.Sprintf("| %s | %d | %s | %s |\n", types.Name, types.Base, abBaseString, abModString)
		s = append(s, profRow)
	}

	return s
}

func (c *Character) BuildSkills() []string {
	s := make([]string, 0, len(c.Skills)+10)
	skillHeader := "*Skills*\n\n"
	s = append(s, skillHeader)

	skillTopRow := "| Skill | Ability | Modifier |\n"
	skillSpacer := "| --- | --- | --- |\n"
	s = append(s, skillTopRow)
	s = append(s, skillSpacer)

	for _, skill := range c.Skills {
		skillModifierString := ""
		if skill.SkillModifier > 0 {
			skillModifierString = "+"
		}

		skillModifierString = fmt.Sprintf("%s%d", skillModifierString, skill.SkillModifier)
		skillRow := fmt.Sprintf("| %s | %s | %s |\n", skill.Name, skill.Ability, skillModifierString)
		s = append(s, skillRow)
	}

	return s
}

func (c *Character) BuildSpells() []string {
	s := make([]string, 0, len(c.Spells)+20)

	// if the character doesn't have a spell save dc, they can't cast any spell and we don't need
	// to polute the markdown file with a bunch of empty spell data
	if c.SpellSaveDC == 0 {
		return s
	}

	nl := "\n"
	spellHeader := "*Spells*\n\n"
	s = append(s, spellHeader)

	spellTopRow := "| Slot Level | Ritual | Spell | IsPrepared |\n"
	spellSpacer := "| --- | --- | --- | --- |\n"
	s = append(s, spellTopRow)
	s = append(s, spellSpacer)

	for _, spell := range c.Spells {
		rString := " "
		if spell.IsRitual {
			rString = "*"
		}

		pString := " "
		if spell.IsPrepared {
			pString = "*"
		}

		spellRow := fmt.Sprintf("| %d | %s | %s | %s |\n", spell.SlotLevel, rString, spell.Name, pString)
		s = append(s, spellRow)
	}
	s = append(s, nl)

	spellSlots := "- Spell Slots\n"
	s = append(s, spellSlots)

	for _, spellSlot := range c.SpellSlots {
		if spellSlot.Maximum == 0 {
			continue
		}

		fullCircle := strings.Repeat("● ", spellSlot.Available)
		hollowCircle := strings.Repeat("○ ", (spellSlot.Maximum - spellSlot.Available))
		slotRow := fmt.Sprintf("	- Level %d: %s%s\n", spellSlot.Level, fullCircle, hollowCircle)
		s = append(s, slotRow)
	}

	return s
}

func (c *Character) BuildWeapons() []string {
	s := make([]string, 0, len(c.Weapons)+10)
	weaponsHeader := "*Weapons*\n\n"
	s = append(s, weaponsHeader)

	weaponTopRow := "| Weapon | Bonus | Damage | Type | Properties | Equipped |\n"
	weaponSpacer := "| --- | --- | --- | --- | --- | --- |\n"
	s = append(s, weaponTopRow)
	s = append(s, weaponSpacer)

	// this is basically only used in cases where you have two of the same kind of Weapons
	// and both of them are equipped
	primaryEquippedChecked := false

	for _, weapon := range c.Weapons {
		wBonusString := ""
		if weapon.Bonus > 0 {
			wBonusString = "+"
		}

		equippedString := ""

		if strings.EqualFold(c.PrimaryEquipped, weapon.Name) && !primaryEquippedChecked {
			equippedString = "Primary"
			primaryEquippedChecked = true
		} else if strings.EqualFold(c.SecondaryEquipped, weapon.Name) {
			equippedString = "Secondary"
		}

		wBonusString = fmt.Sprintf("%s%d", wBonusString, weapon.Bonus)
		weaponRow := fmt.Sprintf(
			"| %s | %s | %s | %s | %s | %s |\n",
			weapon.Name,
			wBonusString,
			weapon.Damage,
			weapon.Type,
			strings.Join(weapon.Properties, ", "),
			equippedString)

		s = append(s, weaponRow)
	}

	return s
}

func (c *Character) BuildEquipment() []string {
	s := []string{}
	equipmentHeader := "*Equipment*\n\n"
	bodyEquipment := "- Body Equipment\n"
	s = append(s, equipmentHeader)
	s = append(s, bodyEquipment)

	if c.WornEquipment.Head != "" {
		s = append(s, fmt.Sprintf("	- Head: %s\n", c.WornEquipment.Head))
	}
	if c.WornEquipment.Amulet != "" {
		s = append(s, fmt.Sprintf("	- Amulet: %s\n", c.WornEquipment.Amulet))
	}
	if c.WornEquipment.Cloak != "" {
		s = append(s, fmt.Sprintf("	- Cloak: %s\n", c.WornEquipment.Cloak))
	}
	if c.WornEquipment.Armor.Name != "" {
		s = append(s, fmt.Sprintf("	- Armor: %s\n", c.WornEquipment.Armor.Name))
	}
	if c.WornEquipment.HandsArms != "" {
		s = append(s, fmt.Sprintf("	- Hands: %s\n", c.WornEquipment.HandsArms))
	}
	if c.WornEquipment.Ring != "" {
		s = append(s, fmt.Sprintf("	- Ring: %s\n", c.WornEquipment.Ring))
	}
	if c.WornEquipment.Ring2 != "" {
		s = append(s, fmt.Sprintf("	- Ring: %s\n", c.WornEquipment.Ring2))
	}
	if c.WornEquipment.Belt != "" {
		s = append(s, fmt.Sprintf("	- Belt: %s\n", c.WornEquipment.Belt))
	}
	if c.WornEquipment.Boots != "" {
		s = append(s, fmt.Sprintf("	- Boots: %s\n", c.WornEquipment.Boots))
	}
	if c.WornEquipment.Shield != "" {
		shield := fmt.Sprintf("	- Shield: %s", c.WornEquipment.Shield)
		if strings.EqualFold(c.WornEquipment.Shield, c.PrimaryEquipped) ||
			strings.EqualFold(c.WornEquipment.Shield, c.SecondaryEquipped) {
			shield += " (equipped)"
		}
		s = append(s, shield+"\n")
	}

	return s
}

func (c *Character) BuildBackpack() []string {
	s := make([]string, 0, len(c.Weapons)+10)
	packHeader := "*Backpack*\n\n"
	s = append(s, packHeader)

	itemTopRow := "| Item | Quantity |\n"
	itemSpacer := "| --- | --- |\n"
	s = append(s, itemTopRow)
	s = append(s, itemSpacer)

	for _, item := range c.Backpack {
		itemRow := fmt.Sprintf("| %s | %d |\n", item.Name, item.Quantity)
		s = append(s, itemRow)
	}

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

		abilityScoreImprovementHeader := "Ability Score Improvement\n"
		s = append(s, abilityScoreImprovementHeader)

		// TODO: Make this more sophistocated so we don't need to loop through this twice
		for _, ability := range c.AbilityScoreImprovement {
			if ability.Ability == "" {
				continue
			}

			expLine := fmt.Sprintf("- %s: +%d\n", ability.Ability, ability.Bonus)
			s = append(s, expLine)
		}
		s = append(s, "\n")
	}

	return s
}

func GetSlots(available int, max int) string {
	// using non breaking spaces for how lipgloss trims regular spaces at the end of strings
	fullCircle := strings.Repeat("●\u00A0", available)
	hollowCircle := strings.Repeat("○\u00A0", (max - available))

	return fmt.Sprintf("%s%s", fullCircle, hollowCircle)
}

// CLI Actions

func (c *Character) AddItemToPack(item string, quantity int) {
	for i, packItem := range c.Backpack {
		if strings.EqualFold(packItem.Name, item) {
			c.Backpack[i].Quantity += quantity
			return
		}
	}

	newItem := shared.BackpackItem{
		Name:     item,
		Quantity: quantity,
	}

	c.Backpack = append(c.Backpack, newItem)
}

func (c *Character) RemoveItemFromPack(item string, quantity int) error {
	var err error
	for i, packItem := range c.Backpack {
		if strings.EqualFold(packItem.Name, item) {
			if packItem.Quantity < quantity {
				err = fmt.Errorf("Quantity to remove (%d) greater than quantity in pack (%d), set to 0",
					quantity,
					packItem.Quantity)

				c.Backpack[i].Quantity = 0
				return err
			}

			c.Backpack[i].Quantity -= quantity
			return nil
		}
	}

	err = fmt.Errorf("Item %s not found in pack", item)

	return err
}

func (c *Character) AddLanguage(language string) {
	c.Languages = append(c.Languages, language)
}

func (c *Character) AddEquipment(equipmentType string, equipmentName string) {
	equipmentName = strings.ToLower(equipmentName)
	switch equipmentType {
	case shared.WornEquipmentHead:
		c.WornEquipment.Head = equipmentName
	case shared.WornEquipmentAmulet:
		c.WornEquipment.Amulet = equipmentName
	case shared.WornEquipmentCloak:
		c.WornEquipment.Cloak = equipmentName
	case shared.WornEquipmentArmor:
		c.WornEquipment.Armor.Name = equipmentName
	case shared.WornEquipmentHandsArms:
		c.WornEquipment.HandsArms = equipmentName
	case shared.WornEquipmentRing:
		c.WornEquipment.Ring = equipmentName
	case shared.WornEquipmentRing2:
		c.WornEquipment.Ring2 = equipmentName
	case shared.WornEquipmentBelt:
		c.WornEquipment.Belt = equipmentName
	case shared.WornEquipmentBoots:
		c.WornEquipment.Boots = equipmentName
	default:
		info := fmt.Sprintf("Invalid Equipment Type: %s", equipmentType)
		logger.Info(info)
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
		logger.Info("Character had no health left")
		return
	}

	if c.HPTemp > 0 {
		c.HPTemp -= hpDecr

		if c.HPTemp > 0 {
			return
		}

		hpDecr = (c.HPTemp * -1)
		c.HPTemp = 0
	}

	c.HPCurrent -= hpDecr

	// reset to zero if the decremented amount is greater than remaining health
	if c.HPCurrent < 0 {
		c.HPCurrent = 0
	}
}

func (c *Character) AddTempHp(tempHP int) {
	c.HPTemp += tempHP
}

func (c *Character) AddSubClass(subClass string) {
	if c.Class != nil {
		c.Class.SetSubClass(subClass)
	}
}

// Adds an ability score to an existing item if found, or creates a new record with the bonus quantity
func (c *Character) AddAbilityScoreImprovementItem(quantity int, ability string) error {
	if quantity != 1 && quantity != 2 {
		return fmt.Errorf("When adding an Ability Score Improvement Item, the bonus must be 1 or 2")
	}

	if !isValidAbilityName(ability) {
		return fmt.Errorf("Ability name '%s' is not valid", ability)
	}

	for i, item := range c.AbilityScoreImprovement {
		if strings.EqualFold(item.Ability, ability) {
			c.AbilityScoreImprovement[i].Bonus += quantity
			fmt.Println(c.AbilityScoreImprovement[i].Bonus)
			return nil
		}
	}

	abi := shared.AbilityScoreImprovementItem{
		Ability: ability,
		Bonus:   quantity,
	}

	c.AbilityScoreImprovement = append(c.AbilityScoreImprovement, abi)

	return nil
}

// Modifies the bonus of an existing ability score improvement item if found, otherwise returns an error
func (c *Character) ModifyAbilityScoreImprovementItem(quantity int, ability string) error {
	if !isValidAbilityName(ability) {
		return fmt.Errorf("Ability name '%s' is not valid", ability)
	}

	for i, item := range c.AbilityScoreImprovement {
		if strings.EqualFold(item.Ability, ability) {
			c.AbilityScoreImprovement[i].Bonus = quantity
			return nil
		}
	}

	return fmt.Errorf("Ability Score Improvement Item for '%s' was not found. Try adding new item", ability)
}

func isValidAbilityName(abilityName string) bool {
	switch strings.ToLower(abilityName) {
	case shared.AbilityStrength:
		return true
	case shared.AbilityDexterity:
		return true
	case shared.AbilityConstitution:
		return true
	case shared.AbilityIntelligence:
		return true
	case shared.AbilityWisdom:
		return true
	case shared.AbilityCharisma:
		return true
	}

	return false
}

// TODO: Rename functionality will have to change with the support of multiple character files
// // RenameCharacter updates the character's name.
// func (c *Character) RenameCharacter(newName string) {
// 	c.Name = newName
// }

func (c *Character) UseSpellSlot(level int) {
	for i := range c.SpellSlots {
		if c.SpellSlots[i].Level == level {
			if c.SpellSlots[i].Available <= 0 {
				info := fmt.Sprintf("Spell Slot Level %d: already at zero", level)
				logger.Info(info)

				return
			}

			c.SpellSlots[i].Available--
			return
		}
	}

	logger.Info("Invalid level, must be 1-9")
}

func (c *Character) RecoverSpellSlots(level int, quantity int) {
	for i := range c.SpellSlots {
		if c.SpellSlots[i].Level == level {
			if c.SpellSlots[i].Maximum == 0 {
				logger.Warnf("Slot level '%d' cannot be recovered because the maximum 0", c.SpellSlots[i].Maximum)
				return
			} else if quantity == 0 {
				c.SpellSlots[i].Available = c.SpellSlots[i].Maximum
				return
			}

			c.SpellSlots[i].Available += quantity
		}
	}
}

func (c *Character) Recover() {
	c.HPCurrent = c.HPMax

	for i := range c.SpellSlots {
		c.SpellSlots[i].Available = c.SpellSlots[i].Maximum
	}

	if c.Class == nil {
		return
	}

	c.RecoverClassTokens("", 0)
}

func (c *Character) UseClassTokens(name string, quantity int) {
	if tokenClass, ok := c.Class.(TokenClass); ok {
		tokenClass.UseClassTokens(name, quantity)
	}
}

func (c *Character) RecoverClassTokens(name string, quantity int) {
	c.CalculateCharacterStats()

	if tokenClass, ok := c.Class.(TokenClass); ok {
		if postCalculater, ok := c.Class.(PostCalculator); ok {
			postCalculater.ExecutePostCalculateMethods(c)
		}

		tokenClass.RecoverClassTokens("", quantity)
	}
}

func (c *Character) Equip(isPrimary bool, name string) error {
	lowerName := strings.ToLower(name)
	equipmentFound := false

	weaponCount := 0
	for _, weapon := range c.Weapons {
		if strings.ToLower(weapon.Name) == lowerName {
			equipmentFound = true
			weaponCount++
			continue // not breaking here to verify if using the same weapon in each hand
		}
	}

	if !equipmentFound {
		if strings.ToLower(c.WornEquipment.Shield) == lowerName {
			equipmentFound = true
		}
	}

	if !equipmentFound {
		return fmt.Errorf("Weapon or shield '%s' not found, check spelling", name)
	}

	if isPrimary {
		// unequip secondary weapon if they only have one and it's already equipped
		if strings.ToLower(c.SecondaryEquipped) == lowerName && weaponCount < 2 {
			c.SecondaryEquipped = ""
		}

		c.PrimaryEquipped = name
	} else {
		// unequip primary weapon if they only have one and it's already equipped
		if strings.ToLower(c.PrimaryEquipped) == lowerName && weaponCount < 2 {
			c.PrimaryEquipped = ""
		}

		c.SecondaryEquipped = name
	}

	return nil
}

func (c *Character) Unequip(isPrimary bool) {
	if isPrimary {
		c.PrimaryEquipped = ""
	} else {
		c.SecondaryEquipped = ""
	}
}
