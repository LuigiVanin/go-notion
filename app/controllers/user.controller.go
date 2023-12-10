package controllers

import (
	"main/common/dto"
	"main/models"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

func FetchUser(ctx *fiber.Ctx) error {
	auth := ctx.Locals("auth").(*models.User)

	if auth == nil {
		return &fiber.Error{
			Code:    fiber.ErrUnauthorized.Code,
			Message: "User not found",
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"email":     auth.Email,
		"id":        auth.ID,
		"name":      auth.Name,
		"createdAt": auth.CreatedAt,
		"updatedAt": auth.UpdatedAt,
	})
}

func UpdateUser(ctx *fiber.Ctx) error {
	auth := ctx.Locals("auth").(*models.User)

	if auth == nil {
		return &fiber.Error{
			Code:    fiber.ErrUnauthorized.Code,
			Message: "User not found",
		}
	}

	updatedUser := ctx.Locals("json").(*dto.UpdateUser)

	err := services.UpdateUser(auth.ID, dto.UpdateUser{Name: updatedUser.Name})

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: "User update failed",
		}
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "User updated successfully",
	})
}
