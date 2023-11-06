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
	document.
		Get(
			"/",
			middleware.JwtSecurity,
			controllers.FetchDocuments,
		).
		Name("Fetch Documents")
	document.Get(
		"/:id",
		middleware.JwtSecurity,
		controllers.FetchDocument,
	).Name("Fetch Document By Id")
	document.Patch(
		"/:id",
		middleware.JwtSecurity,
		middleware.JsonValidatorHandler[dto.UpdateDocument](),
		controllers.UpdateDocument,
	).Name("Patch Document")

}
