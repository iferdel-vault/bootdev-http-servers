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
SELECT id, created_at, updated_at, email, hashed_password, is_chirpy_red
FROM users
WHERE email = $1;

-- name: UpdateUserEmailAndPassword :one
UPDATE users
SET (email, hashed_password, updated_at) = ($1, $2, NOW())
WHERE id = $3
RETURNING *;

-- name: UpdateUserIsChirpyRed :one
UPDATE users
SET (is_chirpy_red, updated_at) = ($1, NOW())
WHERE id = $2
RETURNING *;
