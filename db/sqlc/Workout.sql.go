// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: Workout.sql

package db

import (
	"context"
	"time"
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
workout_date = $2, 
workout_duration = $3, 
notes = $4,
fatigue_level = $5, 
total_sets =$6,
total_distance=$7,
total_repetitions=$8,
total_weight_lifted=$9,
total_calories_burned =$10
WHERE workout_id = $1
RETURNING workout_id, username, workout_date, workout_duration, fatigue_level, notes, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
`

type UpdateWorkoutParams struct {
	WorkoutID           int64        `json:"workout_id"`
	WorkoutDate         time.Time    `json:"workout_date"`
	WorkoutDuration     string       `json:"workout_duration"`
	Notes               string       `json:"notes"`
	FatigueLevel        Fatiguelevel `json:"fatigue_level"`
	TotalSets           int32        `json:"total_sets"`
	TotalDistance       int32        `json:"total_distance"`
	TotalRepetitions    int32        `json:"total_repetitions"`
	TotalWeightLifted   int32        `json:"total_weight_lifted"`
	TotalCaloriesBurned int32        `json:"total_calories_burned"`
}

func (q *Queries) UpdateWorkout(ctx context.Context, arg UpdateWorkoutParams) (Workout, error) {
	row := q.db.QueryRow(ctx, updateWorkout,
		arg.WorkoutID,
		arg.WorkoutDate,
		arg.WorkoutDuration,
		arg.Notes,
		arg.FatigueLevel,
		arg.TotalSets,
		arg.TotalDistance,
		arg.TotalRepetitions,
		arg.TotalWeightLifted,
		arg.TotalCaloriesBurned,
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
