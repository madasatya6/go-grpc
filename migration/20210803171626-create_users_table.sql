-- +migrate Up
CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL NOT NULL PRIMARY KEY,
    company_id INTEGER,
    name VARCHAR(100),
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(100),
    display_picture_url VARCHAR(100),
    identity_id VARCHAR(100),
    invite_notification_flag BOOLEAN,
    is_active_flag BOOLEAN,
    created_at TIMESTAMP NOT NULL,
    created_by INT,
    updated_at TIMESTAMP NOT NULL,
    updated_by INT
);

-- +migrate Down
DROP TABLE IF EXISTS users;
