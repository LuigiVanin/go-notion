package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"

	config "main/config"
	controllers "main/controllers"
	"main/middleware"
)


func main() {
	fmt.Println("Hello, World!")

	env, err := config.LoadEnvData()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := env.DatabaseURL

	db, err	:= config.CreateConnection(url)

	fmt.Println("DATABASE: ", config.Database)
	
	if err != nil {
		fmt.Println(db)
		log.Fatal("Error connecting to database")
		return
	}

	migrationErr := config.Migrate(db)

	if migrationErr != nil {
		fmt.Println("Error migrating database")
	}
	app:= fiber.New()
	app.Use(middleware.Json)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!")
	})
	app.Post("/user", controllers.CreateUser)
	app.Get("/user", controllers.GetCurrentUser)

	log.Fatal(
		app.Listen(fmt.Sprintf(":%d", env.Port)),
	)

}