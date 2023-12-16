
-- +migrate Up
CREATE TYPE TODO_STATUS AS ENUM ('new','in_progress', 'completed', 'cancelled');

CREATE TABLE IF NOT EXISTS todos (
	id SERIAL PRIMARY KEY,
	name VARCHAR(100) NOT NULL,
	status TODO_STATUS,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

-- +migrate Down
drop table todos;
drop type TODO_STATUS;