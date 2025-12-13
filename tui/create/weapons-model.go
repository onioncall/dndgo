package create

import "github.com/charmbracelet/bubbles/textinput"

const (
	weaponNameInput = iota
	damageInput
	proficientWeaponInput
	rangedInput
	normalRangeInput
	longRangeInput
	typeInput
	propertiesInput
)

func weaponInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 8)

	inputs[weaponNameInput] = textinput.New()
	inputs[weaponNameInput].Placeholder = "Shortsword"
	inputs[weaponNameInput].Focus()
	inputs[weaponNameInput].Width = 40
	inputs[weaponNameInput].Prompt = ""
	inputs[weaponNameInput].TextStyle = tertiaryStyle
	inputs[weaponNameInput].Cursor.Style = tertiaryStyle

	inputs[damageInput] = textinput.New()
	inputs[damageInput].Placeholder = "1d6"
	inputs[damageInput].Width = 40
	inputs[damageInput].Prompt = ""
	inputs[damageInput].TextStyle = tertiaryStyle
	inputs[damageInput].Cursor.Style = tertiaryStyle

	inputs[proficientWeaponInput] = textinput.New()
	inputs[proficientWeaponInput].Placeholder = "false"
	inputs[proficientWeaponInput].Width = 40
	inputs[proficientWeaponInput].Prompt = ""
	inputs[proficientWeaponInput].TextStyle = tertiaryStyle
	inputs[proficientWeaponInput].Cursor.Style = tertiaryStyle

	inputs[rangedInput] = textinput.New()
	inputs[rangedInput].Placeholder = "false"
	inputs[rangedInput].Width = 40
	inputs[rangedInput].Prompt = ""
	inputs[rangedInput].TextStyle = tertiaryStyle
	inputs[rangedInput].Cursor.Style = tertiaryStyle

	inputs[normalRangeInput] = textinput.New()
	inputs[normalRangeInput].Placeholder = "0"
	inputs[normalRangeInput].Width = 40
	inputs[normalRangeInput].Prompt = ""
	inputs[normalRangeInput].TextStyle = tertiaryStyle
	inputs[normalRangeInput].Cursor.Style = tertiaryStyle

	inputs[longRangeInput] = textinput.New()
	inputs[longRangeInput].Placeholder = "0"
	inputs[longRangeInput].Width = 40
	inputs[longRangeInput].Prompt = ""
	inputs[longRangeInput].TextStyle = tertiaryStyle
	inputs[longRangeInput].Cursor.Style = tertiaryStyle

	inputs[typeInput] = textinput.New()
	inputs[typeInput].Placeholder = "Piercing"
	inputs[typeInput].Width = 40
	inputs[typeInput].Prompt = ""
	inputs[typeInput].TextStyle = tertiaryStyle
	inputs[typeInput].Cursor.Style = tertiaryStyle

	inputs[propertiesInput] = textinput.New()
	inputs[propertiesInput].Placeholder = "Finesse, Light"
	inputs[propertiesInput].Width = 40
	inputs[propertiesInput].Prompt = ""
	inputs[propertiesInput].TextStyle = tertiaryStyle
	inputs[propertiesInput].Cursor.Style = tertiaryStyle

	return inputs
}
