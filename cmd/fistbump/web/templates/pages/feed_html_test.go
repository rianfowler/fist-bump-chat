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

	// Create a new view model with test data
	viewModel := viewmodels.FeedViewModel{
		Messages: []viewmodels.MessageViewModel{{
			Author:  "user1",
			Content: "Hello World",
		},
		},
	}

	// Execute the template with the view model
	var tpl bytes.Buffer
	if err := tmpl.Execute(&tpl, viewModel); err != nil {
		t.Fatalf("Executing template failed: %v", err)
	}

	// Convert the buffer to a string for easier searching
	result := tpl.String()

	// Check that the result contains expected data
	if !strings.Contains(result, viewModel.Messages[0].Author) {
		t.Errorf("Output does not contain the Author: %s", viewModel.Messages[0].Author)
	}

	if !strings.Contains(result, viewModel.Messages[0].Content) {
		t.Errorf("Output does not contain the content: %s", viewModel.Messages[0].Content)
	}

}
