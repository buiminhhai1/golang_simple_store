-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  email,
  role,
  full_name
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserByUserName :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUser :one
UPDATE users SET full_name = $2
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;