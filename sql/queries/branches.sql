-- name: CreateBranch :one
insert into branches (
    id, created_at, updated_at, name, address, company_id
) values (
    $1, $2, $3, $4, $5, $6
) returning *;

-- name: GetBranchesByCompany :many
select
    id,
    created_at,
    updated_at,
    name,
    address
from branches
where company_id = $1;

-- name: UpdateBranch :one
update branches set
    updated_at = $1,
    name = $2,
    address = $3
where id = $4
returning *;

-- name: DeleteBranch :exec
delete from branches where id = $1;
