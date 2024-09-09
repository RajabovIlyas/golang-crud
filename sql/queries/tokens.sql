-- name: CreateToken :one
insert into tokens (id, access_token_key, user_id, created_at,  updated_at)
values (gen_random_uuid(), gen_random_uuid(), $1, now(), now())
returning id, access_token_key, user_id;

-- name: FindTokenById :one
select id, access_token_key, user_id from tokens where id = $1;

-- name: FindTokenByAccessKey :one
select id, access_token_key, user_id from tokens where access_token_key = $1;

-- name: UpdateTokenById :one
update tokens set access_token_key = gen_random_uuid(), updated_at = now() where id = $1
returning id, access_token_key, user_id;

-- name: DeleteTokenById :one
delete from tokens where id = $1
returning id, access_token_key, user_id;

-- name: DeleteTokenByAccessKey :one
delete from tokens where access_token_key = $1
returning id, access_token_key, user_id;

-- name: DeleteOldTokens :many
delete from tokens where updated_at <= NOW() - INTERVAL '2 days'
returning id, access_token_key, user_id;