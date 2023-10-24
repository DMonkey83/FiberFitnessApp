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
	createSetRequest struct {
		ExerciseName  string `json:"exercise_name" validate:"required"`
		SetNmber      int32  `json:"set_number" validate:"required"`
		Weight        int32  `json:"weight" validate:"required"`
		Notes         string `json:"notes" validate:"required"`
		RepsCompleted int32  `json:"reps_completed" validate:"required"`
		RestDuration  string `json:"rest_duration" validate:"required"`
	}

	getSetRequest struct {
		SetID int64 `uri:"set_id" binding:"required,min=1"`
	}

	updateSetRequest struct {
		SetID        int64  `json:"set_id" validate:"required"`
		SetNumber    int32  `json:"set_number"`
		Weight       int32  `json:"weight"`
		Notes        string `json:"notes"`
		RestDuration string `json:"rest_duration"`
	}

	listSetsRequest struct {
		Limit        int32  `form:"limit" validate:"required,min=5,max=20"`
		Offset       int32  `form:"offset" validate:"required,min=1"`
		ExerciseName string `form:"exercise_name"`
	}
)

func (server *Server) createSet(ctx *fiber.Ctx) error {
	var req createSetRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateSetParams{
		ExerciseName:  req.ExerciseName,
		SetNumber:     req.SetNmber,
		Weight:        req.Weight,
		Notes:         req.Notes,
		RestDuration:  req.RestDuration,
		RepsCompleted: req.RepsCompleted,
	}
	set, err := server.store.CreateSet(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(set)
}

func (server *Server) getSet(ctx *fiber.Ctx) error {
	var req getSetRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	set, err := server.store.GetSet(ctx.Context(), req.SetID)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(set)
}

func (server *Server) updateSet(ctx *fiber.Ctx) error {
	var req updateSetRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateSetParams{
		SetID:        req.SetID,
		SetNumber:    pgtype.Int4{Int32: req.SetNumber, Valid: req.SetNumber != 0},
		Weight:       pgtype.Int4{Int32: req.Weight, Valid: req.Weight != 0},
		Notes:        pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
		RestDuration: pgtype.Text{String: req.RestDuration, Valid: req.RestDuration != ""},
	}
	set, err := server.store.UpdateSet(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(set)
}

func (server *Server) listSets(ctx *fiber.Ctx) error {
	var req listSetsRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListSetsParams{
		Limit:        req.Limit,
		Offset:       req.Offset,
		ExerciseName: req.ExerciseName,
	}
	set, err := server.store.ListSets(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(set)
}
