-- name: CreateUser :one
INSERT INTO users (id, created_at, updated_at, email, hashed_password)
VALUES (
    gen_random_uuid(),
		NOW(),
		NOW(),
		$1,
		$2
)
RETURNING *;

-- name: DeleteUsers :exec
DELETE FROM users;

-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, email, hashed_password
FROM users
WHERE email = $1;

-- name: UpdateUserEmailAndPassword :one
UPDATE users
SET (email, hashed_password) = ($1, $2)
WHERE id = $1
RETURNING *;
