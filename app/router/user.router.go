package router

import (
	controllers "main/app/controllers"
	middleware "main/app/middleware"

	"github.com/gofiber/fiber/v2"
)

func addUserRoutes(router *fiber.App) {
	user := router.Group("/user")

	user.Get(
		"/",
		middleware.JwtSecurity,
		controllers.FetchUser,
	).Name("Fetch User")
}
