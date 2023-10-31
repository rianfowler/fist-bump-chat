package templates_test

import (
	"bytes"
	"html/template"
	"strings"
	"testing"

	"github.com/rianfowler/fist-bump-chat/internal/viewmodels"
)

// Define a mock view model that your template expects
type PageViewModel struct {
	Title   string
	Content string
}

func TestTemplateExecution(t *testing.T) {
	// Initialize the template
	tmpl, err := template.ParseFiles("./feed.html", "../partials/message.html")

	if err != nil {
		t.Fatalf("Parsing template failed: %v", err)
	}

	tests := []struct {
		name      string
		viewModel viewmodels.FeedViewModel
		want      string
	}{
		{
			name: "single message",
			viewModel: viewmodels.FeedViewModel{
				Messages: []viewmodels.MessageViewModel{{
					Author: "user1", Content: "Hello World",
				}},
			},
			want: "Hello World",
		},
		{
			name: "multiple messages",
			viewModel: viewmodels.FeedViewModel{
				Messages: []viewmodels.MessageViewModel{
					{
						Author: "user1", Content: "Hello World",
					}, {
						Author: "user2", Content: "Hello World too",
					},
				},
			},
			want: "Hello World too",
		},
		{
			name: "no messages ",
			viewModel: viewmodels.FeedViewModel{
				Messages: []viewmodels.MessageViewModel{},
			},
			want: "No messages. Try adding one",
		},
	}

	for _, tt := range tests {
		// Execute the template with the view model
		var tpl bytes.Buffer
		if err := tmpl.Execute(&tpl, tt.viewModel); err != nil {
			t.Fatalf("Executing template failed: %v", err)
		}

		// Convert the buffer to a string for easier searching
		result := tpl.String()

		// Check that the result contains expected data
		if !strings.Contains(result, tt.want) {
			t.Errorf("Output does not contain %s", tt.want)
		}
	}
}
