package cli

import (
	"fmt"
	"sort"

	"github.com/onioncall/dndgo/api/responses"
	"github.com/onioncall/wrapt"
)

func PrintSpellSingle(spell responses.Spell, termWidth int) {
	fmt.Printf("%s\n\n", spell.Name)
	for _, description := range spell.Description {
		fmt.Printf("%s\n\n", wrapt.Wrap(description, termWidth))	
	}

	if spell.AreaOfEffect.Size != 0 {
		fmt.Printf("Area of Effect: %v %v\n", spell.AreaOfEffect.Type, spell.AreaOfEffect.Size)
	}

	fmt.Printf("Range: %v\n", spell.Range)
	fmt.Printf("Casting Time: %v\n", spell.CastingTime)
	fmt.Printf("Duration: %v\n", spell.Duration)

	if spell.Damage != nil {
		fmt.Println()
		fmt.Printf("Damage Type: %s\n", spell.Damage.DamageType.Name)

		if spell.Damage.DamageAtSlotLevel != nil {
			fmt.Printf("Damage By Slot Level: \n\n")
			// Because maps aren't sortable, we have to do this to print the damage by slot level nicely
			printDamageBySlotLevel(spell.Damage.DamageAtSlotLevel)
		}
	}
}

func printDamageBySlotLevel(dmg map[int]string) {
    keys := make([]int, 0, len(dmg))
	for k := range dmg {
		keys = append(keys, k)
	}

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("	%v: %v\n", k, dmg[k])
	}		
}

func PrintSpellList(spellList responses.SpellList) {
	fmt.Println("Spell Name | Level")
	for _, spell := range spellList.ListItems {
		fmt.Printf("%s: %d\n", spell.Name, spell.Level)
	}
}
