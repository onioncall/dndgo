package manage

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
)

type EquipmentModel struct {
	wornEquipmentViewport viewport.Model
	backpackViewport      viewport.Model
	weaponsViewport       viewport.Model
}

func NewEquipmentModel(character *models.Character) EquipmentModel {
	wornEquipmentVp := viewport.New(0, 0)
	wornEquipmentContent := getWornEquipmentContent(character)
	wornEquipmentVp.SetContent(wornEquipmentContent)

	backpackVp := viewport.New(0, 0)
	backpackContent := getBackpackContent(character)
	backpackVp.SetContent(backpackContent)

	weaponsVp := viewport.New(0, 0)
	weaponsContent := getWeaponsContent(character)
	weaponsVp.SetContent(weaponsContent)

	return EquipmentModel{
		wornEquipmentViewport: wornEquipmentVp,
		backpackViewport:      backpackVp,
		weaponsViewport:       weaponsVp,
	}
}

func getBackpackContent(character *models.Character) string {
	backpackContent := "Backpack\n\n"
	maxLength := 8 //length of header

	var contentWithoutSpacers []string
	for _, item := range character.Backpack {
		itemStr := fmt.Sprintf("%d - %s", item.Quantity, item.Name)
		contentWithoutSpacers = append(contentWithoutSpacers, itemStr)
		maxLength = max(maxLength, utf8.RuneCountInString(itemStr))
	}

	for _, item := range contentWithoutSpacers {
		backpackContent += fmt.Sprintf("%s%s\n", item, strings.Repeat("\u00A0", maxLength-utf8.RuneCountInString(item)))
	}

	return backpackContent
}

func getWornEquipmentContent(character *models.Character) string {
	equipmentContent := "Worn Equipment\n\n"
	headerLen := 14

	amuletStr := fmt.Sprintf("Amulet: %s", character.WornEquipment.Amulet)
	amuletLen := utf8.RuneCountInString(amuletStr)
	beltStr := fmt.Sprintf("Belt: %s", character.WornEquipment.Belt)
	beltLen := utf8.RuneCountInString(beltStr)
	bootsStr := fmt.Sprintf("Boots: %s", character.WornEquipment.Boots)
	bootsLen := utf8.RuneCountInString(bootsStr)
	cloakStr := fmt.Sprintf("Cloak: %s", character.WornEquipment.Cloak)
	cloakLen := utf8.RuneCountInString(cloakStr)
	headStr := fmt.Sprintf("Helmet: %s", character.WornEquipment.Head)
	headLen := utf8.RuneCountInString(headStr)
	ringStr := fmt.Sprintf("Ring: %s", character.WornEquipment.Ring)
	ringLen := utf8.RuneCountInString(ringStr)
	ring2Str := fmt.Sprintf("Ring2: %s", character.WornEquipment.Ring2)
	ring2Len := utf8.RuneCountInString(ring2Str)
	armorStr := fmt.Sprintf("Armor: %s", character.WornEquipment.Armor.Name)
	armorLen := utf8.RuneCountInString(armorStr)

	maxLen := max(headerLen, amuletLen, beltLen, bootsLen, cloakLen, ring2Len, headLen, ringLen, armorLen)

	equipmentContent += fmt.Sprintf("%s%s\n", armorStr, strings.Repeat("\u00A0", maxLen-armorLen))
	equipmentContent += fmt.Sprintf("%s%s\n", amuletStr, strings.Repeat("\u00A0", maxLen-amuletLen))
	equipmentContent += fmt.Sprintf("%s%s\n", cloakStr, strings.Repeat("\u00A0", maxLen-cloakLen))
	equipmentContent += fmt.Sprintf("%s%s\n", headStr, strings.Repeat("\u00A0", maxLen-headLen))
	equipmentContent += fmt.Sprintf("%s%s\n", beltStr, strings.Repeat("\u00A0", maxLen-beltLen))
	equipmentContent += fmt.Sprintf("%s%s\n", bootsStr, strings.Repeat("\u00A0", maxLen-bootsLen))
	equipmentContent += fmt.Sprintf("%s%s\n", ring2Str, strings.Repeat("\u00A0", maxLen-ring2Len))
	equipmentContent += fmt.Sprintf("%s%s\n", ringStr, strings.Repeat("\u00A0", maxLen-ringLen))

	return equipmentContent
}

