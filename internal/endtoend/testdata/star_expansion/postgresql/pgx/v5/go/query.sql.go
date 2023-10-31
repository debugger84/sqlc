// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: query.sql

package querytest

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const starExpansion = `-- name: StarExpansion :many
SELECT a, b, a, b, foo.a, foo.b FROM foo
`

type StarExpansionRow struct {
	A   pgtype.Text
	B   pgtype.Text
	A_2 pgtype.Text
	B_2 pgtype.Text
	A_3 pgtype.Text
	B_3 pgtype.Text
}

func (q *Queries) StarExpansion(ctx context.Context) ([]StarExpansionRow, error) {
	rows, err := q.db.Query(ctx, starExpansion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []StarExpansionRow
	for rows.Next() {
		var i StarExpansionRow
		if err := rows.Scan(
			&i.A,
			&i.B,
			&i.A_2,
			&i.B_2,
			&i.A_3,
			&i.B_3,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const starQuotedExpansion = `-- name: StarQuotedExpansion :many
SELECT t.a, t.b FROM foo "t"
`

func (q *Queries) StarQuotedExpansion(ctx context.Context) ([]Foo, error) {
	rows, err := q.db.Query(ctx, starQuotedExpansion)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.A, &i.B); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
