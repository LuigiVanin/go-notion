package services

import (
	"main/common/dto"
	conn "main/config/database"
	"main/models"

	"github.com/gofiber/fiber/v2"
)

func CreateDocument(userId uint, data *dto.CreateDocument) error {
	if data != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "User not found",
		}
	}

	conn.Database.Create(&models.Document{
		UserId: userId,
		Title:  data.Title,
		Text:   data.Text,
		Status: data.Status,
	})

	return nil
}
