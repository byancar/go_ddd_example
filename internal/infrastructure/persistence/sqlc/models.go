// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package users

import (
	"database/sql"
)

type Address struct {
	ID         int32
	UserID     sql.NullInt64
	Street     sql.NullString
	City       sql.NullString
	State      sql.NullString
	Country    sql.NullString
	PostalCode sql.NullString
}

type User struct {
	ID       int32
	Username sql.NullString
	Email    sql.NullString
}
