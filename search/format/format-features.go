package cli

import (
	"fmt"
	"strings"

	"github.com/onioncall/dndgo/search/api/responses"
	"github.com/onioncall/wrapt"
)

func FormatFeatureSingle(feature responses.Feature, termWidth int) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("%s\n\n", feature.Name))
	builder.WriteString(fmt.Sprintf("Class: %s\n\n", feature.Class.Name))
	for _, description := range feature.Desc {
		builder.WriteString(fmt.Sprintf("%s\n\n", wrapt.Wrap(description, termWidth)))
	}

	return builder.String()
}

func FormatFeatureList(featureList responses.FeatureList) string {
	var builder strings.Builder

	fmt.Print("Feature Name\n\n")
	for _, feature := range featureList.ListItems {
		builder.WriteString(fmt.Sprintf("%s - %s\n", feature.Name, feature.Index))
	}

	return builder.String()
}
