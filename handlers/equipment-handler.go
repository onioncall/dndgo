package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/api/responses"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/logger"
)

type EquipmentRequest api.BaseRequest

const EquipmentType api.PathType = "equipment"

func HandleEquipmentRequest(equipmentQuery string, termWidth int) error {
	r := EquipmentRequest{
		Name:     equipmentQuery,
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
	r := EquipmentRequest{
		Name:     "",
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

func (s *EquipmentRequest) GetList() (responses.EquipmentList, error) {
	equipmentList, err := api.ExecuteGetRequest[responses.EquipmentList](EquipmentType, "")
	if err != nil {
		logErr := fmt.Errorf("Failed to search Equipment (list)")
		logger.HandleError(err, logErr)

		return equipmentList, err
	}

	return equipmentList, nil
}

func (e *EquipmentRequest) GetSingle() (responses.Equipment, error) {
	e.Name = strings.ReplaceAll(e.Name, " ", "-")

	equipment := responses.Equipment{}
	equipment, err := api.ExecuteGetRequest[responses.Equipment](EquipmentType, e.Name)
	if err != nil {
		logErr := fmt.Errorf("Failed to search Equipment (single): %s", e.Name)
		logger.HandleError(err, logErr)

		return equipment, err
	}

	return equipment, nil
}
