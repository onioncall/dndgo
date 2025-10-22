package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
)

type Druid struct {
	WildShape     WildShape              `json:"wild-shape"`
	OtherFeatures []models.ClassFeatures `json:"other-features"`
}

type WildShape struct {
	Available int `json:"available"`
	Maximum   int `json:"maximum"`
}

func LoadDruid(data []byte) (*Druid, error) {
	var druid Druid
	if err := json.Unmarshal(data, &druid); err != nil {
		return &druid, fmt.Errorf("Failed to parse class data: %w", err)
	}

	return &druid, nil
}

func (d *Druid) LoadMethods() {
}

func (d *Druid) ExecutePostCalculateMethods(c *models.Character) {
	for _, m := range models.PostCalculateMethods {
		m(c)
	}
}

func (d *Druid) ExecutePreCalculateMethods(c *models.Character) {
	for _, m := range models.PreCalculateMethods {
		m(c)
	}
}

func (d *Druid) ExecuteArchDruid(c *models.Character) {
	if c.Level < 20 {
		return
	}

	// These are now unlimited, no need to track them anymore
	d.WildShape.Available = 0
	d.WildShape.Maximum = 0
}

// CLI

func (d *Druid) UseClassTokens(tokenName string) {
	// We only really need slot name for classes that have multiple slots
	// since druid only has wild shape, we won't check the slot name value
	if d.WildShape.Available <= 0 {
		logger.HandleInfo("No Bardic Inspiration tokens left")
		return
	}

	d.WildShape.Available--
}

func (d *Druid) RecoverClassTokens(tokenName string, quantity int) {
	// We only really need slot name for classes that have multiple slots
	// since druid only has wild shape, we won't check the slot name value
	d.WildShape.Available += quantity

	// if no quantity is provided, or the new value exceeds the max we will perform a full recover
	if quantity == 0 || d.WildShape.Available > d.WildShape.Maximum {
		d.WildShape.Available = d.WildShape.Maximum
	}
}
