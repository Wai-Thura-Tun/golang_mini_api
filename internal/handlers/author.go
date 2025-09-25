package handlers

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/services"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
	"github.com/gofiber/fiber/v2"
)

type AuthorHandler struct {
	service *services.AuthorService
}

func NewAuthorHandler(service *services.AuthorService) *AuthorHandler {
	return &AuthorHandler{service: service}
}

func (h *AuthorHandler) RegisterAuthor(c *fiber.Ctx) error {
	var req dto.AuthorRegisterRequest
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

	response := h.service.RegisterAuthor(c, &req)
	return c.Status(response.Code).JSON(response.Obj)
}
