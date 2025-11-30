package db

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/onioncall/dndgo/character-management/models"
	c "github.com/ostafen/clover/v2"
	cdocument "github.com/ostafen/clover/v2/document"
	cquery "github.com/ostafen/clover/v2/query"
)

const character_collection = "characters"
const class_collection = "classes"
const db_dirname = "dndgo"

type Repo struct {
	db *c.DB
}

func NewRepo() (Repo, error) {
	xdgData := os.Getenv("XDG_DATA_HOME")
	if xdgData == "" {
		home := os.Getenv("HOME")
		xdgData = filepath.Join(home, ".local", "share")
	}

	dbDir := filepath.Join(xdgData, db_dirname)

	err := os.MkdirAll(dbDir, 0755)
	if err != nil {
		return Repo{}, fmt.Errorf("Failed to create data dir at %v:\n%w", dbDir, err)
	}

	db, err := c.Open(dbDir)
	if err != nil {
		return Repo{}, fmt.Errorf("Failed to open dndgo db: %w", err)
	}
	return Repo{
		db: db,
	}, nil
}

func (r Repo) Deinit() error {
	return r.db.Close()
}

// InsertCharacter creates a new character in the db
// Returns the inserted character ID
func (r Repo) InsertCharacter(character models.Character) (string, error) {
	if err := r.db.CreateCollection(character_collection); err != nil {
		return "", fmt.Errorf("Failed to create or locate characters collection: %w", err)
	}

	doc := cdocument.NewDocumentOf(character)
	docId, err := r.db.InsertOne(character_collection, doc)
	if err != nil {
		return "", fmt.Errorf("Failed to insert new character to db: %w", err)
	}

	return docId, nil
}

// GetCharacter gets the default character
func (r Repo) GetCharacter() (*models.Character, error) {
	doc, err := r.db.FindFirst(
		cquery.NewQuery(character_collection).Where(cquery.Field("default").Eq(true)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve character from db:\n%w", err)
	}

	res := models.Character{}
	if err = doc.Unmarshal(&res); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal db record into character struct:\n%w", err)
	}

	return &res, nil
}

// SyncCharacter will update the "Default" character with all
// properties in 'character'
func (r Repo) SyncCharacter(character models.Character) error {
	bytes, err := json.Marshal(character)
	if err != nil {
		return fmt.Errorf("Failed to marshal character:\n%w", err)
	}

	updates := make(map[string]interface{})
	err = json.Unmarshal(bytes, &updates)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal character to generic map:\n%w", err)
	}

	if err = r.db.Update(cquery.NewQuery(character_collection).Where(cquery.Field("default").Eq(true)), updates); err != nil {
		return fmt.Errorf("Failed to update character doc:\n%w", err)
	}

	return nil
}

// InsertClass creates a new class record in the db
func (r Repo) InsertClass(class models.Class) error {
	if err := r.db.CreateCollection(class_collection); err != nil {
		return fmt.Errorf("Failed to create or locate classes collection: %w", err)
	}

	doc := cdocument.NewDocumentOf(class)
	_, err := r.db.InsertOne(class_collection, doc)
	if err != nil {
		return fmt.Errorf("Failed to insert new class record to db: %w", err)
	}

	return nil
}

// GetClass retrieves a class from the db based on the ID of the corresponding character
// Returns a []byte json representation of the class
func (r Repo) GetClass(id string) ([]byte, error) {
	doc, err := r.db.FindFirst(
		cquery.NewQuery(class_collection).Where(cquery.Field("character_id").Eq(id)),
	)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve class from db: %w", err)
	}

	return cdocument.Encode(doc)
}

// SyncClass updates the class doc based on class.CharacterID
func (r Repo) SyncClass(class models.Class) error {
	bytes, err := json.Marshal(class)
	if err != nil {
		return fmt.Errorf("Failed to marshal class:\n%w", err)
	}

	updates := make(map[string]interface{})
	err = json.Unmarshal(bytes, &updates)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal class to generic map:\n%w", err)
	}

	if err = r.db.Update(cquery.NewQuery(class_collection).Where(cquery.Field("character_id").Eq(class.GetCharacterId())), updates); err != nil {
		return fmt.Errorf("Failed to update class doc:\n%w", err)
	}

	return nil
}
