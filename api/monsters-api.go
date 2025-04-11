package api

import (
	"strings"

	"github.com/onioncall/dndgo/models"
)

type MonsterRequest BaseRequest

func (m *MonsterRequest) GetList() models.MonsterList {
	monsterList, err := ExecuteGetRequest[models.MonsterList](MonsterType, "")
	if err != nil {
		panic(err)
	}

	return monsterList
}

func (m *MonsterRequest) IsList() bool {
    if m.Name == "list" || m.Name == "l" {
		return true
	}

	return false
}

func (m *MonsterRequest) GetSingle() models.Monster {
	m.Name = strings.ReplaceAll(m.Name, " ", "-")

	monster, err := ExecuteGetRequest[models.Monster](MonsterType, m.Name)
	if err != nil {
		panic(err)
	}
	
	return monster
}

