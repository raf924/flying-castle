package dao

import (
	"database/sql"
	"flying-castle/requests"
	"github.com/jmoiron/sqlx"
	"time"
)

type PathDAO struct {
	Id         int64         `db:"id"`
	ParentId   sql.NullInt64 `db:"parent_id"`
	CreatedAt  time.Time     `db:"created_at"`
	ModifiedAt time.Time     `db:"modified_at"`
	AccessedAt time.Time     `db:"accessed_at"`
	Name       string        `db:"name"`
}

type PathRepository struct {
	tx *sqlx.Tx
}

func NewPathRepository(tx *sqlx.Tx) PathRepository {
	return PathRepository{tx: tx}
}

func (pathRepo *PathRepository) FindByParentAndName(parent int64, name string) *PathDAO {
	var pathDAO = PathDAO{}
	var req = requests.MustAsset("requests/find_by_name_in_parent.sql")
	var err = pathRepo.tx.Get(&pathDAO, string(req), parent, name)
	if err != nil {
		return nil
	}
	return &pathDAO
}

func (pathRepo *PathRepository) GetById(id int64) PathDAO {
	var pathDAO = PathDAO{}
	var err = pathRepo.tx.Get(&pathDAO, "SELECT * FROM path where id = ?", id)
	if err != nil {
		panic(err)
	}
	return pathDAO
}
