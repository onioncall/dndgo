package types

type SpellSlot struct {
	Level     int `json:"level"`
	Maximum   int `json:"maximum"`
	Available int `json:"available"`
}

type CharacterSpell struct {
	IsCaltrop bool   `json:"is-caltrop"`
	SlotLevel int    `json:"slot-level"`
	IsRitual  bool   `json:"ritual"`
	Name      string `json:"name"`
}
