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
	createOneOfExerciseRequest struct {
		WorkoutID       int64               `json:"workout_id" validate:"required"`
		ExerciseName    string              `json:"exercise_name" validate:"required"`
		Description     string              `json:"description" validate:"required"`
		MuscleGroupName val.Musclegroupenum `json:"muscle_group_name" validate:"required"`
	}

	getOneOfExerciseRequest struct {
		WorkoutID int64 `uri:"workout_id" binding:"required,min=1"`
		ID        int32 `uri:"id" binding:"required,min=1"`
	}

	updateOneOfExerciseRequest struct {
		WorkoutID       int64               `json:"workout_id" validate:"required"`
		ID              int32               `json:"id" validate:"required"`
		Description     string              `json:"description"`
		MuscleGroupName val.Musclegroupenum `json:"muscle_group_name"`
	}

	listAllOneOfExercisesRequest struct {
		Limit     int32 `form:"limit" validate:"required,min=5,max=20"`
		Offset    int32 `form:"offset" validate:"required,min=1"`
		WorkoutID int64 `form:"workout_id"`
	}
)

func (server *Server) createOneOfExercise(ctx *fiber.Ctx) error {
	var req createOneOfExerciseRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateOneOffWorkoutExerciseParams{
		ExerciseName:    req.ExerciseName,
		WorkoutID:       req.WorkoutID,
		Description:     req.Description,
		MuscleGroupName: db.Musclegroupenum(req.MuscleGroupName),
	}
	exercise, err := server.store.CreateOneOffWorkoutExercise(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercise)
}

func (server *Server) getOneOfExercise(ctx *fiber.Ctx) error {
	var req getOneOfExerciseRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.GetOneOffWorkoutExerciseParams{
		WorkoutID: req.WorkoutID,
		ID:        req.ID,
	}

	exercise, err := server.store.GetOneOffWorkoutExercise(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exercise)
}

func (server *Server) updateOneOfExercise(ctx *fiber.Ctx) error {
	var req updateOneOfExerciseRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateOneOffWorkoutExerciseParams{
		WorkoutID:       req.WorkoutID,
		ID:              req.ID,
		Description:     pgtype.Text{String: req.Description, Valid: req.Description != ""},
		MuscleGroupName: db.NullMusclegroupenum{Musclegroupenum: db.Musclegroupenum(req.MuscleGroupName), Valid: req.MuscleGroupName != ""},
	}

	exercise, err := server.store.UpdateOneOffWorkoutExercise(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exercise)
}

func (server *Server) listAllOneOfExercises(ctx *fiber.Ctx) error {
	var req listAllOneOfExercisesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListAllOneOffWorkoutExercisesParams{
		Limit:     req.Limit,
		Offset:    req.Offset,
		WorkoutID: req.WorkoutID,
	}
	exercises, err := server.store.ListAllOneOffWorkoutExercises(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercises)
}
