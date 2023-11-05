package router

import (
	controllers "main/app/controllers"
	"main/app/middleware"
	"main/common/dto"

	"github.com/gofiber/fiber/v2"
)

func addDocumentRoutes(router *fiber.App) {
	document := router.Group("/document")

	document.Post(
		"/",
		middleware.JwtSecurity,
		middleware.JsonValidatorHandler[dto.CreateDocument](),
		controllers.CreateDocument,
	).Name("Create Document")
}
