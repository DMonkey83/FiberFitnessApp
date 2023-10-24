// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: ExerciseSet.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createExerciseSet = `-- name: CreateExerciseSet :one
INSERT INTO ExerciseSet (exercise_log_id, set_number, weight_lifted, repetitions_completed)
VALUES ($1, $2, $3, $4)
RETURNING set_id, exercise_log_id, set_number, weight_lifted, repetitions_completed
`

type CreateExerciseSetParams struct {
	ExerciseLogID        int64 `json:"exercise_log_id"`
	SetNumber            int32 `json:"set_number"`
	WeightLifted         int32 `json:"weight_lifted"`
	RepetitionsCompleted int32 `json:"repetitions_completed"`
}

func (q *Queries) CreateExerciseSet(ctx context.Context, arg CreateExerciseSetParams) (Exerciseset, error) {
	row := q.db.QueryRow(ctx, createExerciseSet,
		arg.ExerciseLogID,
		arg.SetNumber,
		arg.WeightLifted,
		arg.RepetitionsCompleted,
	)
	var i Exerciseset
	err := row.Scan(
		&i.SetID,
		&i.ExerciseLogID,
		&i.SetNumber,
		&i.WeightLifted,
		&i.RepetitionsCompleted,
	)
	return i, err
}

const deleteExerciseSet = `-- name: DeleteExerciseSet :exec
DELETE FROM ExerciseSet
WHERE set_id = $1
`

func (q *Queries) DeleteExerciseSet(ctx context.Context, setID int64) error {
	_, err := q.db.Exec(ctx, deleteExerciseSet, setID)
	return err
}

const getExerciseSet = `-- name: GetExerciseSet :one
SELECT set_id, exercise_log_id, set_number, weight_lifted, repetitions_completed
FROM ExerciseSet
WHERE set_id = $1
`

func (q *Queries) GetExerciseSet(ctx context.Context, setID int64) (Exerciseset, error) {
	row := q.db.QueryRow(ctx, getExerciseSet, setID)
	var i Exerciseset
	err := row.Scan(
		&i.SetID,
		&i.ExerciseLogID,
		&i.SetNumber,
		&i.WeightLifted,
		&i.RepetitionsCompleted,
	)
	return i, err
}

const listExerciseSets = `-- name: ListExerciseSets :many
SELECT set_id, exercise_log_id, set_number, weight_lifted, repetitions_completed
FROM ExerciseSet
WHERE exercise_log_id = $1
ORDER BY set_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListExerciseSetsParams struct {
	ExerciseLogID int64 `json:"exercise_log_id"`
	Limit         int32 `json:"limit"`
	Offset        int32 `json:"offset"`
}

func (q *Queries) ListExerciseSets(ctx context.Context, arg ListExerciseSetsParams) ([]Exerciseset, error) {
	rows, err := q.db.Query(ctx, listExerciseSets, arg.ExerciseLogID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Exerciseset{}
	for rows.Next() {
		var i Exerciseset
		if err := rows.Scan(
			&i.SetID,
			&i.ExerciseLogID,
			&i.SetNumber,
			&i.WeightLifted,
			&i.RepetitionsCompleted,
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

const updateExerciseSet = `-- name: UpdateExerciseSet :one
UPDATE ExerciseSet
SET 
set_number = COALESCE($1,set_number),
weight_lifted = COALESCE($2,weight_lifted),
repetitions_completed = COALESCE($3,repetitions_completed)
WHERE set_id = $4
RETURNING set_id, exercise_log_id, set_number, weight_lifted, repetitions_completed
`

type UpdateExerciseSetParams struct {
	SetNumber            pgtype.Int4 `json:"set_number"`
	WeightLifted         pgtype.Int4 `json:"weight_lifted"`
	RepetitionsCompleted pgtype.Int4 `json:"repetitions_completed"`
	SetID                int64       `json:"set_id"`
}

func (q *Queries) UpdateExerciseSet(ctx context.Context, arg UpdateExerciseSetParams) (Exerciseset, error) {
	row := q.db.QueryRow(ctx, updateExerciseSet,
		arg.SetNumber,
		arg.WeightLifted,
		arg.RepetitionsCompleted,
		arg.SetID,
	)
	var i Exerciseset
	err := row.Scan(
		&i.SetID,
		&i.ExerciseLogID,
		&i.SetNumber,
		&i.WeightLifted,
		&i.RepetitionsCompleted,
	)
	return i, err
}
