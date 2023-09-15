package token

import (
	"time"

	"github.com/gofiber/fiber/v2"

	"github.com/DMonkey83/FiberFitnessApp/config"
)

type Maker interface {
	CreateToken(username string, duration time.Duration) (string, *Payload, error)
	GenerateAccessCookie(token string, config config.Config, Username string, ctx *fiber.Ctx) error
	GenerateRefreshCookie(token string, config config.Config, Username string, ctx *fiber.Ctx) error
	VerifyToken(token string) (*Payload, error)
}
