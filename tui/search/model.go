package search

type Model struct {
	width    int
	height   int
	pageText string
}

func NewModel() Model {
	return Model{
		pageText: "Search is currently under construction",
	}
}
