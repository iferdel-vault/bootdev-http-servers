-- name: CreateRefreshToken :one
INSERT INTO refresh_tokens (token, created_at, updated_at, user_id, expires_at, revoked_at)
VALUES (
    $1,
		NOW(),
		NOW(),
		$2,
		$3,
		$4
)
RETURNING *;

-- name: GetRefreshToken :one
SELECT token, created_at, updated_at, user_id, expires_at, revoked_at
FROM refresh_tokens
WHERE token = $1;

-- name: RevokeRefreshToken :exec
UPDATE refresh_tokens
SET (revoked_at, updated_at) = (NOW(), NOW())
WHERE token = $1;

-- name: GetUserFromRefreshToken :one
SELECT u.*
FROM refresh_tokens rt
INNER JOIN users u
ON rt.user_id = u.id
WHERE rt.token = $1;
