package middleware

import (
	"main/internals/errors"

	"github.com/gofiber/fiber/v2"
)

func ErrorMiddleware(ctx *fiber.Ctx, err error) error {
	if appErr, ok := err.(*errors.AppError); ok {
		return ctx.Status(appErr.Code).JSON(fiber.Map{
			"code":    appErr.Code,
			"message": appErr.Message,
			"fields":  appErr.Fields,
		})
	}

	if appErr, ok := err.(*fiber.Error); ok {
		return ctx.Status(appErr.Code).JSON(fiber.Map{
			"code":    appErr.Code,
			"message": appErr.Message,
		})
	}

	return ctx.Status(500).JSON(fiber.Map{
		"code":    500,
		"message": "Internal Server Error",
	})
}
