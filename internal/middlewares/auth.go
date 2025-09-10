package middlewares

import (
	"strings"
	"time"

	"github.com/Wai-Thura-Tun/golang_book_api/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Auth(secret string, db *database.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Missing or invalid authorization header",
			})
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.NewError(fiber.StatusUnauthorized, "Invalid token signing method")
			}
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid or expired token",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if userID, ok := claims["user_id"].(float64); ok {
				isExist, err := db.CheckUserExists(uint(userID))
				if err != nil || !isExist {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error": "User does not exist",
					})
				}
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid user id in token claims",
				})
			}

			if exp, ok := claims["exp"].(float64); ok {
				expTime := time.Unix(int64(exp), 0)
				if time.Now().After(expTime) {
					return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
						"error": "Token has expired",
					})
				}
			} else {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid expiration in token claims",
				})
			}

		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		return c.Next()
	}
}
