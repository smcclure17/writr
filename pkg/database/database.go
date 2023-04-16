package database

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/smcclure17/writr/pkg/models"
	"golang.org/x/exp/slices"
)

// TODO: Use enums or env vars.
const messagesTable = "Messages"

// Database is the main database instance.
type Database struct {
	Client *dynamodb.Client
}

// NewDatabase creates a new database instance.
func NewDatabase() *Database {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"))
	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
	}
	return &Database{
		Client: dynamodb.NewFromConfig(cfg),
	}
}

// InsertData inserts message data into the database. For now, only the messages table is supported.
func (d *Database) InsertData(tableName string, data models.Message) {
	if tableName != messagesTable {
		panic("Table name must be 'messages' for now.")
	}

	tableNames := d.getTableNames()
	if !slices.Contains(tableNames, tableName) {
		panic("Table doesn't exist. Make sure AWS is configured correctly and terraform has been run.")
	}

	out, err := d.Client.PutItem(context.TODO(), &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item: map[string]types.AttributeValue{
			"Name":    &types.AttributeValueMemberS{Value: data.Document}, // TODO: update
			"Content": &types.AttributeValueMemberS{Value: data.Message},
		},
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(out)
}

// GetTableNames returns a list of table names in the database.
func (d *Database) getTableNames() []string {
	result, err := d.Client.ListTables(context.TODO(), &dynamodb.ListTablesInput{})
	if err != nil {
		fmt.Printf("Failed to list tables, %v", err)
	}
	return result.TableNames
}
