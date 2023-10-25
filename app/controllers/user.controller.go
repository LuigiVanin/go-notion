package controllers

import (
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func FetchUser(ctx *fiber.Ctx) error {
	auth := ctx.Locals("auth").(*models.User)

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"Email":     auth.Email,
		"ID":        auth.ID,
		"Name":      auth.Name,
		"CreatedAt": auth.CreatedAt,
		"UpdatedAt": auth.UpdatedAt,
	})
}
