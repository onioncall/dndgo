package shared

type SpellSlot struct {
	Level     int `json:"level" clover:"level"`
	Maximum   int `json:"maximum" clover:"maximum"`
	Available int `json:"available" clover:"available"`
}

type CharacterSpell struct {
	SlotLevel  int    `json:"slot-level" clover:"slot-level"`
	IsRitual   bool   `json:"ritual" clover:"ritual"`
	Name       string `json:"name" clover:"name"`
	IsPrepared bool   `json:"-" clover:"-"`
}

type Token struct {
	Maximum   int `json:"maximum" clover:"maximum"`
	Available int `json:"available" clover:"available"`
}

type PreSetToken struct {
	Maximum   int `json:"-" clover:"-"`
	Available int `json:"available" clover:"available"`
}

type NamedToken struct {
	Name      string `json:"name" clover:"name"`
	Maximum   int    `json:"-" clover:"-"`
	Available int    `json:"available" clover:"available"`
	Level     int    `json:"level" clover:"level"`
}
