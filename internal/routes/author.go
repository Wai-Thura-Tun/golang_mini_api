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

func SetupAuthorRoutes(router fiber.Router, config *config.Config, db *database.DB) {
	authorRepo := repos.NewAuthorRepo(db.Conn())
	authorService := services.NewAuthorService(authorRepo)
	authorHandler := handlers.NewAuthorHandler(authorService)

	authorRouter := router.Group("author", middlewares.Auth(config.JWTSecret, db))

	authorRouter.Post("/register", authorHandler.RegisterAuthor)
}
