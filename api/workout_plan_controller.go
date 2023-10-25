package api

import (
	"log"
	"net/http"
	"time"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	"github.com/DMonkey83/FiberFitnessApp/middleware"
	"github.com/DMonkey83/FiberFitnessApp/token"
	val "github.com/DMonkey83/FiberFitnessApp/util/Validate"
	res "github.com/DMonkey83/FiberFitnessApp/util/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	createWorkoutPlanRequest struct {
		Username    string              `json:"username" validate:"required"`
		PlanName    string              `json:"plan_name" validate:"required"`
		Description string              `json:"description" validate:"required"`
		StartDate   time.Time           `json:"start_date" validate:"required"`
		EndDate     time.Time           `json:"end_date" validate:"required"`
		Goal        val.Workoutgoalenum `json:"goal" validate:"required"`
		Difficulty  val.Difficulty      `json:"difficulty" validate:"required"`
		IsPlublic   val.Visibility      `json:"is_public" validate:"required"`
	}

	getWorkoutPlanRequest struct {
		PlanID   int64  `uri:"plan_id" binding:"required,min=1"`
		Username string `uri:"username" binding:"required"`
	}

	updateWorkoutPlanRequest struct {
		PlanID      int64               `json:"plan_id" validate:"required"`
		PlanName    string              `json:"plan_name" validate:"required"`
		Username    string              `json:"username" validate:"required"`
		Description string              `json:"description"`
		Startdate   time.Time           `json:"start_date"`
		EndDate     time.Time           `json:"end_date"`
		Goal        val.Workoutgoalenum `json:"goal"`
		IsPublic    val.Visibility      `json:"is_public"`
		Difficulty  val.Difficulty      `json:"difficulty"`
	}
)

func (server *Server) createWorkoutPlan(ctx *fiber.Ctx) error {
	var req createWorkoutPlanRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.CreatePlanParams{
		Username:    authPayload.Username,
		PlanName:    req.PlanName,
		Description: req.Description,
		StartDate:   req.StartDate,
		EndDate:     req.EndDate,
		Goal:        db.Workoutgoalenum(req.Goal),
		Difficulty:  db.Difficulty(req.Difficulty),
		IsPublic:    db.Visibility(req.IsPlublic),
	}

	plan, err := server.store.CreatePlan(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(plan)
}

func (server *Server) getWorkoutPlan(ctx *fiber.Ctx) error {
	var req getWorkoutPlanRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.GetPlanParams{
		PlanID:   req.PlanID,
		Username: authPayload.Username,
	}

	plan, err := server.store.GetPlan(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(plan)
}

func (server *Server) updateWorkoutPlan(ctx *fiber.Ctx) error {
	var req updateWorkoutPlanRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.UpdatePlanParams{
		PlanID:      req.PlanID,
		Username:    authPayload.Username,
		PlanName:    pgtype.Text{String: req.PlanName, Valid: req.PlanName != ""},
		Goal:        db.NullWorkoutgoalenum{Workoutgoalenum: db.Workoutgoalenum(req.Goal), Valid: req.Goal != ""},
		Difficulty:  db.NullDifficulty{Difficulty: db.Difficulty(req.Difficulty), Valid: req.Difficulty != ""},
		IsPublic:    db.NullVisibility{Visibility: db.Visibility(req.IsPublic), Valid: req.IsPublic != ""},
		Description: pgtype.Text{String: req.Description, Valid: req.Description != ""},
		StartDate:   pgtype.Timestamptz{Time: req.Startdate, Valid: req.Startdate != time.Time{}},
		EndDate:     pgtype.Timestamptz{Time: req.EndDate, Valid: req.EndDate != time.Time{}},
	}
	plan, err := server.store.UpdatePlan(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(plan)
}
