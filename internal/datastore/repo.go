package datastore

import (
	"bytes"
	"context"
	"fistbump/ent"
	"fmt"
	"html/template"
	"os"
)

type Repository interface {
	CreateMessage(ctx context.Context, message Message) (*ent.Message, error)
	GetUser(ctx context.Context, GithubID int) (*ent.User, error)
	ListMessages(ctx context.Context) ([]Message, error)
	UpsertUser(ctx context.Context, userInput UserInput) (*ent.User, error)
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
