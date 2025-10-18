package responses

import "time"

type Feature struct {
    Index         string     `json:"index"`
    Class         *Reference `json:"class,omitempty"`
    Subclass      *Reference `json:"subclass,omitempty"`
    Name          string     `json:"name"`
    Level         int        `json:"level"`
    Prerequisites []string   `json:"prerequisites"`
    Desc          []string   `json:"desc"`
    URL           string     `json:"url"`
    UpdatedAt     time.Time  `json:"updated_at"`
}

type FeatureList struct {
	ListItems []Reference `json:"result"`
}
