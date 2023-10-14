package controllers

import (
	"fmt"
	"main/common/dto"
	"main/services"

	"github.com/gofiber/fiber/v2"
)

func Signup(ctx *fiber.Ctx) error {
	singupUserData := ctx.Locals("json").(*dto.SignupUser)
	fmt.Println(singupUserData)

	err := services.CreateUser(*singupUserData)

	if err != nil {
		return err
	}

	return ctx.Status(200).JSON(singupUserData)
}
