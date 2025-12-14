package class

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/wrapt"
)

type ClassModel struct {
	SubClassViewPort      viewport.Model
	DetailViewport        viewport.Model
	OtherFeaturesViewPort viewport.Model
	contentSet            bool
}

func NewClassModel() ClassModel {
	subClassViewPort := viewport.New(0, 0)
	detailViewPort := viewport.New(0, 0)
	otherFeaturesViewPort := viewport.New(0, 0)

	return ClassModel{
		SubClassViewPort:      subClassViewPort,
		DetailViewport:        detailViewPort,
		OtherFeaturesViewPort: otherFeaturesViewPort,
	}
}

func GetClassFeatures(c models.Character) string {
	return c.Class.ClassFeatures(c.Level)
}

func GetClassDetails(c models.Character) string {
	return c.Class.ClassDetails(c.Level)
}

func GetSubClass(c models.Character) string {
	return c.Class.SubClass(c.Level)
}

func (m ClassModel) UpdateSize(innerWidth, availableHeight int, character models.Character) ClassModel {
	// Column 1: 50% width, split 50/50 vertically
	col1Width := innerWidth / 2

	subClassHeight := (availableHeight * 15) / 100
	detailHeight := availableHeight - subClassHeight

	subClassInnerWidth := col1Width - 2
	subClassInnerHeight := subClassHeight - 2
	detailInnerWidth := col1Width - 2
	detailInnerHeight := detailHeight - 2

	m.SubClassViewPort.Width = subClassInnerWidth
	m.SubClassViewPort.Height = subClassInnerHeight
	m.DetailViewport.Width = detailInnerWidth
	m.DetailViewport.Height = detailInnerHeight

	// Column 2: 50% width, full height
	col2Width := innerWidth / 2
	otherFeaturesInnerWidth := col2Width - 6
	otherFeaturesInnerHeight := availableHeight - 2

	m.OtherFeaturesViewPort.Width = otherFeaturesInnerWidth
	m.OtherFeaturesViewPort.Height = otherFeaturesInnerHeight

	if !m.contentSet {
		subClassContent := GetSubClass(character)
		subClassContent = wrapt.Wrap(fmt.Sprintf("Sub-Class: %s", subClassContent), m.SubClassViewPort.Width)
		if subClassContent == "" {
			subClassContent = "Class has no sub class to show"
		}
		m.SubClassViewPort.SetContent(subClassContent)

		classFeaturesContent := "Class Features\n\n"
		classFeaturesContent += GetClassFeatures(character)
		classFeaturesContent = strings.ReplaceAll(classFeaturesContent, "---", "")
		classFeaturesContent = wrapt.Wrap(classFeaturesContent, m.OtherFeaturesViewPort.Width)
		if classFeaturesContent == "" {
			classFeaturesContent = "Class has no features yet"
		}
		m.OtherFeaturesViewPort.SetContent(classFeaturesContent)

		classDetailsContent := "Class Details\n\n"
		classDetailsContent += GetClassDetails(character)
		classDetailsContent = wrapt.Wrap(classDetailsContent, m.DetailViewport.Width)
		if classDetailsContent == "" {
			classDetailsContent = "Class has no details to show"
		}
		m.DetailViewport.SetContent(classDetailsContent)
	}

	return m
}
