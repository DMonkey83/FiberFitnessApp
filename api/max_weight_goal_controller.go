package api

import (
	"log"
	"net/http"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	val "github.com/DMonkey83/FiberFitnessApp/util/Validate"
	res "github.com/DMonkey83/FiberFitnessApp/util/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	createMaxWeightGoalRequest struct {
		Username     string `json:"username" validate:"required"`
		ExerciseName string `json:"exercise_name" validate:"required"`
		GoalWeight   int32  `json:"goal_weight" validate:"required"`
		Notes        string `json:"notes" validate:"required"`
	}

	getMaxWeightGoalRequest struct {
		GoalID       int64  `uri:"goal_id" binding:"required,min=1"`
		Username     string `uri:"username" binding:"required"`
		ExerciseName string `uri:"exercise_name" binding:"required"`
	}

	updateMaxWeightGoalRequest struct {
		GoalID       int64  `json:"goal_id" binding:"required,min=1"`
		Username     string `json:"username" validate:"required"`
		ExerciseName string `json:"exercise_name" validate:"required"`
		GoalWeight   int32  `json:"goal_weight" validate:"required"`
		Notes        string `json:"notes" validate:"required"`
	}

	listMaxWeightGoalsRequest struct {
		ExerciseName string `form:"exercise_name" binding:"required,min=1"`
		Username     string `form:"username" binding:"required,min=1"`
		Limit        int32  `form:"limit" validate:"required,min=1"`
		Offset       int32  `form:"offset" validate:"required"`
	}
)

func (server *Server) createMaxWeightGoal(ctx *fiber.Ctx) error {
	var req createMaxWeightGoalRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateMaxWeightGoalParams{
		Username:     req.Username,
		ExerciseName: req.ExerciseName,
		Notes:        req.Notes,
		GoalWeight:   req.GoalWeight,
	}
	goal, err := server.store.CreateMaxWeightGoal(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(goal)
}

func (server *Server) getMaxWeightGoal(ctx *fiber.Ctx) error {
	var req getMaxWeightGoalRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.GetMaxWeightGoalParams{
		GoalID:       req.GoalID,
		Username:     req.Username,
		ExerciseName: req.ExerciseName,
	}

	goal, err := server.store.GetMaxWeightGoal(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(goal)
}

func (server *Server) updateMaxWeightGoal(ctx *fiber.Ctx) error {
	var req updateMaxWeightGoalRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateMaxWeightGoalParams{
		GoalID:       req.GoalID,
		Username:     req.Username,
		ExerciseName: req.ExerciseName,
		GoalWeight:   pgtype.Int4{Int32: req.GoalWeight, Valid: req.GoalWeight != 0},
		Notes:        pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
	}

	goal, err := server.store.UpdateMaxWeightGoal(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(goal)
}

func (server *Server) listMaxWeightGoals(ctx *fiber.Ctx) error {
	var req listMaxWeightGoalsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListMaxWeightGoalsParams{
		Limit:        req.Limit,
		Offset:       req.Offset,
		Username:     req.Username,
		ExerciseName: req.ExerciseName,
	}
	goals, err := server.store.ListMaxWeightGoals(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(goals)
}
