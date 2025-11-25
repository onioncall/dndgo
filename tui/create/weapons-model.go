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

	inputs[damageInput] = textinput.New()
	inputs[damageInput].Placeholder = "1d6"
	inputs[damageInput].Width = 40
	inputs[damageInput].Prompt = ""

	inputs[proficientWeaponInput] = textinput.New()
	inputs[proficientWeaponInput].Placeholder = "false"
	inputs[proficientWeaponInput].Width = 40
	inputs[proficientWeaponInput].Prompt = ""

	inputs[rangedInput] = textinput.New()
	inputs[rangedInput].Placeholder = "false"
	inputs[rangedInput].Width = 40
	inputs[rangedInput].Prompt = ""

	inputs[normalRangeInput] = textinput.New()
	inputs[normalRangeInput].Placeholder = "0"
	inputs[normalRangeInput].Width = 40
	inputs[normalRangeInput].Prompt = ""

	inputs[longRangeInput] = textinput.New()
	inputs[longRangeInput].Placeholder = "0"
	inputs[longRangeInput].Width = 40
	inputs[longRangeInput].Prompt = ""

	inputs[typeInput] = textinput.New()
	inputs[typeInput].Placeholder = "Piercing"
	inputs[typeInput].Width = 40
	inputs[typeInput].Prompt = ""

	inputs[propertiesInput] = textinput.New()
	inputs[propertiesInput].Placeholder = "Finesse, Light"
	inputs[propertiesInput].Width = 40
	inputs[propertiesInput].Prompt = ""

	return inputs
}
