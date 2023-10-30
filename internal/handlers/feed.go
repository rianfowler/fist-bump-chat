package handlers

import (
	"log"

	"github.com/rianfowler/fist-bump-chat/internal/datastore"

	"github.com/gofiber/fiber/v2"
)

type FeedHandler struct {
	repo datastore.MessageRepository
}

func NewFeedHandler(repo datastore.MessageRepository) *FeedHandler {
	return &FeedHandler{repo: repo}
}

func (fh *FeedHandler) GetFeedPage(c *fiber.Ctx) error {
	msgs, err := fh.repo.ListMessages(c.Context())

	if err != nil {
		log.Fatalf("Could not retrieve messages")
	}
	return c.Render("web/templates/pages/feed", fiber.Map{
		"messages": msgs,
	}, "web/templates/layouts/main")
}

func (fh *FeedHandler) GetNewMessagePage(c *fiber.Ctx) error {
	return c.Render("web/templates/pages/new-message", nil, "web/templates/layouts/main")
}

func (fh *FeedHandler) PostNewMessage(c *fiber.Ctx) error {
	message := c.FormValue("message")

	user, ok := c.Locals("user").(datastore.User)

	if !ok {
		log.Fatalf("postNewMessage could not retrieve user from context")
		return c.Status(fiber.StatusInternalServerError).Render("web/templates/pages/404", nil, "web/templates/layouts/main")
	}

	if len(message) > 0 {
		err := fh.repo.CreateMessage(c.Context(), datastore.Message{
			Message:  message,
			UserId:   user.Id,
			Username: user.Name,
		})
		if err != nil {
			log.Fatalf("postNewMessage could not add message to storage")
			return c.Status(fiber.StatusInternalServerError).Render("web/templates/pages/404", nil, "web/templates/layouts/main")
		}
	}

	return c.Redirect("/feed", fiber.StatusMovedPermanently)
}
