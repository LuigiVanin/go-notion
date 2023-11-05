package router

import (
	middleware "main/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Setup(router *fiber.App) {
	router.Use(middleware.Json)

	addAuthRoutes(router)
	addUserRoutes(router)
	addDocumentRoutes(router)
	addSwaggerRoutes(router)

	router.Use(func(ctx *fiber.Ctx) error {
		return ctx.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}
