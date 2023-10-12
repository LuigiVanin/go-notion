package controllers

import (
	"fmt"
	"main/common/dto"

	"github.com/gofiber/fiber/v2"
)




func Signup(ctx *fiber.Ctx) error {
	user := ctx.Locals("json").(*dto.SignupUser)
	fmt.Println(user)
	
	return ctx.Status(200).JSON(user)
}
