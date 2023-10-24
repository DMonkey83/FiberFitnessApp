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
SET 
sets_completed = COALESCE(sqlc.narg(sets_completed),sets_completed),
repetitions_completed = COALESCE(sqlc.narg(repetitions_completed),repetitions_completed),
weight_lifted = COALESCE(sqlc.narg(weight_lifted),weight_lifted),
notes = COALESCE(sqlc.narg(notes),notes)
WHERE exercise_log_id = @exercise_log_id
RETURNING *;

-- name: ListExerciseLog :many
SELECT *
FROM ExerciseLog
WHERE log_id = $1
ORDER BY exercise_log_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
