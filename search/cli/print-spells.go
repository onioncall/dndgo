package cli

import (
	"fmt"
	"sort"
	"strings"

	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/wrapt"
)

func PrintSpellSingle(spell responses.Spell, termWidth int) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%s\n\n", spell.Name))
	for _, description := range spell.Description {
		builder.WriteString(fmt.Sprintf("%s\n\n", wrapt.Wrap(description, termWidth)))
	}

	if spell.AreaOfEffect.Size != 0 {
		builder.WriteString(fmt.Sprintf("Area of Effect: %v %v\n", spell.AreaOfEffect.Type, spell.AreaOfEffect.Size))
	}

	builder.WriteString(fmt.Sprintf("Range: %v\n", spell.Range))
	builder.WriteString(fmt.Sprintf("Casting Time: %v\n", spell.CastingTime))
	builder.WriteString(fmt.Sprintf("Duration: %v\n", spell.Duration))

	if spell.Damage != nil {
		fmt.Println()
		builder.WriteString(fmt.Sprintf("Damage Type: %s\n", spell.Damage.DamageType.Name))

		if spell.Damage.DamageAtSlotLevel != nil {
			builder.WriteString(fmt.Sprintf("Damage By Slot Level: \n\n"))
			// Because maps aren't sortable, we have to do this to print the damage by slot level nicely
			builder.WriteString(damageBySlotLevel(spell.Damage.DamageAtSlotLevel))
		}
	}

	return builder.String()
}

func damageBySlotLevel(dmg map[int]string) string {
	var builder strings.Builder

	keys := make([]int, 0, len(dmg))
	for k := range dmg {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		builder.WriteString(fmt.Sprintf("	%v: %v\n", k, dmg[k]))
	}

	return builder.String()
}

func PrintSpellList(spellList responses.SpellList) string {
	var builder strings.Builder
	builder.WriteString("Spell Name | Level")
	for _, spell := range spellList.ListItems {
		builder.WriteString(fmt.Sprintf("%s: %d\n", spell.Name, spell.Level))
	}

	return builder.String()
}
