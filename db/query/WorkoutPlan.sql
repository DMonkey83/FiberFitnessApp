
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
plan_name = $3, 
description = $4,
start_date = $5,
end_date = $6,
goal = $7,
difficulty = $8,
is_public = $9
WHERE plan_id = $1 AND username = $2
RETURNING *;
