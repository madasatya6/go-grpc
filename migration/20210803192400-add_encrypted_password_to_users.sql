-- +migrate Up
ALTER TABLE users ADD COLUMN encrypted_password VARCHAR(100);

-- +migrate Down
ALTER TABLE users DROP COLUMN encrypted_password;
