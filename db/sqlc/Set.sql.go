// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: Set.sql

package db

import (
	"context"
)

const createSet = `-- name: CreateSet :one
INSERT INTO Set (exercise_name, set_number, weight, rest_duration, notes, reps_completed)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING set_id, exercise_name, set_number, weight, rest_duration, reps_completed, notes, created_at
`

type CreateSetParams struct {
	ExerciseName  string `json:"exercise_name"`
	SetNumber     int32  `json:"set_number"`
	Weight        int32  `json:"weight"`
	RestDuration  string `json:"rest_duration"`
	Notes         string `json:"notes"`
	RepsCompleted int32  `json:"reps_completed"`
}

func (q *Queries) CreateSet(ctx context.Context, arg CreateSetParams) (Set, error) {
	row := q.db.QueryRow(ctx, createSet,
		arg.ExerciseName,
		arg.SetNumber,
		arg.Weight,
		arg.RestDuration,
		arg.Notes,
		arg.RepsCompleted,
	)
	var i Set
	err := row.Scan(
		&i.SetID,
		&i.ExerciseName,
		&i.SetNumber,
		&i.Weight,
		&i.RestDuration,
		&i.RepsCompleted,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const deleteSet = `-- name: DeleteSet :exec
DELETE FROM Set
WHERE set_id = $1
`

func (q *Queries) DeleteSet(ctx context.Context, setID int64) error {
	_, err := q.db.Exec(ctx, deleteSet, setID)
	return err
}

const getSet = `-- name: GetSet :one
SELECT set_id, exercise_name, set_number, weight, rest_duration, reps_completed, notes, created_at
FROM Set
WHERE set_id = $1
`

func (q *Queries) GetSet(ctx context.Context, setID int64) (Set, error) {
	row := q.db.QueryRow(ctx, getSet, setID)
	var i Set
	err := row.Scan(
		&i.SetID,
		&i.ExerciseName,
		&i.SetNumber,
		&i.Weight,
		&i.RestDuration,
		&i.RepsCompleted,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const listSets = `-- name: ListSets :many
SELECT set_id, exercise_name, set_number, weight, rest_duration, reps_completed, notes, created_at
FROM Set
WHERE exercise_name = $1
ORDER BY set_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListSetsParams struct {
	ExerciseName string `json:"exercise_name"`
	Limit        int32  `json:"limit"`
	Offset       int32  `json:"offset"`
}

func (q *Queries) ListSets(ctx context.Context, arg ListSetsParams) ([]Set, error) {
	rows, err := q.db.Query(ctx, listSets, arg.ExerciseName, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Set{}
	for rows.Next() {
		var i Set
		if err := rows.Scan(
			&i.SetID,
			&i.ExerciseName,
			&i.SetNumber,
			&i.Weight,
			&i.RestDuration,
			&i.RepsCompleted,
			&i.Notes,
			&i.CreatedAt,
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

const updateSet = `-- name: UpdateSet :one
UPDATE Set
SET set_number = $2, weight = $3, rest_duration = $4, notes = $5
WHERE set_id = $1
RETURNING set_id, exercise_name, set_number, weight, rest_duration, reps_completed, notes, created_at
`

type UpdateSetParams struct {
	SetID        int64  `json:"set_id"`
	SetNumber    int32  `json:"set_number"`
	Weight       int32  `json:"weight"`
	RestDuration string `json:"rest_duration"`
	Notes        string `json:"notes"`
}

func (q *Queries) UpdateSet(ctx context.Context, arg UpdateSetParams) (Set, error) {
	row := q.db.QueryRow(ctx, updateSet,
		arg.SetID,
		arg.SetNumber,
		arg.Weight,
		arg.RestDuration,
		arg.Notes,
	)
	var i Set
	err := row.Scan(
		&i.SetID,
		&i.ExerciseName,
		&i.SetNumber,
		&i.Weight,
		&i.RestDuration,
		&i.RepsCompleted,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}