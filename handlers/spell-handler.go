package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/api/responses"
	"github.com/onioncall/dndgo/cli"
)

type SpellRequest api.BaseRequest

const SpellType api.PathType = "spells"

func HandleSpellRequest(spellQuery string, termWidth int) error {
	r := SpellRequest{
		Name:     spellQuery,
		PathType: SpellType,
	}

	s, err := r.GetSingle()
	if err != nil {
		return fmt.Errorf("Failed to handle spell request (%s): %w", spellQuery, err)
	}

	cli.PrintSpellSingle(s, termWidth)
	return nil
}

func HandleSpellListRequest() error {
	r := SpellRequest{
		Name:     "",
		PathType: SpellType,
	}

	sl, err := r.GetList()
	if err != nil {
		return fmt.Errorf("Failed to handle spell request list: %w", err)
	}

	cli.PrintSpellList(sl)
	return nil
}

func (s *SpellRequest) GetList() (responses.SpellList, error) {
	spellList, err := api.ExecuteGetRequest[responses.SpellList](SpellType, "")
	if err != nil {
		return spellList, fmt.Errorf("Failed to get spell (list): %w", err)
	}

	return spellList, nil
}

func (s *SpellRequest) GetSingle() (responses.Spell, error) {
	s.Name = strings.ReplaceAll(s.Name, " ", "-")

	spell := responses.Spell{}
	spell, err := api.ExecuteGetRequest[responses.Spell](SpellType, s.Name)
	if err != nil {
		return spell, fmt.Errorf("Failed to get spell (%s): %w", s.Name, err)
	}

	return spell, nil
}
