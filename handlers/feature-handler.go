package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/api/responses"
	"github.com/onioncall/dndgo/cli"
)

type FeatureRequest api.BaseRequest

const FeatureType api.PathType = "features"

func HandleFeatureRequest(featureQuery string, termWidth int) error {
	r := FeatureRequest{
		Name:     featureQuery,
		PathType: FeatureType,
	}

	f, err := r.GetSingle()
	if err != nil {
		return fmt.Errorf("Failed to handle feature request (%s): %w", featureQuery, err)
	}

	cli.PrintFeatureSingle(f, termWidth)
	return nil
}

func HandleFeatureListRequest() error {
	r := FeatureRequest{
		PathType: FeatureType,
	}

	fl, err := r.GetList()
	if err != nil {
		return fmt.Errorf("Failed to handle feature request list: %w", err)
	}

	cli.PrintFeatureList(fl)
	return nil
}

func (f *FeatureRequest) GetList() (responses.FeatureList, error) {
	featureList, err := api.ExecuteGetRequest[responses.FeatureList](FeatureType, "")
	if err != nil {
		return featureList, fmt.Errorf("Failed to search feature list: %w", err)
	}

	return featureList, nil
}

func (f *FeatureRequest) GetSingle() (responses.Feature, error) {
	f.Name = strings.ReplaceAll(f.Name, " ", "-")

	feature := responses.Feature{}
	feature, err := api.ExecuteGetRequest[responses.Feature](EquipmentType, f.Name)
	if err != nil {
		return feature, fmt.Errorf("Failed to search feature (%s): %w", f.Name, err)
	}

	return feature, nil
}
