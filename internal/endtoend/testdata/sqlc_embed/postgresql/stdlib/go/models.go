// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"database/sql"
)

type BazUser struct {
	ID   int32
	Name string
}

type Post struct {
	ID     int32
	UserID int32
	Likes  []int32
}

type User struct {
	ID   int32
	Name string
	Age  sql.NullInt32
}