func getWeaponsContent(character *models.Character) string {
	// As a general note, any weirdness around how we're handling primary weapons is probably related to
	// handling primary and secondary when both are the same weapon name

	maxNameLength := 4
	maxTypeLength := 4
	maxPropertiesLength := 10

	for _, w := range character.Weapons {
		nameStr := w.Name
		if strings.ToLower(character.PrimaryEquipped) == strings.ToLower(w.Name) {
			nameStr += "*"
		} else if strings.ToLower(character.SecondaryEquipped) == strings.ToLower(w.Name) {
			nameStr += "**"
		}

		maxNameLength = max(maxNameLength, utf8.RuneCountInString(nameStr))
		maxTypeLength = max(maxTypeLength, utf8.RuneCountInString(w.Type))
		maxPropertiesLength = max(maxPropertiesLength, utf8.RuneCountInString(strings.Join(w.Properties, ", ")))
	}

	weaponsHeader := fmt.Sprintf("Name%s - Proficient - Bonus - Damage - Normal/Long - Type%s - Properties%s\n",
		strings.Repeat(" ", maxNameLength-4),
		strings.Repeat(" ", maxTypeLength-4),
		strings.Repeat("\u00A0", maxPropertiesLength-10))

	weaponsContent := weaponsHeader
	lineWidth := utf8.RuneCountInString(weaponsHeader)
	weaponsContent += fmt.Sprintf("%s\n", strings.Repeat("─", lineWidth))

	for _, w := range character.Weapons {
		normalLongStr := fmt.Sprintf("%d/%d", w.Range.NormalRange, w.Range.LongRange)
		prof := strings.Repeat(" ", 10)
		if w.Proficient {
			prof = "Proficient"
		}
		bonusStr := fmt.Sprintf(" %d", w.Bonus)
		if w.Bonus >= 0 {
			bonusStr = fmt.Sprintf("%s%s", "+", bonusStr)
		}
		propertiesStr := strings.Join(w.Properties, ", ")

		nameStr := w.Name
		if strings.ToLower(character.PrimaryEquipped) == strings.ToLower(w.Name) {
			nameStr += "*"
		} else if strings.ToLower(character.SecondaryEquipped) == strings.ToLower(w.Name) {
			nameStr += "**"
		}

		nameLen := utf8.RuneCountInString(nameStr)
		bonusLen := utf8.RuneCountInString(bonusStr)
		damageLen := utf8.RuneCountInString(w.Damage)
		rangeLen := utf8.RuneCountInString(normalLongStr)
		typeLen := utf8.RuneCountInString(w.Type)
		propertiesLen := utf8.RuneCountInString(propertiesStr)

		// For any concerned citizens, this is so we can correctly align the spacing of the headers with the rows.
		// The reason we are subtracting these numbers from a min function is so that we don't get a repeater 0
		// error if the row is longer than the header. Since we already set the header to the max length of the
		// string, the spacer is 0. If it's longer than the header title, we add the appropriate spacer
		nameSpacer := strings.Repeat(" ", maxNameLength-min(nameLen, maxNameLength))
		bonusSpacer := strings.Repeat(" ", 5-min(bonusLen, 5))
		damageSpacer := strings.Repeat(" ", 6-min(damageLen, 6))
		rangeSpacer := strings.Repeat(" ", 11-min(rangeLen, 11))
		typeSpacer := strings.Repeat(" ", maxTypeLength-min(typeLen, maxTypeLength))
		propertiesSpacer := strings.Repeat("\u00A0", maxPropertiesLength-min(propertiesLen, maxPropertiesLength))

		weaponsContent += fmt.Sprintf("%s%s - %s - %s%s - %s%s - %s%s - %s%s - %s%s\n",
			nameStr,
			nameSpacer,
			prof,
			bonusStr,
			bonusSpacer,
			w.Damage,
			damageSpacer,
			normalLongStr,
			rangeSpacer,
			w.Type,
			typeSpacer,
			propertiesStr,
			propertiesSpacer)
	}

	return weaponsContent
}

func (m EquipmentModel) UpdateSize(innerWidth, availableHeight int, character *models.Character) EquipmentModel {
	// Row 1: 50/50 horizontal split, taking 50% of height
	row1Height := availableHeight / 2
	row2Height := availableHeight - row1Height

	// Worn Equipment viewport (left side of row 1)
	wornWidth := innerWidth / 2
	wornInnerWidth := wornWidth - 2
	wornInnerHeight := row1Height - 2

	m.wornEquipmentViewport.Width = wornInnerWidth
	m.wornEquipmentViewport.Height = wornInnerHeight

	// Backpack viewport (right side of row 1)
	backpackWidth := innerWidth - wornWidth
	backpackInnerWidth := backpackWidth - 2
	backpackInnerHeight := row1Height - 2

	m.backpackViewport.Width = backpackInnerWidth
	m.backpackViewport.Height = backpackInnerHeight

	// Weapons viewport (row 2, full width)
	weaponsInnerWidth := innerWidth - 2
	weaponsInnerHeight := row2Height - 2

	m.weaponsViewport.Width = weaponsInnerWidth
	m.weaponsViewport.Height = weaponsInnerHeight

	return m
}
