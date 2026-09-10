-- +goose Up
create table if not exists users (
    id uuid not null primary key,
    created_at timestamptz not null,
    updated_at timestamptz not null,
    name text not null,
    last_name text not null,
    email text not null,
    password_hash text not null,
    unique(email)
);

-- +goose Down
drop table if exists users;
