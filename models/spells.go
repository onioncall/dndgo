package models 

import (
	"time"
)

type Spell struct {
	HigherLevel    []string    `json:"higher_level"`
	Index         string      `json:"index"`
	Name          string      `json:"name"`
	Description   []string    `json:"desc"`
	Range         string      `json:"range"`
	Components    []string    `json:"components"`
	Ritual        bool        `json:"ritual"`
	Duration      string      `json:"duration"`
	Concentration bool        `json:"concentration"`
	CastingTime   string      `json:"casting_time"`
	Level         int         `json:"level"`
	AreaOfEffect  AreaOfEffect `json:"area_of_effect"`
	School        School      `json:"school"`
	Classes       []Reference `json:"classes"`
	Subclasses    []Reference `json:"subclasses"`
	URL           string      `json:"url"`
	UpdatedAt     time.Time   `json:"updated_at"`
	Damage        *SpellDamage `json:"damage,omitempty"`
}

type SpellListItem struct {
	Index 	string	`json:"index"`
	Name	string	`json:"name"`
	Level	int		`json:"level"`
	Url		string	`json:"url"`
}

type SpellList struct {
	ListItems	[]SpellListItem 	`json:"results"`
}

type AreaOfEffect struct {
	Type string `json:"type"`
	Size int    `json:"size"`
}

type School struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type Reference struct {
	Index string `json:"index"`
	Name  string `json:"name"`
	URL   string `json:"url"`
}

type SpellDamage struct {
	DamageType        Reference          `json:"damage_type"`
	DamageAtSlotLevel map[int]string `json:"damage_at_slot_level"`
}

