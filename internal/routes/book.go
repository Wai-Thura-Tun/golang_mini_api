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

func SetupBookRoutes(router fiber.Router, config *config.Config, db *database.DB) {

	bookRepo := repos.NewBookRepo(db.Conn())
	bookService := services.NewBookService(bookRepo)
	bookHandler := handlers.NewBookHandler(bookService)
	bookRouter := router.Group("/book", middlewares.Auth(config.JWTSecret, db))

	bookRouter.Get("/list", bookHandler.GetBooks)
	bookRouter.Post("/register", bookHandler.RegisterBook)
}
