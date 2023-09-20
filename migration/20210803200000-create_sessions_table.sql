-- +migrate Up
CREATE TABLE IF NOT EXISTS "sessions"(
    session_id SERIAL NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    token VARCHAR(100) NOT NULL,
    expires_at TIMESTAMP,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE INDEX sessions_access_token ON "sessions" (token);

-- +migrate Down
DROP TABLE IF EXISTS "sessions";
