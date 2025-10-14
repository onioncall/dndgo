package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
)

type EquipmentRequest api.BaseRequest
const EquipmentType api.PathType = "equipment"

func HandleEquipmentRequest(equipmentQuery string, termWidth int) error {
	r := EquipmentRequest {
		Name: equipmentQuery,
		PathType: EquipmentType,
	}		

	e, err := r.GetSingle()	
	if err != nil {
		logErr := fmt.Errorf("Failed to Handle Equipment Request (single)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintEquipmentSingle(e, termWidth)
	return nil
}

func HandleEquipmentListRequest() error {
	r := EquipmentRequest {
		Name: "",
		PathType: EquipmentType,
	}		

	el, err := r.GetList()	
	if err != nil {
		logErr := fmt.Errorf("Failed to Handle Equipment Request (list)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintEquipmentList(el)
	return nil
}

func (s *EquipmentRequest) GetList() (models.EquipmentList, error) {
	equipmentList, err := api.ExecuteGetRequest[models.EquipmentList](EquipmentType, "")
	if err != nil {
		logErr := fmt.Errorf("Failed to search Equipment (list)")
		logger.HandleError(err, logErr)

		return equipmentList, err
	}

	return equipmentList, nil
}

func (e *EquipmentRequest) GetSingle() (models.Equipment, error) {
	e.Name = strings.ReplaceAll(e.Name, " ", "-")

	equipment := models.Equipment{}
	equipment, err := api.ExecuteGetRequest[models.Equipment](EquipmentType, e.Name)
	if err != nil {
		logErr := fmt.Errorf("Failed to search Equipment (single): %s", e.Name)
		logger.HandleError(err, logErr)

		return equipment, err
	}

	return equipment, nil
}
