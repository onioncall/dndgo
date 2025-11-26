package cli

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/wrapt"
)

func FormatMonsterSingle(monster responses.Monster, termWidth int) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%s\n\n", monster.Name))
	builder.WriteString(fmt.Sprintf("Hit Points: %d\n", monster.HitPoints))
	builder.WriteString(fmt.Sprintf("Strength: %d\n", monster.Strength))
	builder.WriteString(fmt.Sprintf("Dexterity: %d\n", monster.Dexterity))
	builder.WriteString(fmt.Sprintf("Consitution: %d\n", monster.Constitution))
	builder.WriteString(fmt.Sprintf("Intelligence: %d\n", monster.Intelligence))
	builder.WriteString(fmt.Sprintf("Wisdom: %d\n", monster.Wisdom))
	builder.WriteString(fmt.Sprintf("Charisma: %d\n", monster.Charisma))
	builder.WriteString("\n")

	if len(monster.SpecialAbilities) > 0 {
		formatSpecialAbilities(monster.SpecialAbilities, termWidth)
	}

	return builder.String()
}

func FormatMonsterList(monsterList responses.MonsterList) string {
	var builder strings.Builder

	fmt.Print("Monster Name\n\n")
	for _, monster := range monsterList.ListItems {
		builder.WriteString(fmt.Sprintf("%s - %s\n", monster.Name, monster.Index))
	}

	return builder.String()
}

func formatSpecialAbilities(abilities []responses.SpecialAbility, termWidth int) string {
	var builder strings.Builder

	builder.WriteString("Special Abilities:\n")

	for _, ability := range abilities {
		tab := "    "
		builder.WriteString("\n")
		builder.WriteString(fmt.Sprintf("%s%s\n", tab, ability.Name))
		for _, s := range wrapt.WrapArray(ability.Desc, len(tab), termWidth) {
			builder.WriteString(fmt.Sprintf("%s%s\n", tab, s))
		}
		if ability.Usage != nil {
			builder.WriteString(fmt.Sprintf("%sType: %s\n", tab, ability.Usage.Type))
			builder.WriteString(fmt.Sprintf("%sUsage Times: %d\n", tab, ability.Usage.Times))
		}
	}

	return builder.String()
}
