package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/api/responses"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/logger"
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
		logErr := fmt.Errorf("Failed to Handle Spell Request (single)")
		logger.HandleError(err, logErr)

		return err
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
		logErr := fmt.Errorf("Failed to Handle Spell Request (list)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintSpellList(sl)
	return nil
}

func (s *SpellRequest) GetList() (responses.SpellList, error) {
	spellList, err := api.ExecuteGetRequest[responses.SpellList](SpellType, "")
	if err != nil {
		logErr := fmt.Errorf("Failed to search Spell (list)")
		logger.HandleError(err, logErr)

		return spellList, err
	}

	return spellList, nil
}

func (s *SpellRequest) GetSingle() (responses.Spell, error) {
	s.Name = strings.ReplaceAll(s.Name, " ", "-")

	spell := responses.Spell{}
	spell, err := api.ExecuteGetRequest[responses.Spell](SpellType, s.Name)
	if err != nil {
		logErr := fmt.Errorf("Failed to search Spell (single): %s", s.Name)
		logger.HandleError(err, logErr)

		return spell, err
	}

	return spell, nil
}
