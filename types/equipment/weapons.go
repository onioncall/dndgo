package equipment

type Weapon struct {
	Name   		string 		`json:"name"`
	Bonus  		int    		`json:"bonus"`
	Damage 		string 		`json:"damage"`
	Range		string 		`json:"range"`
	Type   		string 		`json:"type"`
	Properties 	[]string 	`json:"properties"`
}

// Properties
const (
	Ammunition		string = "ammunition"
	Finesse			string = "finesse"
	Heavy			string = "heavy"
	Light			string = "light"
	Loading			string = "loading"
	Reach			string = "reach"
	Special			string = "special"
	Thrown			string = "thrown"
	TwoHanded		string = "two-handed"
	Versatile		string = "versatile"
	Monk			string = "monk"
)
