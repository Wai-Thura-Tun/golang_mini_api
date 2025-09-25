package routes

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/config"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/database"
	"github.com/gofiber/fiber/v2"
)

func SetUpRoutes(app *fiber.App, config *config.Config, db *database.DB) {
	api := app.Group("/api")

	v1 := api.Group("/v1")

	SetUpAuthRoutes(v1, config, db)
	SetupAuthorRoutes(v1, config, db)
	SetupCategoryRoutes(v1, config, db)
	SetupBookRoutes(v1, config, db)
}
