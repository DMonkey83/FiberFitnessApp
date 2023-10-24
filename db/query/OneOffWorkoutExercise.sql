-- name: CreateOneOffWorkoutExercise :one
INSERT INTO OneOffWorkoutExercise (
  workout_id,
  exercise_name,
  description,
  muscle_group_name
  )
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetOneOffWorkoutExercise :one
SELECT *
FROM OneOffWorkoutExercise
WHERE id = $1 AND workout_id = $2;

-- name: DeleteOneOffWorkoutExercise :exec
DELETE FROM OneOffWorkoutExercise
WHERE id = $1 AND workout_id = $2;

-- name: UpdateOneOffWorkoutExercise :one
UPDATE OneOffWorkoutExercise
SET 
description = COALESCE(sqlc.narg(description),description),
muscle_group_name = COALESCE(sqlc.narg(muscle_group_name),muscle_group_name)
WHERE id = @id AND workout_id = @workout_id
RETURNING *;

-- name: ListAllOneOffWorkoutExercises :many
SELECT *
FROM OneOffWorkoutExercise
WHERE workout_id = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
