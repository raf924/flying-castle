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

func (chunkRepo *ChunkRepository) GetById(id int64) ChunkDAO {
	var chunkDAO = ChunkDAO{}
	var err = chunkRepo.tx.Get(&chunkDAO, "SELECT * FROM chunk where id = ?", id)
	if err != nil {
		panic(err)
	}
	return chunkDAO
}

func (chunkRepo *ChunkRepository) Create(dao ChunkDAO) {
	chunkRepo.tx.MustExec("INSERT INTO chunk(id, path, next_chunk) VALUES (?, ?, ?)", dao.Id, dao.Path, dao.NextChunk)
}

func (chunkRepo *ChunkRepository) UpdatePath(id int64, path string) {
	chunkRepo.tx.MustExec("UPDATE chunk SET path = ? WHERE id = ?", path, id)
}
