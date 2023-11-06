package controllers

import (
	"main/common/dto"
	"main/common/helpers"
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
	err := services.CreateDocument(auth.ID, data)

	if err != nil {
		return err
	}

	return ctx.Status(fiber.StatusCreated).JSON(data)
}

func FetchDocuments(ctx *fiber.Ctx) error {
	type FetchDocumentsQuery struct {
		helpers.QueryListBase
		Status *string `query:"status"`
	}

	auth := ctx.Locals("auth").(*models.User)
	query := &FetchDocumentsQuery{}
	ctx.QueryParser(query)

	if query.Amount == 0 {
		query.Amount = 10
	}

	docs, err := services.FetchDocumentList(
		&models.Document{UserId: auth.ID},
		query.Amount,
		query.Page,
		query.Status,
	)

	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"docs": docs,
		})
}

func FetchDocument(ctx *fiber.Ctx) error {
	auth := ctx.Locals("auth").(*models.User)

	if auth == nil {
		return &fiber.Error{
			Code:    fiber.ErrForbidden.Code,
			Message: "Forbidden access",
		}
	}

	docId := ctx.Params("id")

	if docId == "" {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Document id is required",
		}
	}
	if document, err := services.FetchDocumentById(docId, auth.ID); err != nil {
		return err
	} else {
		return ctx.
			Status(fiber.StatusOK).
			JSON(document)
	}
}

func UpdateDocument(ctx *fiber.Ctx) error {
	auth := ctx.Locals("auth").(*models.User)
	data := ctx.Locals("json").(*dto.UpdateDocument)

	if auth == nil || data == nil {
		return &fiber.Error{
			Code:    fiber.ErrForbidden.Code,
			Message: "Forbidden access",
		}
	}

	docId := ctx.Params("id")

	if docId == "" {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Document id is required",
		}
	}

	if document, err := services.UpdateDocumentById(docId, auth.ID, data); err != nil {
		return err
	} else {
		return ctx.
			Status(fiber.StatusCreated).
			JSON(document)

	}
}
