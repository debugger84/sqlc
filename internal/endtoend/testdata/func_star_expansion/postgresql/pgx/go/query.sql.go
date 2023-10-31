// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const testFuncGetTime = `-- name: TestFuncGetTime :one
select test_func_get_time from test_func_get_time()
`

func (q *Queries) TestFuncGetTime(ctx context.Context) (pgtype.Timestamp, error) {
	row := q.db.QueryRow(ctx, testFuncGetTime)
	var test_func_get_time pgtype.Timestamp
	err := row.Scan(&test_func_get_time)
	return test_func_get_time, err
}

const testFuncSelectBlog = `-- name: TestFuncSelectBlog :one
select test_func_select_blog from test_func_select_blog($1)
`

func (q *Queries) TestFuncSelectBlog(ctx context.Context, pID int32) (interface{}, error) {
	row := q.db.QueryRow(ctx, testFuncSelectBlog, pID)
	var test_func_select_blog interface{}
	err := row.Scan(&test_func_select_blog)
	return test_func_select_blog, err
}

const testFuncSelectString = `-- name: TestFuncSelectString :one
select test_func_select_string from test_func_select_string($1)
`

func (q *Queries) TestFuncSelectString(ctx context.Context, pString string) (pgtype.Text, error) {
	row := q.db.QueryRow(ctx, testFuncSelectString, pString)
	var test_func_select_string pgtype.Text
	err := row.Scan(&test_func_select_string)
	return test_func_select_string, err
}
