-- +goose Up
CREATE TABLE refresh_tokens (
	token TEXT PRIMARY KEY,
	created_at TIMESTAMPTZ NOT NULL,
	updated_at TIMESTAMPTZ NOT NULL,
	user_id UUID NOT NULL,
	expires_at TIMESTAMPTZ NOT NULL,
	revoked_at TIMESTAMPTZ DEFAULT NULL,
	CONSTRAINT fk_user
		FOREIGN KEY(user_id)
			REFERENCES users(id)
			ON DELETE CASCADE
);

-- +goose Down
DROP TABLE refresh_tokens;
