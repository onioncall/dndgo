package handlers

import (
	"strings"

	"github.com/onioncall/dndgo/api"
	"github.com/onioncall/dndgo/cli"
	"github.com/onioncall/dndgo/models"
)

type FeatureRequest api.BaseRequest
const FeatureType api.PathType = "features"

func HandleFeatureRequest(featureQuery string, termWidth int) {
	r := FeatureRequest {
		Name: featureQuery,
		PathType: FeatureType,
	}		

	f := r.GetSingle()
	cli.PrintFeatureSingle(f, termWidth)
}

func HandleFeatureListRequest() {
	r := FeatureRequest {
		PathType: FeatureType,
	}		

	fl := r.GetList()	
	cli.PrintFeatureList(fl)
}

func (f *FeatureRequest) GetList() models.FeatureList {
	featureList, err := api.ExecuteGetRequest[models.FeatureList](FeatureType, "")
	if err != nil {
		panic(err)
	}

	return featureList
}

func (f *FeatureRequest) GetSingle() models.Feature {
	f.Name = strings.ReplaceAll(f.Name, " ", "-")

	feature := models.Feature{}
	feature, err := api.ExecuteGetRequest[models.Feature](EquipmentType, f.Name)
	if err != nil {
		panic(err)
	}

	return feature
}
