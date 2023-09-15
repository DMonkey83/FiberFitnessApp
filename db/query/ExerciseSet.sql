-- name: CreateExerciseSet :one
INSERT INTO ExerciseSet (exercise_log_id, set_number, weight_lifted, repetitions_completed)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetExerciseSet :one
SELECT *
FROM ExerciseSet
WHERE set_id = $1;

-- name: DeleteExerciseSet :exec
DELETE FROM ExerciseSet
WHERE set_id = $1;

-- name: UpdateExerciseSet :one
UPDATE ExerciseSet
SET weight_lifted = $2, repetitions_completed = $3
WHERE set_id = $1
RETURNING *;

-- name: ListExerciseSets :many
SELECT *
FROM ExerciseSet
WHERE exercise_log_id = $1
ORDER BY set_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
