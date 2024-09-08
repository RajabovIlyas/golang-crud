-- name: CreateUser :one
insert into users (id, username, password, created_at,  updated_at)
values (gen_random_uuid(), $1, $2, now(), now())
returning id, username;

-- name: FindUsers :many
select id, username from users;

-- name: FindUserById :one
select id, username from users where id = $1;

-- name: UpdateUserById :one
update users set username = $2, updated_at = now() where id = $1
    returning id, username;

-- name: DeleteUserById :exec
delete from users where id = $1 ;

-- name: FindUserByUsername :one
select id, username, password from users where username = $1 ;

-- name: UpdateUserPasswordById :one
update users set password = $2, updated_at = now() where id = $1
returning id, username;