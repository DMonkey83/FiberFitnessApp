-- name: CreateExerciseLog :one
INSERT INTO ExerciseLog (
  log_id,
  exercise_name, 
  sets_completed, 
  repetitions_completed,
  weight_lifted,
  notes
  )
VALUES ($1, $2, $3, $4,$5,$6)
RETURNING *;

-- name: GetExerciseLog :one
SELECT *
FROM ExerciseLog
WHERE exercise_log_id = $1;

-- name: DeleteExerciseLog :exec
DELETE FROM ExerciseLog
WHERE exercise_log_id = $1;

-- name: UpdateExerciseLog :one
UPDATE ExerciseLog
SET sets_completed = $2, repetitions_completed = $3, weight_lifted = $4, notes = $5
WHERE exercise_log_id = $1
RETURNING *;

-- name: ListExerciseLog :many
SELECT *
FROM ExerciseLog
WHERE log_id = $1
ORDER BY exercise_log_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
