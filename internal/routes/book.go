package routes

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/config"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/database"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetupBookRoutes(router fiber.Router, config *config.Config, db *database.DB) {
	bookRouter := router.Group("/book", middlewares.Auth(config.JWTSecret, db))

	bookRouter.Get("/list")
	bookRouter.Post("/register")
}
