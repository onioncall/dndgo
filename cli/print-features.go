package cli

import (
	"fmt"

	"github.com/onioncall/dndgo/models"
	"github.com/onioncall/wrapt"
)

func PrintFeatureSingle(feature models.Feature) {
	fmt.Printf("%s\n\n", feature.Name)
	fmt.Printf("Class: %s\n\n", feature.Class.Name)
	for _, description := range feature.Desc {
		fmt.Printf("%s\n\n", wrapt.Wrap(description))	
	}
}

func PrintFeatureList(featureList models.FeatureList) {
	fmt.Print("Feature Name\n\n")
	for _, feature := range featureList.ListItems {
		fmt.Printf("%s - %s\n", feature.Name, feature.Index)	
	}
}
