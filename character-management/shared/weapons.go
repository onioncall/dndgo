package shared

type Weapon struct {
	Name        string      `json:"name"`
	Bonus       int         `json:"-"`
	CustomBonus int         `json:"custom-bonus"`
	Proficient  bool        `json:"proficient"`
	Damage      string      `json:"damage"`
	Ranged      bool        `json:"ranged"`
	Range       WeaponRange `json:"range"`
	Type        string      `json:"type"`
	Properties  []string    `json:"properties"`
}

type WeaponRange struct {
	NormalRange int `json:"normal-range"`
	LongRange   int `json:"long-range"`
}

// Properties
const (
	WeaponPropertyAmmunition string = "ammunition"
	WeaponPropertyFinesse    string = "finesse"
	WeaponPropertyHeavy      string = "heavy"
	WeaponPropertyLight      string = "light"
	WeaponPropertyLoading    string = "loading"
	WeaponPropertyReach      string = "reach"
	WeaponPropertySpecial    string = "special"
	WeaponPropertyThrown     string = "thrown"
	WeaponPropertyTwoHanded  string = "two-handed"
	WeaponPropertyVersatile  string = "versatile"
	WeaponPropertyMonk       string = "monk"
)
