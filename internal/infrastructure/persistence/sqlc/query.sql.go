// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: query.sql

package users

import (
	"context"
	"database/sql"
)

const deleteAddress = `-- name: DeleteAddress :exec
DELETE FROM addresses WHERE id = $1
`

func (q *Queries) DeleteAddress(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteAddress, id)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1
`

func (q *Queries) DeleteUser(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteUser, id)
	return err
}

const getAddressesByUserID = `-- name: GetAddressesByUserID :one
SELECT id, user_id, street, city, state, country, postal_code FROM addresses WHERE user_id = $1
`

func (q *Queries) GetAddressesByUserID(ctx context.Context, userID sql.NullInt64) (Address, error) {
	row := q.db.QueryRowContext(ctx, getAddressesByUserID, userID)
	var i Address
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Street,
		&i.City,
		&i.State,
		&i.Country,
		&i.PostalCode,
	)
	return i, err
}

const getAllUsers = `-- name: GetAllUsers :many
SELECT id, username, email FROM users
`

func (q *Queries) GetAllUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, getAllUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(&i.ID, &i.Username, &i.Email); err != nil {
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

const getUserByID = `-- name: GetUserByID :one
SELECT id, username, email FROM users WHERE id = $1
`

func (q *Queries) GetUserByID(ctx context.Context, id int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByID, id)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Email)
	return i, err
}

const insertAddress = `-- name: InsertAddress :exec
INSERT INTO addresses (user_id, street, city, state, country, postal_code)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING id
`

type InsertAddressParams struct {
	UserID     sql.NullInt64
	Street     sql.NullString
	City       sql.NullString
	State      sql.NullString
	Country    sql.NullString
	PostalCode sql.NullString
}

func (q *Queries) InsertAddress(ctx context.Context, arg InsertAddressParams) error {
	_, err := q.db.ExecContext(ctx, insertAddress,
		arg.UserID,
		arg.Street,
		arg.City,
		arg.State,
		arg.Country,
		arg.PostalCode,
	)
	return err
}

const insertUser = `-- name: InsertUser :exec
INSERT INTO users (username, email) VALUES ($1, $2) RETURNING id
`

type InsertUserParams struct {
	Username sql.NullString
	Email    sql.NullString
}

func (q *Queries) InsertUser(ctx context.Context, arg InsertUserParams) error {
	_, err := q.db.ExecContext(ctx, insertUser, arg.Username, arg.Email)
	return err
}

const updateAddress = `-- name: UpdateAddress :exec
UPDATE addresses
SET street = $2, city = $3, state = $4, country = $5, postal_code = $6
WHERE id = $1
`

type UpdateAddressParams struct {
	ID         int32
	Street     sql.NullString
	City       sql.NullString
	State      sql.NullString
	Country    sql.NullString
	PostalCode sql.NullString
}

func (q *Queries) UpdateAddress(ctx context.Context, arg UpdateAddressParams) error {
	_, err := q.db.ExecContext(ctx, updateAddress,
		arg.ID,
		arg.Street,
		arg.City,
		arg.State,
		arg.Country,
		arg.PostalCode,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users SET username = $2, email = $3 WHERE id = $1
`

type UpdateUserParams struct {
	ID       int32
	Username sql.NullString
	Email    sql.NullString
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.db.ExecContext(ctx, updateUser, arg.ID, arg.Username, arg.Email)
	return err
}