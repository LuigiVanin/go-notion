package services

import (
	"main/common/dto"
	conn "main/config/database"
	"main/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateDocument(userId uint, data *dto.CreateDocument) error {
	document := &models.Document{}
	queryRes := conn.
		Database.
		Where(&models.Document{UserId: userId, Title: data.Title}).
		First(document)

	if queryRes.Error == nil {
		return &fiber.Error{
			Code:    fiber.ErrConflict.Code,
			Message: "Document with title already exists",
		}
	}

	res := conn.Database.Create(&models.Document{
		UserId: userId,
		Title:  data.Title,
		Text:   data.Text,
		Status: data.Status,
	})

	return res.Error
}

func FetchDocumentList(
	fields *models.Document,
	amount int,
	page int,
	status *string,
) ([]models.Document, error) {

	var documents []models.Document
	if status != nil {
		fields.Status = *status
	}
	res := conn.
		Database.
		Preload("User").
		Where(fields).
		Limit(amount).
		Offset((page - 1) * amount).
		Find(&documents)

	if res.Error != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: res.Error.Error(),
		}
	}

	return documents, nil
}

func FetchDocumentById(id string, userId uint) (*models.Document, error) {
	idUint, err := strconv.Atoi(id)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	document := &models.Document{}

	if res := conn.
		Database.
		Preload("User").
		Where(
			&models.Document{
				Base: models.Base{
					ID: uint(idUint),
				},
			},
		).
		First(document, id); res.Error != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: res.Error.Error(),
		}
	}

	if document.UserId != userId {
		return nil, &fiber.Error{
			Code:    fiber.ErrForbidden.Code,
			Message: "Forbidden access",
		}
	}

	return document, nil
}

func UpdateDocumentById(
	id string,
	userId uint,
	data *dto.UpdateDocument,
) (*models.Document, error) {
	idUint, err := strconv.Atoi(id)
	if err != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}
	query := &models.Document{
		Base: models.Base{
			ID: uint(idUint),
		},
	}

	document := &models.Document{}
	if res := conn.Database.Where(query).First(document); document.UserId != userId || res.Error != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrForbidden.Code,
			Message: "Forbidden access",
		}
	}

	res := conn.
		Database.
		Preload("User").
		Where(query).
		Updates(&models.Document{
			Title:  data.Title,
			Text:   data.Text,
			Status: data.Status,
		})

	return document, res.Error
}
