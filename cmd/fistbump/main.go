package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/rianfowler/fist-bump-chat/ent"
	"github.com/rianfowler/fist-bump-chat/internal/config"
	"github.com/rianfowler/fist-bump-chat/internal/datastore"
	"github.com/rianfowler/fist-bump-chat/internal/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/session"

	"github.com/gofiber/template/html/v2"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed web/*
var templatesFS embed.FS

func main() {
	cfg := config.New()

	// Example usage
	port := cfg.GetString("POSTGRES_PORT")
	fmt.Println("Running on port:", port)

	dbDriver := os.Getenv("DB_DRIVER")
	var dbClient *ent.Client
	var err error

	if dbDriver == "postgres" {
		pgurl, err := datastore.MakePostgresUrl()

		if err != nil {
			log.Fatalf("Failed to make Postgres connection string: %v", err)
		}
		dbClient, err = ent.Open("postgres", pgurl)
		if err != nil {
			log.Fatalf("failed opening connection to postgres: %v", err)
		}

	} else {
		// Initialize SQLite3
		dbClient, err = ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
		if err != nil {
			log.Fatalf("failed opening connection to sqlite: %v", err)
		}
	}

	defer dbClient.Close()
	if err := dbClient.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	repo := datastore.NewEntRepository(dbClient)
	store := session.New(session.Config{
		Expiration: 24 * time.Hour,
	})

	feedHandler := handlers.NewFeedHandler(repo)
	authHandler := handlers.NewAuthHandler(repo, store, cfg)

	engine := html.NewFileSystem(http.FS(templatesFS), ".html")
	engine.Reload(true)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", handlers.GetRootPage)
	app.Get("/login", handlers.GetLoginPage)
	app.Get("auth/login/google", authHandler.GoogleLoginStart)
	app.Get("auth/callback/google", authHandler.GoogleOAuth2Callback)
	staticFS, err := fs.Sub(templatesFS, "web/static")
	if err != nil {
		log.Fatalf("failed to create static fs directory: %v", err)
	}
	app.Use("/static", filesystem.New(filesystem.Config{
		Root: http.FS(staticFS),
	}))
	// app.Static("/static", "./web/static")

	// add mock user if SIGNIN_ENABLED is false
	if !cfg.GetBool("SIGNIN_ENABLED") {
		app.Use(handlers.NewMockUserMiddleware(repo, store).Middleware)
	}
	app.Use(handlers.NewUserMiddleware(repo, store).Middleware)

	protected := app.Group("/", handlers.EnsureAuthenticated)
	protected.Get("/feed", feedHandler.GetFeedPage)
	protected.Get("/new-message", feedHandler.GetNewMessagePage)
	protected.Post("/new-message", feedHandler.PostNewMessage)

	app.Use(handlers.Get404Page)
	log.Fatal(app.Listen(":8080"))
}
