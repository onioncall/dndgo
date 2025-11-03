package types

type WornEquipment struct {
	Head      string `json:"head"`
	Amulet    string `json:"amulet"`
	Cloak     string `json:"cloak"`
	Armor     string `json:"armour"`
	HandsArms string `json:"hands-arms"`
	Ring      string `json:"ring"`
	Ring2     string `json:"ring2"`
	Belt      string `json:"belt"`
	Boots     string `json:"boots"`
}

type BackpackItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

const (
	WornEquipmentHead      string = "head"
	WornEquipmentAmulet    string = "amulet"
	WornEquipmentCloak     string = "cloak"
	WornEquipmentArmour    string = "armor"
	WornEquipmentHandsArms string = "hands-arms"
	WornEquipmentRing      string = "ring"
	WornEquipmentRing2     string = "ring2"
	WornEquipmentBelt      string = "belt"
	WornEquipmentBoots     string = "boots"
)
