package main

import (
	"flag"
	"fmt"

	"github.com/onioncall/dndgo/handlers"
)

func main() {
	tui := flag.Bool("tui", false, "Use TUI instead of CLI")
	spell := flag.String("s", "", "Search Spells, use l or list as an argument to return all spells")
	monster := flag.String("m", "", "Search Monsters, use l or list as an argument to return all monsters")
	character := flag.String("c", "", "Create or Update Character Markdown")

	flag.Parse()

	// This is temporary, and will be removed when the tui is implemented
	if *tui {
		fmt.Println("the dndgo tui is under construction, please try it another time")
		return
	}
	
	switch {
	case *spell != "":
		handlers.HandleSpellRequest(*spell)
	case *monster != "":
		handlers.HandleMonsterRequest(*monster)
	case *character != "":
		handlers.HandleCharacter(*character)
	}
}
