-- name: CreateWeightEntry :one
INSERT INTO WeightEntry (username, entry_date, weight_kg, weight_lb, notes)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetWeightEntry :one
SELECT *
FROM WeightEntry
WHERE weight_entry_id = $1 AND username = $2;

-- name: DeleteWeightEntry :exec
DELETE FROM WeightEntry
WHERE weight_entry_id = $1 AND username = $2;

-- name: UpdateWeightEntry :one
UPDATE WeightEntry
SET 
entry_date = COALESCE(sqlc.narg(entry_date),entry_date),
weight_kg = COALESCE(sqlc.narg(weight_kg),weight_kg),
weight_lb = COALESCE(sqlc.narg(weight_lb),weight_lb),
notes = COALESCE(sqlc.narg(notes),notes)
WHERE weight_entry_id = @weight_entry_id AND username = @username
RETURNING *;

-- name: ListWeightEntries :many
SELECT *
FROM WeightEntry
WHERE username = $1
ORDER BY weight_entry_id -- You can change the ORDER BY clause to order by a different column if needed
LIMIT $2
OFFSET $3;
