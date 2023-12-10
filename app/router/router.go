package router

import (
	"fmt"
	middleware "main/app/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Setup(router *fiber.App) {
	router.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))
	router.Use(middleware.Json)
	router.Use(middleware.Logger())

	router.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{
			"message": "Server is Running!",
		})
	})

	addAuthRoutes(router)
	addUserRoutes(router)
	addDocumentRoutes(router)

	fmt.Println("\nRoutes:")
	for _, route := range router.GetRoutes() {
		if route.Path != "/" && route.Method != "HEAD" {
			fmt.Println("", route.Method, route.Path, " -> ", route.Name)
		}
	}

	router.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}
