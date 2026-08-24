-- +goose Up
CREATE TABLE IF NOT EXISTS companies (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    name TEXT NOT NULL,
    UNIQUE(name)
);

-- +goose Down
DROP TABLE IF EXISTS companies;
