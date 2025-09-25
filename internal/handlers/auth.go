package handlers

import (
	"github.com/Wai-Thura-Tun/golang_book_api/internal/config"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/dto"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/services"
	"github.com/Wai-Thura-Tun/golang_book_api/internal/util"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		service: authService,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx, config *config.Config) error {
	var req dto.LoginRequest
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
	response := h.service.Login(c, &req, config.JWTSecret)
	return c.Status(response.Code).JSON(response.Obj)
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest

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
	response := h.service.CreateUser(c, &req)
	return c.Status(response.Code).JSON(response.Obj)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx, config *config.Config) error {
	var req dto.RefreshRequest
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
	response := h.service.RefreshAccessToken(c, &req, config.JWTSecret)
	return c.Status(response.Code).JSON(response.Obj)
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(uint64)
	response := h.service.Logout(c, userID)
	return c.Status(response.Code).JSON(response.Obj)
}
