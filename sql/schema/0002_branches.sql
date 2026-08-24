-- +goose Up
CREATE TABLE IF NOT EXISTS branches (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    name TEXT NOT NULL,
    address TEXT NOT NULL,
    company_id UUID NOT NULL,
    FOREIGN KEY (company_id) REFERENCES companies(id)
);

-- +goose Down
DROP TABLE IF EXISTS branches;
