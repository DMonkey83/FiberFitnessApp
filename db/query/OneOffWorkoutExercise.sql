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
SET description = $3, muscle_group_name = $4
WHERE id = $1 AND workout_id = $2
RETURNING *;

-- name: ListAllOneOffWorkoutExercises :many
SELECT *
FROM OneOffWorkoutExercise
WHERE workout_id = $1
ORDER BY exercise_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
