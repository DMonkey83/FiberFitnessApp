-- name: CreateWorkoutLog :one
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
RETURNING *;

-- name: GetWorkoutLog :one
SELECT *
FROM WorkoutLog
WHERE log_id = $1;

-- name: DeleteWorkoutLog :exec
DELETE FROM WorkoutLog
WHERE log_id = $1;

-- name: UpdateWorkoutLog :one
UPDATE WorkoutLog
SET 
log_date = COALESCE(sqlc.narg(log_date),log_date),
workout_duration = COALESCE(sqlc.narg(workout_duration),workout_duration),
comments = COALESCE(sqlc.narg(comments),comments),
fatigue_level = COALESCE(sqlc.narg(fatigue_level),fatigue_level),
total_sets =COALESCE(sqlc.narg(total_sets),total_sets),
total_distance=COALESCE(sqlc.narg(total_distance),total_distance),
total_repetitions= COALESCE(sqlc.narg(total_repetitions),total_repetitions),
total_weight_lifted= COALESCE(sqlc.narg(total_weight_lifted),total_weight_lifted),
total_calories_burned = COALESCE(sqlc.narg(total_calories_burned),total_calories_burned),
rating = COALESCE(sqlc.narg(rating),rating),
overall_feeling = COALESCE(sqlc.narg(overall_feeling),overall_feeling)
WHERE log_id = @log_id
RETURNING *;

-- name: ListWorkoutLogs :many
SELECT * FROM WorkoutLog
WHERE plan_id = $1
ORDER BY log_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3; 
