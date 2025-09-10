package handlers

import (
	"fmt"

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

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	return nil
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var req dto.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err,
		})
	}
	if err := util.Validate.Struct(req); err != nil {
		fmt.Print(err)
		return c.SendString("")
	}

	return h.service.CreateUser(c.Context(), req)
}

func (h *AuthHandler) Refresh(c *fiber.Ctx) error {
	return nil
}

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return nil
}
