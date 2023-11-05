package controllers

import (
	"main/common/dto"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

func CreateDocument(ctx *fiber.Ctx) error {
	auth := ctx.Locals("auth").(*models.User)

	if auth == nil {
		return &fiber.Error{
			Code:    fiber.ErrUnauthorized.Code,
			Message: "User not found",
		}
	}

	data := ctx.Locals("json").(*dto.CreateDocument)

	services.CreateDocument(auth.ID, data)

	return ctx.JSON(fiber.Map{
		"message": "Document created successfully",
		"data":    data,
	})
}
