-- name: CreateUser :one
insert into users (
    id, created_at, updated_at, name, last_name, email, password_hash
) values (
    $1, $2, $3, $4, $5, $6, $7
) returning
    id,
    created_at,
    updated_at,
    name,
    last_name,
    email;
