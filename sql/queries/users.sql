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

-- TODO: Limit & pagination
-- name: ListUsers :many
select
    id,
    created_at,
    updated_at,
    name,
    last_name,
    email
from users;

-- name: GetUserByEmail :one
select
    id,
    created_at,
    updated_at,
    name,
    last_name,
    email
from users
where email = $1;

-- name: GetUserByID :one
select
    id,
    created_at,
    updated_at,
    name,
    last_name,
    email
from users
where id = $1;

-- name: GetUserByEmailForAuth :one
select
    id,
    email,
    password_hash
from users
where email = $1;

-- name: UpdateUser :one
update users set
    updated_at = $1,
    name = $2,
    last_name = $3,
    email = $4
where id = $5
returning
    id,
    created_at,
    updated_at,
    name,
    last_name,
    email;

-- name: DeleteUser :exec
delete from users where id = $1;
