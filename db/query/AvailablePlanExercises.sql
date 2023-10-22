-- name: CreateAvailablePlanExercise :one
INSERT INTO AvailablePlanExercises (
    "exercise_name",
    "plan_id",
    "sets",
    "rest_duration",
    "notes"
)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAvailablePlanExercise :one
SELECT *
FROM AvailablePlanExercises
WHERE "id" = $1;

-- name: DeleteAvailablePlanExercise :exec
DELETE FROM AvailablePlanExercises
WHERE "id" = $1;

-- name: UpdateAvailablePlanExercise :one
UPDATE AvailablePlanExercises
SET 
notes = COALESCE(sqlc.narg(notes),notes),
sets = COALESCE(sqlc.narg(sets),sets),
rest_duration = COALESCE(sqlc.narg(rest_duration),rest_duration),
exercise_name = COALESCE(sqlc.narg(exercise_name),exercise_name)
WHERE id = @id
RETURNING *;

-- name: ListAllAvailablePlanExercises :many
SELECT *
FROM AvailablePlanExercises
ORDER BY "exercise_name"
LIMIT $1
OFFSET $2;
