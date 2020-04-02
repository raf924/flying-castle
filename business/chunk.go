package business

import (
	"flying-castle/db"
	"flying-castle/db/dao"
	"flying-castle/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ChunkBusiness struct {
	db *sqlx.DB
}

var getOrphansQuery = `SELECT c.* from chunk c 
left outer join chunk c1 on c.id = c1.next_chunk
left outer join file f on c.id = f.first_chunk_id
 where c1.next_chunk is null and f.first_chunk_id is null`

func NewChunkBusiness() *ChunkBusiness {
	return &ChunkBusiness{db: sqlx.NewDb(db.GetDB().GetDB(), db.GetDB().GetDriver())}
}

func (cB *ChunkBusiness) MustGetChunkById(id int64) model.Chunk {
	chunk, err := cB.GetChunkById(id)
	if err != nil {
		panic(err)
	}
	return chunk
}

func findOrphanChunks(tx *sqlx.Tx) ([]model.Chunk, error) {
	rows, err := tx.Query(getOrphansQuery)
	if err != nil {
		return nil, err
	}
	var chunks = []model.Chunk{}
	for rows.Next() {
		var chunkDao = dao.ChunkDAO{}
		err = rows.Scan(&chunkDao)
		if err != nil {
			return nil, err
		}
		var nextChunk = new(int64)
		if chunkDao.NextChunk.Valid {
			*nextChunk = chunkDao.NextChunk.Int64
		}
		chunks = append(chunks, model.Chunk{
			Id:        chunkDao.Id,
			Key:       chunkDao.Path,
			NextChunk: nextChunk,
		})
	}
	return chunks, nil
}

func (cB *ChunkBusiness) DeleteOrphanChunks() ([]model.Chunk, error) {
	var deletedChunks = []model.Chunk{}
	var tx = cB.db.MustBegin()
	for {
		chunks, err := findOrphanChunks(tx)
		if err != nil {
			return nil, model.DatabaseError
		}
		if len(chunks) == 0 {
			return nil, nil
		}
		params := strings.Join(strings.Split(strings.Repeat("?", len(chunks)), ""), ",")
		deleteQuery := fmt.Sprintf("DELETE FROM chunk WHERE id IN (%s)", params)
		res, err := tx.Exec(deleteQuery, chunks)
		if err != nil {
			return nil, model.DatabaseError
		}
		var rows int64
		if rows, err = res.RowsAffected(); err != nil {
			return nil, model.DatabaseError
		}
		if rows == 0 {
			break
		}
		deletedChunks = append(deletedChunks, chunks...)
	}
	return deletedChunks, tx.Commit()
}

func (cB *ChunkBusiness) GetChunkById(id int64) (model.Chunk, error) {
	var tx = cB.db.MustBegin()
	var chunkRepo = dao.NewChunkRepository(tx)
	var chunkDao = chunkRepo.GetById(id)
	var chunk = model.Chunk{
		Id:        chunkDao.Id,
		Key:       chunkDao.Path,
		NextChunk: &chunkDao.NextChunk.Int64,
	}
	err := tx.Commit()
	if err != nil {
		return chunk, model.DatabaseError
	}
	return chunk, nil
}
