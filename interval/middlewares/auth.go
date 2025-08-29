package middlewares

import (
	"go/token"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Auth(secret string) fiber.Handler {
	return func (c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid Authorization header",
			})
		}

	    tokenString := strings.TrimPrefix(authHeader, "Bearer")
		jwt.Parse(tokenString, func (token *jwt.Token) (interface{}, error) {
			
		})
	}
}