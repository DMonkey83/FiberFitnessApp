-- name: CreateAvailablePlan :one
INSERT INTO AvailableWorkoutPlans (
  plan_name, 
  description, 
  goal, 
  difficulty,
  is_public,
  creator_username
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetAvailablePlan :one
SELECT *
FROM AvailableWorkoutPlans
WHERE plan_id = $1;

-- name: DeleteAvailablePlan :exec
DELETE FROM AvailableWorkoutPlans
WHERE plan_id = $1;

-- name: UpdateAvailablePlan :one
UPDATE AvailableWorkoutPlans
SET 
description = $2, 
plan_name = $3, 
goal = $4,
difficulty = $5,
is_public = $6
WHERE plan_id = $1
RETURNING *;

-- name: ListAvailablePlansByCreator :many
SELECT *
FROM AvailableWorkoutPlans
WHERE creator_username =$1
ORDER BY plan_name -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;

-- name: ListAllAvailablePlans :many
SELECT *
FROM AvailableWorkoutPlans
ORDER BY plan_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $1
OFFSET $2;
