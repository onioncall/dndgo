package class

import (
	"encoding/json"
	"fmt"

	"github.com/onioncall/dndgo/logger"
	"github.com/onioncall/dndgo/models"
)

type Rogue struct {
	OtherFeatures []models.ClassFeatures `json:"other-features"`
}

func LoadRogue(data []byte) (*Rogue, error) {
	var ranger Rogue
	if err := json.Unmarshal(data, &ranger); err != nil {
		errLog := fmt.Errorf("Failed to parse class data: %w", err)
		logger.HandleError(errLog, err)

		return nil, err
	}

	return &ranger, nil
}

func (r *Rogue) LoadMethods() {
}

func (r *Rogue) ExecutePostCalculateMethods(c *models.Character) {
	for _, m := range models.PostCalculateMethods {
		m(c)
	}
}

func (r *Rogue) ExecutePreCalculateMethods(c *models.Character) {
	for _, m := range models.PreCalculateMethods {
		m(c)
	}
}

func (r *Rogue) PrintClassDetails(c *models.Character) []string {
	s := c.BuildClassDetailsHeader()

	if len(r.OtherFeatures) > 0 {
		for _, detail := range r.OtherFeatures {
			if detail.Level > c.Level {
				continue
			}

			name := fmt.Sprintf("---\n**%s**\n", detail.Name)
			s = append(s, name)
			detail := fmt.Sprintf("%s\n", detail.Details)
			s = append(s, detail)
		}
	}

	return s
}
