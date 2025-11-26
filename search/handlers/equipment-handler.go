package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/search/api"
	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/dndgo/search/cli"
)

type EquipmentRequest api.BaseRequest

const EquipmentType api.PathType = "equipment"

func HandleEquipmentRequest(equipmentQuery string, termWidth int) (string, error) {
	r := EquipmentRequest{
		Name:     equipmentQuery,
		PathType: EquipmentType,
	}

	e, err := r.GetSingle()
	if err != nil {
		return "", fmt.Errorf("Failed to handle equipment request (%s): %w", equipmentQuery, err)
	}

	result := cli.FormatEquipmentSingle(e, termWidth)
	return result, nil
}

func HandleEquipmentListRequest() (string, error) {
	r := EquipmentRequest{
		Name:     "",
		PathType: EquipmentType,
	}

	el, err := r.GetList()
	if err != nil {
		return "", fmt.Errorf("Failed to handle equipment request list: %w", err)
	}

	result := cli.FormatEquipmentList(el)
	return result, nil
}

func (s *EquipmentRequest) GetList() (responses.EquipmentList, error) {
	equipmentList, err := api.ExecuteGetRequest[responses.EquipmentList](EquipmentType, "")
	if err != nil {
		return equipmentList, fmt.Errorf("Failed to get equipment request list: %w", err)
	}

	return equipmentList, nil
}

func (e *EquipmentRequest) GetSingle() (responses.Equipment, error) {
	e.Name = strings.ReplaceAll(e.Name, " ", "-")

	equipment := responses.Equipment{}
	equipment, err := api.ExecuteGetRequest[responses.Equipment](EquipmentType, e.Name)
	if err != nil {
		return equipment, fmt.Errorf("Failed to get equipment (%s): %w", e.Name, err)
	}

	return equipment, nil
}
