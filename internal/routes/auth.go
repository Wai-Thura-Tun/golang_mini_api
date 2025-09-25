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

	authRouter := router.Group("/auth")

	// Public routes
	authRouter.Post("/register", authHandler.Register)
	authRouter.Post("/login", func(c *fiber.Ctx) error {
		return authHandler.Login(c, config)
	})
	authRouter.Post("/refresh", func(c *fiber.Ctx) error {
		return authHandler.Refresh(c, config)
	})

	// Protected route
	authRouter.Get("/logout", middlewares.Auth(config.JWTSecret, db), authHandler.Logout)
}
