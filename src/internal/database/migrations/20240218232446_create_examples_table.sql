-- +goose Up

-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION set_updated_at_column() RETURNS TRIGGER AS $$
  BEGIN
   NEW.updated_at = NOW();
   RETURN NEW;
  END;
$$ language 'plpgsql';
-- +goose StatementEnd

-- +goose StatementBegin
CREATE TABLE examples (
    id              VARCHAR(32)     PRIMARY KEY NOT NULL,
    description     TEXT            NOT NULL,
    created_at      TIMESTAMP(6)    NOT NULL DEFAULT NOW(),
    updated_at      TIMESTAMP(6)    NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE examples;
-- +goose StatementEnd
