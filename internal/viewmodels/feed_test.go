package viewmodels_test

import (
	"testing"

	"github.com/rianfowler/fist-bump-chat/internal/datastore"
	"github.com/rianfowler/fist-bump-chat/internal/viewmodels"
)

func TestMessagesToFeedView(t *testing.T) {
	tests := []struct {
		name     string
		messages []datastore.Message
		want     viewmodels.FeedViewModel
		wantErr  bool
	}{
		{
			name:     "no messages",
			messages: []datastore.Message{},
			want:     viewmodels.FeedViewModel{Messages: []viewmodels.MessageViewModel{}},
			wantErr:  false,
		},
		{
			name: "empty message",
			messages: []datastore.Message{
				{Username: "user1", Message: ""},
			},
			want:    viewmodels.FeedViewModel{Messages: []viewmodels.MessageViewModel{}},
			wantErr: true,
		},
		{
			name: "empty username",
			messages: []datastore.Message{
				{Username: "", Message: "message"},
			},
			want:    viewmodels.FeedViewModel{Messages: []viewmodels.MessageViewModel{}},
			wantErr: true,
		},
		{
			name: "empty username and message ",
			messages: []datastore.Message{
				{Username: "", Message: ""},
			},
			want:    viewmodels.FeedViewModel{Messages: []viewmodels.MessageViewModel{}},
			wantErr: true,
		},
		{
			name: "single message",
			messages: []datastore.Message{
				{Username: "user1", Message: "Hello World"},
			},
			want: viewmodels.FeedViewModel{
				Messages: []viewmodels.MessageViewModel{
					{Author: "user1", Content: "Hello World"},
				},
			},
			wantErr: false,
		},
		{
			name: "multiple messages",
			messages: []datastore.Message{
				{Username: "user1", Message: "Hello World"},
				{Username: "user2", Message: "Goodbye World"},
			},
			want: viewmodels.FeedViewModel{
				Messages: []viewmodels.MessageViewModel{
					{Author: "user1", Content: "Hello World"},
					{Author: "user2", Content: "Goodbye World"},
				},
			},
			wantErr: false,
		},
		// Add more test cases as necessary.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := viewmodels.MessagesToFeedView(tt.messages)
			if (gotErr != nil) != tt.wantErr {
				t.Errorf("MessagesToFeedView().err = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if len(got.Messages) != len(tt.want.Messages) {
				t.Errorf("MessagesToFeedView() = %v, want %v", got, tt.want)
				return
			}
			for i, gotMessage := range got.Messages {
				if gotMessage != tt.want.Messages[i] {
					t.Errorf("MessagesToFeedView()[%d] = %v, want %v", i, gotMessage, tt.want.Messages[i])
				}
			}
		})
	}
}
