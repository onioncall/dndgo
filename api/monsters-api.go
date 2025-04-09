package api

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/models"
)

type MonsterRequest BaseRequest

func (m *MonsterRequest) GetList() models.MonsterList {
	monsterList, err := ExecuteGetRequest[models.MonsterList](MonsterType, "")
	if err != nil {
		panic(err)
	}

	return monsterList
}

func (m *MonsterRequest) IsList() bool {
    if m.Name == "list" || m.Name == "l" {
		return true
	}

	return false
}

func (m *MonsterRequest) GetSingle() models.Monster {
	m.Name = strings.ReplaceAll(m.Name, " ", "-")

	monster, err := ExecuteGetRequest[models.Monster](MonsterType, m.Name)
	if err != nil {
		panic(err)
	}
	
	return monster
}

func (m *MonsterRequest) PrintSingle(monster models.Monster) {

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
		printSpecialAbilities(monster.SpecialAbilities)
	}
	fmt.Println()
}

func (m *MonsterRequest) PrintList(monsterList models.MonsterList) {
	fmt.Print("Monster Name\n\n")
	for _, monster := range monsterList.ListItems {
		fmt.Printf("%s\n", monster.Name)
	}
}

func printSpecialAbilities(abilities []models.SpecialAbility) {
	fmt.Print("Special Abilities:\n\n")
	
	for _, ability := range abilities {
		fmt.Printf("	%s\n", ability.Name)
		fmt.Printf("	%s\n", ability.Desc)
		if ability.Usage != nil {
			fmt.Printf("	Usage Type: %s\n", ability.Usage.Type)
			fmt.Printf("	Usage Times: %d\n", ability.Usage.Times)
		}
		fmt.Println()
	}
}
