
-- +migrate Up
ALTER TABLE users ADD password varchar(255) NULL;
-- +migrate Down
ALTER TABLE users DROP COLUMN password;
