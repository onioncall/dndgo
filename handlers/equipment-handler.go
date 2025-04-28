package handlers

import (
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/models"
)

type EquipmentRequest api.BaseRequest
const EquipmentType api.PathType = "equipment"

func HandleEquipmentRequest(equipmentQuery string) {
	r := EquipmentRequest {
		Name: equipmentQuery,
		PathType: EquipmentType,
	}		

	e := r.GetSingle()
	cli.PrintEquipmentSingle(e)
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

func (s *EquipmentRequest) GetSingle() models.Equipment {
	s.Name = strings.ReplaceAll(s.Name, " ", "-")

	equipment := models.Equipment{}
	equipment, err := api.ExecuteGetRequest[models.Equipment](EquipmentType, s.Name)
	if err != nil {
		panic(err)
	}

	return equipment
}
