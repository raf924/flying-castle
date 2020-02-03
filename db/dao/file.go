package dao

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type FileDAO struct {
	Id           int64     `db:"id"`
	FirstChunkId int64     `db:"first_chunk_id"`
	Name         string    `db:"name"`
	PathId       int64     `db:"path_id"`
	Size         int64     `db:"size"`
	ModifiedAt   time.Time `db:"modified_at"`
}

type FileRepository struct {
	tx *sqlx.Tx
}

func NewFileRepository(tx *sqlx.Tx) FileRepository {
	return FileRepository{tx: tx}
}

func (fileRepo *FileRepository) GetById(id int64) FileDAO {
	var fileDAO = FileDAO{}
	var err = fileRepo.tx.Get(&fileDAO, "SELECT * FROM fileDAO WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return fileDAO
}
