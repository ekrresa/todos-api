
-- +migrate Up
ALTER TABLE todos
ALTER COLUMN updated_at
SET DEFAULT CURRENT_TIMESTAMP;

-- +migrate Down
ALTER TABLE todos
ALTER COLUMN updated_at
DROP DEFAULT;