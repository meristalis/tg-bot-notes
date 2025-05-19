package middleware

import (
	"crypto/rsa"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func JWTMiddleware(publicKey *rsa.PublicKey) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Получаем токен из заголовка Authorization
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Missing Authorization header")
		}

		// Ожидаем формат "Bearer <token>"
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token format")
		}

		// Проверка JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверяем, что алгоритм подписи совпадает с ожидаемым
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return publicKey, nil
		})

		// Проверка ошибок валидации
		if err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token: "+err.Error())
		}

		// Проверка на срок действия токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			if exp, ok := claims["exp"].(float64); ok {
				if time.Now().Unix() > int64(exp) {
					return fiber.NewError(fiber.StatusUnauthorized, "Token has expired")
				}
			}
		} else {
			return fiber.NewError(fiber.StatusUnauthorized, "Invalid token claims")
		}

		// Если токен валиден, передаем запрос дальше
		return c.Next()
	}
}
