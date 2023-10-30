package datastore

import (
	"context"
	"fistbump/ent"
	"fistbump/ent/user"
	"fmt"
	"log"

	"github.com/rianfowler/fist-bump-chat/ent"
	"github.com/rianfowler/fist-bump-chat/ent/user"

	"entgo.io/ent/dialect/sql"
)

type UserInput struct {
	Email    string
	Name     string
	Username string
}

type User struct {
	Email    string
	Id       int
	Name     string
	Username string
}

func (er *EntRepository) CreateUser(ctx context.Context, userInput UserInput) (*ent.User, error) {
	u, err := er.client.User.
		Create().
		SetName(userInput.Name).
		SetEmail(userInput.Email).
		Save(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}

	return u, nil
}

func (er *EntRepository) GetUser(ctx context.Context, ID int) (*ent.User, error) {
	u, err := er.client.User.Get(ctx, ID)

	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}

	return u, nil
}

func (er *EntRepository) UpsertUser(ctx context.Context, userInput UserInput) (*ent.User, error) {
	id, err := er.client.User.Create().
		SetName(userInput.Name).
		SetEmail(userInput.Email).
		OnConflict(sql.ConflictColumns(user.FieldEmail)).
		UpdateNewValues().
		ID(ctx)

	if err != nil {
		// Handle error. For example, log it and return.
		log.Printf("Error during user upsert: %v", err)
		return nil, err
	}

	u, err := er.client.User.Get(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("failed upserting user: %v", err)
	}

	return u, nil
}
