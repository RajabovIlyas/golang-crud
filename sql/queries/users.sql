-- name: CreateUser :one
insert into users (id, name, created_at,  updated_at)
values (gen_random_uuid(), $1, CURRENT_DATE, CURRENT_DATE)
returning *;

-- name: GetUsers :many
select * from users;

-- name: GetUserById :one
select * from users where id = $1;

-- name: ChangeUserById :one
update users set name = $2, updated_at = CURRENT_DATE where id = $1
    returning *;

-- name: DeleteUserById :exec
delete from users where id = $1 ;