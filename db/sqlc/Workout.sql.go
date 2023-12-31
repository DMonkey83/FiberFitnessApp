// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: Workout.sql

package db

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

const createWorkout = `-- name: CreateWorkout :one
INSERT INTO Workout 
  (
  username, 
  workout_date, 
  workout_duration, 
  notes, 
  fatigue_level,
  total_calories_burned,
  total_distance,
  total_repetitions,
  total_sets,
  total_weight_lifted
  )
VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10)
RETURNING workout_id, username, workout_date, workout_duration, fatigue_level, notes, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
`

type CreateWorkoutParams struct {
	Username            string       `json:"username"`
	WorkoutDate         time.Time    `json:"workout_date"`
	WorkoutDuration     string       `json:"workout_duration"`
	Notes               string       `json:"notes"`
	FatigueLevel        Fatiguelevel `json:"fatigue_level"`
	TotalCaloriesBurned int32        `json:"total_calories_burned"`
	TotalDistance       int32        `json:"total_distance"`
	TotalRepetitions    int32        `json:"total_repetitions"`
	TotalSets           int32        `json:"total_sets"`
	TotalWeightLifted   int32        `json:"total_weight_lifted"`
}

func (q *Queries) CreateWorkout(ctx context.Context, arg CreateWorkoutParams) (Workout, error) {
	row := q.db.QueryRow(ctx, createWorkout,
		arg.Username,
		arg.WorkoutDate,
		arg.WorkoutDuration,
		arg.Notes,
		arg.FatigueLevel,
		arg.TotalCaloriesBurned,
		arg.TotalDistance,
		arg.TotalRepetitions,
		arg.TotalSets,
		arg.TotalWeightLifted,
	)
	var i Workout
	err := row.Scan(
		&i.WorkoutID,
		&i.Username,
		&i.WorkoutDate,
		&i.WorkoutDuration,
		&i.FatigueLevel,
		&i.Notes,
		&i.TotalCaloriesBurned,
		&i.TotalDistance,
		&i.TotalRepetitions,
		&i.TotalSets,
		&i.TotalWeightLifted,
		&i.CreatedAt,
	)
	return i, err
}

const deleteWorkout = `-- name: DeleteWorkout :exec
DELETE FROM Workout
WHERE workout_id = $1
`

func (q *Queries) DeleteWorkout(ctx context.Context, workoutID int64) error {
	_, err := q.db.Exec(ctx, deleteWorkout, workoutID)
	return err
}

const getWorkout = `-- name: GetWorkout :one
SELECT workout_id, username, workout_date, workout_duration, fatigue_level, notes, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
FROM Workout
WHERE workout_id = $1
`

func (q *Queries) GetWorkout(ctx context.Context, workoutID int64) (Workout, error) {
	row := q.db.QueryRow(ctx, getWorkout, workoutID)
	var i Workout
	err := row.Scan(
		&i.WorkoutID,
		&i.Username,
		&i.WorkoutDate,
		&i.WorkoutDuration,
		&i.FatigueLevel,
		&i.Notes,
		&i.TotalCaloriesBurned,
		&i.TotalDistance,
		&i.TotalRepetitions,
		&i.TotalSets,
		&i.TotalWeightLifted,
		&i.CreatedAt,
	)
	return i, err
}

const listWorkouts = `-- name: ListWorkouts :many
SELECT workout_id, username, workout_date, workout_duration, fatigue_level, notes, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at FROM Workout
WHERE username = $1
ORDER BY workout_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListWorkoutsParams struct {
	Username string `json:"username"`
	Limit    int32  `json:"limit"`
	Offset   int32  `json:"offset"`
}

func (q *Queries) ListWorkouts(ctx context.Context, arg ListWorkoutsParams) ([]Workout, error) {
	rows, err := q.db.Query(ctx, listWorkouts, arg.Username, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Workout{}
	for rows.Next() {
		var i Workout
		if err := rows.Scan(
			&i.WorkoutID,
			&i.Username,
			&i.WorkoutDate,
			&i.WorkoutDuration,
			&i.FatigueLevel,
			&i.Notes,
			&i.TotalCaloriesBurned,
			&i.TotalDistance,
			&i.TotalRepetitions,
			&i.TotalSets,
			&i.TotalWeightLifted,
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

const updateWorkout = `-- name: UpdateWorkout :one
UPDATE Workout
SET 
workout_date = COALESCE($1,workout_date),
workout_duration = COALESCE($2,workout_duration),
notes = COALESCE($3,notes),
fatigue_level = COALESCE($4,fatigue_level),
total_sets = COALESCE($5,total_sets),
total_distance=COALESCE($6,total_distance),
total_repetitions=COALESCE($7,total_repetitions),
total_weight_lifted=COALESCE($8,total_weight_lifted),
total_calories_burned =COALESCE($9,total_calories_burned)
WHERE workout_id = $10
RETURNING workout_id, username, workout_date, workout_duration, fatigue_level, notes, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
`

type UpdateWorkoutParams struct {
	WorkoutDate         pgtype.Timestamptz `json:"workout_date"`
	WorkoutDuration     pgtype.Text        `json:"workout_duration"`
	Notes               pgtype.Text        `json:"notes"`
	FatigueLevel        NullFatiguelevel   `json:"fatigue_level"`
	TotalSets           pgtype.Int4        `json:"total_sets"`
	TotalDistance       pgtype.Int4        `json:"total_distance"`
	TotalRepetitions    pgtype.Int4        `json:"total_repetitions"`
	TotalWeightLifted   pgtype.Int4        `json:"total_weight_lifted"`
	TotalCaloriesBurned pgtype.Int4        `json:"total_calories_burned"`
	WorkoutID           int64              `json:"workout_id"`
}

func (q *Queries) UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (Workout, error) {
	row := q.db.QueryRow(ctx, updateWorkout,
		arg.WorkoutDate,
		arg.WorkoutDuration,
		arg.Notes,
		arg.FatigueLevel,
		arg.TotalSets,
		arg.TotalDistance,
		arg.TotalRepetitions,
		arg.TotalWeightLifted,
		arg.TotalCaloriesBurned,
		arg.WorkoutID,
	)
	var i Workout
	err := row.Scan(
		&i.WorkoutID,
		&i.Username,
		&i.WorkoutDate,
		&i.WorkoutDuration,
		&i.FatigueLevel,
		&i.Notes,
		&i.TotalCaloriesBurned,
		&i.TotalDistance,
		&i.TotalRepetitions,
		&i.TotalSets,
		&i.TotalWeightLifted,
		&i.CreatedAt,
	)
	return i, err
}
