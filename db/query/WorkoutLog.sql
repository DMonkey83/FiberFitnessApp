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
RETURNING *;

-- name: ListWorkoutLogs :many
SELECT * FROM WorkoutLog
WHERE plan_id = $1
ORDER BY log_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3; 
