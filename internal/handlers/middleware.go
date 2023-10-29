package handlers

import (
	"fistbump/internal/datastore"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

type UserMiddleware struct {
	repo  datastore.Repository
	store *session.Store
}

func NewUserMiddleware(repo datastore.Repository, store *session.Store) *UserMiddleware {
	return &UserMiddleware{repo: repo, store: store}
}

func (um *UserMiddleware) Middleware(c *fiber.Ctx) error {
	// Get session store
	session, _ := um.store.Get(c)

	// Try to get user ID from session
	userID, exists := session.Get("userID").(int)
	if !exists {
		// If user not in session, just proceed without doing anything
		return c.Next()
	}

	// Look up the user in your database or whatever storage mechanism you're using
	user, err := um.repo.GetUser(c.Context(), userID)
	if err != nil {
		// Handle error (e.g., log it, send a response, etc.)
		return c.Status(500).SendString("Failed to retrieve user")
	}

	// Attach the user to the Fiber context
	c.Locals("user", datastore.User{
		Email:    user.Email,
		Id:       user.ID,
		Name:     user.Name,
		Username: user.Name,
	})

	return c.Next()
}

func EnsureAuthenticated(c *fiber.Ctx) error {
	// You can adjust this according to where and how you store the user data in the context.
	user := c.Locals("user")
	if user == nil {
		return c.Redirect("/login") // Redirect to the login page if not authenticated.
	}
	return c.Next() // Continue processing if authenticated.
}

type MockUserMiddleware struct {
	repo  datastore.Repository
	store *session.Store
}

func NewMockUserMiddleware(repo datastore.Repository, store *session.Store) *MockUserMiddleware {
	return &MockUserMiddleware{repo: repo, store: store}
}

func (mum *MockUserMiddleware) Middleware(c *fiber.Ctx) error {
	// Get sess store
	sess, _ := mum.store.Get(c)

	// Try to get user ID from session
	_, exists := sess.Get("userID").(int)
	if !exists {

		u, err := mum.repo.UpsertUser(c.Context(), datastore.UserInput{
			Email:    "mock@email.com",
			Name:     "Barrington Florence",
			Username: "mock@email.com",
		})

		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		sess.Set("userID", u.ID)
		sess.Save()
	}
	return c.Next()
}
