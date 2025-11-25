package create

import (
	"github.com/charmbracelet/bubbles/textinput"
)

const (
	itemNameInput = iota
	itemQuantityInput
)

func backpackInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[itemNameInput] = textinput.New()
	inputs[itemNameInput].Placeholder = "Gold"
	inputs[itemNameInput].Focus()
	inputs[itemNameInput].Width = 40
	inputs[itemNameInput].Prompt = ""

	inputs[itemQuantityInput] = textinput.New()
	inputs[itemQuantityInput].Placeholder = "15"
	inputs[itemQuantityInput].Width = 40
	inputs[itemQuantityInput].Prompt = ""

	return inputs
}
