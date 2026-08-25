-- +goose Up
create table if not exists warehouses (
    id uuid not null primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    name text not null
);

-- +goose Down
drop table if exists warehouses;
