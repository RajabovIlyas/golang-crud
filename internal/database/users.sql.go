// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
insert into users (id, username, password, created_at,  updated_at)
values (gen_random_uuid(), $1, $2, now(), now())
returning id, username
`

type CreateUserParams struct {
	Username string
	Password string
}

type CreateUserRow struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password)
	var i CreateUserRow
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const deleteUserById = `-- name: DeleteUserById :exec
delete from users where id = $1
`

func (q *Queries) DeleteUserById(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.ExecContext(ctx, deleteUserById, id)
	return err
}

const findUserById = `-- name: FindUserById :one
select id, username from users where id = $1
`

type FindUserByIdRow struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) FindUserById(ctx context.Context, id uuid.UUID) (FindUserByIdRow, error) {
	row := q.db.QueryRowContext(ctx, findUserById, id)
	var i FindUserByIdRow
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const findUserByUsername = `-- name: FindUserByUsername :one
select id, username, password from users where username = $1
`

type FindUserByUsernameRow struct {
	ID       uuid.UUID
	Username string
	Password string
}

func (q *Queries) FindUserByUsername(ctx context.Context, username string) (FindUserByUsernameRow, error) {
	row := q.db.QueryRowContext(ctx, findUserByUsername, username)
	var i FindUserByUsernameRow
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const findUsers = `-- name: FindUsers :many
select id, username from users
`

type FindUsersRow struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) FindUsers(ctx context.Context) ([]FindUsersRow, error) {
	rows, err := q.db.QueryContext(ctx, findUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []FindUsersRow
	for rows.Next() {
		var i FindUsersRow
		if err := rows.Scan(&i.ID, &i.Username); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUserById = `-- name: UpdateUserById :one
update users set username = $2, updated_at = now() where id = $1
    returning id, username
`

type UpdateUserByIdParams struct {
	ID       uuid.UUID
	Username string
}

type UpdateUserByIdRow struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) UpdateUserById(ctx context.Context, arg UpdateUserByIdParams) (UpdateUserByIdRow, error) {
	row := q.db.QueryRowContext(ctx, updateUserById, arg.ID, arg.Username)
	var i UpdateUserByIdRow
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}

const updateUserPasswordById = `-- name: UpdateUserPasswordById :one
update users set password = $2, updated_at = now() where id = $1
returning id, username
`

type UpdateUserPasswordByIdParams struct {
	ID       uuid.UUID
	Password string
}

type UpdateUserPasswordByIdRow struct {
	ID       uuid.UUID
	Username string
}

func (q *Queries) UpdateUserPasswordById(ctx context.Context, arg UpdateUserPasswordByIdParams) (UpdateUserPasswordByIdRow, error) {
	row := q.db.QueryRowContext(ctx, updateUserPasswordById, arg.ID, arg.Password)
	var i UpdateUserPasswordByIdRow
	err := row.Scan(&i.ID, &i.Username)
	return i, err
}
