// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: WorkoutLog.sql

package db

import (
	"context"
	"time"
)

const createWorkoutLog = `-- name: CreateWorkoutLog :one
INSERT INTO WorkoutLog 
  (
  username, 
  plan_id,
  log_date, 
  rating,
  fatigue_level,
  overall_feeling,
  comments,
  workout_duration,
  total_calories_burned,
  total_distance,
  total_repetitions,
  total_sets,
  total_weight_lifted
  )
VALUES ($1, $2, $3, $4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
RETURNING log_id, username, plan_id, log_date, rating, fatigue_level, overall_feeling, comments, workout_duration, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
`

type CreateWorkoutLogParams struct {
	Username            string       `json:"username"`
	PlanID              int64        `json:"plan_id"`
	LogDate             time.Time    `json:"log_date"`
	Rating              Rating       `json:"rating"`
	FatigueLevel        Fatiguelevel `json:"fatigue_level"`
	OverallFeeling      string       `json:"overall_feeling"`
	Comments            string       `json:"comments"`
	WorkoutDuration     string       `json:"workout_duration"`
	TotalCaloriesBurned int32        `json:"total_calories_burned"`
	TotalDistance       int32        `json:"total_distance"`
	TotalRepetitions    int32        `json:"total_repetitions"`
	TotalSets           int32        `json:"total_sets"`
	TotalWeightLifted   int32        `json:"total_weight_lifted"`
}

func (q *Queries) CreateWorkoutLog(ctx context.Context, arg CreateWorkoutLogParams) (Workoutlog, error) {
	row := q.db.QueryRow(ctx, createWorkoutLog,
		arg.Username,
		arg.PlanID,
		arg.LogDate,
		arg.Rating,
		arg.FatigueLevel,
		arg.OverallFeeling,
		arg.Comments,
		arg.WorkoutDuration,
		arg.TotalCaloriesBurned,
		arg.TotalDistance,
		arg.TotalRepetitions,
		arg.TotalSets,
		arg.TotalWeightLifted,
	)
	var i Workoutlog
	err := row.Scan(
		&i.LogID,
		&i.Username,
		&i.PlanID,
		&i.LogDate,
		&i.Rating,
		&i.FatigueLevel,
		&i.OverallFeeling,
		&i.Comments,
		&i.WorkoutDuration,
		&i.TotalCaloriesBurned,
		&i.TotalDistance,
		&i.TotalRepetitions,
		&i.TotalSets,
		&i.TotalWeightLifted,
		&i.CreatedAt,
	)
	return i, err
}

const deleteWorkoutLog = `-- name: DeleteWorkoutLog :exec
DELETE FROM WorkoutLog
WHERE log_id = $1
`

func (q *Queries) DeleteWorkoutLog(ctx context.Context, logID int64) error {
	_, err := q.db.Exec(ctx, deleteWorkoutLog, logID)
	return err
}

const getWorkoutLog = `-- name: GetWorkoutLog :one
SELECT log_id, username, plan_id, log_date, rating, fatigue_level, overall_feeling, comments, workout_duration, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
FROM WorkoutLog
WHERE log_id = $1
`

func (q *Queries) GetWorkoutLog(ctx context.Context, logID int64) (Workoutlog, error) {
	row := q.db.QueryRow(ctx, getWorkoutLog, logID)
	var i Workoutlog
	err := row.Scan(
		&i.LogID,
		&i.Username,
		&i.PlanID,
		&i.LogDate,
		&i.Rating,
		&i.FatigueLevel,
		&i.OverallFeeling,
		&i.Comments,
		&i.WorkoutDuration,
		&i.TotalCaloriesBurned,
		&i.TotalDistance,
		&i.TotalRepetitions,
		&i.TotalSets,
		&i.TotalWeightLifted,
		&i.CreatedAt,
	)
	return i, err
}

const listWorkoutLogs = `-- name: ListWorkoutLogs :many
SELECT log_id, username, plan_id, log_date, rating, fatigue_level, overall_feeling, comments, workout_duration, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at FROM WorkoutLog
WHERE plan_id = $1
ORDER BY log_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3
`

type ListWorkoutLogsParams struct {
	PlanID int64 `json:"plan_id"`
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListWorkoutLogs(ctx context.Context, arg ListWorkoutLogsParams) ([]Workoutlog, error) {
	rows, err := q.db.Query(ctx, listWorkoutLogs, arg.PlanID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Workoutlog{}
	for rows.Next() {
		var i Workoutlog
		if err := rows.Scan(
			&i.LogID,
			&i.Username,
			&i.PlanID,
			&i.LogDate,
			&i.Rating,
			&i.FatigueLevel,
			&i.OverallFeeling,
			&i.Comments,
			&i.WorkoutDuration,
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

const updateWorkoutLog = `-- name: UpdateWorkoutLog :one
UPDATE WorkoutLog
SET 
log_date = $2, 
workout_duration = $3, 
comments = $4,
fatigue_level = $5, 
total_sets =$6,
total_distance=$7,
total_repetitions=$8,
total_weight_lifted=$9,
total_calories_burned =$10,
rating = $11,
overall_feeling = $12
WHERE log_id = $1
RETURNING log_id, username, plan_id, log_date, rating, fatigue_level, overall_feeling, comments, workout_duration, total_calories_burned, total_distance, total_repetitions, total_sets, total_weight_lifted, created_at
`

type UpdateWorkoutLogParams struct {
	LogID               int64        `json:"log_id"`
	LogDate             time.Time    `json:"log_date"`
	WorkoutDuration     string       `json:"workout_duration"`
	Comments            string       `json:"comments"`
	FatigueLevel        Fatiguelevel `json:"fatigue_level"`
	TotalSets           int32        `json:"total_sets"`
	TotalDistance       int32        `json:"total_distance"`
	TotalRepetitions    int32        `json:"total_repetitions"`
	TotalWeightLifted   int32        `json:"total_weight_lifted"`
	TotalCaloriesBurned int32        `json:"total_calories_burned"`
	Rating              Rating       `json:"rating"`
	OverallFeeling      string       `json:"overall_feeling"`
}

func (q *Queries) UpdateWorkoutLog(ctx context.Context, arg UpdateWorkoutLogParams) (Workoutlog, error) {
	row := q.db.QueryRow(ctx, updateWorkoutLog,
		arg.LogID,
		arg.LogDate,
		arg.WorkoutDuration,
		arg.Comments,
		arg.FatigueLevel,
		arg.TotalSets,
		arg.TotalDistance,
		arg.TotalRepetitions,
		arg.TotalWeightLifted,
		arg.TotalCaloriesBurned,
		arg.Rating,
		arg.OverallFeeling,
	)
	var i Workoutlog
	err := row.Scan(
		&i.LogID,
		&i.Username,
		&i.PlanID,
		&i.LogDate,
		&i.Rating,
		&i.FatigueLevel,
		&i.OverallFeeling,
		&i.Comments,
		&i.WorkoutDuration,
		&i.TotalCaloriesBurned,
		&i.TotalDistance,
		&i.TotalRepetitions,
		&i.TotalSets,
		&i.TotalWeightLifted,
		&i.CreatedAt,
	)
	return i, err
}
