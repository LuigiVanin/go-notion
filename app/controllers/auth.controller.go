package controllers

import (
	"main/common/dto"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {
	singupUserData := ctx.Locals("json").(*dto.SignupUser)
	err := services.CreateUser(*singupUserData)

	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(singupUserData)
}

func Login(ctx *fiber.Ctx) error {
	loginUserData := ctx.Locals("json").(*dto.LoginUser)
	result, err := services.LoginUser(*loginUserData)

	if err != nil {
		return err
	}

	return ctx.
		Status(fiber.StatusCreated).
		JSON(result)
}
