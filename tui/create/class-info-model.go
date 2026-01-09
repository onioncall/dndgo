package create

import "github.com/charmbracelet/bubbles/textinput"

const (
	classTypeInput = iota
	classLevelInput
)

func classInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 2)

	inputs[classTypeInput] = textinput.New()
	inputs[classTypeInput].Placeholder = "Barbarian"
	inputs[classTypeInput].Focus()
	inputs[classTypeInput].Width = 40
	inputs[classTypeInput].Prompt = ""
	inputs[classTypeInput].TextStyle = tertiaryStyle
	inputs[classTypeInput].Cursor.Style = tertiaryStyle

	inputs[classLevelInput] = textinput.New()
	inputs[classLevelInput].Placeholder = "0"
	inputs[classLevelInput].Width = 40
	inputs[classLevelInput].Prompt = ""
	inputs[classLevelInput].TextStyle = tertiaryStyle
	inputs[classLevelInput].Cursor.Style = tertiaryStyle

	return inputs
}
