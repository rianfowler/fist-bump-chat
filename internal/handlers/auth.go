package handlers

import (
	"encoding/json"
	"fistbump/internal/config"
	"fistbump/internal/datastore"
	"io"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHandler struct {
	repo  datastore.Repository
	store *session.Store
	cfg   *config.Configuration
}

func NewAuthHandler(repo datastore.Repository, store *session.Store, cfg *config.Configuration) *AuthHandler {
	return &AuthHandler{repo: repo, store: store, cfg: cfg}
}

type GoogleUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

var googleOAuth2Config = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  "",
	Scopes:       []string{"profile", "email"},
	Endpoint:     google.Endpoint,
}

func (ah *AuthHandler) GoogleLoginStart(c *fiber.Ctx) error {
	log.Println("Start Google Auth called")
	googleOAuth2Config.ClientID = ah.cfg.GetString("GOOGLE_CLIENT_ID")
	googleOAuth2Config.ClientSecret = ah.cfg.GetString("GOOGLE_CLIENT_SECRET")
	googleOAuth2Config.RedirectURL = ah.cfg.GetString("GOOGLE_REDIRECT_URL")
	// Generate OAuth2 URL and redirect to it
	url := googleOAuth2Config.AuthCodeURL("random-state")
	log.Printf("redirecting to: %s", url)
	return c.Redirect(url)
}

func (ah *AuthHandler) GoogleOAuth2Callback(c *fiber.Ctx) error {

	// Get the Google code from the callback
	code := c.Query("code")

	log.Printf("code: %s", code)

	token, err := googleOAuth2Config.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// Get the user's info using the access token
	client := googleOAuth2Config.Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	var user GoogleUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	u, err := ah.repo.UpsertUser(c.Context(), datastore.UserInput{
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Email,
	})

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	sess, err := ah.store.Get(c)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	sess.Set("userID", u.ID)
	sess.Save()
	log.Println("Logged in as user", u.ID)

	return c.Redirect("/feed", fiber.StatusMovedPermanently)
}

func GetLoginPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("web/templates/pages/login", nil, "web/templates/layouts/main")
}
