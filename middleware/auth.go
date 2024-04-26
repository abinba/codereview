package middleware

import (
	"fmt"
	"strings"

	"github.com/abinba/codereview/config"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"

	"time"
)

var jwtSecretKey = []byte(config.Config("JWT_SECRET_KEY"))

func GenerateJWT(UserID uuid.UUID) (string, error) {
	claims := jwt.MapClaims{
		"user_id": UserID,
		"exp":     time.Now().Add(time.Minute * 180).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecretKey)
}

func JWTProtected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")

		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Missing token"})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token format"})
		}
		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return jwtSecretKey, nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Locals("user_id", claims["user_id"])
			if c.Params("id") != "" && claims["user_id"] != c.Params("id") {
				return c.Status(fiber.StatusForbidden).JSON(fiber.Map{"status": "error", "message": "Access denied"})
			}

			return c.Next()
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid token", "data": err})
		}
	}
}
