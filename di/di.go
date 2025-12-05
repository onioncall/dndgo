package di

import (
	"github.com/onioncall/dndgo/character-management/db"
	"github.com/onioncall/dndgo/character-management/handlers"
)

var (
	CharacterHandler handlers.CharacterHandler
	Repo             db.Repo
)

func Init() error {
	Repo, err := db.NewRepo()
	if err != nil {
		return err
	}

	CharacterHandler = handlers.NewCharacterHandler(Repo)
	return nil
}

func Deinit() error {
	return Repo.Deinit()
}
