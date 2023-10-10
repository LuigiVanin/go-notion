package controllers

import (
	"fmt"
	config "main/config"
	models "main/models"

	"github.com/gofiber/fiber/v2"
)


func GetCurrentUser(ctx *fiber.Ctx) error {
	result := config.Database.Find(&models.User{})
	rows, err := result.Rows()
	if err != nil {
		fmt.Println(err)
	}
 	fmt.Println("Rows: ", rows)

	ctx.Status(200).JSON(models.User{
		Email: "luisfvanin",
	})
	return nil
}

type SignupUser struct {
	Email string `json:"email"`
}

func CreateUser(ctx *fiber.Ctx) error {
	// db := config.Database
	user := new(models.User)
	err := ctx.BodyParser(user)
	if err != nil {
		fmt.Println("Error!")
		fmt.Println(err)
	}
	fmt.Println(user)
	
	
	return nil
}