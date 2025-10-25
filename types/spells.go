package types

type SpellSlot struct {
	Level     int `json:"level"`
	Maximum   int `json:"maximum"`
	Available int `json:"available"`
}

type CharacterSpell struct {
	SlotLevel int    `json:"slot-level"`
	IsRitual  bool   `json:"ritual"`
	Name      string `json:"name"`
}

type Token struct {
	Maximum   int `json:"maximum"`
	Available int `json:"available"`
}
