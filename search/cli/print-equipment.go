package cli

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/wrapt"
)

func PrintEquipmentSingle(equipment responses.Equipment, termWidth int) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%s\n\n", equipment.Name))

	for _, description := range equipment.Desc {
		builder.WriteString(fmt.Sprintf("%s\n\n", wrapt.Wrap(description, termWidth)))
	}

	builder.WriteString(fmt.Sprintf("Equipment Category: %s\n", equipment.EquipmentCategory.Name))
	if equipment.GearCategory != nil {
		builder.WriteString(fmt.Sprintf("Gear Category: %s\n", equipment.GearCategory.Name))
	} else if equipment.WeaponCategory != "" {
		builder.WriteString(fmt.Sprintf("Weapon Category: %s\n", equipment.WeaponCategory))
		builder.WriteString(fmt.Sprintf("Weapon Range: %s - %d\n", equipment.WeaponRange, equipment.Range.Normal))
		builder.WriteString(fmt.Sprintf("Damage: %s\n", equipment.Damage.DamageDice))
		builder.WriteString(fmt.Sprintf("Damage Type: %s\n", equipment.Damage.DamageType.Name))
	}

	builder.WriteString(fmt.Sprintf("Cost: %d%s\n", equipment.Cost.Quantity, equipment.Cost.Unit))

	return builder.String()
}

func PrintEquipmentList(equipmentList responses.EquipmentList) string {
	var builder strings.Builder
	builder.WriteString("Equipment Name\n\n")
	for _, equpment := range equipmentList.ListItems {
		builder.WriteString(fmt.Sprintf("%s - %s\n", equpment.Name, equpment.Index))
	}

	return builder.String()
}
