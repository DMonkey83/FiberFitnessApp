
-- name: CreateMaxWeightGoal :one
INSERT INTO MaxWeightGoal (username, exercise_name, goal_weight, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMaxWeightGoal :one
SELECT *
FROM MaxWeightGoal
WHERE exercise_name = $1 AND username = $2 AND goal_id = $3;

-- name: DeleteMaxWeightGoal :exec
DELETE FROM MaxWeightGoal
WHERE exercise_name = $1 AND username = $2 AND goal_id = $3;

-- name: UpdateMaxWeightGoal :one
UPDATE MaxWeightGoal
SET goal_weight = $4, notes = $5
WHERE exercise_name = $1 AND username = $2 AND goal_id = $3
RETURNING *;

-- name: ListMaxWeightGoals :many
SELECT *
FROM MaxWeightGoal
WHERE exercise_name = $1 AND username = $2
ORDER BY goal_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $3
OFFSET $4;
