package app

import (
	"fmt"
	"log"
	"main/app/router"
	config "main/config"

	"github.com/gofiber/fiber/v2"
)

func Run(env *config.EnvData) {
	app := fiber.New()

	router.Setup(app)

	fmt.Println("\nRoutes:")
	for _, route := range app.GetRoutes() {
		if route.Path != "/" {
			fmt.Println("", route.Method, route.Path, " -> ", route.Name)
		}
	}

	log.Fatal(
		app.Listen(fmt.Sprintf(":%d", env.Port)),
	)

}
