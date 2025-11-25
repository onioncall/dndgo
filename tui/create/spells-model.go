package create

import "github.com/charmbracelet/bubbles/textinput"

const (
	spellNameInput = iota
	isRitualInput
	slotLevelInput
)

func spellInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 3)

	inputs[spellNameInput] = textinput.New()
	inputs[spellNameInput].Placeholder = "Silvery Barbs"
	inputs[spellNameInput].Focus()
	inputs[spellNameInput].Width = 40
	inputs[spellNameInput].Prompt = ""

	inputs[isRitualInput] = textinput.New()
	inputs[isRitualInput].Placeholder = "false"
	inputs[isRitualInput].Width = 40
	inputs[isRitualInput].Prompt = ""

	inputs[slotLevelInput] = textinput.New()
	inputs[slotLevelInput].Placeholder = "1"
	inputs[slotLevelInput].Width = 40
	inputs[slotLevelInput].Prompt = ""

	return inputs
}
