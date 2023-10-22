package api

import (
	"github.com/gofiber/fiber/v2"

	auth "github.com/DMonkey83/FiberFitnessApp/middleware"
)

func (server *Server) SetupRouter(app *fiber.App) {
	routes := app.Group("/api")
	// user SetupRouter
	routes.Post("users", server.createUser)
	routes.Post("users/login", server.LoginUser)
	routes.Post("/token/refresh", server.renewAccessToken)
	routes.Post("user_profile", server.createUserProfile)

	authRoutes := routes.Use(auth.AuthMiddleware(server.tokenMaker))

	// user_profile SetupRouter
	authRoutes.Get("user_profile/:username", server.getUserProfile)
	authRoutes.Patch("user_profile", server.updateUserProfile)

	server.app = app
}
