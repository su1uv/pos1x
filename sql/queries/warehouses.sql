-- name: CreateWarehouse :one
insert into warehouses (
    id, created_at, updated_at, name
) values (
    $1, $2, $3, $4
) returning *;

-- name: GetWarehouses :many
select
    id,
    created_at,
    updated_at,
    name
from warehouses;

-- name: UpdateWarehouse :one
update warehouses set
    updated_at = $1,
    name = $2
where id = $3
returning *;

-- name: DeleteWarehouse :exec
delete from warehouses where id = $1;
