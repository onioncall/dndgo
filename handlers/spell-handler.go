package handlers

import (
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/models"
)

type SpellRequest api.BaseRequest
const SpellType	api.PathType = "spells"

func HandleSpellRequest(spellQuery string, termWidth int) {
	r := SpellRequest {
		Name: spellQuery,
		PathType: SpellType,
	}		
	
	s := r.GetSingle()
	cli.PrintSpellSingle(s, termWidth)
}

func HandleSpellListRequest() {
	r := SpellRequest {
		Name: "",
		PathType: SpellType,
	}		

	sl := r.GetList()	
	cli.PrintSpellList(sl)
}

func (s *SpellRequest) GetList() models.SpellList {
	spellList, err := api.ExecuteGetRequest[models.SpellList](SpellType, "")
	if err != nil {
		panic(err)
	}

	return spellList
}

func (s *SpellRequest) GetSingle() models.Spell {
	s.Name = strings.ReplaceAll(s.Name, " ", "-")

	spell := models.Spell{}
	spell, err := api.ExecuteGetRequest[models.Spell](SpellType, s.Name)
	if err != nil {
		panic(err)
	}

	return spell
}
