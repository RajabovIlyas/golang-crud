// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: tokens.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createToken = `-- name: CreateToken :one
insert into tokens (id, access_token_key, user_id, created_at,  updated_at)
values (gen_random_uuid(), gen_random_uuid(), $1, now(), now())
returning id, access_token_key, user_id
`

type CreateTokenRow struct {
	ID             uuid.UUID
	AccessTokenKey uuid.UUID
	UserID         uuid.UUID
}

func (q *Queries) CreateToken(ctx context.Context, userID uuid.UUID) (CreateTokenRow, error) {
	row := q.db.QueryRowContext(ctx, createToken, userID)
	var i CreateTokenRow
	err := row.Scan(&i.ID, &i.AccessTokenKey, &i.UserID)
	return i, err
}

const deleteOldTokens = `-- name: DeleteOldTokens :exec
delete from tokens where updated_at <= NOW() - INTERVAL '2 days'
`

func (q *Queries) DeleteOldTokens(ctx context.Context) error {
	_, err := q.db.ExecContext(ctx, deleteOldTokens)
	return err
}

const deleteTokenByAccessKey = `-- name: DeleteTokenByAccessKey :exec
delete from tokens where access_token_key = $1
`

func (q *Queries) DeleteTokenByAccessKey(ctx context.Context, accessTokenKey uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTokenByAccessKey, accessTokenKey)
	return err
}

const deleteTokenById = `-- name: DeleteTokenById :exec
delete from tokens where id = $1
`

func (q *Queries) DeleteTokenById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteTokenById, id)
	return err
}

const findTokenByAccessKey = `-- name: FindTokenByAccessKey :one
select id, access_token_key, user_id from tokens where access_token_key = $1
`

type FindTokenByAccessKeyRow struct {
	ID             uuid.UUID
	AccessTokenKey uuid.UUID
	UserID         uuid.UUID
}

func (q *Queries) FindTokenByAccessKey(ctx context.Context, accessTokenKey uuid.UUID) (FindTokenByAccessKeyRow, error) {
	row := q.db.QueryRowContext(ctx, findTokenByAccessKey, accessTokenKey)
	var i FindTokenByAccessKeyRow
	err := row.Scan(&i.ID, &i.AccessTokenKey, &i.UserID)
	return i, err
}

const findTokenById = `-- name: FindTokenById :one
select id, access_token_key, user_id from tokens where id = $1
`

type FindTokenByIdRow struct {
	ID             uuid.UUID
	AccessTokenKey uuid.UUID
	UserID         uuid.UUID
}

func (q *Queries) FindTokenById(ctx context.Context, id uuid.UUID) (FindTokenByIdRow, error) {
	row := q.db.QueryRowContext(ctx, findTokenById, id)
	var i FindTokenByIdRow
	err := row.Scan(&i.ID, &i.AccessTokenKey, &i.UserID)
	return i, err
}

const updateTokenById = `-- name: UpdateTokenById :one
update tokens set access_token_key = gen_random_uuid(), updated_at = now() where id = $1
returning id, access_token_key, user_id
`

type UpdateTokenByIdRow struct {
	ID             uuid.UUID
	AccessTokenKey uuid.UUID
	UserID         uuid.UUID
}

func (q *Queries) UpdateTokenById(ctx context.Context, id uuid.UUID) (UpdateTokenByIdRow, error) {
	row := q.db.QueryRowContext(ctx, updateTokenById, id)
	var i UpdateTokenByIdRow
	err := row.Scan(&i.ID, &i.AccessTokenKey, &i.UserID)
	return i, err
}
