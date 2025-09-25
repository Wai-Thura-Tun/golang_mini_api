package main

import (
	"context"
	"log"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/config"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/database"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
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

	app := fiber.New(fiber.Config{
		AppName: "BooK API",
	})

	app.Use(recover.New())

	app.Get("/test", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Golang!")
	})

	routes.SetUpRoutes(app, config, dbConn)

	log.Fatal(app.Listen(":3000"))
}
