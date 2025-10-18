package responses

import "time"

type Equipment struct {
    Desc              []string          `json:"desc"`
    Special           []string          `json:"special"`
    Index             string            `json:"index"`
    Name              string            `json:"name"`
    EquipmentCategory *Reference		`json:"equipment_category,omitempty"`
    GearCategory      *Reference	   	`json:"gear_category,omitempty"`
    WeaponCategory    string         	`json:"weapon_category,omitempty"`
    WeaponRange       string            `json:"weapon_range,omitempty"`
    CategoryRange     string            `json:"category_range,omitempty"`
    Cost              Cost              `json:"cost"`
    Damage            *Damage           `json:"damage,omitempty"`
    Range             *Range            `json:"range,omitempty"`
    Weight            float64           `json:"weight"`
    Properties        []Reference        `json:"properties"`
    URL               string            `json:"url"`
    UpdatedAt         time.Time         `json:"updated_at"`
    Contents          []any			    `json:"contents"`
}

type Cost struct {
    Quantity int    `json:"quantity"`
    Unit     string `json:"unit"`
}

type Damage struct {
    DamageDice string     `json:"damage_dice"`
    DamageType Reference  `json:"damage_type"`
}

type Range struct {
    Normal int `json:"normal"`
    Long   int `json:"long,omitempty"`
}

type EquipmentList struct {
	ListItems []Reference `json:"results"`
}
