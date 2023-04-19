package database

import (
	"context"
	"fmt"

	"cloud.google.com/go/datastore"
	"github.com/smcclure17/writr/pkg/models"
)

// TODO: Use enums or env vars.
const messagesTable = "documents"

// Database is the main database instance.
type Database struct {
	Client *datastore.Client
}

// NewDatabase creates a new database instance.
func NewDatabase() *Database {
	ctx := context.Background()
	client, err := datastore.NewClient(ctx, "writr-dev-384017")
	if err != nil {
		panic(err)
	}

	return &Database{
		Client: client,
	}
}

// InsertData inserts message data into the database. For now, only the messages table is supported.
func (d *Database) InsertData(tableName string, data models.Message) (bool, error) {
	ctx := context.Background()

	if tableName != messagesTable {
		return false, fmt.Errorf("table %s not supported", tableName)
	}

	_, err := d.Client.Put(ctx, datastore.NameKey(messagesTable, "doc-1", nil), &data)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (d *Database) GetData(tableName string, documentName string) (models.Message, error) {
	ctx := context.Background()

	if tableName != messagesTable {
		return models.Message{}, fmt.Errorf("table %s not supported", tableName)
	}

	var message models.Message
	key := datastore.NameKey(messagesTable, documentName, nil)

	err := d.Client.Get(ctx, key, &message)
	if err != nil {
		return models.Message{}, err
	}

	return message, nil
}
