
-- name: CreateSet :one
INSERT INTO Set (exercise_name, set_number, weight, rest_duration, notes, reps_completed)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetSet :one
SELECT *
FROM Set
WHERE set_id = $1;

-- name: DeleteSet :exec
DELETE FROM Set
WHERE set_id = $1;

-- name: UpdateSet :one
UPDATE Set
SET 
set_number = COALESCE(sqlc.narg(set_number),set_number), 
weight = COALESCE(sqlc.narg(weight),weight), 
rest_duration = COALESCE(sqlc.narg(rest_duration),rest_duration), 
notes =COALESCE(sqlc.narg(notes),notes)
WHERE set_id = @set_id
RETURNING *;

-- name: ListSets :many
SELECT *
FROM Set
WHERE exercise_name = $1
ORDER BY set_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
