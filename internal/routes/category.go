package routes

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/config"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/database"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/handlers"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/middlewares"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/repos"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/services"
	"github.com/gofiber/fiber/v2"
)

func SetupCategoryRoutes(router fiber.Router, config *config.Config, db *database.DB) {
	categoryRepo := repos.NewCategoryRepo(db.Conn())
	categoryService := services.NewCategoryService(categoryRepo)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	categoryRouter := router.Group("/category", middlewares.Auth(config.JWTSecret, db))

	categoryRouter.Post("/register", categoryHandler.RegisterCategory)
}
