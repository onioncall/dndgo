package models

import "time"

type Equipment struct {
    Desc              []string          `json:"desc"`
    Special           []string          `json:"special"`
    Index             string            `json:"index"`
    Name              string            `json:"name"`
    EquipmentCategory *Category			`json:"equipment_category,omitempty"`
    GearCategory      *Category	    	`json:"gear_category,omitempty"`
    WeaponCategory    string         	`json:"weapon_category,omitempty"`
    WeaponRange       string            `json:"weapon_range,omitempty"`
    CategoryRange     string            `json:"category_range,omitempty"`
    Cost              Cost              `json:"cost"`
    Damage            *Damage           `json:"damage,omitempty"`
    Range             *Range            `json:"range,omitempty"`
    Weight            float64           `json:"weight"`
    Properties        []Property        `json:"properties"`
    URL               string            `json:"url"`
    UpdatedAt         time.Time         `json:"updated_at"`
    Contents          []any			    `json:"contents"`
}

type Category struct {
    Index string `json:"index"`
    Name  string `json:"name"`
    URL   string `json:"url"`
}

type Cost struct {
    Quantity int    `json:"quantity"`
    Unit     string `json:"unit"`
}

type Damage struct {
    DamageDice string     `json:"damage_dice"`
    DamageType DamageType `json:"damage_type"`
}

type DamageType struct {
    Index string `json:"index"`
    Name  string `json:"name"`
    URL   string `json:"url"`
}

type Range struct {
    Normal int `json:"normal"`
    Long   int `json:"long,omitempty"`
}

type Property struct {
    Index string `json:"index"`
    Name  string `json:"name"`
    URL   string `json:"url"`
}

type EquipmentList struct {
	ListItems []EquipmentListItem `json:"results"`
}

type EquipmentListItem struct {
	Index 	string	`json:"index"`
	Name	string	`json:"name"`
	Url		string	`json:"url"`
}
