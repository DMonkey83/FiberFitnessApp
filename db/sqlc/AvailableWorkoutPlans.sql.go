// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: AvailableWorkoutPlans.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createAvailablePlan = `-- name: CreateAvailablePlan :one
INSERT INTO AvailableWorkoutPlans (
  plan_name, 
  description, 
  goal, 
  difficulty,
  is_public,
  creator_username
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING plan_id, plan_name, description, goal, difficulty, is_public, created_at, updated_at, creator_username
`

type CreateAvailablePlanParams struct {
	PlanName        string          `json:"plan_name"`
	Description     string          `json:"description"`
	Goal            Workoutgoalenum `json:"goal"`
	Difficulty      Difficulty      `json:"difficulty"`
	IsPublic        Visibility      `json:"is_public"`
	CreatorUsername string          `json:"creator_username"`
}

func (q *Queries) CreateAvailablePlan(ctx context.Context, arg CreateAvailablePlanParams) (Availableworkoutplan, error) {
	row := q.db.QueryRow(ctx, createAvailablePlan,
		arg.PlanName,
		arg.Description,
		arg.Goal,
		arg.Difficulty,
		arg.IsPublic,
		arg.CreatorUsername,
	)
	var i Availableworkoutplan
	err := row.Scan(
		&i.PlanID,
		&i.PlanName,
		&i.Description,
		&i.Goal,
		&i.Difficulty,
		&i.IsPublic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatorUsername,
	)
	return i, err
}

const deleteAvailablePlan = `-- name: DeleteAvailablePlan :exec
DELETE FROM AvailableWorkoutPlans
WHERE plan_id = $1
`

func (q *Queries) DeleteAvailablePlan(ctx context.Context, planID int64) error {
	_, err := q.db.Exec(ctx, deleteAvailablePlan, planID)
	return err
}

const getAvailablePlan = `-- name: GetAvailablePlan :one
SELECT plan_id, plan_name, description, goal, difficulty, is_public, created_at, updated_at, creator_username
FROM AvailableWorkoutPlans
WHERE plan_id = $1
`

func (q *Queries) GetAvailablePlan(ctx context.Context, planID int64) (Availableworkoutplan, error) {
	row := q.db.QueryRow(ctx, getAvailablePlan, planID)
	var i Availableworkoutplan
	err := row.Scan(
		&i.PlanID,
		&i.PlanName,
		&i.Description,
		&i.Goal,
		&i.Difficulty,
		&i.IsPublic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatorUsername,
	)
	return i, err
}

const listAllAvailablePlans = `-- name: ListAllAvailablePlans :many
SELECT plan_id, plan_name, description, goal, difficulty, is_public, created_at, updated_at, creator_username
FROM AvailableWorkoutPlans
ORDER BY plan_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2
`

type ListAllAvailablePlansParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListAllAvailablePlans(ctx context.Context, arg ListAllAvailablePlansParams) ([]Availableworkoutplan, error) {
	rows, err := q.db.Query(ctx, listAllAvailablePlans, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Availableworkoutplan{}
	for rows.Next() {
		var i Availableworkoutplan
		if err := rows.Scan(
			&i.PlanID,
			&i.PlanName,
			&i.Description,
			&i.Goal,
			&i.Difficulty,
			&i.IsPublic,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatorUsername,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listAvailablePlansByCreator = `-- name: ListAvailablePlansByCreator :many
SELECT plan_id, plan_name, description, goal, difficulty, is_public, created_at, updated_at, creator_username
FROM AvailableWorkoutPlans
WHERE creator_username =$1
ORDER BY plan_name DESC -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListAvailablePlansByCreatorParams struct {
	CreatorUsername string `json:"creator_username"`
	Limit           int32  `json:"limit"`
	Offset          int32  `json:"offset"`
}

func (q *Queries) ListAvailablePlansByCreator(ctx context.Context, arg ListAvailablePlansByCreatorParams) ([]Availableworkoutplan, error) {
	rows, err := q.db.Query(ctx, listAvailablePlansByCreator, arg.CreatorUsername, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Availableworkoutplan{}
	for rows.Next() {
		var i Availableworkoutplan
		if err := rows.Scan(
			&i.PlanID,
			&i.PlanName,
			&i.Description,
			&i.Goal,
			&i.Difficulty,
			&i.IsPublic,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.CreatorUsername,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateAvailablePlan = `-- name: UpdateAvailablePlan :one
UPDATE AvailableWorkoutPlans
SET 
description = COALESCE($1,description),
plan_name = COALESCE($2,plan_name),
goal = COALESCE($3,goal),
difficulty = COALESCE($4,difficulty),
is_public = COALESCE($5,is_public)
WHERE creator_username = $6
RETURNING plan_id, plan_name, description, goal, difficulty, is_public, created_at, updated_at, creator_username
`

type UpdateAvailablePlanParams struct {
	Description     pgtype.Text         `json:"description"`
	PlanName        pgtype.Text         `json:"plan_name"`
	Goal            NullWorkoutgoalenum `json:"goal"`
	Difficulty      NullDifficulty      `json:"difficulty"`
	IsPublic        NullVisibility      `json:"is_public"`
	CreatorUsername string              `json:"creator_username"`
}

func (q *Queries) UpdateAvailablePlan(ctx context.Context, arg UpdateAvailablePlanParams) (Availableworkoutplan, error) {
	row := q.db.QueryRow(ctx, updateAvailablePlan,
		arg.Description,
		arg.PlanName,
		arg.Goal,
		arg.Difficulty,
		arg.IsPublic,
		arg.CreatorUsername,
	)
	var i Availableworkoutplan
	err := row.Scan(
		&i.PlanID,
		&i.PlanName,
		&i.Description,
		&i.Goal,
		&i.Difficulty,
		&i.IsPublic,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.CreatorUsername,
	)
	return i, err
}
