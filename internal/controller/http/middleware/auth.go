package middleware

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/meristalis/tg-bot-notes/pkg/logger"
)

func JWTMiddleware(publicKey *rsa.PublicKey, logger logger.Interface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: "+err.Error())
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
		}

		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return fiber.NewError(fiber.StatusUnauthorized, "Token has expired")
			}
		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token expiration")
		}
		// Извлекаем user_id
		email, ok := claims["email"].(string)
		if !ok || email == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "email not found in token")
		}
		// Извлекаем user_id
		userID, ok := claims["uid"].(string)
		if !ok || userID == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "User ID not found in token")
		}
		c.Locals("user_id", userID)
		return c.Next()
	}
}
