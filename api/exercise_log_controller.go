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
	createExerciseLogRequest struct {
		LogID                int64  `json:"log_id" validate:"required"`
		ExerciseName         string `json:"exercise_name" validate:"required"`
		Notes                string `json:"notes" validate:"required"`
		SetsCompleted        int32  `json:"sets_completed" validate:"required"`
		RepetitionsCompleted int32  `json:"repetitions_completed" validate:"required"`
		WeightLifted         int32  `json:"weights_lifted" validate:"required"`
	}

	getExerciseLogRequest struct {
		ExerciseLogID int64 `uri:"exercise_log_id" binding:"required,min=1"`
	}

	updateExerciseLogRequest struct {
		ExerciseLogID        int64  `json:"exercise_log_id" binding:"required,min=1"`
		Notes                string `json:"notes"`
		SetsCompleted        int32  `json:"sets_completed"`
		RepetitionsCompleted int32  `json:"repetitions_completed"`
		WeightLifted         int32  `json:"weights_lifted"`
	}

	listAllExerciseLogsRequest struct {
		LogID  int64 `form:"exercise_log_id" binding:"required,min=1"`
		Limit  int32 `form:"limit" validate:"required,min=1"`
		Offset int32 `form:"offset" validate:"required"`
	}
)

func (server *Server) createExerciseLog(ctx *fiber.Ctx) error {
	var req createExerciseLogRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateExerciseLogParams{
		ExerciseName:         req.ExerciseName,
		LogID:                req.LogID,
		Notes:                req.Notes,
		SetsCompleted:        req.SetsCompleted,
		WeightLifted:         req.WeightLifted,
		RepetitionsCompleted: req.RepetitionsCompleted,
	}
	exerciseLog, err := server.store.CreateExerciseLog(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exerciseLog)
}

func (server *Server) getExerciseLog(ctx *fiber.Ctx) error {
	var req getExerciseLogRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	exerciseLog, err := server.store.GetExerciseLog(ctx.Context(), req.ExerciseLogID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exerciseLog)
}

func (server *Server) updateExerciseLog(ctx *fiber.Ctx) error {
	var req updateExerciseLogRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateExerciseLogParams{
		SetsCompleted:        pgtype.Int4{Int32: req.SetsCompleted, Valid: req.SetsCompleted != 0},
		Notes:                pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
		WeightLifted:         pgtype.Int4{Int32: req.WeightLifted, Valid: req.WeightLifted != 0},
		ExerciseLogID:        req.ExerciseLogID,
		RepetitionsCompleted: pgtype.Int4{Int32: req.RepetitionsCompleted, Valid: req.RepetitionsCompleted != 0},
	}

	exerciseLog, err := server.store.UpdateExerciseLog(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exerciseLog)
}

func (server *Server) listAllExerciseLogs(ctx *fiber.Ctx) error {
	var req listAllExerciseLogsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListExerciseLogParams{
		Limit:  req.Limit,
		Offset: req.Offset,
		LogID:  req.LogID,
	}
	exercisesLog, err := server.store.ListExerciseLog(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercisesLog)
}
