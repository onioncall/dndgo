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
	inputs[strengthBaseInput].TextStyle = tertiaryStyle
	inputs[strengthBaseInput].Cursor.Style = tertiaryStyle

	inputs[strengthProficientInput] = textinput.New()
	inputs[strengthProficientInput].Placeholder = "false"
	inputs[strengthProficientInput].Width = 20
	inputs[strengthProficientInput].Prompt = ""
	inputs[strengthProficientInput].TextStyle = tertiaryStyle
	inputs[strengthProficientInput].Cursor.Style = tertiaryStyle

	inputs[dexterityBaseInput] = textinput.New()
	inputs[dexterityBaseInput].Placeholder = "10"
	inputs[dexterityBaseInput].Width = 20
	inputs[dexterityBaseInput].Prompt = ""
	inputs[dexterityBaseInput].TextStyle = tertiaryStyle
	inputs[dexterityBaseInput].Cursor.Style = tertiaryStyle

	inputs[dexterityProficientInput] = textinput.New()
	inputs[dexterityProficientInput].Placeholder = "false"
	inputs[dexterityProficientInput].Width = 20
	inputs[dexterityProficientInput].Prompt = ""
	inputs[dexterityProficientInput].TextStyle = tertiaryStyle
	inputs[dexterityProficientInput].Cursor.Style = tertiaryStyle

	inputs[constitutionBaseInput] = textinput.New()
	inputs[constitutionBaseInput].Placeholder = "10"
	inputs[constitutionBaseInput].Width = 20
	inputs[constitutionBaseInput].Prompt = ""
	inputs[constitutionBaseInput].TextStyle = tertiaryStyle
	inputs[constitutionBaseInput].Cursor.Style = tertiaryStyle

	inputs[constitutionProficientInput] = textinput.New()
	inputs[constitutionProficientInput].Placeholder = "false"
	inputs[constitutionProficientInput].Width = 20
	inputs[constitutionProficientInput].Prompt = ""
	inputs[constitutionProficientInput].TextStyle = tertiaryStyle
	inputs[constitutionProficientInput].Cursor.Style = tertiaryStyle

	inputs[intelligenceBaseInput] = textinput.New()
	inputs[intelligenceBaseInput].Placeholder = "10"
	inputs[intelligenceBaseInput].Width = 20
	inputs[intelligenceBaseInput].Prompt = ""
	inputs[intelligenceBaseInput].TextStyle = tertiaryStyle
	inputs[intelligenceBaseInput].Cursor.Style = tertiaryStyle

	inputs[intelligenceProficientInput] = textinput.New()
	inputs[intelligenceProficientInput].Placeholder = "false"
	inputs[intelligenceProficientInput].Width = 20
	inputs[intelligenceProficientInput].Prompt = ""
	inputs[intelligenceProficientInput].TextStyle = tertiaryStyle
	inputs[intelligenceProficientInput].Cursor.Style = tertiaryStyle

	inputs[wisdomBaseInput] = textinput.New()
	inputs[wisdomBaseInput].Placeholder = "10"
	inputs[wisdomBaseInput].Width = 20
	inputs[wisdomBaseInput].Prompt = ""
	inputs[wisdomBaseInput].TextStyle = tertiaryStyle
	inputs[wisdomBaseInput].Cursor.Style = tertiaryStyle

	inputs[wisdomProficientInput] = textinput.New()
	inputs[wisdomProficientInput].Placeholder = "false"
	inputs[wisdomProficientInput].Width = 20
	inputs[wisdomProficientInput].Prompt = ""
	inputs[wisdomProficientInput].TextStyle = tertiaryStyle
	inputs[wisdomProficientInput].Cursor.Style = tertiaryStyle

	inputs[charismaBaseInput] = textinput.New()
	inputs[charismaBaseInput].Placeholder = "10"
	inputs[charismaBaseInput].Width = 20
	inputs[charismaBaseInput].Prompt = ""
	inputs[charismaBaseInput].TextStyle = tertiaryStyle
	inputs[charismaBaseInput].Cursor.Style = tertiaryStyle

	inputs[charismaProficientInput] = textinput.New()
	inputs[charismaProficientInput].Placeholder = "false"
	inputs[charismaProficientInput].Width = 20
	inputs[charismaProficientInput].Prompt = ""
	inputs[charismaProficientInput].TextStyle = tertiaryStyle
	inputs[charismaProficientInput].Cursor.Style = tertiaryStyle

	return inputs
}
