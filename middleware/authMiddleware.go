package middleware

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"

	"github.com/DMonkey83/FiberFitnessApp/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

func AuthMiddleware(tokenMaker token.Maker) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authorizationHeader := ctx.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			return ctx.Status(http.StatusUnauthorized).JSON(err.Error())
		}

		ctx.Locals(AuthorizationPayloadKey, payload)
		return ctx.Next()
	}
}
