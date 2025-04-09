package api

import (
	"fmt"
	"sort"
	"strings"

	"github.com/onioncall/dndgo/models"
)

type SpellRequest BaseRequest

func (s *SpellRequest) GetList() models.SpellList {
	spellList, err := ExecuteGetRequest[models.SpellList](SpellType, "")
	if err != nil {
		panic(err)
	}

	return spellList
}

func (s *SpellRequest) IsList() bool {
    if s.Name == "list" || s.Name == "l" {
		return true
	}

	return false
}

func (s *SpellRequest) GetSingle() models.Spell {
	s.Name = strings.ReplaceAll(s.Name, " ", "-")

	spell := models.Spell{}
	spell, err := ExecuteGetRequest[models.Spell](SpellType, s.Name)
	if err != nil {
		panic(err)
	}

	return spell
}

func (s *SpellRequest) PrintSingle(spell models.Spell) {
	fmt.Printf("%s\n\n", spell.Name)
	for _, description := range spell.Description {
		fmt.Printf("%s\n\n", description)	
	}

	if spell.AreaOfEffect.Size != 0 {
		fmt.Printf("Area of Effect: %v %v\n", spell.AreaOfEffect.Type, spell.AreaOfEffect.Size)
	}

	fmt.Printf("Range: %v\n", spell.Range)
	fmt.Printf("Casting Time: %v\n", spell.CastingTime)

	if spell.Damage != nil {
		fmt.Println()
		fmt.Printf("Damage By Slot Level: \n\n")
		// Because maps aren't sortable, we have to do this to print the damage by slot level nicely
		printDamageBySlotLevel(spell.Damage.DamageAtSlotLevel)
	}

	fmt.Println()
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

func (s *SpellRequest) PrintList(spellList models.SpellList) {
	fmt.Println("Spell Name | Level")
	for _, spell := range spellList.ListItems {
		fmt.Printf("%s: %d\n", spell.Name, spell.Level)
	}
}
