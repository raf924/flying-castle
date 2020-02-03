package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

func MustGetLastIdFrom(tx *sqlx.Tx, tableName string) int64 {
	id, err := GetLastIdFrom(tx, tableName)
	if err != nil {
		panic(err)
	}
	return id
}

func GetLastIdFrom(tx *sqlx.Tx, tableName string) (int64, error) {
	var lastRow = struct {
		Id int64 `db:"id"`
	}{}
	var err = tx.Get(&lastRow, fmt.Sprintf("SELECT MAX(id) as id from \"%s\"", tableName))
	if err != nil {
		return 0, err
	}
	return lastRow.Id, nil
}
