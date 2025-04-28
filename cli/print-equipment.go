package cli

import (
	"fmt"

	"github.com/onioncall/dndgo/models"
	"github.com/onioncall/wrapt"
)

func PrintEquipmentSingle(equipment models.Equipment) {
	fmt.Printf("%s\n\n", equipment.Name)

	for _, description := range equipment.Desc {
		fmt.Printf("%s\n\n", wrapt.Wrap(description))	
	}
	
	fmt.Printf("Equipment Category: %s\n", equipment.EquipmentCategory.Name)
	if equipment.GearCategory != nil {
		fmt.Printf("Gear Category: %s\n", equipment.GearCategory.Name)
	} else if equipment.WeaponCategory != "" {
		fmt.Printf("Weapon Category: %s\n", equipment.WeaponCategory)
		fmt.Printf("Weapon Range: %s - %d\n", equipment.WeaponRange, equipment.Range.Normal)
		fmt.Printf("Damage: %s\n", equipment.Damage.DamageDice)
		fmt.Printf("Damage Type: %s\n", equipment.Damage.DamageType.Name)
	}

	fmt.Printf("Cost: %d%s\n", equipment.Cost.Quantity, equipment.Cost.Unit)
}

func PrintEquipmentList(equipmentList models.EquipmentList) {
	fmt.Print("Equipment Name\n\n")
	for _, equpment := range equipmentList.ListItems {
		fmt.Printf("%s - %s\n", equpment.Name, equpment.Index)
	}
}
