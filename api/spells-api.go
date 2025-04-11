package api

import (
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

