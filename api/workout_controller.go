package api

import (
	"log"
	"net/http"
	"time"

	db "github.com/DMonkey83/FiberFitnessApp/db/sqlc"
	val "github.com/DMonkey83/FiberFitnessApp/util/Validate"
	res "github.com/DMonkey83/FiberFitnessApp/util/response"
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgtype"
)

type (
	createWorkoutRequest struct {
		Username            string           `json:"username" validate:"required"`
		WorkoutDate         time.Time        `json:"workout_date" validate:"required"`
		WorkoutDuration     string           `json:"workout_duration" validate:"required"`
		Notes               string           `json:"notes" validate:"required"`
		FatigueLevel        val.Fatiguelevel `json:"fatigue_level" validate:"required"`
		TotalCaloriesBurned int32            `json:"total_calories_burned" validate:"required"`
		TotalDistance       int32            `json:"total_distance" validate:"required"`
		TotalRepetitions    int32            `json:"total_repetitions" validate:"required"`
		TotalSets           int32            `json:"total_sets" validate:"required"`
		TotalWeightLifted   int32            `json:"total_weight_lifted" validate:"required"`
	}

	getWorkoutRequest struct {
		WorkoutID int64 `uri:"workout_id" binding:"required,min=1"`
	}

	updateWorkoutRequest struct {
		WokroutID           int64            `json:"workout_id" validate:"required"`
		WorkoutDate         time.Time        `json:"workout_date"`
		WorkoutDuration     string           `json:"workout_duration"`
		Notes               string           `json:"notes"`
		FatigueLevel        val.Fatiguelevel `json:"fatigue_level"`
		TotalCaloriesBurned int32            `json:"total_calories_burned"`
		TotalDistance       int32            `json:"total_distance"`
		TotalRepetitions    int32            `json:"total_repetitions"`
		TotalSets           int32            `json:"total_sets"`
		TotalWeightLifted   int32            `json:"total_weight_lifted"`
	}

	listWorkoutRequest struct {
		Limit    int32  `form:"limit" validate:"required,min=5,max=20"`
		Offset   int32  `form:"offset" validate:"required,min=1"`
		Username string `form:"username" validate:"required"`
	}
)

func (server *Server) createWorkout(ctx *fiber.Ctx) error {
	var req createWorkoutRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateWorkoutParams{
		Username:            req.Username,
		WorkoutDate:         req.WorkoutDate,
		Notes:               req.Notes,
		TotalSets:           req.TotalSets,
		TotalDistance:       req.TotalDistance,
		TotalRepetitions:    req.TotalRepetitions,
		TotalWeightLifted:   req.TotalWeightLifted,
		TotalCaloriesBurned: req.TotalCaloriesBurned,
		WorkoutDuration:     req.WorkoutDuration,
		FatigueLevel:        db.Fatiguelevel(req.FatigueLevel),
	}
	workout, err := server.store.CreateWorkout(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(workout)
}

func (server *Server) getWorkout(ctx *fiber.Ctx) error {
	var req getWorkoutRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	workout, err := server.store.GetWorkout(ctx.Context(), req.WorkoutID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(workout)
}

func (server *Server) updateWorkout(ctx *fiber.Ctx) error {
	var req updateWorkoutRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateWorkoutParams{
		WorkoutID:           req.WokroutID,
		WorkoutDate:         pgtype.Timestamptz{Time: req.WorkoutDate, Valid: req.WorkoutDate != time.Time{}},
		WorkoutDuration:     pgtype.Text{String: req.WorkoutDuration, Valid: req.WorkoutDuration != ""},
		Notes:               pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
		FatigueLevel:        db.NullFatiguelevel{Fatiguelevel: db.Fatiguelevel(req.FatigueLevel), Valid: req.FatigueLevel != ""},
		TotalCaloriesBurned: pgtype.Int4{Int32: req.TotalCaloriesBurned, Valid: req.TotalCaloriesBurned != 0},
		TotalDistance:       pgtype.Int4{Int32: req.TotalDistance, Valid: req.TotalDistance != 0},
		TotalRepetitions:    pgtype.Int4{Int32: req.TotalRepetitions, Valid: req.TotalRepetitions != 0},
		TotalSets:           pgtype.Int4{Int32: req.TotalSets, Valid: req.TotalSets != 0},
		TotalWeightLifted:   pgtype.Int4{Int32: req.TotalWeightLifted, Valid: req.TotalWeightLifted != 0},
	}
	workout, err := server.store.UpdateWorkout(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(workout)
}

func (server *Server) listWorkout(ctx *fiber.Ctx) error {
	var req listWorkoutRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListWorkoutsParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		Username: req.Username,
	}
	workout, err := server.store.ListWorkouts(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(workout)
}
