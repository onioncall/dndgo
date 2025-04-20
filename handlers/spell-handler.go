package handlers

import (
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/models"
)

type SpellRequest api.BaseRequest
const SpellType	api.PathType = "spells"

func HandleSpellRequest(spellQuery string) {
	isList := false
	if spellQuery == "list" || spellQuery == "l" {
		isList = true
		spellQuery = ""
	}

	r := SpellRequest {
		Name: spellQuery,
		PathType: SpellType,
	}		
	
	if isList {
		sl := r.GetList()	
		cli.PrintSpellList(sl)
	} else {
		s := r.GetSingle()
		cli.PrintSpellSingle(s)
	}
}

func (s *SpellRequest) GetList() models.SpellList {
	spellList, err := api.ExecuteGetRequest[models.SpellList](SpellType, "")
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
	spell, err := api.ExecuteGetRequest[models.Spell](SpellType, s.Name)
	if err != nil {
		panic(err)
	}

	return spell
}
