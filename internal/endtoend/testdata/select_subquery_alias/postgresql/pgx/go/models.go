// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0

package querytest

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Wallet struct {
	ID      int64
	Address string
	Type    string
	Balance pgtype.Numeric
}
