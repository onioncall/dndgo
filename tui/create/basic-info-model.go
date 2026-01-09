package create

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type (
	errMsg error
)

const (
	nameInput = iota
	raceInput
	backgroundInput
	languagesInput
	hpInput
	speedInput
)

func basicInfoInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 6)

	inputs[nameInput] = textinput.New()
	inputs[nameInput].Placeholder = "Nim the Bold"
	inputs[nameInput].Focus()
	inputs[nameInput].Width = 40
	inputs[nameInput].Prompt = ""
	inputs[nameInput].TextStyle = tertiaryStyle
	inputs[nameInput].Cursor.Style = tertiaryStyle

	inputs[raceInput] = textinput.New()
	inputs[raceInput].Placeholder = "Dragonborn"
	inputs[raceInput].Width = 40
	inputs[raceInput].Prompt = ""
	inputs[raceInput].TextStyle = tertiaryStyle
	inputs[raceInput].Cursor.Style = tertiaryStyle

	inputs[backgroundInput] = textinput.New()
	inputs[backgroundInput].Placeholder = "Optional"
	inputs[backgroundInput].Width = 40
	inputs[backgroundInput].Prompt = ""
	inputs[backgroundInput].TextStyle = tertiaryStyle
	inputs[backgroundInput].Cursor.Style = tertiaryStyle

	inputs[languagesInput] = textinput.New()
	inputs[languagesInput].Placeholder = "Common, Elvish"
	inputs[languagesInput].Width = 40
	inputs[languagesInput].Prompt = ""
	inputs[languagesInput].TextStyle = tertiaryStyle
	inputs[languagesInput].Cursor.Style = tertiaryStyle

	inputs[hpInput] = textinput.New()
	inputs[hpInput].Placeholder = "10"
	inputs[hpInput].Width = 40
	inputs[hpInput].Prompt = ""
	inputs[hpInput].TextStyle = tertiaryStyle
	inputs[hpInput].Cursor.Style = tertiaryStyle

	inputs[speedInput] = textinput.New()
	inputs[speedInput].Placeholder = "30"
	inputs[speedInput].Width = 40
	inputs[speedInput].Prompt = ""
	inputs[speedInput].TextStyle = tertiaryStyle
	inputs[speedInput].Cursor.Style = tertiaryStyle

	return inputs
}
