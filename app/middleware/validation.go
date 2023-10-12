package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	h "main/common/helpers"
)

func JsonValidatorHandler[T interface{}]() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		data, err := h.ParseJsonBody[T](ctx)
		if err != nil {
			return &fiber.Error{
				Code: fiber.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}
		if err := h.ValidateJson(data); err != nil {
			fmt.Println(err.Error())
			return &fiber.Error{
				Code: fiber.ErrBadRequest.Code,
				Message: err.Error(),
			}
		}
		ctx.Locals("json", data)
		return ctx.Next()
	}
}