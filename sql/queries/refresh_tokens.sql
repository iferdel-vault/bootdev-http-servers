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

-- name: GetUserFromRefreshToken :one
SELECT u.*
FROM refresh_tokens rt
INNER JOIN users u
ON rt.user_id = u.id
WHERE rt.token = $1;
