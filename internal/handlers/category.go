package handlers

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/services"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	service *services.CategoryService
}

func NewCategoryHandler(service *services.CategoryService) *CategoryHandler {
	return &CategoryHandler{
		service: service,
	}
}

func (h *CategoryHandler) RegisterCategory(c *fiber.Ctx) error {
	var req dto.CategoryRegisterRequest
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
	response := h.service.RegisterCategory(c, &req)
	return c.Status(response.Code).JSON(response.Obj)
}
