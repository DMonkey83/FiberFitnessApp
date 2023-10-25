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
	createMaxRepGoalRequest struct {
		Username     string `json:"username" validate:"required"`
		ExerciseName string `json:"exercise_name" validate:"required"`
		GoalReps     int32  `json:"goal_reps" validate:"required"`
		Notes        string `json:"notes" validate:"required"`
	}

	getMaxRepGoalRequest struct {
		GoalID       int64  `uri:"goal_id" binding:"required,min=1"`
		Username     string `uri:"username" binding:"required"`
		ExerciseName string `uri:"exercise_name" binding:"required"`
	}

	updateMaxRepGoalRequest struct {
		GoalID       int64  `json:"goal_id" binding:"required,min=1"`
		Username     string `json:"username" validate:"required"`
		ExerciseName string `json:"exercise_name" validate:"required"`
		GoalReps     int32  `json:"goal_reps" validate:"required"`
		Notes        string `json:"notes" validate:"required"`
	}

	listMaxRepGoalsRequest struct {
		ExerciseName string `form:"exercise_name" binding:"required,min=1"`
		Username     string `form:"username" binding:"required,min=1"`
		Limit        int32  `form:"limit" validate:"required,min=1"`
		Offset       int32  `form:"offset" validate:"required"`
	}
)

func (server *Server) createMaxRepGoal(ctx *fiber.Ctx) error {
	var req createMaxRepGoalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.CreateMaxRepGoalParams{
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
		Notes:        req.Notes,
		GoalReps:     req.GoalReps,
	}
	goal, err := server.store.CreateMaxRepGoal(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(goal)
}

func (server *Server) getMaxRepGoal(ctx *fiber.Ctx) error {
	var req getMaxRepGoalRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.GetMaxRepGoalParams{
		GoalID:       req.GoalID,
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
	}

	goal, err := server.store.GetMaxRepGoal(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(goal)
}

func (server *Server) updateMaxRepGoal(ctx *fiber.Ctx) error {
	var req updateMaxRepGoalRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.UpdateMaxRepGoalParams{
		GoalID:       req.GoalID,
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
		GoalReps:     pgtype.Int4{Int32: req.GoalReps, Valid: req.GoalReps != 0},
		Notes:        pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
	}

	goal, err := server.store.UpdateMaxRepGoal(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(goal)
}

func (server *Server) listMaxRepGoals(ctx *fiber.Ctx) error {
	var req listMaxRepGoalsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	authPayload := ctx.Locals(middleware.AuthorizationPayloadKey).(*token.Payload)
	arg := db.ListMaxRepGoalsParams{
		Limit:        req.Limit,
		Offset:       req.Offset,
		Username:     authPayload.Username,
		ExerciseName: req.ExerciseName,
	}
	goals, err := server.store.ListMaxRepGoals(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(goals)
}
