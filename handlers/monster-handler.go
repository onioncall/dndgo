package handlers

import (
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/models"
)

type MonsterRequest api.BaseRequest
const MonsterType api.PathType = "monsters"

func HandleMonsterRequest(monsterQuery string) {
	isList := false
	if monsterQuery == "list" || monsterQuery == "l" {
		isList = true
		monsterQuery = ""
	}

	r := MonsterRequest {
		Name: monsterQuery,
		PathType: MonsterType,
	}		
	
	if isList {
		ml := r.GetList()	
		cli.PrintMonsterList(ml)
	} else {
		m := r.GetSingle()
		cli.PrintMonsterSingle(m)
	}
}

func (m *MonsterRequest) GetList() models.MonsterList {
	monsterList, err := api.ExecuteGetRequest[models.MonsterList](MonsterType, "")
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

	monster, err := api.ExecuteGetRequest[models.Monster](MonsterType, m.Name)
	if err != nil {
		panic(err)
	}
	
	return monster
}
