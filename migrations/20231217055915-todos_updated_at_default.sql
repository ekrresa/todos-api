
-- +migrate Up

-- +migrate StatementBegin
CREATE OR REPLACE FUNCTION trigger_set_timestamp()
RETURNS TRIGGER AS $$
BEGIN
  NEW.updated_at = CURRENT_TIMESTAMP;
  RETURN NEW;
END;
$$
LANGUAGE plpgsql;
-- +migrate StatementEnd

CREATE TRIGGER set_timestamp
BEFORE UPDATE ON todos
FOR EACH ROW
EXECUTE PROCEDURE trigger_set_timestamp();

-- +migrate Down
DROP TRIGGER IF EXISTS set_timestamp ON todos;
DROP FUNCTION IF EXISTS trigger_set_timestamp();