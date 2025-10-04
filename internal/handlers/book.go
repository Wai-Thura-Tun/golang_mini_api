package handlers

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/services"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
	"github.com/gofiber/fiber/v2"
)

type BookHandler struct {
	service *services.BookService
}

func NewBookHandler(service *services.BookService) *BookHandler {
	return &BookHandler{
		service: service,
	}
}

func (h *BookHandler) GetBooks(c *fiber.Ctx) error {
	response := h.service.GetBooks(c)
	return c.Status(response.Code).JSON(response.Obj)
}

func (h *BookHandler) RegisterBook(c *fiber.Ctx) error {
	var req dto.BookRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":  "Invalid request payload",
			"detail": err.Error(),
		})
	}

	if err := util.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":             "Validation Errors",
			"validation_errors": util.ValidationErrorsToMap(err),
		})
	}
	response := h.service.RegisterBook(c, &req)
	return c.Status(response.Code).JSON(response.Obj)
}
