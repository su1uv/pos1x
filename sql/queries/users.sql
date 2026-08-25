-- name: CreateUser :one
with inserted_user as (
    insert into users (
        id, created_at, updated_at, email, password_hash, branch_id
    ) values (
        $1, $2, $3, $4, $5, $6
    ) returning *
) select
    id,
    created_at,
    updated_at,
    email,
    branch_id
from inserted_user;

-- name: GetUsersByCompany :many
select
    u.id,
    u.created_at,
    u.updated_at,
    u.email,
    u.branch_id
from users u
inner join branches b on b.id = u.branch_id
where b.company_id = $1;

-- name: UpdateUser :one
update users set
    updated_at = $1,
    email = $2,
    branch_id = $3
where id = $4
returning *;

-- name: DeleteUser :exec
delete from users where id = $1;
