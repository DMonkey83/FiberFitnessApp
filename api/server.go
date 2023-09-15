package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/DMonkey83/FiberFitnessApp/config"
	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	"github.com/DMonkey83/FiberFitnessApp/token"
)

type Server struct {
	config     config.Config
	store      db.Store
	tokenMaker token.Maker
	app        *fiber.App
}

// NewServer function
func NewServer(config config.Config, store db.Store) (*Server, error) {
	tokenMaker, error := token.NewJWTMaker(config.TokenSymmetricKey)
	if error != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", error)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	return server, nil
}

func (server *Server) Start(address string) error {
	app := fiber.New()

	config := cors.Config{
		AllowOrigins:     "*", // Replace with your frontend's URL
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "X-Requested-With,Content-Type,Authorization,Access-Control-Allow-Credentials,Access-Control-Allow-Origin",
		AllowCredentials: true,
	}

	app.Use(
		cors.New(config),
		logger.New(),
	)
	app.Options("/*path", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNoContent)
	})

	server.app = app
	server.SetupRouter(app)
	return app.Listen(address)
}
