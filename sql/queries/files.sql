-- name: CreateFile :one
insert into files (id, file_name, path, format, user_id, size, created_at,  updated_at)
values (gen_random_uuid(), $1, $2, $3, $4, $5, now(), now())
returning id, file_name, path, format, user_id;

-- name: FindFileById :one
select id, file_name, path, format, user_id, size from files where id = $1;

-- name: FindFileByFileName :one
select id, file_name, path, format, user_id, size from files where file_name = $1;

-- name: FindFiles :many
select id, file_name, path, format, user_id, size from files;

-- name: UpdateFileById :one
update files set file_name = $2, path = $3, format = $4, updated_at = now() where id = $1
returning id, file_name, path, format, user_id, size;

-- name: DeleteFileById :exec
delete from files where id = $1 ;


