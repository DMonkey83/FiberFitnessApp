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
	createAvailablePlanExerciseRequest struct {
		ExerciseName string `json:"exercise_name" validate:"required"`
		PlanID       int64  `json:"plan_id" validate:"required"`
		Sets         int32  `json:"sets" validate:"required"`
		RestDuration string `json:"rest_duration" validate:"required"`
		Notes        string `json:"notes" validate:"required"`
	}

	getAvailablePlanExerciseRequest struct {
		ID int64 `uri:"id" binding:"required,min=1"`
	}

	updateAvailablePlanExerciseRequest struct {
		ID           int64  `json:"id" validate:"required,min=1"`
		ExerciseName string `json:"exercise_name"`
		Sets         int32  `json:"sets"`
		Notes        string `json:"notes"`
		RestDuration string `json:"rest_duration"`
	}

	listAllAvailablePlansExercisesRequest struct {
		Limit       int32  `form:"limit" validate:"required,min=1"`
		Offset      int32  `form:"offset" validate:"required"`
		ExercseName string `form:"exercise_name" validate:"required"`
	}
)

func (server *Server) createAvailablePlanExercise(ctx *fiber.Ctx) error {
	var req createAvailablePlanExerciseRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateAvailablePlanExerciseParams{}
	availablePlanExercise, err := server.store.CreateAvailablePlanExercise(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(availablePlanExercise)
}

func (server *Server) getAvailablePlanExercise(ctx *fiber.Ctx) error {
	var req getAvailablePlanExerciseRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	availablePlanExercise, err := server.store.GetAvailablePlanExercise(ctx.Context(), req.ID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(availablePlanExercise)
}

func (server *Server) updateAvailablePlanExercise(ctx *fiber.Ctx) error {
	var req updateAvailablePlanExerciseRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateAvailablePlanExerciseParams{
		Notes:        pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
		Sets:         pgtype.Int4{Int32: req.Sets, Valid: req.Sets != 0},
		RestDuration: pgtype.Text{String: req.RestDuration, Valid: req.RestDuration != ""},
		ID:           req.ID,
	}

	availablePlanExercise, err := server.store.UpdateAvailablePlanExercise(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(availablePlanExercise)
}

func (server *Server) listAllAvailablePlansExercises(ctx *fiber.Ctx) error {
	var req listAllAvailablePlansExercisesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListAllAvailablePlanExercisesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}
	availablePlanExercises, err := server.store.ListAllAvailablePlanExercises(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(availablePlanExercises)
}
