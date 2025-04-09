package main

import (
	"flag"
	"fmt"

	"github.com/onioncall/dndgo/api"
)

func main() {
	tui := flag.Bool("tui", false, "Use TUI instead of CLI")
	spell := flag.String("s", "", "Search Spells, use l or list as an argument to return all spells")
	monster := flag.String("m", "", "Search Monsters, use l or list as an argument to return all monsters")

	flag.Parse()

	// This is temporary, and will be removed when the tui is implemented
	if *tui {
		fmt.Println("the dndgo tui is under construction, please try it another time")
		return
	}
	
	// Once I have a third one of these, I'll abstract this logic into a generic function
	// Leaving it for now...
	switch {
	case *spell != "":
		r := api.RequestFactory(*spell, api.SpellRequest{}, api.SpellType)
		if !r.IsList() {
			s := r.GetSingle()
			if !*tui {
				r.PrintSingle(s)		
			}
		} else {
			s := r.GetList()
			if !*tui {
				r.PrintList(s)
			}
		}
	case *monster != "":
		r := api.RequestFactory(*monster, api.MonsterRequest{}, api.MonsterType)
		if !r.IsList() {
			m := r.GetSingle()
			if !*tui {
				r.PrintSingle(m)		
			}
		} else {
			m := r.GetList()
			if !*tui {
				r.PrintList(m)
			}
		}
	}
}
