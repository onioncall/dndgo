package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/api/responses"
)

type MonsterRequest api.BaseRequest
const MonsterType api.PathType = "monsters"

func HandleMonsterRequest(monsterQuery string, termWidth int) error {
	r := MonsterRequest {
		Name: monsterQuery,
		PathType: MonsterType,
	}		
	
	m, err := r.GetSingle()
	if err != nil {
		logErr := fmt.Errorf("Failed to Handle Monster Request (single)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintMonsterSingle(m, termWidth)
	return nil
}

func HandleMonsterListRequest() error {
	r := MonsterRequest {
		Name: "",
		PathType: MonsterType,
	}		

	ml, err := r.GetList()
	if err != nil {
		logErr := fmt.Errorf("Failed to Handle Monster Request (list)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintMonsterList(ml)
	return nil
}

func (m *MonsterRequest) GetList() (responses.MonsterList, error) {
	monsterList, err := api.ExecuteGetRequest[responses.MonsterList](MonsterType, "")
	if err != nil {
		logErr := fmt.Errorf("Failed to search Monster (list)")
		logger.HandleError(err, logErr)

		return monsterList, err
	}

	return monsterList, nil
}

func (m *MonsterRequest) GetSingle() (responses.Monster, error) {
	m.Name = strings.ReplaceAll(m.Name, " ", "-")

	monster, err := api.ExecuteGetRequest[responses.Monster](MonsterType, m.Name)
	if err != nil {
		logErr := fmt.Errorf("Failed to search Monster (single): %s", m.Name)
		logger.HandleError(err, logErr)

		return monster, err
	}
	
	return monster, nil
}
