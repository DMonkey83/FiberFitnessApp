package api

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	"github.com/DMonkey83/FiberFitnessApp/middleware"
	"github.com/DMonkey83/FiberFitnessApp/token"
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

	userProfileResponse struct {
		Username      string         `uri:"username"`
		FullName      string         `json:"full_name"`
		Email         string         `json:"email"`
		Age           int32          `json:"age"`
		Gender        string         `json:"gender"`
		HeightCm      int32          `json:"height_cm"`
		HeightFtIn    string         `json:"height_ft_in"`
		PreferredUnit val.Weightunit `json:"preferred_unit"`
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

// newUserResponse function

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
		return res.ResponseError(ctx, nil, "Error while creating user profile")
	}

	return ctx.Status(http.StatusOK).JSON(userProfile)
}

func (server *Server) getUserProfile(ctx *fiber.Ctx) error {
	var req getUserProfileRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	log.Print(ctx.Cookies("refresh_token"))

	userProfile, err := server.store.GetUserProfile(ctx.Context(), req.Username)
	if err != nil {
		if err == db.ErrRecordNotFound {
			log.Print("getUserProfile", req.Username)
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	rsp := userProfileResponse{
		Username:      authPayload.Username,
		FullName:      userProfile.Userprofile.FullName,
		Email:         userProfile.User.Email,
		Age:           userProfile.Userprofile.Age,
		Gender:        userProfile.Userprofile.Gender,
		HeightCm:      userProfile.Userprofile.HeightCm,
		HeightFtIn:    userProfile.Userprofile.HeightFtIn,
		PreferredUnit: val.Weightunit(userProfile.Userprofile.PreferredUnit),
	}

	return ctx.Status(http.StatusOK).JSON(rsp)
}

func (server *Server) updateUserProfile(ctx *fiber.Ctx) error {
	var req updateUserProfileRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateUserProfileParams{
		Username:      req.Username,
		FullName:      pgtype.Text{String: req.FullName, Valid: req.FullName != ""},
		Age:           pgtype.Int4{Int32: req.Age, Valid: req.Age == 0},
		Gender:        pgtype.Text{String: req.Gender, Valid: req.Gender != ""},
		HeightCm:      pgtype.Int4{Int32: req.HeightCm, Valid: req.HeightCm == 0},
		HeightFtIn:    pgtype.Text{String: req.HeightFtIn, Valid: req.HeightFtIn != ""},
		PreferredUnit: db.NullWeightunit{Weightunit: db.Weightunit(req.PreferredUnit), Valid: req.PreferredUnit != ""},
	}

	userProfile, err := server.store.UpdateUserProfile(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(userProfile)
}
