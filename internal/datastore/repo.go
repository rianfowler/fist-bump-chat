package datastore

import (
	"bytes"
	"context"
	"fmt"
	"html/template"
	"os"

	"github.com/rianfowler/fist-bump-chat/ent"
)

type UserRepository interface {
	GetUser(ctx context.Context, GithubID int) (*ent.User, error)
	UpsertUser(ctx context.Context, userInput UserInput) (*ent.User, error)
}

type MessageRepository interface {
	CreateMessage(ctx context.Context, message Message) error
	ListMessages(ctx context.Context) ([]Message, error)
}

type EntRepository struct {
	client *ent.Client
}

func NewEntRepository(client *ent.Client) *EntRepository {
	return &EntRepository{client: client}
}

func MakePostgresUrl() (string, error) {
	tmplStr := "postgres://{{.USERNAME}}:{{.PASSWORD}}@{{.HOST}}:{{.PORT}}/{{.DBNAME}}"

	// Parse the template
	tmpl, err := template.New("envTemplate").Parse(tmplStr)
	if err != nil {
		return "", fmt.Errorf("unable to create postgres connection template")
	}

	// Create a map or struct with the values from environment variables
	data := map[string]string{
		"USERNAME": os.Getenv("POSTGRES_USERNAME"),
		"PASSWORD": os.Getenv("POSTGRES_PASSWORD"),
		"HOST":     os.Getenv("POSTGRES_HOST"),
		"PORT":     os.Getenv("POSTGRES_PORT"),
		"DBNAME":   os.Getenv("POSTGRES_DBNAME"),
	}

	for key, val := range data {
		if val == "" {
			return "", fmt.Errorf("unable to get env var %s", key)
		}
	}

	// Execute the template
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("unable to template postgres connection string")
	}

	// Print the populated template
	return buf.String(), nil
}
