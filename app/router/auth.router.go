package router

import (
	controllers "main/app/controllers"
	middleware "main/app/middleware"
	"main/common/dto"

	"github.com/gofiber/fiber/v2"
)

func addAuthRoutes(router *fiber.App) {
	auth := router.Group("/auth")

	auth.Post(
		"/signup", 
		middleware.JsonValidatorHandler[dto.SignupUser](),
		controllers.Signup,
	).Name("User Signup")
}