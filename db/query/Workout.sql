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
RETURNING *;

-- name: ListWorkouts :many
SELECT * FROM Workout
WHERE username = $1
ORDER BY workout_date -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3; 
