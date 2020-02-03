package dao

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
)

type ChunkDAO struct {
	Id        int64         `db:"id"`
	Path      string        `db:"path"`
	NextChunk sql.NullInt64 `db:"next_chunk"`
}

type ChunkRepository struct {
	tx *sqlx.Tx
}

func NewChunkRepository(tx *sqlx.Tx) ChunkRepository {
	return ChunkRepository{tx: tx}
}

func (cr *ChunkRepository) GetById(id int64) ChunkDAO {
	var chunkDAO = ChunkDAO{}
	var err = cr.tx.Get(&chunkDAO, "SELECT * FROM chunk where id = ?", id)
	if err != nil {
		panic(err)
	}
	return chunkDAO
}
