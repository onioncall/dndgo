package handlers

import (
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/models"
)

type EquipmentRequest api.BaseRequest
const EquipmentType api.PathType = "equipment"

func HandleEquipmentRequest(equipmentQuery string, termWidth int) {
	r := EquipmentRequest {
		Name: equipmentQuery,
		PathType: EquipmentType,
	}		

	e := r.GetSingle()
	cli.PrintEquipmentSingle(e, termWidth)
}

func HandleEquipmentListRequest() {
	r := EquipmentRequest {
		Name: "",
		PathType: EquipmentType,
	}		

	el := r.GetList()	
	cli.PrintEquipmentList(el)
}

func (s *EquipmentRequest) GetList() models.EquipmentList {
	equipmentList, err := api.ExecuteGetRequest[models.EquipmentList](EquipmentType, "")
	if err != nil {
		panic(err)
	}

	return equipmentList
}

func (e *EquipmentRequest) GetSingle() models.Equipment {
	e.Name = strings.ReplaceAll(e.Name, " ", "-")

	equipment := models.Equipment{}
	equipment, err := api.ExecuteGetRequest[models.Equipment](EquipmentType, e.Name)
	if err != nil {
		panic(err)
	}

	return equipment
}
