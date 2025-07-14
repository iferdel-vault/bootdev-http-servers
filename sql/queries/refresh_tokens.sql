-- name: GetUserFromRefreshToken :one
SELECT u.*
FROM refresh_tokens rt
INNER JOIN users u
ON rt.user_id = u.id
WHERE rt.token = $1;
