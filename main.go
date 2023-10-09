package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)


func firstPost(ctx *fiber.Ctx) error  {
	db := Database
	user := User{
		Email: "luisfvanin@gmail.com",
		Password: "123456",
		Name: "Luis",
	}
	db.Create(&user)
	
	return ctx.SendString("User Criado!!!")
}

func main() {
	fmt.Println("Hello, World!")

	env, err := loadEnvData()

	if err != nil {
		log.Fatal("Error loading .env file")
		return
	}

	url := env.DatabaseURL

	fmt.Println("URL DO BANCO", url)

	db, err	:= createConnection(url)
	fmt.Println(Database)
	
	if err != nil {
		fmt.Println(db)
		log.Fatal("Error connecting to database")
		return
	}

	migrationErr := migrate(db)

	if migrationErr != nil {
		fmt.Println("Error migrating database")
	}
	app:= fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!")
	})
	app.Post("/user", firstPost)
	log.Fatal(
		app.Listen(fmt.Sprintf(":%d", env.Port)),
	)

}