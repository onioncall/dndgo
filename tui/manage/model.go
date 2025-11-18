package manage

type Model struct {
	width    int
	height   int
	pageText string
}

func NewModel() Model {
	return Model{
		pageText: "Manage Character is currently under construction",
	}
}
