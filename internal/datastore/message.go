package datastore

import (
	"context"
	"fmt"
	"log"

	"github.com/rianfowler/fist-bump-chat/ent"
)

type Message struct {
	Username string
	Message  string
	UserId   int
}

func (er *EntRepository) CreateMessage(ctx context.Context, message Message) error {
	_, err := er.client.Message.
		Create().
		SetMessage(message.Message).
		SetUserID(message.UserId).
		Save(ctx)

	if err != nil {
		return fmt.Errorf("failed creating message: %v", err)
	}

	return nil
}

func (er *EntRepository) ListMessages(ctx context.Context) ([]Message, error) {
	// for tomorrow:
	// look into querying this so that it also grabs the edges for the user
	entMsgs, err := er.client.Message.Query().WithUser().All(ctx)

	if err != nil {
		return nil, fmt.Errorf("failed querying messages: %v", err)
	}

	var msgs []Message
	for _, entMsg := range entMsgs {
		msgs = append(msgs, FromEntMessage(entMsg))
	}

	return msgs, nil
}

func FromEntMessage(e *ent.Message) Message {
	log.Printf("FromEntMessage: %v", e)
	return Message{
		Username: e.Edges.User.Name,
		Message:  e.Message,
	}
}
