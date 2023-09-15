package api

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	"github.com/gofiber/fiber/v2"

	"github.com/DMonkey83/FiberFitnessApp/util/response"
)

type renewAccessTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type renewAccessTokenResponse struct {
	AccessToken          string    `json:"access_token"`
	AccessTokenExpiresAt time.Time `json:"access_token_expires_at"`
}

func (server *Server) renewAccessToken(ctx *fiber.Ctx) error {
	var req renewAccessTokenRequest
	if err := ctx.BodyParser(&req); err != nil {
		return response.ResponseValidationError(ctx, nil, err.Error())
	}

	refreshPayload, err := server.tokenMaker.VerifyToken(req.RefreshToken)
	if err != nil {
		return response.ResponseUnauthenticated(ctx, nil, err.Error())
	}

	session, err := server.store.GetSession(ctx.Context(), refreshPayload.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			return response.ResponseNotFound(ctx, nil, err.Error())
		}
		return response.ResponseError(ctx, nil, err.Error())
	}

	if session.IsBlocked {
		err := fmt.Errorf("blocked session")
		return response.ResponseUnauthenticated(ctx, nil, err.Error())
	}

	if session.Username != refreshPayload.Username {
		err := fmt.Errorf("incorrect session user")
		return response.ResponseUnauthenticated(ctx, nil, err.Error())
	}

	if session.RefreshToken != req.RefreshToken {
		err := fmt.Errorf("mismatched session token")
		return response.ResponseUnauthenticated(ctx, nil, err.Error())
	}

	if time.Now().After(session.ExpiresAt) {
		err := fmt.Errorf("expired session")
		return response.ResponseUnauthenticated(ctx, nil, err.Error())
	}

	accessToken, accessPayload, err := server.tokenMaker.CreateToken(
		refreshPayload.Username,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		return response.ResponseError(ctx, nil, err.Error())
	}

	rsp := renewAccessTokenResponse{
		AccessToken:          accessToken,
		AccessTokenExpiresAt: accessPayload.ExpiredAt,
	}
	return ctx.Status(http.StatusOK).JSON(rsp)
}
