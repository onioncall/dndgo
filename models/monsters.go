package models

import "time"

type Monster struct {
	Index                  string                `json:"index"`
	Name                   string                `json:"name"`
	Size                   string                `json:"size"`
	Type                   string                `json:"type"`
	Alignment              string                `json:"alignment"`
	ArmorClass             []ArmorClass          `json:"armor_class"`
	HitPoints              int                   `json:"hit_points"`
	HitDice                string                `json:"hit_dice"`
	HitPointsRoll          string                `json:"hit_points_roll"`
	Speed                  Speed                 `json:"speed"`
	Strength               int                   `json:"strength"`
	Dexterity              int                   `json:"dexterity"`
	Constitution           int                   `json:"constitution"`
	Intelligence           int                   `json:"intelligence"`
	Wisdom                 int                   `json:"wisdom"`
	Charisma               int                   `json:"charisma"`
	Proficiencies          []Proficiency         `json:"proficiencies"`
	DamageVulnerabilities  []string              `json:"damage_vulnerabilities"`
	DamageResistances      []string              `json:"damage_resistances"`
	DamageImmunities       []string              `json:"damage_immunities"`
	ConditionImmunities    []string              `json:"condition_immunities"`
	Senses                 Senses                `json:"senses"`
	Languages              string                `json:"languages"`
	ChallengeRating        float64               `json:"challenge_rating"`
	ProficiencyBonus       int                   `json:"proficiency_bonus"`
	XP                     int                   `json:"xp"`
	SpecialAbilities       []SpecialAbility      `json:"special_abilities"`
	Actions                []Action              `json:"actions"`
	LegendaryActions       []LegendaryAction     `json:"legendary_actions,omitempty"`
	Image                  string                `json:"image"`
	URL                    string                `json:"url"`
	UpdatedAt              time.Time             `json:"updated_at"`
}

type MonsterList struct {
	ListItems []Reference `json:"results"`
}

type ArmorClass struct {
	Type  string `json:"type"`
	Value int    `json:"value"`
}

type Speed struct {
	Walk string `json:"walk,omitempty"`
	Fly  string `json:"fly,omitempty"`
	Swim string `json:"swim,omitempty"`
}

type Proficiency struct {
	Value       int        `json:"value"`
	Proficiency Reference `json:"proficiency"`
}

type Senses struct {
	Blindsight        string `json:"blindsight,omitempty"`
	Darkvision        string `json:"darkvision,omitempty"`
	PassivePerception int    `json:"passive_perception"`
}

type SpecialAbility struct {
	Name   string    `json:"name"`
	Desc   string    `json:"desc"`
	Usage  *Usage    `json:"usage,omitempty"`
}

type Action struct {
	Name            string           `json:"name"`
	MultiattackType string           `json:"multiattack_type,omitempty"`
	Desc            string           `json:"desc"`
	AttackBonus     int              `json:"attack_bonus,omitempty"`
	Damage          []DamageDetails  `json:"damage,omitempty"`
	DC              *DC              `json:"dc,omitempty"`
	Usage           *Usage           `json:"usage,omitempty"`
	Actions         []SubAction      `json:"actions,omitempty"`
}

type LegendaryAction struct {
	Name   string `json:"name"`
	Desc   string `json:"desc"`
	DC     *DC    `json:"dc,omitempty"`
	Damage []DamageDetails `json:"damage,omitempty"`
}

type SubAction struct {
	ActionName string `json:"action_name"`
	Count      string `json:"count"`
	Type       string `json:"type"`
}

type DamageDetails struct {
	DamageType Reference `json:"damage_type"`
	DamageDice string    `json:"damage_dice"`
}

type DC struct {
	DCType      Reference `json:"dc_type"`
	DCValue     int       `json:"dc_value"`
	SuccessType string    `json:"success_type"`
}

type Usage struct {
	Type      string   `json:"type"`
	Times     int      `json:"times,omitempty"`
	Dice      string   `json:"dice,omitempty"`
	MinValue  int      `json:"min_value,omitempty"`
	RestTypes []string `json:"rest_types,omitempty"`
}
