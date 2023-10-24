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
	createWeightEntryRequest struct {
		Username  string    `json:"username" validate:"required"`
		EntryDate time.Time `json:"entry_date" validate:"required"`
		WeightKg  int32     `json:"weight_kg" validate:"required"`
		WeightLb  int32     `json:"weight_lb" validate:"required"`
		Notes     string    `json:"notes" validate:"required"`
	}

	getWeightEntryRequest struct {
		WeightEntryID int64  `uri:"weight_entry_id" binding:"required,min=1"`
		Username      string `uri:"username" binding:"required"`
	}

	updateWeightEntryRequest struct {
		Username      string    `json:"username" validate:"required"`
		WeightEntryID int64     `json:"weight_entry_id" validate:"required"`
		EntryDate     time.Time `json:"entry_date"`
		WeightKg      int32     `json:"weight_kg"`
		WeightLb      int32     `json:"weight_lb"`
		Notes         string    `json:"notes"`
	}

	listWeightEntriesRequest struct {
		Limit    int32  `form:"limit" validate:"required,min=5,max=20"`
		Offset   int32  `form:"offset" validate:"required,min=1"`
		Username string `form:"username" validate:"required"`
	}
)

func (server *Server) createWeightEntry(ctx *fiber.Ctx) error {
	var req createWeightEntryRequest
	if err := ctx.BodyParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	if err := val.ValidatePayload(ctx, &req); err != nil {
		log.Print(err, req)
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.CreateWeightEntryParams{
		Username:  req.Username,
		EntryDate: req.EntryDate,
		WeightKg:  req.WeightKg,
		WeightLb:  req.WeightLb,
		Notes:     req.Notes,
	}
	weight_entry, err := server.store.CreateWeightEntry(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(weight_entry)
}

func (server *Server) getWeightEntry(ctx *fiber.Ctx) error {
	var req getWeightEntryRequest
	if err := ctx.ParamsParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.GetWeightEntryParams{
		WeightEntryID: req.WeightEntryID,
		Username:      req.Username,
	}

	weight_entry, err := server.store.GetWeightEntry(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(weight_entry)
}

func (server *Server) updateWeightEntry(ctx *fiber.Ctx) error {
	var req updateWeightEntryRequest
	if err := val.ValidatePayload(ctx, &req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}

	arg := db.UpdateWeightEntryParams{
		EntryDate:     pgtype.Timestamptz{Time: req.EntryDate, Valid: req.EntryDate != time.Time{}},
		WeightKg:      pgtype.Int4{Int32: req.WeightKg, Valid: req.WeightKg != 0},
		WeightLb:      pgtype.Int4{Int32: req.WeightLb, Valid: req.WeightLb != 0},
		Notes:         pgtype.Text{String: req.Notes, Valid: req.Notes != ""},
		Username:      req.Username,
		WeightEntryID: req.WeightEntryID,
	}
	weight_entry, err := server.store.UpdateWeightEntry(ctx.Context(), arg)
	if err != nil {
		if err == db.ErrRecordNotFound {
			return res.ResponseNotFound(ctx, nil, err.Error())
		}
		return res.ResponseError(ctx, nil, err.Error())
	}

	return ctx.Status(http.StatusOK).JSON(weight_entry)
}

func (server *Server) listWeightEntries(ctx *fiber.Ctx) error {
	var req listWeightEntriesRequest
	if err := ctx.QueryParser(&req); err != nil {
		return res.ResponseValidationError(ctx, nil, err.Error())
	}
	arg := db.ListWeightEntriesParams{
		Limit:    req.Limit,
		Offset:   req.Offset,
		Username: req.Username,
	}
	weight_entry, err := server.store.ListWeightEntries(ctx.Context(), arg)
	if err != nil {
		return res.ResponseError(ctx, nil, err.Error())
	}
	return ctx.Status(http.StatusOK).JSON(weight_entry)
}
