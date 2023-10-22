package api

import (
	"log"
	"net/http"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	"github.com/DMonkey83/FiberFitnessApp/middleware"
	"github.com/DMonkey83/FiberFitnessApp/token"
	val "github.com/DMonkey83/FiberFitnessApp/util/Validate"
	res "github.com/DMonkey83/FiberFitnessApp/util/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	createAvailablePlanRequest struct {
		PlanName        string              `json:"plan_name" validate:"required"`
		Description     string              `json:"description" validate:"required"`
		Goal            val.Workoutgoalenum `json:"goal" validate:"required"`
		Difficulty      val.Difficulty      `json:"difficulty" validate:"required"`
		IsPublic        val.Visibility      `json:"is_public" validate:"required"`
		CreatorUsername string              `json:"creator_username" validate:"required"`
	}

	getAvailablePlanRequest struct {
		PlanID int64 `uri:"plan_id" binding:"required,min=1"`
	}

	updateAvailablePlanRequest struct {
		CreatorUsername string              `json:"creator_username" validate:"required"`
		PlanName        string              `json:"plan_name"`
		Description     string              `json:"description"`
		Goal            val.Workoutgoalenum `json:"goal"`
		Difficulty      val.Difficulty      `json:"difficulty"`
		IsPublic        val.Visibility      `json:"is_public"`
	}

	listAllAvailablePlansRequest struct {
		Limit  int32 `form:"limit" validate:"required,min=1"`
		Offset int32 `form:"offset" validate:"required"`
	}

	listAllAvailablePlansByCreatorRequest struct {
		Limit           int32  `form:"limit" validate:"required,min=5,max=20"`
		Offset          int32  `form:"offset" validate:"required,min=1"`
		CreatorUsername string `form:"creator" validate:"required"`
	}
)

func (server *Server) createAvailablePlan(ctx *fiber.Ctx) error {
	var req createAvailablePlanRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.CreateAvailablePlanParams{
		PlanName:        req.PlanName,
		Description:     req.Description,
		Goal:            db.Workoutgoalenum(req.Goal),
		Difficulty:      db.Difficulty(req.Difficulty),
		IsPublic:        db.Visibility(req.IsPublic),
		CreatorUsername: authPayload.Username,
	}
	availablePlan, err := server.store.CreateAvailablePlan(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(availablePlan)
}

func (server *Server) getAvailablePlan(ctx *fiber.Ctx) error {
	var req getAvailablePlanRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	availablePlan, err := server.store.GetAvailablePlan(ctx.Context(), req.PlanID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(availablePlan)
}

func (server *Server) updateAvailablePlan(ctx *fiber.Ctx) error {
	var req updateAvailablePlanRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.UpdateAvailablePlanParams{
		CreatorUsername: authPayload.Username,
		PlanName:        pgtype.Text{String: req.PlanName, Valid: true},
		Description:     pgtype.Text{String: req.Description, Valid: true},
		Goal:            db.NullWorkoutgoalenum{Workoutgoalenum: db.Workoutgoalenum(req.Goal), Valid: true},
		Difficulty:      db.NullDifficulty{Difficulty: db.Difficulty(req.Difficulty), Valid: true},
		IsPublic:        db.NullVisibility{Visibility: db.Visibility(req.IsPublic), Valid: true},
	}

	availablePlan, err := server.store.UpdateAvailablePlan(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(availablePlan)
}

func (server *Server) listAllAvailablePlansByCreator(ctx *fiber.Ctx) error {
	var req listAllAvailablePlansByCreatorRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListAvailablePlansByCreatorParams{
		Limit:           req.Limit,
		Offset:          req.Offset,
		CreatorUsername: req.CreatorUsername,
	}
	availablePlans, err := server.store.ListAvailablePlansByCreator(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(availablePlans)
}

func (server *Server) listAllAvailablePlans(ctx *fiber.Ctx) error {
	var req listAllAvailablePlansRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListAvailablePlansByCreatorParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	availablePlans, err := server.store.ListAvailablePlansByCreator(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(availablePlans)
}
