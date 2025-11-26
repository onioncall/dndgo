package handlers

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/search/api"
	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/dndgo/search/format"
)

type FeatureRequest api.BaseRequest

const FeatureType api.PathType = "features"

func HandleFeatureRequest(featureQuery string, termWidth int) (string, error) {
	r := FeatureRequest{
		Name:     featureQuery,
		PathType: FeatureType,
	}

	f, err := r.GetSingle()
	if err != nil {
		return "", fmt.Errorf("Failed to handle feature request (%s): %w", featureQuery, err)
	}

	result := format.FormatFeatureSingle(f, termWidth)
	return result, nil
}

func HandleFeatureListRequest() (string, error) {
	r := FeatureRequest{
		PathType: FeatureType,
	}

	fl, err := r.GetList()
	if err != nil {
		return "", fmt.Errorf("Failed to handle feature request list: %w", err)
	}

	result := format.FormatFeatureList(fl)
	return result, nil
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
