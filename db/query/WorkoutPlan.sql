
-- name: CreatePlan :one
INSERT INTO WorkoutPlan (
  username, 
  plan_name,
  description,
  start_date,
  end_date,
  goal,
  difficulty,
  is_public
  )
VALUES ($1, $2, $3, $4,$5,$6,$7,$8)
RETURNING *;

-- name: GetPlan :one
SELECT *
FROM WorkoutPlan
WHERE plan_id = $1 AND username = $2;

-- name: DeletePlan :exec
DELETE FROM WorkoutPlan
WHERE plan_id = $1 AND username =$2;

-- name: UpdatePlan :one
UPDATE WorkoutPlan
SET 
plan_name = COALESCE(sqlc.narg(plan_name),plan_name),
description = COALESCE(sqlc.narg(description),description),
start_date = COALESCE(sqlc.narg(start_date),start_date),
end_date = COALESCE(sqlc.narg(end_date),end_date),
goal = COALESCE(sqlc.narg(goal),goal),
difficulty = COALESCE(sqlc.narg(difficulty),difficulty),
is_public = COALESCE(sqlc.narg(is_public),is_public)
WHERE plan_id = @plan_id AND username = @username
RETURNING *;
