package shared

type Weapon struct {
	Name        string      `json:"name" clover:"name"`
	Bonus       int         `json:"-" clover:"-"`
	CustomBonus int         `json:"custom-bonus" clover:"custom-bonus"`
	Proficient  bool        `json:"proficient" clover:"proficient"`
	Damage      string      `json:"damage" clover:"damage"`
	Ranged      bool        `json:"ranged" clover:"ranged"`
	Range       WeaponRange `json:"range" clover:"range"`
	Type        string      `json:"type" clover:"type"`
	Properties  []string    `json:"properties" clover:"properties"`
}

type WeaponRange struct {
	NormalRange int `json:"normal-range" clover:"normal-range"`
	LongRange   int `json:"long-range" clover:"long-range"`
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
