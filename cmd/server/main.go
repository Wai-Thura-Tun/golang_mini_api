package main

import (
	"context"
	"log"

	"github.com/Wai-Thura-Tun/golang_book_api/interval/config"
	"github.com/Wai-Thura-Tun/golang_book_api/interval/database"
	"github.com/gofiber/fiber/v2"
)

func main() {

	ctx := context.Background()

	// Load configuration
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// Connect database
	dbConn, err := database.NewDB(ctx, config.Database_URL)

	if err != nil {
		log.Fatal(err)
	}

	defer dbConn.Close()

	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Run database migrations
	if err := database.RunMigrations(config.Database_URL, dbConn); err != nil {
		log.Fatal("Failed to run database migrations:", err)
	}

	app := fiber.New(fiber.Config{
		AppName: "BooK API",
	})

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Golang!")
	})

	log.Fatal(app.Listen(":3000"))
}
