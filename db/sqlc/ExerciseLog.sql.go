// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: ExerciseLog.sql

package db

import (
	"context"
)

const createExerciseLog = `-- name: CreateExerciseLog :one
INSERT INTO ExerciseLog (
  log_id,
  exercise_name, 
  sets_completed, 
  repetitions_completed,
  weight_lifted,
  notes
  )
VALUES ($1, $2, $3, $4,$5,$6)
RETURNING exercise_log_id, log_id, exercise_name, sets_completed, repetitions_completed, weight_lifted, notes, created_at
`

type CreateExerciseLogParams struct {
	LogID                int64  `json:"log_id"`
	ExerciseName         string `json:"exercise_name"`
	SetsCompleted        int32  `json:"sets_completed"`
	RepetitionsCompleted int32  `json:"repetitions_completed"`
	WeightLifted         int32  `json:"weight_lifted"`
	Notes                string `json:"notes"`
}

func (q *Queries) CreateExerciseLog(ctx context.Context, arg CreateExerciseLogParams) (Exerciselog, error) {
	row := q.db.QueryRow(ctx, createExerciseLog,
		arg.LogID,
		arg.ExerciseName,
		arg.SetsCompleted,
		arg.RepetitionsCompleted,
		arg.WeightLifted,
		arg.Notes,
	)
	var i Exerciselog
	err := row.Scan(
		&i.ExerciseLogID,
		&i.LogID,
		&i.ExerciseName,
		&i.SetsCompleted,
		&i.RepetitionsCompleted,
		&i.WeightLifted,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const deleteExerciseLog = `-- name: DeleteExerciseLog :exec
DELETE FROM ExerciseLog
WHERE exercise_log_id = $1
`

func (q *Queries) DeleteExerciseLog(ctx context.Context, exerciseLogID int64) error {
	_, err := q.db.Exec(ctx, deleteExerciseLog, exerciseLogID)
	return err
}

const getExerciseLog = `-- name: GetExerciseLog :one
SELECT exercise_log_id, log_id, exercise_name, sets_completed, repetitions_completed, weight_lifted, notes, created_at
FROM ExerciseLog
WHERE exercise_log_id = $1
`

func (q *Queries) GetExerciseLog(ctx context.Context, exerciseLogID int64) (Exerciselog, error) {
	row := q.db.QueryRow(ctx, getExerciseLog, exerciseLogID)
	var i Exerciselog
	err := row.Scan(
		&i.ExerciseLogID,
		&i.LogID,
		&i.ExerciseName,
		&i.SetsCompleted,
		&i.RepetitionsCompleted,
		&i.WeightLifted,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}

const listExerciseLog = `-- name: ListExerciseLog :many
SELECT exercise_log_id, log_id, exercise_name, sets_completed, repetitions_completed, weight_lifted, notes, created_at
FROM ExerciseLog
WHERE log_id = $1
ORDER BY exercise_log_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListExerciseLogParams struct {
	LogID  int64 `json:"log_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListExerciseLog(ctx context.Context, arg ListExerciseLogParams) ([]Exerciselog, error) {
	rows, err := q.db.Query(ctx, listExerciseLog, arg.LogID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Exerciselog{}
	for rows.Next() {
		var i Exerciselog
		if err := rows.Scan(
			&i.ExerciseLogID,
			&i.LogID,
			&i.ExerciseName,
			&i.SetsCompleted,
			&i.RepetitionsCompleted,
			&i.WeightLifted,
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

const updateExerciseLog = `-- name: UpdateExerciseLog :one
UPDATE ExerciseLog
SET sets_completed = $2, repetitions_completed = $3, weight_lifted = $4, notes = $5
WHERE exercise_log_id = $1
RETURNING exercise_log_id, log_id, exercise_name, sets_completed, repetitions_completed, weight_lifted, notes, created_at
`

type UpdateExerciseLogParams struct {
	ExerciseLogID        int64  `json:"exercise_log_id"`
	SetsCompleted        int32  `json:"sets_completed"`
	RepetitionsCompleted int32  `json:"repetitions_completed"`
	WeightLifted         int32  `json:"weight_lifted"`
	Notes                string `json:"notes"`
}

func (q *Queries) UpdateExerciseLog(ctx context.Context, arg UpdateExerciseLogParams) (Exerciselog, error) {
	row := q.db.QueryRow(ctx, updateExerciseLog,
		arg.ExerciseLogID,
		arg.SetsCompleted,
		arg.RepetitionsCompleted,
		arg.WeightLifted,
		arg.Notes,
	)
	var i Exerciselog
	err := row.Scan(
		&i.ExerciseLogID,
		&i.LogID,
		&i.ExerciseName,
		&i.SetsCompleted,
		&i.RepetitionsCompleted,
		&i.WeightLifted,
		&i.Notes,
		&i.CreatedAt,
	)
	return i, err
}
