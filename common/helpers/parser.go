package helpers

import "github.com/gofiber/fiber/v2"

func ParseJsonBody[T interface{}](ctx *fiber.Ctx) (*T, *fiber.Error) {
	var data T
	if err := ctx.BodyParser(&data); err != nil {
		return nil, &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "Invalid request",
		}
	}
	return &data, nil
}
