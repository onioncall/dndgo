package cli

import (
	"fmt"

	"github.com/onioncall/dndgo/models"
	"github.com/onioncall/wrapt"
)

func PrintMonsterSingle(monster models.Monster, termWidth int) {

	fmt.Printf("%s\n\n", monster.Name)
	fmt.Printf("Hit Points: %d\n", monster.HitPoints)
	fmt.Printf("Strength: %d\n", monster.Strength)
	fmt.Printf("Dexterity: %d\n", monster.Dexterity)
	fmt.Printf("Consitution: %d\n", monster.Constitution)
	fmt.Printf("Intelligence: %d\n", monster.Intelligence)
	fmt.Printf("Wisdom: %d\n", monster.Wisdom)
	fmt.Printf("Charisma: %d\n", monster.Charisma)
	fmt.Println()

	if len(monster.SpecialAbilities) > 0 {
		printSpecialAbilities(monster.SpecialAbilities, termWidth)
	}
}

func PrintMonsterList(monsterList models.MonsterList) {
	fmt.Print("Monster Name\n\n")
	for _, monster := range monsterList.ListItems {
		fmt.Printf("%s - %s\n", monster.Name, monster.Index)
	}
}

func printSpecialAbilities(abilities []models.SpecialAbility, termWidth int) {
	fmt.Print("Special Abilities:\n")
	
	for _, ability := range abilities {
		tab := "    "
		fmt.Println()
		fmt.Printf("%s%s\n", tab, ability.Name)
		for _, s := range wrapt.WrapArray(ability.Desc, len(tab), termWidth) {
			fmt.Printf("%s%s\n", tab, s)
		}
		if ability.Usage != nil {
			fmt.Printf("%sType: %s\n", tab, ability.Usage.Type)
			fmt.Printf("%sUsage Times: %d\n", tab, ability.Usage.Times)
		}
	}
}
