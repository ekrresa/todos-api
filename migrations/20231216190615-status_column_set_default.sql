
-- +migrate Up
ALTER TABLE todos
ALTER COLUMN status
SET DEFAULT 'new';

-- +migrate Down
ALTER TABLE todos
ALTER COLUMN status
DROP DEFAULT;
