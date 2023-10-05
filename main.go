package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func firstGet(ctx *fiber.Ctx) error  {
	ctx.JSON(ctx.Body())
	return ctx.SendString("Hello, World!!!")
}

func main() {
	fmt.Println("Hello, World!")

	env, err := loadEnvData()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	url := env.DatabaseURL

	fmt.Println("URL DO BANCO", url)

	app:= fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!")
	})
	app.Get("/api", firstGet)
	app.Listen(fmt.Sprint(":%d", env.Port))

}