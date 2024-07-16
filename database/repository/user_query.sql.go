// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: user_query.sql

package repository

import (
	"context"
	"database/sql"
	"encoding/json"
)

const createUser = `-- name: CreateUser :execresult
INSERT INTO user (uuid, roles, password, email, username, created_at, updated_at) VALUES (?, ?, ?, ?, ?, NOW(), NOW())
`

type CreateUserParams struct {
	Uuid     string
	Roles    json.RawMessage
	Password string
	Email    string
	Username sql.NullString
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.createUserStmt, createUser,
		arg.Uuid,
		arg.Roles,
		arg.Password,
		arg.Email,
		arg.Username,
	)
}

const deleteUser = `-- name: DeleteUser :execresult
UPDATE user SET deleted_at = NOW() WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) (sql.Result, error) {
	return q.exec(ctx, q.deleteUserStmt, deleteUser, id)
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, created_at, updated_at, uuid, roles, password, email, username FROM user WHERE email = ? AND deleted_at IS NULL
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Uuid,
		&i.Roles,
		&i.Password,
		&i.Email,
		&i.Username,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, created_at, updated_at, uuid, roles, password, email, username FROM user WHERE id = ? AND deleted_at IS NULL
`

func (q *Queries) GetUserById(ctx context.Context, id int32) (User, error) {
	row := q.queryRow(ctx, q.getUserByIdStmt, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Uuid,
		&i.Roles,
		&i.Password,
		&i.Email,
		&i.Username,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, created_at, updated_at, uuid, roles, password, email, username FROM user WHERE deleted_at IS NULL
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Uuid,
			&i.Roles,
			&i.Password,
			&i.Email,
			&i.Username,
		); err != nil {
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

const updateUser = `-- name: UpdateUser :execresult
UPDATE user SET roles = ?, password = ?, email = ?, username = ?, updated_at = NOW() WHERE id = ? AND deleted_at IS NULL
`

type UpdateUserParams struct {
	Roles    json.RawMessage
	Password string
	Email    string
	Username sql.NullString
	ID       int32
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (sql.Result, error) {
	return q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Roles,
		arg.Password,
		arg.Email,
		arg.Username,
		arg.ID,
	)
}