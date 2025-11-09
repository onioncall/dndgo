package types

type SpellSlot struct {
	Level     int `json:"level"`
	Maximum   int `json:"maximum"`
	Available int `json:"available"`
}

type CharacterSpell struct {
	SlotLevel  int    `json:"slot-level"`
	IsRitual   bool   `json:"ritual"`
	Name       string `json:"name"`
	IsPrepared bool   `json:"-"`
}

type Token struct {
	Maximum   int `json:"maximum"`
	Available int `json:"available"`
}

type PreSetToken struct {
	Maximum   int `json:"-"`
	Available int `json:"available"`
}

type NamedToken struct {
	Name      string `json:"name"`
	Maximum   int    `json:"-"`
	Available int    `json:"available"`
	Level     int    `json:"level"`
}
