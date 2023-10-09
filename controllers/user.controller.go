package controllers

import (
	"fmt"
	i "main/internal"

	"github.com/gofiber/fiber/v2"
)


func getCurrentUser(ctx *fiber.Ctx) error {
	result := i.Database.Find(&i.User{})
	rows, err := result.Rows()
	if err != nil {
		fmt.Println(err)
	}
 	fmt.Println(rows)

	ctx.Status(200).JSON(i.User{
		Email: "luisfvanin",
	})
	return nil
}