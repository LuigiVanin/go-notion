package app

import (
	"fmt"
	"log"
	"main/app/middleware"
	"main/app/router"
	config "main/config"

	"github.com/gofiber/fiber/v2"
)

func Run(env *config.EnvData) {
	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorMiddleware,
	})

	router.Setup(app)

	log.Fatal(
		app.Listen(fmt.Sprintf(":%d", env.Port)),
	)

}
