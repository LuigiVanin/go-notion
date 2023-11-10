package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Logger() func(*fiber.Ctx) error {
	info := "[${ip}]:${port} ${status} - ${method} ${path} ${error}\n"
	body := "${body}\n"
	return logger.New(logger.Config{
		Format: info + body,
	})
}
