-- +goose Up
CREATE TABLE IF NOT EXISTS users (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    email TEXT NOT NULL,
    password_hash TEXT NOT NULL,
    branch_id UUID NOT NULL,
    UNIQUE (email),
    FOREIGN KEY (branch_id) REFERENCES branches(id) on delete cascade
);

-- +goose Down
DROP TABLE IF EXISTS users;
