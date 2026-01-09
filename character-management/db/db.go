package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/onioncall/dndgo/character-management/models"
	c "github.com/ostafen/clover/v2"
	cdocument "github.com/ostafen/clover/v2/document"
	cquery "github.com/ostafen/clover/v2/query"
)

const (
	characterCollection = "characters"
	classCollection     = "classes"
	dbDirname           = "dndgo"
)

type Repository struct {
	db *c.DB
}

var (
	Repo *Repository
	once sync.Once
)

func Init() error {
	var initErr error
	once.Do(func() {
		r, err := newRepository()
		if err != nil {
			initErr = err
		}

		Repo = r
	})
	return initErr
}

func newRepository() (*Repository, error) {
	xdgData := os.Getenv("XDG_DATA_HOME")
	if xdgData == "" {
		home := os.Getenv("HOME")
		xdgData = filepath.Join(home, ".local", "share")
	}

	dbDir := filepath.Join(xdgData, dbDirname)

	err := os.MkdirAll(dbDir, 0o755)
	if err != nil {
		return nil, fmt.Errorf("Failed to create data dir at %v:\n%w", dbDir, err)
	}

	db, err := c.Open(dbDir)
	if err != nil {
		return nil, fmt.Errorf("Failed to open dndgo db: %w", err)
	}

	r := &Repository{
		db: db,
	}

	if err := r.createCollection(characterCollection); err != nil {
		return nil, fmt.Errorf("Failed to create or locate characters collection: %w", err)
	}
	if err := r.createCollection(classCollection); err != nil {
		return nil, fmt.Errorf("Failed to create or locate class collection: %w", err)
	}

	return r, nil
}

func (r Repository) Deinit() error {
	return r.db.Close()
}

// InsertCharacter creates a new character in the db
// Returns the inserted character ID
func (r Repository) InsertCharacter(character models.Character) (string, error) {
	doc := cdocument.NewDocumentOf(character)
	docId, err := r.db.InsertOne(characterCollection, doc)
	if err != nil {
		return "", fmt.Errorf("Failed to insert new character to db: %w", err)
	}

	return docId, nil
}

// GetCharacter gets the default character
func (r Repository) GetCharacter() (*models.Character, error) {
	doc, err := r.db.FindFirst(
		cquery.NewQuery(characterCollection).Where(cquery.Field("default").Eq(true)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve character from db:\n%w", err)
	}

	res := models.Character{}

	// No default character found. May be expected, caller should handle nil.
	if doc == nil {
		return nil, nil
	}

	if err = doc.Unmarshal(&res); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal db record into character struct:\n%w", err)
	}
	res.ID = doc.ObjectId()

	return &res, nil
}

// GetCharacterById gets the character with the specified ID
func (r Repository) GetCharacterById(id string) (*models.Character, error) {
	doc, err := r.db.FindFirst(
		cquery.NewQuery(characterCollection).Where(cquery.Field("_id").Eq(id)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve character from db:\n%w", err)
	}

	res := models.Character{}

	// No character found. May be expected, caller should handle nil.
	if doc == nil {
		return nil, nil
	}

	if err = doc.Unmarshal(&res); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal db record into character struct:\n%w", err)
	}
	res.ID = doc.ObjectId()

	return &res, nil
}

// GetCharacterByName gets the character with the provided name
func (r Repository) GetCharacterByName(name string) (*models.Character, error) {
	doc, err := r.db.FindFirst(
		cquery.NewQuery(characterCollection).Where(cquery.Field("name").Eq(name)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve character from db:\n%w", err)
	}

	if doc == nil {
		return nil, fmt.Errorf("Character with name '%v' not found", name)
	}

	res := models.Character{}
	if err = doc.Unmarshal(&res); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal db record into character struct:\n%w", err)
	}
	res.ID = doc.ObjectId()

	return &res, nil
}

// Returns a slice of default characters
// There should only ever be one default character. But in the event something gets out of whack, we'll want to return all defaults.
func (r Repository) GetDefaultCharacters() ([]models.Character, error) {
	docs, err := r.db.FindAll(
		cquery.NewQuery(characterCollection).Where(cquery.Field("default").Eq(true)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve character(s) from db:\n%w", err)
	}

	// No default characters found
	if docs == nil {
		return nil, nil
	}

	result := []models.Character{}
	for _, doc := range docs {
		character := models.Character{}
		if err = doc.Unmarshal(&character); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal db record into character struct:\n%w", err)
		}

		result = append(result, character)
	}

	return result, nil
}

// Get all character names
func (r Repository) GetCharacterNames() ([]string, error) {
	docs, err := r.db.FindAll(cquery.NewQuery(characterCollection))
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve characters from db:\n%w", err)
	}

	result := []string{}

	for _, doc := range docs {
		character := models.Character{}
		if err = doc.Unmarshal(&character); err != nil {
			return nil, fmt.Errorf("Failed to unmarshal db record into character struct:\n%w", err)
		}

		result = append(result, character.Name)
	}

	return result, nil
}

