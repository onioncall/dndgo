package equipment

type BodyEquipment struct {
	Head      string `json:"head"`
	Amulet    string `json:"amulet"`
	Cloak     string `json:"cloak"`
	Armour    string `json:"armour"`
	HandsArms string `json:"hands-arms"`
	Ring      string `json:"ring"`
	Ring2     string `json:"ring2"`
	Belt      string `json:"belt"`
	Boots     string `json:"boots"`
}

type BackpackItem struct {
	Name	 string `json:"name"`
	Quantity int 	`json:"quantity"`
}

const (
	Head      string = "head"
	Amulet    string = "amulet"
	Cloak     string = "cloak"
	Armour    string = "armor"
	HandsArms string = "hands-arms"
	Ring      string = "ring"
	Ring2     string = "ring2"
	Belt      string = "belt"
	Boots     string = "boots"
)

