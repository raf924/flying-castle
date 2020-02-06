package db

import (
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
)

func MustGetLastIdFrom(tx *sqlx.Tx, tableName string) int64 {
	id, err := GetLastIdFrom(tx, tableName)
	if err != nil {
		panic(err)
	}
	return id.Int64
}

func GetLastIdFrom(tx *sqlx.Tx, tableName string) (sql.NullInt64, error) {
	var lastRow = struct {
		Id sql.NullInt64 `db:"id"`
	}{}
	var err = tx.Get(&lastRow, fmt.Sprintf("SELECT MAX(id) as id from \"%s\"", tableName))
	if err != nil {
		return sql.NullInt64{
			Int64: 0,
			Valid: false,
		}, err
	}
	return lastRow.Id, nil
}

func MustCommit(tx *sqlx.Tx) {
	var err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
