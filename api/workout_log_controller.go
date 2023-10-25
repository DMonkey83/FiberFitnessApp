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
	createWorkoutLogRequest struct {
		Username            string           `json:"username" validate:"required"`
		PlanID              int64            `json:"plan_id" validate:"required"`
		LogDate             time.Time        `json:"log_date" validate:"required"`
		OverallFeeling      string           `json:"overall_feeling" validate:"required"`
		WorkoutDuration     string           `json:"workout_duration" validate:"required"`
		FatigueLevel        val.Fatiguelevel `json:"fatigue_level" validate:"required"`
		Rating              val.Rating       `json:"rating" validate:"required"`
		Comments            string           `json:"comments" validate:"required"`
		TotalCaloriesBurned int32            `json:"total_calories_burned" validate:"required"`
		TotalDistance       int32            `json:"total_distance" validate:"required"`
		TotalRepetitions    int32            `json:"total_repetitions" validate:"required"`
		TotalSets           int32            `json:"total_sets" validate:"required"`
		TotalWeightLifted   int32            `json:"total_weight_lifted" validate:"required"`
	}

	getWorkoutLogRequest struct {
		LogID int64 `uri:"log_id" binding:"required,min=1"`
	}

	updateWorkoutLogRequest struct {
		LogID               int64            `json:"log_id" validate:"required"`
		LogDate             time.Time        `json:"log_date"`
		WorkoutDuration     string           `json:"workout_duration"`
		Comments            string           `json:"comments"`
		OverallFeeling      string           `json:"overall_feeling"`
		FatigueLevel        val.Fatiguelevel `json:"fatigue_level"`
		Rating              val.Rating       `json:"rating"`
		TotalCaloriesBurned int32            `json:"total_calories_burned"`
		TotalDistance       int32            `json:"total_distance"`
		TotalRepetitions    int32            `json:"total_repetitions"`
		TotalSets           int32            `json:"total_sets"`
		TotalWeightLifted   int32            `json:"total_weight_lifted"`
	}

	listWorkoutLogRequest struct {
		Limit  int32 `form:"limit" validate:"required,min=5,max=20"`
		Offset int32 `form:"offset" validate:"required,min=1"`
		PlanID int64 `form:"plan_id"`
	}
)

func (server *Server) createWorkoutLog(ctx *fiber.Ctx) error {
	var req createWorkoutLogRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateWorkoutLogParams{
		Username:            req.Username,
		TotalSets:           req.TotalSets,
		TotalDistance:       req.TotalDistance,
		TotalRepetitions:    req.TotalRepetitions,
		TotalWeightLifted:   req.TotalWeightLifted,
		TotalCaloriesBurned: req.TotalCaloriesBurned,
		WorkoutDuration:     req.WorkoutDuration,
		FatigueLevel:        db.Fatiguelevel(req.FatigueLevel),
		OverallFeeling:      req.OverallFeeling,
		Comments:            req.Comments,
		Rating:              db.Rating(req.Rating),
		LogDate:             req.LogDate,
		PlanID:              req.PlanID,
	}
	log, err := server.store.CreateWorkoutLog(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(log)
}

func (server *Server) getWorkoutLog(ctx *fiber.Ctx) error {
	var req getWorkoutLogRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	log, err := server.store.GetWorkoutLog(ctx.Context(), req.LogID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(log)
}

func (server *Server) updateWorkoutLog(ctx *fiber.Ctx) error {
	var req updateWorkoutLogRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateWorkoutLogParams{
		LogID:               req.LogID,
		Comments:            pgtype.Text{String: req.Comments, Valid: req.Comments != ""},
		OverallFeeling:      pgtype.Text{String: req.OverallFeeling, Valid: req.OverallFeeling != ""},
		Rating:              db.NullRating{Rating: db.Rating(req.Rating), Valid: req.Rating != ""},
		LogDate:             pgtype.Timestamptz{Time: req.LogDate, Valid: req.LogDate != time.Time{}},
		WorkoutDuration:     pgtype.Text{String: req.WorkoutDuration, Valid: req.WorkoutDuration != ""},
		FatigueLevel:        db.NullFatiguelevel{Fatiguelevel: db.Fatiguelevel(req.FatigueLevel), Valid: req.FatigueLevel != ""},
		TotalCaloriesBurned: pgtype.Int4{Int32: req.TotalCaloriesBurned, Valid: req.TotalCaloriesBurned != 0},
		TotalDistance:       pgtype.Int4{Int32: req.TotalDistance, Valid: req.TotalDistance != 0},
		TotalRepetitions:    pgtype.Int4{Int32: req.TotalRepetitions, Valid: req.TotalRepetitions != 0},
		TotalSets:           pgtype.Int4{Int32: req.TotalSets, Valid: req.TotalSets != 0},
		TotalWeightLifted:   pgtype.Int4{Int32: req.TotalWeightLifted, Valid: req.TotalWeightLifted != 0},
	}
	log, err := server.store.UpdateWorkoutLog(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(log)
}

func (server *Server) listWorkoutLogs(ctx *fiber.Ctx) error {
	var req listWorkoutLogRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListWorkoutLogsParams{
		Limit:  req.Limit,
		Offset: req.Offset,
		PlanID: req.PlanID,
	}
	log, err := server.store.ListWorkoutLogs(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(log)
}