// Deletes character by Id
func (r Repository) DeleteCharacter(characterId string) error {
	err := r.db.DeleteById(characterCollection, characterId)
	if err != nil {
		return fmt.Errorf("Failed to delete character with Id '%s':\n%w", characterId, err)
	}

	return nil
}

// Deletes all classes tied to a specific characterId, and returns an error
func (r Repository) DeleteClassesByCharacterId(characterId string) error {
	err := r.db.Delete(
		cquery.NewQuery(classCollection).Where(cquery.Field("character-id").Eq(characterId)),
	)
	if err != nil {
		return fmt.Errorf("Failed to delete class with Id '%s':\n%w", characterId, err)
	}

	return nil
}

// SyncCharacter will update the provided Character with all
// properties in 'character'
// Update record is based on character.id
func (r Repository) SyncCharacter(character models.Character) error {
	bytes, err := json.Marshal(character)
	if err != nil {
		return fmt.Errorf("Failed to marshal character:\n%w", err)
	}

	updates := make(map[string]any)
	err = json.Unmarshal(bytes, &updates)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal character to generic map:\n%w", err)
	}

	if err = r.db.Update(cquery.NewQuery(characterCollection).Where(cquery.Field("_id").Eq(character.ID)), updates); err != nil {
		return fmt.Errorf("Failed to update character doc:\n%w", err)
	}

	return nil
}

// InsertClass creates a new class record in the db
func (r Repository) InsertClass(class models.Class) error {
	doc := cdocument.NewDocumentOf(class)
	_, err := r.db.InsertOne(classCollection, doc)
	if err != nil {
		return fmt.Errorf("Failed to insert new class record to db: %w", err)
	}

	return nil
}

// GetClass retrieves a class from the db based on the ID of the corresponding character, and the class type
// The result will be unmarshaled into `obj`
func (r Repository) GetClass(chid string, obj models.Class, classType string) error {
	doc, err := r.db.FindFirst(
		cquery.NewQuery(classCollection).
			Where(cquery.Field("character-id").Eq(chid).
				And(cquery.Field("class-type").Eq(strings.ToLower(classType)))),
	)
	if err != nil {
		return fmt.Errorf("Failed to retrieve class from db: %w", err)
	}

	if doc == nil {
		return nil
	}

	if err = doc.Unmarshal(obj); err != nil {
		return err
	}

	return nil
}

// Retrieves all classes from the db based on the ID of the corresponding character
// The result will be unmarshaled into `objs`
func (r Repository) GetClassesByCharacterId(chid string, objs []models.Class) error {
	docs, err := r.db.FindAll(
		cquery.NewQuery(classCollection).Where(cquery.Field("character-id").Eq(chid)),
	)
	if err != nil {
		return fmt.Errorf("Failed to retrieve class from db: %w", err)
	}

	if docs == nil {
		return nil
	}

	for i, doc := range docs {
		if err = doc.Unmarshal(objs[i]); err != nil {
			return err
		}
	}

	return nil
}

// SyncClass updates the class doc based on class.CharacterID
func (r Repository) SyncClass(class models.Class) error {
	bytes, err := json.Marshal(class)
	if err != nil {
		return fmt.Errorf("Failed to marshal class:\n%w", err)
	}

	updates := make(map[string]any)
	err = json.Unmarshal(bytes, &updates)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal class to generic map:\n%w", err)
	}

	if err = r.db.Update(cquery.NewQuery(classCollection).
		Where(cquery.Field("character-id").Eq(class.GetCharacterId()).
			And(cquery.Field("class-type").Eq(strings.ToLower(class.GetClassType())))), updates); err != nil {
		return fmt.Errorf("failed to update class doc:\n%w", err)
	}

	return nil
}

// // SyncClasses updates all class docs based on class.CharacterID
// func (r Repository) SyncClasses(classes []models.Class) error {
// 	for _, class := range classes {
// 		bytes, err := json.Marshal(class)
// 		if err != nil {
// 			return fmt.Errorf("Failed to marshal class:\n%w", err)
// 		}
//
// 		updates := make(map[string]any)
// 		err = json.Unmarshal(bytes, &updates)
// 		if err != nil {
// 			return fmt.Errorf("Failed to unmarshal class to generic map:\n%w", err)
// 		}
//
// 		if err = r.db.Update(cquery.NewQuery(classCollection).
// 		Where(cquery.Field("character-id").Eq(class.GetCharacterId()).
// 		And(cquery.Field("class-type").Eq(class.GetClassType()))), updates); err != nil {
// 			return fmt.Errorf("Failed to update class doc:\n%w", err)
// 		}
// 	}
//
// 	return nil
// }

func (r Repository) createCollection(collection string) error {
	exist, err := r.db.HasCollection(collection)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	return r.db.CreateCollection(collection)
}
