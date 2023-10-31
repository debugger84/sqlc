// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const getFooBar = `-- name: GetFooBar :exec
SELECT name FROM foo.bar
`

func (q *Queries) GetFooBar(ctx context.Context) error {
	_, err := q.db.Exec(ctx, getFooBar)
	return err
}

const updateFooBar = `-- name: UpdateFooBar :exec
UPDATE foo.bar SET name = $1
`

func (q *Queries) UpdateFooBar(ctx context.Context, name sql.NullString) error {
	_, err := q.db.Exec(ctx, updateFooBar, name)
	return err
}
