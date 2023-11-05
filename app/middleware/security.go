package middleware

import (
	"main/services"

	"github.com/gofiber/fiber/v2"
)

func JwtSecurity(ctx *fiber.Ctx) error {
	type AuthHeader struct {
		Authorization string `reqHeader:"Authorization"`
	}

	header := AuthHeader{}

	if err := ctx.ReqHeaderParser(&header); err != nil {
		return err
	}

	jwtToken := header.Authorization

	data, err := services.ParseJwtToken(jwtToken)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: err.Error(),
		}
	}

	user, err := services.FetchUser(data.UserId)

	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrUnauthorized.Code,
			Message: "User already exists",
		}
	}

	ctx.Locals("auth", user)
	return ctx.Next()
}
