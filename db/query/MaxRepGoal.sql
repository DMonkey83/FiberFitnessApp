
-- name: CreateMaxRepGoal :one
INSERT INTO MaxRepGoal (username, exercise_name, goal_reps, notes)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetMaxRepGoal :one
SELECT *
FROM MaxRepGoal
WHERE exercise_name = $1 AND username = $2 AND goal_id = $3;

-- name: DeleteMaxRepGoal :exec
DELETE FROM MaxRepGoal
WHERE exercise_name = $1 AND username = $2 AND goal_id = $3;

-- name: UpdateMaxRepGoal :one
UPDATE MaxRepGoal
SET 
goal_reps = COALESCE(sqlc.narg(goal_reps),goal_reps), 
notes = COALESCE(sqlc.narg(notes),notes)
WHERE exercise_name = @exercise_name AND username = @username AND goal_id = @goal_id
RETURNING *;

-- name: ListMaxRepGoals :many
SELECT *
FROM MaxRepGoal
WHERE exercise_name = $1 AND username = $2
ORDER BY goal_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $3
OFFSET $4;
