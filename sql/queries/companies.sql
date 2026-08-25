-- name: CreateCompany :one
insert into companies (
    id, created_at, updated_at, name
) values (
    $1, $2, $3, $4
) returning *;

-- name: GetCompanies :many
select
    id,
    created_at,
    updated_at,
    name
from companies;

-- name: UpdateCompany :one
update companies set
    updated_at = $1,
    name = $2
where id = $3
returning *;


-- name: DeleteCompany :exec
delete from companies where id = $1;
