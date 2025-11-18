package create

type Model struct {
	width    int
	height   int
	pageText string
}

func NewModel() Model {
	return Model{
		pageText: "Create Character is currently under construction",
	}
}
