package types

type Weapon struct {
	Name       string   `json:"name"`
	Bonus      int      `json:"bonus"`
	Damage     string   `json:"damage"`
	Range      string   `json:"range"`
	Type       string   `json:"type"`
	Properties []string `json:"properties"`
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
