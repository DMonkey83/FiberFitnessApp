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
SET 
full_name = COALESCE(sqlc.narg(full_name),full_name),
age = COALESCE(sqlc.narg(age),age),
gender = COALESCE(sqlc.narg(gender),gender), 
height_cm = COALESCE(sqlc.narg(height_cm),height_cm),
height_ft_in = COALESCE(sqlc.narg(height_ft_in),height_ft_in),
preferred_unit = COALESCE(sqlc.narg(preferred_unit),preferred_unit)
WHERE username = @username
RETURNING *;

-- name: DeleteUserProfile :exec
DELETE FROM UserProfile
WHERE username = $1;

