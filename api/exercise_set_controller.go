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
	createExerciseSetRequest struct {
		ExerciseLogID        int64 `json:"exercise_log_id" validate:"required"`
		SetNumber            int32 `json:"set_number" validate:"required"`
		RepetitionsCompleted int32 `json:"repetitions_completed" validate:"required"`
		WeightLifted         int32 `json:"weights_lifted" validate:"required"`
	}

	getExerciseSetRequest struct {
		SetID int64 `uri:"set_id" binding:"required,min=1"`
	}

	updateExerciseSetRequest struct {
		SetID                int64 `json:"set_id" binding:"required,min=1"`
		SetNumber            int32 `json:"set_number"`
		RepetitionsCompleted int32 `json:"repetitions_completed"`
		WeightLifted         int32 `json:"weights_lifted"`
	}

	listAllExerciseSetRequest struct {
		ExerciseLogID int64 `form:"exercise_log_id" binding:"required,min=1"`
		Limit         int32 `form:"limit" validate:"required,min=1"`
		Offset        int32 `form:"offset" validate:"required"`
	}
)

func (server *Server) createExerciseSet(ctx *fiber.Ctx) error {
	var req createExerciseSetRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateExerciseSetParams{
		ExerciseLogID:        req.ExerciseLogID,
		SetNumber:            req.SetNumber,
		WeightLifted:         req.WeightLifted,
		RepetitionsCompleted: req.RepetitionsCompleted,
	}
	exerciseSet, err := server.store.CreateExerciseSet(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exerciseSet)
}

func (server *Server) getExerciseSet(ctx *fiber.Ctx) error {
	var req getExerciseSetRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	exerciseSet, err := server.store.GetExerciseSet(ctx.Context(), req.SetID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exerciseSet)
}

func (server *Server) updateExerciseSet(ctx *fiber.Ctx) error {
	var req updateExerciseSetRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateExerciseSetParams{
		SetID:                req.SetID,
		WeightLifted:         pgtype.Int4{Int32: req.WeightLifted, Valid: req.WeightLifted != 0},
		RepetitionsCompleted: pgtype.Int4{Int32: req.RepetitionsCompleted, Valid: req.RepetitionsCompleted != 0},
		SetNumber:            pgtype.Int4{Int32: req.SetNumber, Valid: req.SetNumber != 0},
	}

	exerciseSet, err := server.store.UpdateExerciseSet(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exerciseSet)
}

func (server *Server) listAllExerciseSets(ctx *fiber.Ctx) error {
	var req listAllExerciseSetRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListExerciseSetsParams{
		Limit:         req.Limit,
		Offset:        req.Offset,
		ExerciseLogID: req.ExerciseLogID,
	}
	exercisesSet, err := server.store.ListExerciseSets(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercisesSet)
}
