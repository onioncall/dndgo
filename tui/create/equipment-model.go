package create

import (
	"github.com/charmbracelet/bubbles/textinput"
)

const (
	headInput = iota
	amuletInput
	cloakInput
	handsArmsInput
	ringInput
	ring2Input
	beltInput
	bootsInput
	shieldInput
	armorInput
	armorProficientInput
	armorClassInput
	armorTypeInput
)

func equipmentInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 13)

	inputs[headInput] = textinput.New()
	inputs[headInput].Placeholder = "Circlet of Blasting"
	inputs[headInput].Focus()
	inputs[headInput].Width = 40
	inputs[headInput].Prompt = ""

	inputs[amuletInput] = textinput.New()
	inputs[amuletInput].Placeholder = "Amulet of Health"
	inputs[amuletInput].Width = 40
	inputs[amuletInput].Prompt = ""

	inputs[cloakInput] = textinput.New()
	inputs[cloakInput].Placeholder = "Cloak of Protection"
	inputs[cloakInput].Width = 40
	inputs[cloakInput].Prompt = ""

	inputs[handsArmsInput] = textinput.New()
	inputs[handsArmsInput].Placeholder = "Gauntlets of Ogre Power"
	inputs[handsArmsInput].Width = 40
	inputs[handsArmsInput].Prompt = ""

	inputs[ringInput] = textinput.New()
	inputs[ringInput].Placeholder = "Ring of Protection"
	inputs[ringInput].Width = 40
	inputs[ringInput].Prompt = ""

	inputs[ring2Input] = textinput.New()
	inputs[ring2Input].Placeholder = "Ring of Spell Storing"
	inputs[ring2Input].Width = 40
	inputs[ring2Input].Prompt = ""

	inputs[beltInput] = textinput.New()
	inputs[beltInput].Placeholder = "Belt of Giant Strength"
	inputs[beltInput].Width = 40
	inputs[beltInput].Prompt = ""

	inputs[bootsInput] = textinput.New()
	inputs[bootsInput].Placeholder = "Boots of Speed"
	inputs[bootsInput].Width = 40
	inputs[bootsInput].Prompt = ""

	inputs[shieldInput] = textinput.New()
	inputs[shieldInput].Placeholder = "Shield"
	inputs[shieldInput].Width = 40
	inputs[shieldInput].Prompt = ""

	inputs[armorInput] = textinput.New()
	inputs[armorInput].Placeholder = "Leather Armor"
	inputs[armorInput].Width = 40
	inputs[armorInput].Prompt = ""

	inputs[armorProficientInput] = textinput.New()
	inputs[armorProficientInput].Placeholder = "true"
	inputs[armorProficientInput].Width = 40
	inputs[armorProficientInput].Prompt = ""

	inputs[armorClassInput] = textinput.New()
	inputs[armorClassInput].Placeholder = "11"
	inputs[armorClassInput].Width = 40
	inputs[armorClassInput].Prompt = ""

	inputs[armorTypeInput] = textinput.New()
	inputs[armorTypeInput].Placeholder = "Light"
	inputs[armorTypeInput].Width = 40
	inputs[armorTypeInput].Prompt = ""

	return inputs
}
