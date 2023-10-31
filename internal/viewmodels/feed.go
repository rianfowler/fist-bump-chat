package viewmodels

import (
	"fmt"

	"github.com/rianfowler/fist-bump-chat/internal/datastore"
)

type FeedViewModel struct {
	Messages []MessageViewModel
}

type MessageViewModel struct {
	Author  string
	Content string
}

func MessagesToFeedView(messages []datastore.Message) (FeedViewModel, error) {
	var fvm FeedViewModel
	for _, message := range messages {
		// Example validation: ensure that Username and Message are not empty
		if message.Username == "" || message.Message == "" {
			return FeedViewModel{}, fmt.Errorf("empty message or username")
		}
		fvm.Messages = append(fvm.Messages, MessageViewModel{
			Author:  message.Username,
			Content: message.Message,
		})
	}

	return fvm, nil
}
