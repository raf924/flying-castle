package dao

import (
	"database/sql"
	"time"
)

type Path struct {
	Id         int64         `db:"id"`
	ParentId   sql.NullInt64 `db:"parent_id"`
	CreatedAt  time.Time     `db:"created_at"`
	ModifiedAt time.Time     `db:"modified_at"`
	AccessedAt time.Time     `db:"accessed_at"`
}
