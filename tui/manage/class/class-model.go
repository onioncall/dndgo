package class

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/wrapt"
)

type ClassModel struct {
	SubClassViewport      viewport.Model
	DetailViewport        viewport.Model
	OtherFeaturesViewport viewport.Model
	contentSet            bool
}

func NewClassModel() ClassModel {
	subClassViewPort := viewport.New(0, 0)
	detailViewPort := viewport.New(0, 0)
	otherFeaturesViewPort := viewport.New(0, 0)

	return ClassModel{
		SubClassViewport:      subClassViewPort,
		DetailViewport:        detailViewPort,
		OtherFeaturesViewport: otherFeaturesViewPort,
	}
}

func GetClassFeatures(c models.Character) string {
	features := ""

	for _, class := range c.Classes {
		features += fmt.Sprintf("%s\n\n", class.GetClassFeatures())
	}

	return features
}

func GetClassDetails(currentClass string, c models.Character) string {
	classDetails := ""

	if len(c.ClassTypes) > 1 {
		classDetails += fmt.Sprintf("Current Class: %s\n", currentClass)
	}

	for _, class := range c.Classes {
		classDetails += class.ClassDetails()
	}

	return classDetails
}

func GetSubClass(c models.Character) string {
	subClassLine := "Sub-Class | "
	subClassCount := 0
	for i, class := range c.Classes {
		subClass := class.GetSubClass()

		if subClass == "" {
			continue
		}

		subClassLine += subClass
		if i < len(c.Classes)-1 {
			subClassLine += ", "
		}

		subClassCount++
	}

	if subClassCount == 0 {
		subClassLine = "No sub-classes configured for character"
	}

	return subClassLine
}

func (m ClassModel) UpdateSize(innerWidth, availableHeight int, currentClass string, character models.Character) ClassModel {
	// Column 1: 50% width, split 50/50 vertically
	col1Width := innerWidth / 2

	subClassHeight := (availableHeight * 15) / 100
	detailHeight := availableHeight - subClassHeight

	subClassInnerWidth := col1Width - 2
	subClassInnerHeight := subClassHeight - 2
	detailInnerWidth := col1Width - 2
	detailInnerHeight := detailHeight - 2

	m.SubClassViewport.Width = subClassInnerWidth
	m.SubClassViewport.Height = subClassInnerHeight
	m.DetailViewport.Width = detailInnerWidth
	m.DetailViewport.Height = detailInnerHeight

	// Column 2: 50% width, full height
	col2Width := innerWidth / 2
	otherFeaturesInnerWidth := col2Width - 6
	otherFeaturesInnerHeight := availableHeight - 2

	m.OtherFeaturesViewport.Width = otherFeaturesInnerWidth
	m.OtherFeaturesViewport.Height = otherFeaturesInnerHeight

	if !m.contentSet {
		subClassContent := GetSubClass(character)
		subClassContent = wrapt.Wrap(subClassContent, m.SubClassViewport.Width)
		if subClassContent == "" {
			subClassContent = "Class has no sub class to show"
		}
		m.SubClassViewport.SetContent(subClassContent)

		classFeaturesContent := "Class Features\n\n"
		classFeaturesContent += GetClassFeatures(character)
		classFeaturesContent = strings.ReplaceAll(classFeaturesContent, "---", "")
		classFeaturesContent = wrapt.Wrap(classFeaturesContent, m.OtherFeaturesViewport.Width)
		if classFeaturesContent == "" {
			classFeaturesContent = "Class has no features yet"
		}
		m.OtherFeaturesViewport.SetContent(classFeaturesContent)

		classDetailsContent := "Class Details\n\n"
		classDetailsContent += GetClassDetails(currentClass, character)
		classDetailsContent = wrapt.Wrap(classDetailsContent, m.DetailViewport.Width)
		if classDetailsContent == "" {
			classDetailsContent = "Class has no details to show"
		}
		m.DetailViewport.SetContent(classDetailsContent)
	}

	return m
}
