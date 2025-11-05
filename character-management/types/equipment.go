package types

type WornEquipment struct {
	Head      string `json:"head"`
	Amulet    string `json:"amulet"`
	Cloak     string `json:"cloak"`
	Armor     Armor  `json:"armor"`
	HandsArms string `json:"hands-arms"`
	Ring      string `json:"ring"`
	Ring2     string `json:"ring2"`
	Belt      string `json:"belt"`
	Boots     string `json:"boots"`
	Shield    string `json:"shield"`
}

type Armor struct {
	Name       string `json:"name"`
	Proficient bool   `json:"proficient"`
	Class      int    `json:"class"`
	Type       string `json:"type"`
}

type BackpackItem struct {
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

type Equipped string

const (
	PrimaryEquipped   Equipped = "primary"
	SecondaryEquipped Equipped = "secondary"
	NotEquipped       Equipped = "unequipped"
)

const (
	LightArmor  string = "light"
	MediumArmor string = "medium"
	HeavyArmor  string = "heavy"
)

const (
	WornEquipmentHead      string = "head"
	WornEquipmentAmulet    string = "amulet"
	WornEquipmentCloak     string = "cloak"
	WornEquipmentArmor     string = "armor"
	WornEquipmentHandsArms string = "hands-arms"
	WornEquipmentRing      string = "ring"
	WornEquipmentRing2     string = "ring2"
	WornEquipmentBelt      string = "belt"
	WornEquipmentBoots     string = "boots"
)
