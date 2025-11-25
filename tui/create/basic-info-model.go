package create

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type (
	errMsg error
)

const (
	nameInput = iota
	levelInput
	classInput
	raceInput
	backgroundInput
	languagesInput
	hpInput
	speedInput
)

func basicInfoInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 8)

	inputs[nameInput] = textinput.New()
	inputs[nameInput].Placeholder = "Nim the Bold"
	inputs[nameInput].Focus()
	inputs[nameInput].Width = 40
	inputs[nameInput].Prompt = ""

	inputs[levelInput] = textinput.New()
	inputs[levelInput].Placeholder = "1"
	inputs[levelInput].Width = 40
	inputs[levelInput].Prompt = ""

	inputs[classInput] = textinput.New()
	inputs[classInput].Placeholder = "Bard"
	inputs[classInput].Width = 40
	inputs[classInput].Prompt = ""

	inputs[raceInput] = textinput.New()
	inputs[raceInput].Placeholder = "Dragonborn"
	inputs[raceInput].Width = 40
	inputs[raceInput].Prompt = ""

	inputs[backgroundInput] = textinput.New()
	inputs[backgroundInput].Placeholder = "Optional"
	inputs[backgroundInput].Width = 40
	inputs[backgroundInput].Prompt = ""

	inputs[languagesInput] = textinput.New()
	inputs[languagesInput].Placeholder = "Common, Elvish"
	inputs[languagesInput].Width = 40
	inputs[languagesInput].Prompt = ""

	inputs[hpInput] = textinput.New()
	inputs[hpInput].Placeholder = "10"
	inputs[hpInput].Width = 40
	inputs[hpInput].Prompt = ""

	inputs[speedInput] = textinput.New()
	inputs[speedInput].Placeholder = "30"
	inputs[speedInput].Width = 40
	inputs[speedInput].Prompt = ""

	return inputs
}
