-- name: CreateUserProfile :one
INSERT INTO UserProfile (username, full_name, age, gender, height_cm, height_ft_in, preferred_unit)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserProfile :one
SELECT sqlc.embed(UserProfile),sqlc.embed(users)
FROM UserProfile
JOIN users ON UserProfile.username = users.username
WHERE UserProfile.username = $1 LIMIT 1;

-- name: UpdateUserProfile :one
UPDATE UserProfile
SET full_name = $2, age = $3, gender = $4, height_cm = $5, height_ft_in = $6, preferred_unit = $7
WHERE username = $1
RETURNING *;

-- name: DeleteUserProfile :exec
DELETE FROM UserProfile
WHERE username = $1;

