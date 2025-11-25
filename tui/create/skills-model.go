package create

import (
	"github.com/charmbracelet/bubbles/textinput"
)

type skillInfo struct {
	name    string
	ability string
}

const (
	// Strength
	athleticsInput = iota

	// Dexterity
	acrobaticsInput
	sleightOfHandInput
	stealthInput

	// Constitution has no skills (womp womp)

	// Intelligence
	arcanaInput
	historyInput
	investigationInput
	natureInput
	religionInput

	// Wisdom
	animalHandlingInput
	insightInput
	medicineInput
	perceptionInput
	survivalInput

	// Charisma
	deceptionInput
	intimidationInput
	performanceInput
	persuasionInput
)

var skillToAbility = map[int]skillInfo{
	athleticsInput:      {name: "Athletics", ability: "Strength"},
	acrobaticsInput:     {name: "Acrobatics", ability: "Dexterity"},
	sleightOfHandInput:  {name: "Sleight of Hand", ability: "Dexterity"},
	stealthInput:        {name: "Stealth", ability: "Dexterity"},
	arcanaInput:         {name: "Arcana", ability: "Intelligence"},
	historyInput:        {name: "History", ability: "Intelligence"},
	investigationInput:  {name: "Investigation", ability: "Intelligence"},
	natureInput:         {name: "Nature", ability: "Intelligence"},
	religionInput:       {name: "Religion", ability: "Intelligence"},
	animalHandlingInput: {name: "Animal Handling", ability: "Wisdom"},
	insightInput:        {name: "Insight", ability: "Wisdom"},
	medicineInput:       {name: "Medicine", ability: "Wisdom"},
	perceptionInput:     {name: "Perception", ability: "Wisdom"},
	survivalInput:       {name: "Survival", ability: "Wisdom"},
	deceptionInput:      {name: "Deception", ability: "Charisma"},
	intimidationInput:   {name: "Intimidation", ability: "Charisma"},
	performanceInput:    {name: "Performance", ability: "Charisma"},
	persuasionInput:     {name: "Persuasion", ability: "Charisma"},
}

func skillsInputs() []textinput.Model {
	var inputs []textinput.Model = make([]textinput.Model, 18)

	inputs[acrobaticsInput] = textinput.New()
	inputs[acrobaticsInput].Placeholder = "false"
	inputs[acrobaticsInput].Width = 6
	inputs[acrobaticsInput].Prompt = ""

	inputs[animalHandlingInput] = textinput.New()
	inputs[animalHandlingInput].Placeholder = "false"
	inputs[animalHandlingInput].Width = 6
	inputs[animalHandlingInput].Prompt = ""

	inputs[arcanaInput] = textinput.New()
	inputs[arcanaInput].Placeholder = "false"
	inputs[arcanaInput].Width = 6
	inputs[arcanaInput].Prompt = ""

	inputs[athleticsInput] = textinput.New()
	inputs[athleticsInput].Placeholder = "false"
	inputs[athleticsInput].Focus()
	inputs[athleticsInput].Width = 6
	inputs[athleticsInput].Prompt = ""

	inputs[deceptionInput] = textinput.New()
	inputs[deceptionInput].Placeholder = "false"
	inputs[deceptionInput].Width = 6
	inputs[deceptionInput].Prompt = ""

	inputs[historyInput] = textinput.New()
	inputs[historyInput].Placeholder = "false"
	inputs[historyInput].Width = 6
	inputs[historyInput].Prompt = ""

	inputs[insightInput] = textinput.New()
	inputs[insightInput].Placeholder = "false"
	inputs[insightInput].Width = 6
	inputs[insightInput].Prompt = ""

	inputs[intimidationInput] = textinput.New()
	inputs[intimidationInput].Placeholder = "false"
	inputs[intimidationInput].Width = 6
	inputs[intimidationInput].Prompt = ""

	inputs[investigationInput] = textinput.New()
	inputs[investigationInput].Placeholder = "false"
	inputs[investigationInput].Width = 6
	inputs[investigationInput].Prompt = ""

	inputs[medicineInput] = textinput.New()
	inputs[medicineInput].Placeholder = "false"
	inputs[medicineInput].Width = 6
	inputs[medicineInput].Prompt = ""

	inputs[natureInput] = textinput.New()
	inputs[natureInput].Placeholder = "false"
	inputs[natureInput].Width = 6
	inputs[natureInput].Prompt = ""

	inputs[perceptionInput] = textinput.New()
	inputs[perceptionInput].Placeholder = "false"
	inputs[perceptionInput].Width = 6
	inputs[perceptionInput].Prompt = ""

	inputs[performanceInput] = textinput.New()
	inputs[performanceInput].Placeholder = "false"
	inputs[performanceInput].Width = 6
	inputs[performanceInput].Prompt = ""

	inputs[persuasionInput] = textinput.New()
	inputs[persuasionInput].Placeholder = "false"
	inputs[persuasionInput].Width = 6
	inputs[persuasionInput].Prompt = ""

	inputs[religionInput] = textinput.New()
	inputs[religionInput].Placeholder = "false"
	inputs[religionInput].Width = 6
	inputs[religionInput].Prompt = ""

	inputs[sleightOfHandInput] = textinput.New()
	inputs[sleightOfHandInput].Placeholder = "false"
	inputs[sleightOfHandInput].Width = 6
	inputs[sleightOfHandInput].Prompt = ""

	inputs[stealthInput] = textinput.New()
	inputs[stealthInput].Placeholder = "false"
	inputs[stealthInput].Width = 6
	inputs[stealthInput].Prompt = ""

	inputs[survivalInput] = textinput.New()
	inputs[survivalInput].Placeholder = "false"
	inputs[survivalInput].Width = 6
	inputs[survivalInput].Prompt = ""

	return inputs
}
