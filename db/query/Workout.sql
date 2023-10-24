-- name: CreateWorkout :one
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
RETURNING *;

-- name: GetWorkout :one
SELECT *
FROM Workout
WHERE workout_id = $1;

-- name: DeleteWorkout :exec
DELETE FROM Workout
WHERE workout_id = $1;

-- name: UpdateWorkout :one
UPDATE Workout
SET 
workout_date = COALESCE(sqlc.narg(workout_date),workout_date),
workout_duration = COALESCE(sqlc.narg(workout_duration),workout_duration),
notes = COALESCE(sqlc.narg(notes),notes),
fatigue_level = COALESCE(sqlc.narg(fatigue_level),fatigue_level),
total_sets = COALESCE(sqlc.narg(total_sets),total_sets),
total_distance=COALESCE(sqlc.narg(total_distance),total_distance),
total_repetitions=COALESCE(sqlc.narg(total_repetitions),total_repetitions),
total_weight_lifted=COALESCE(sqlc.narg(total_weight_lifted),total_weight_lifted),
total_calories_burned =COALESCE(sqlc.narg(total_calories_burned),total_calories_burned)
WHERE workout_id = @workout_id
RETURNING *;

-- name: ListWorkouts :many
SELECT * FROM Workout
WHERE username = $1
ORDER BY workout_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3; 
