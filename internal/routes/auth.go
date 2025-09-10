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

func SetUpAuthRoutes(router fiber.Router, config *config.Config, db *database.DB) {
	authRepo := repos.NewAuthRepo(db.Conn())
	authService := services.NewAuthService(authRepo)
	authHandler := handlers.NewAuthHandler(authService)

	// Public routes
	router.Post("/register", authHandler.Register)
	router.Post("/login", nil)
	router.Post("/refresh", nil)

	// Protected route
	router.Post("/logout", middlewares.Auth(config.JWTSecret, db))
}
