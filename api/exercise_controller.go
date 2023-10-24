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
	createExerciseRequest struct {
		ExerciseName      string              `json:"exercise_name" validate:"required"`
		Description       string              `json:"description" validate:"required"`
		EquipmentRequired val.Equipmenttype   `json:"equipment_required" validate:"required"`
		MuscleGroupName   val.Musclegroupenum `json:"muscle_group_name" validate:"required"`
	}

	getExerciseRequest struct {
		ExerciseName string `uri:"exercise_name" binding:"required,min=1"`
	}

	updateExerciseRequest struct {
		ExerciseName      string              `json:"exercise_name" validate:"required"`
		Description       string              `json:"description"`
		EquipmentRequired val.Equipmenttype   `json:"equipment_required"`
		MuscleGroupName   val.Musclegroupenum `json:"muscle_group_name"`
	}

	listAllEquipmentExercisesRequest struct {
		Limit             int32             `form:"limit" validate:"required,min=1"`
		Offset            int32             `form:"offset" validate:"required"`
		ExerciseName      string            `form:"exercise_name" validate:"required"`
		EquipmentRequired val.Equipmenttype `form:"equipment_required" validate:"required"`
	}

	listAllMuscleGroupExercisesRequest struct {
		Limit           int32               `form:"limit" validate:"required,min=5,max=20"`
		Offset          int32               `form:"offset" validate:"required,min=1"`
		MuscreGroupName val.Musclegroupenum `form:"muscle_group_name" validate:"required"`
		ExerciseName    string              `form:"exercise_name" validate:"required"`
	}
)

func (server *Server) createExercise(ctx *fiber.Ctx) error {
	var req createExerciseRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateExerciseParams{
		ExerciseName:      req.ExerciseName,
		Description:       req.Description,
		EquipmentRequired: db.Equipmenttype(req.EquipmentRequired),
		MuscleGroupName:   db.Musclegroupenum(req.MuscleGroupName),
	}
	exercise, err := server.store.CreateExercise(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercise)
}

func (server *Server) getExercise(ctx *fiber.Ctx) error {
	var req getExerciseRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	exercise, err := server.store.GetExercise(ctx.Context(), req.ExerciseName)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exercise)
}

func (server *Server) updateExercise(ctx *fiber.Ctx) error {
	var req updateExerciseRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateExerciseParams{
		ExerciseName:      req.ExerciseName,
		Description:       pgtype.Text{String: req.Description, Valid: req.Description != ""},
		EquipmentRequired: db.NullEquipmenttype{Equipmenttype: db.Equipmenttype(req.EquipmentRequired), Valid: req.EquipmentRequired != ""},
		MuscleGroupName:   db.NullMusclegroupenum{Musclegroupenum: db.Musclegroupenum(req.MuscleGroupName), Valid: req.MuscleGroupName != ""},
	}

	exercise, err := server.store.UpdateExercise(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(exercise)
}

func (server *Server) listAllEquipmentExercises(ctx *fiber.Ctx) error {
	var req listAllEquipmentExercisesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListEquipmentExercisesParams{
		Limit:             req.Limit,
		Offset:            req.Offset,
		EquipmentRequired: db.Equipmenttype(req.EquipmentRequired),
	}
	exercises, err := server.store.ListEquipmentExercises(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercises)
}

func (server *Server) listAllMuscleGroupExercises(ctx *fiber.Ctx) error {
	var req listAllMuscleGroupExercisesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListMuscleGroupExercisesParams{
		Limit:           req.Limit,
		Offset:          req.Offset,
		MuscleGroupName: db.Musclegroupenum(req.MuscreGroupName),
	}
	exercises, err := server.store.ListMuscleGroupExercises(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(exercises)
}
