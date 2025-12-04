package equipment

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/tui/shared"
)

type EquipmentModel struct {
	WornEquipmentViewport viewport.Model
	BackpackViewport      viewport.Model
	WeaponsViewport       viewport.Model
	contentSet            bool
}

func NewEquipmentModel(character *models.Character) EquipmentModel {
	wornEquipmentVp := viewport.New(0, 0)
	backpackVp := viewport.New(0, 0)
	weaponsVp := viewport.New(0, 0)

	return EquipmentModel{
		WornEquipmentViewport: wornEquipmentVp,
		BackpackViewport:      backpackVp,
		WeaponsViewport:       weaponsVp,
	}
}

func GetBackpackContent(character *models.Character, width int) string {
	backpackContent := "Backpack\n"
	width = width - (widthPadding * 2) //padding on both sides
	backpackContent += fmt.Sprintf("%s\n", strings.Repeat("─", width))
	maxLength := 8 //length of header

	var contentWithoutSpacers []string
	for _, item := range character.Backpack {
		itemStr := fmt.Sprintf("%d - %s", item.Quantity, item.Name)
		contentWithoutSpacers = append(contentWithoutSpacers, itemStr)
		maxLength = max(maxLength, utf8.RuneCountInString(itemStr))
	}

	for _, item := range contentWithoutSpacers {
		itemLength := utf8.RuneCountInString(item)
		item := shared.TruncateString(item, width)
		backpackContent += fmt.Sprintf("%s%s\n", item, strings.Repeat("\u00A0", maxLength-itemLength))
	}

	return backpackContent
}

func GetWornEquipmentContent(character *models.Character, width int) string {
	width = width - (widthPadding * 2)
	equipmentContent := "Worn Equipment\n"
	equipmentContent += fmt.Sprintf("%s\n", strings.Repeat("─", width))
	headerLen := 14

	amuletStr := shared.TruncateString(fmt.Sprintf("Amulet: %s", character.WornEquipment.Amulet), width)
	amuletLen := utf8.RuneCountInString(amuletStr)
	beltStr := shared.TruncateString(fmt.Sprintf("Belt: %s", character.WornEquipment.Belt), width)
	beltLen := utf8.RuneCountInString(beltStr)
	bootsStr := shared.TruncateString(fmt.Sprintf("Boots: %s", character.WornEquipment.Boots), width)
	bootsLen := utf8.RuneCountInString(bootsStr)
	cloakStr := shared.TruncateString(fmt.Sprintf("Cloak: %s", character.WornEquipment.Cloak), width)
	cloakLen := utf8.RuneCountInString(cloakStr)
	headStr := shared.TruncateString(fmt.Sprintf("Helmet: %s", character.WornEquipment.Head), width)
	headLen := utf8.RuneCountInString(headStr)
	ringStr := shared.TruncateString(fmt.Sprintf("Ring: %s", character.WornEquipment.Ring), width)
	ringLen := utf8.RuneCountInString(ringStr)
	ring2Str := shared.TruncateString(fmt.Sprintf("Ring2: %s", character.WornEquipment.Ring2), width)
	ring2Len := utf8.RuneCountInString(ring2Str)
	armorStr := shared.TruncateString(fmt.Sprintf("Armor: %s", character.WornEquipment.Armor.Name), width)
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

func GetWeaponsContent(character *models.Character, width int) string {
	// As a general note, any weirdness around how we're handling primary weapons is probably related to
	// handling primary and secondary when both are the same weapon name

	width = width - (widthPadding * 2)
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

	weaponsHeader := fmt.Sprintf("Name%s - Bonus - Damage - Normal/Long - Type%s - Properties%s\n",
		strings.Repeat(" ", maxNameLength-4),
		strings.Repeat(" ", maxTypeLength-4),
		strings.Repeat("\u00A0", maxPropertiesLength-10))

	weaponsContent := weaponsHeader
	weaponsContent += fmt.Sprintf("%s\n", strings.Repeat("─", width-widthPadding))

	for _, w := range character.Weapons {
		normalLongStr := fmt.Sprintf("%d/%d", w.Range.NormalRange, w.Range.LongRange)
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

		weaponsLine := fmt.Sprintf("%s%s - %s%s - %s%s - %s%s - %s%s - %s%s\n",
			nameStr,
			nameSpacer,
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

		weaponsContent += shared.TruncateString(weaponsLine, width)
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

	m.WornEquipmentViewport.Width = wornInnerWidth
	m.WornEquipmentViewport.Height = wornInnerHeight

	// Backpack viewport (right side of row 1)
	backpackWidth := innerWidth - wornWidth
	backpackInnerWidth := backpackWidth - 2
	backpackInnerHeight := row1Height - 2

	m.BackpackViewport.Width = backpackInnerWidth
	m.BackpackViewport.Height = backpackInnerHeight

	// Weapons viewport (row 2, full width)
	weaponsInnerWidth := innerWidth - 2
	weaponsInnerHeight := row2Height - 2

	m.WeaponsViewport.Width = weaponsInnerWidth
	m.WeaponsViewport.Height = weaponsInnerHeight

	if !m.contentSet {
		backpackContent := GetBackpackContent(character, m.BackpackViewport.Width)
		m.BackpackViewport.SetContent(backpackContent)

		equipmentContent := GetWornEquipmentContent(character, m.WornEquipmentViewport.Width)
		m.WornEquipmentViewport.SetContent(equipmentContent)

		weaponsContent := GetWeaponsContent(character, m.WeaponsViewport.Width)
		m.WeaponsViewport.SetContent(weaponsContent)
	}

	return m
}
