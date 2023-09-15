package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	val "github.com/DMonkey83/FiberFitnessApp/util/Validate"
	res "github.com/DMonkey83/FiberFitnessApp/util/response"
)

type (
	createUserProfileRequest struct {
		Username      string         `uri:"username" validate:"required,min=1"`
		FullName      string         `json:"full_name" validate:"required"`
		Age           int32          `json:"age" validate:"required"`
		Gender        string         `json:"gender" validate:"required,oneof=female male"`
		HeightCm      int32          `json:"height_cm" validate:"required"`
		HeightFtIn    string         `json:"height_ft_in"`
		PreferredUnit val.Weightunit `json:"preferred_unit" validate:"required"`
	}

	getUserProfileRequest struct {
		Username string `uri:"username" binding:"required,min=1"`
	}

	updateUserProfileRequest struct {
		Username      string         `uri:"username" binding:"required,min=1"`
		FullName      string         `json:"full_name"`
		Age           int32          `json:"age"`
		Gender        string         `json:"gender" binding:"oneof=female male"`
		HeightCm      int32          `json:"height_cm"`
		HeightFtIn    string         `json:"height_ft_in"`
		PreferredUnit val.Weightunit `json:"preferred_unit" binding:"weight_unit"`
	}
)

func (server *Server) createUserProfile(ctx *fiber.Ctx) error {
	var req createUserProfileRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.CreateUserProfileParams{
		Username:      req.Username,
		FullName:      req.FullName,
		Age:           req.Age,
		Gender:        req.Gender,
		HeightCm:      req.HeightCm,
		HeightFtIn:    req.HeightFtIn,
		PreferredUnit: db.Weightunit(req.PreferredUnit),
	}
	userProfile, err := server.store.CreateUserProfile(ctx.Context(), arg)
	if err != nil {
		errorCode := db.ErrorCode(err)
		if errorCode == db.UniqueViolation {
			return res.ResponseUnauthenticated(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, "Error while creating user")
	}

	return ctx.Status(http.StatusOK).JSON(userProfile)
}

func (server *Server) getUserProfile(ctx *fiber.Ctx) error {
	var req getUserProfileRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	userProfile, err := server.store.GetUserProfile(ctx.Context(), req.Username)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(userProfile)
}

func (server *Server) updateUserProfile(ctx *fiber.Ctx) error {
	var req updateUserProfileRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateUserProfileParams{
		Username:      req.Username,
		FullName:      req.FullName,
		Age:           req.Age,
		Gender:        req.Gender,
		HeightCm:      req.HeightCm,
		HeightFtIn:    req.HeightFtIn,
		PreferredUnit: db.Weightunit(req.PreferredUnit),
	}

	userProfile, err := server.store.UpdateUserProfile(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(userProfile)
}
