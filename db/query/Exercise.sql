-- name: CreateExercise :one
INSERT INTO Exercise (exercise_name, description, equipment_required, muscle_group_name)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetExercise :one
SELECT *
FROM Exercise
WHERE exercise_name = $1;

-- name: DeleteExercise :exec
DELETE FROM Exercise
WHERE exercise_name = $1;

-- name: UpdateExercise :one
UPDATE Exercise
SET 
description = COALESCE(sqlc.narg(description),description), 
equipment_required =COALESCE(sqlc.narg(equipment_required),equipment_required), 
muscle_group_name = COALESCE(sqlc.narg(muscle_group_name),muscle_group_name)
WHERE exercise_name = @exercise_name
RETURNING *;

-- name: ListEquipmentExercises :many
SELECT *
FROM Exercise
WHERE equipment_required = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;

-- name: ListMuscleGroupExercises :many
SELECT *
FROM Exercise
WHERE muscle_group_name = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;

-- name: ListAllExercises :many
SELECT *
FROM Exercise
ORDER BY exercise_name  -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
