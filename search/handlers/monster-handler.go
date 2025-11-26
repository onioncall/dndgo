package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/search/api"
	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/dndgo/search/cli"
)

type MonsterRequest api.BaseRequest

const MonsterType api.PathType = "monsters"

func HandleMonsterRequest(monsterQuery string, termWidth int) (string, error) {
	r := MonsterRequest{
		Name:     monsterQuery,
		PathType: MonsterType,
	}

	m, err := r.GetSingle()
	if err != nil {
		return "", fmt.Errorf("Failed to handle monster request (%s): %w", monsterQuery, err)
	}

	result := cli.FormatMonsterSingle(m, termWidth)
	return result, nil
}

func HandleMonsterListRequest() (string, error) {
	r := MonsterRequest{
		Name:     "",
		PathType: MonsterType,
	}

	ml, err := r.GetList()
	if err != nil {
		return "", fmt.Errorf("Failed to handle monster request list: %w", err)
	}

	result := cli.FormatMonsterList(ml)
	return result, nil
}

func (m *MonsterRequest) GetList() (responses.MonsterList, error) {
	monsterList, err := api.ExecuteGetRequest[responses.MonsterList](MonsterType, "")
	if err != nil {
		return monsterList, fmt.Errorf("Failed to get monster list: %w", err)
	}

	return monsterList, nil
}

func (m *MonsterRequest) GetSingle() (responses.Monster, error) {
	m.Name = strings.ReplaceAll(m.Name, " ", "-")

	monster, err := api.ExecuteGetRequest[responses.Monster](MonsterType, m.Name)
	if err != nil {
		return monster, fmt.Errorf("Failed to get monster (%s): %w", m.Name, err)
	}

	return monster, nil
}
