package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/api/responses"
)

type FeatureRequest api.BaseRequest
const FeatureType api.PathType = "features"

func HandleFeatureRequest(featureQuery string, termWidth int) error {
	r := FeatureRequest {
		Name: featureQuery,
		PathType: FeatureType,
	}		

	f, err := r.GetSingle()
	if err != nil {
		logErr := fmt.Errorf("Failed to Handle Feature Request (single)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintFeatureSingle(f, termWidth)
	return nil
}

func HandleFeatureListRequest() error {
	r := FeatureRequest {
		PathType: FeatureType,
	}		

	fl, err := r.GetList()	
	if err != nil {
		logErr := fmt.Errorf("Failed to Handle Feature Request (list)")
		logger.HandleError(err, logErr)

		return err
	}

	cli.PrintFeatureList(fl)
	return nil
}

func (f *FeatureRequest) GetList() (responses.FeatureList, error) {
	featureList, err := api.ExecuteGetRequest[responses.FeatureList](FeatureType, "")
	if err != nil {
		logErr := fmt.Errorf("Failed to search Feature (list)")
		logger.HandleError(err, logErr)

		return featureList, err
	}

	return featureList, nil
}

func (f *FeatureRequest) GetSingle() (responses.Feature, error) {
	f.Name = strings.ReplaceAll(f.Name, " ", "-")

	feature := responses.Feature{}
	feature, err := api.ExecuteGetRequest[responses.Feature](EquipmentType, f.Name)
	if err != nil {
		logErr := fmt.Errorf("Failed to search Feature (single): %s", f.Name)
		logger.HandleError(err, logErr)

		return feature, err
	}

	return feature, nil
}
