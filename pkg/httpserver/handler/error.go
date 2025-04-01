// pkg/handler/error.go
package handler

import (
	"github.com/gofiber/fiber/v2"
)

// Response - структура для ответа с ошибкой
type Response struct {
	Error string `json:"error" example:"message"`
}

// ErrorResponse - универсальный метод для отправки ошибки в API
func ErrorResponse(ctx *fiber.Ctx, code int, msg string) error {
	return ctx.Status(code).JSON(Response{msg})
}
