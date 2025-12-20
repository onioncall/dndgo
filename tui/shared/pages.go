package shared

type NavigateMsg struct {
	Page string
}

type ReloadCharacterMsg struct{}

const (
	TuiHeader  = "dndgo"
	MenuPage   = "menu"
	InitPage   = "init"
	ManagePage = "manage"
	SearchPage = "search"
)
