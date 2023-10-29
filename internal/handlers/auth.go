package handlers

import (
	"encoding/json"
	"fistbump/internal/config"
	"fistbump/internal/datastore"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
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

type GithubUser struct {
	Email    string `json:"email"`
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"login"`
}

type GithubUserEmails struct {
	Email   string `json:"email"`
	Primary bool   `json:"primary"`
}

type GoogleUser struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

var githubOAuth2Config = &oauth2.Config{
	ClientID:     "",
	ClientSecret: "",
	RedirectURL:  "https://www.fistbump.chat/auth/callback",
	Scopes:       []string{"user:email"},
	Endpoint:     github.Endpoint,
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

func (ah *AuthHandler) GithubLoginStart(c *fiber.Ctx) error {
	log.Println("Start Google Auth called")
	githubOAuth2Config.ClientID = ah.cfg.GetString("GITHUB_CLIENT_ID")
	githubOAuth2Config.ClientSecret = ah.cfg.GetString("GITHUB_CLIENT_SECRET")
	githubOAuth2Config.RedirectURL = "https://www.fistbump.chat/auth/callback"
	// Generate OAuth2 URL and redirect to it
	url := githubOAuth2Config.AuthCodeURL("random-state")
	log.Printf("redirecting to: %s", url)
	return c.Redirect(url)
}

func (ah *AuthHandler) GithubOAuth2Callback(c *fiber.Ctx) error {
	log.Println("handle oauth called")
	code := c.Query("code")
	log.Printf("code: %s", code)
	token, err := githubOAuth2Config.Exchange(c.Context(), code)
	if err != nil {
		return c.Status(500).SendString("failed to exchange token: " + err.Error())
	}

	client := githubOAuth2Config.Client(c.Context(), token)

	// Channels for results and errors
	userChan := make(chan *GithubUser, 1)
	emailChan := make(chan string, 1)
	errChan := make(chan error, 2) // we might have 2 errors simultaneously, one from each goroutine

	// Fetch user information
	go func() {
		user, err := fetchGithubUser(client)
		if err != nil {
			errChan <- err
			return
		}
		userChan <- user
	}()

	// Fetch user email
	go func() {
		email, err := fetchGithubUserEmail(client)
		if err != nil {
			errChan <- err
			return
		}
		emailChan <- email
	}()

	var user *GithubUser
	var email string
	// Counter for received results
	received := 0

	for received < 2 { // we expect 2 results: one for user and one for email
		select {
		case user = <-userChan:
			received++
		case email = <-emailChan:
			received++
		case err = <-errChan:
			return c.Status(500).SendString(err.Error())
		}
	}

	inputEmail := email
	if user.Email != "" {
		inputEmail = user.Email
	}

	u, err := ah.repo.UpsertUser(c.Context(), datastore.UserInput{
		Email:    inputEmail,
		Name:     user.Name,
		Username: user.Username,
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

func fetchGithubUser(client *http.Client) (*GithubUser, error) {
	resp, err := client.Get("https://api.github.com/user")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("failed to read user info: %v", err)
	}

	var user GithubUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	return &user, nil
}

func fetchGithubUserEmail(client *http.Client) (string, error) {

	// Just for the sake of illustration:
	resp, err := client.Get("https://api.github.com/user/emails")
	if err != nil {
		return "", fmt.Errorf("failed to get user email: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read user emails: %v", err)
	}

	var emails []GithubUserEmails
	err = json.Unmarshal(body, &emails)
	if err != nil {
		// handle the error
		return "", fmt.Errorf("failed to decode user emails: %v", err)
	}

	var primaryEmail string
	for _, email := range emails {
		if email.Primary {
			primaryEmail = email.Email
			break
		}
	}

	return primaryEmail, nil
}

func GetLoginPage(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).Render("web/templates/pages/login", nil, "web/templates/layouts/main")
}
