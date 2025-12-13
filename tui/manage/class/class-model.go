package class

import (
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/onioncall/dndgo/character-management/models"
	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/wrapt"
)

type ClassModel struct {
	TokenViewPort         viewport.Model
	DetailViewPort        viewport.Model
	OtherFeaturesViewPort viewport.Model
	contentSet            bool
}

func NewClassModel() ClassModel {
	tokenViewPort := viewport.New(0, 0)
	detailViewPort := viewport.New(0, 0)
	otherFeaturesViewPort := viewport.New(0, 0)

	return ClassModel{
		TokenViewPort:         tokenViewPort,
		DetailViewPort:        detailViewPort,
		OtherFeaturesViewPort: otherFeaturesViewPort,
	}
}

func getClassFeatures(c models.Character) string {
	return c.Class.ClassFeatures(c.Level)
}

func (m ClassModel) UpdateSize(innerWidth, availableHeight int, character models.Character) ClassModel {
	// Column 1: 50% width, split 50/50 vertically
	col1Width := innerWidth / 2
	col1Height := availableHeight / 2

	tokenInnerWidth := col1Width - 2
	tokenInnerHeight := col1Height - 2
	detailInnerWidth := col1Width - 2
	detailInnerHeight := col1Height - 2

	m.TokenViewPort.Width = tokenInnerWidth
	m.TokenViewPort.Height = tokenInnerHeight
	m.DetailViewPort.Width = detailInnerWidth
	m.DetailViewPort.Height = detailInnerHeight

	// Column 2: 50% width, full height
	col2Width := innerWidth / 2
	otherFeaturesInnerWidth := col2Width - 2
	otherFeaturesInnerHeight := availableHeight - 2

	m.OtherFeaturesViewPort.Width = otherFeaturesInnerWidth
	m.OtherFeaturesViewPort.Height = otherFeaturesInnerHeight

	if !m.contentSet {
		classFeaturesContent := getClassFeatures(character)
		if classFeaturesContent == "" {
			classFeaturesContent = "Class has no features yet"
		}
		classFeaturesContent = wrapt.Wrap(classFeaturesContent, m.TokenViewPort.Width)
		logger.Info(classFeaturesContent)

		m.OtherFeaturesViewPort.SetContent(classFeaturesContent)
	}

	return m
}
