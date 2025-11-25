package create

import (
	"github.com/charmbracelet/bubbles/textinput"
)

var abilityInputsMap = map[int]string{
	strength:     "Strength",
	dexterity:    "Dexterity",
	constitution: "Constitution",
	intelligence: "Intelligence",
	wisdom:       "Wisdom",
	charisma:     "Charisma",
}

const (
	strength = iota
	dexterity
	constitution
	intelligence
	wisdom
	charisma
)

const (
	strengthBaseInput = iota
	strengthProficientInput
	dexterityBaseInput
	dexterityProficientInput
	constitutionBaseInput
	constitutionProficientInput
	intelligenceBaseInput
	intelligenceProficientInput
	wisdomBaseInput
	wisdomProficientInput
	charismaBaseInput
	charismaProficientInput
)

func abilitiesInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 12)

	inputs[strengthBaseInput] = textinput.New()
	inputs[strengthBaseInput].Placeholder = "10"
	inputs[strengthBaseInput].Focus()
	inputs[strengthBaseInput].Width = 20
	inputs[strengthBaseInput].Prompt = ""

	inputs[strengthProficientInput] = textinput.New()
	inputs[strengthProficientInput].Placeholder = "false"
	inputs[strengthProficientInput].Width = 20
	inputs[strengthProficientInput].Prompt = ""

	inputs[dexterityBaseInput] = textinput.New()
	inputs[dexterityBaseInput].Placeholder = "10"
	inputs[dexterityBaseInput].Width = 20
	inputs[dexterityBaseInput].Prompt = ""

	inputs[dexterityProficientInput] = textinput.New()
	inputs[dexterityProficientInput].Placeholder = "false"
	inputs[dexterityProficientInput].Width = 20
	inputs[dexterityProficientInput].Prompt = ""

	inputs[constitutionBaseInput] = textinput.New()
	inputs[constitutionBaseInput].Placeholder = "10"
	inputs[constitutionBaseInput].Width = 20
	inputs[constitutionBaseInput].Prompt = ""

	inputs[constitutionProficientInput] = textinput.New()
	inputs[constitutionProficientInput].Placeholder = "false"
	inputs[constitutionProficientInput].Width = 20
	inputs[constitutionProficientInput].Prompt = ""

	inputs[intelligenceBaseInput] = textinput.New()
	inputs[intelligenceBaseInput].Placeholder = "10"
	inputs[intelligenceBaseInput].Width = 20
	inputs[intelligenceBaseInput].Prompt = ""

	inputs[intelligenceProficientInput] = textinput.New()
	inputs[intelligenceProficientInput].Placeholder = "false"
	inputs[intelligenceProficientInput].Width = 20
	inputs[intelligenceProficientInput].Prompt = ""

	inputs[wisdomBaseInput] = textinput.New()
	inputs[wisdomBaseInput].Placeholder = "10"
	inputs[wisdomBaseInput].Width = 20
	inputs[wisdomBaseInput].Prompt = ""

	inputs[wisdomProficientInput] = textinput.New()
	inputs[wisdomProficientInput].Placeholder = "false"
	inputs[wisdomProficientInput].Width = 20
	inputs[wisdomProficientInput].Prompt = ""

	inputs[charismaBaseInput] = textinput.New()
	inputs[charismaBaseInput].Placeholder = "10"
	inputs[charismaBaseInput].Width = 20
	inputs[charismaBaseInput].Prompt = ""

	inputs[charismaProficientInput] = textinput.New()
	inputs[charismaProficientInput].Placeholder = "false"
	inputs[charismaProficientInput].Width = 20
	inputs[charismaProficientInput].Prompt = ""

	return inputs
}
