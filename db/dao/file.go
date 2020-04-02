package dao

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"flying-castle/app"
	"flying-castle/db"
	"flying-castle/encryption"
	"github.com/jmoiron/sqlx"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type FileDAO struct {
	Id           int64     `db:"id"`
	FirstChunkId int64     `db:"first_chunk_id"`
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
	var err = fileRepo.tx.Get(&fileDAO, "SELECT * FROM file WHERE id = ?", id)
	if err != nil {
		panic(err)
	}
	return fileDAO
}

func Reverse(array []byte) []byte {
	newArray := make([]byte, len(array))
	for i, j := 0, len(array)-1; i <= j; i, j = i+1, j-1 {
		newArray[i], newArray[j] = array[j], array[i]
	}
	return newArray
}

func (fileRepo *FileRepository) Create(parent int64, name string, firstChunkId int64, fileSize int64) (FileDAO, error) {
	var fileDAO = FileDAO{}
	lastFileId, err := db.GetLastIdFrom(fileRepo.tx, "file")
	if err != nil {
		return fileDAO, nil
	}
	lastPathId, err := db.GetLastIdFrom(fileRepo.tx, "path")
	if err != nil {
		return fileDAO, nil
	}
	var now = time.Now()
	_, err = fileRepo.tx.Exec(
		"INSERT INTO path(id, parent_id, name, created_at, accessed_at, modified_at) VALUES (?,?,?,?,?,?)",
		lastPathId.Int64+1, parent, name, now, now, now)
	if err != nil {
		return fileDAO, nil
	}
	_, err = fileRepo.tx.Exec(
		"INSERT INTO file (id, first_chunk_id, path_id, \"size\", modified_at) VALUES (?, ?, ?, ?, ?)",
		lastFileId.Int64+1, firstChunkId, lastPathId.Int64+1, fileSize, now)
	if err != nil {
		return fileDAO, nil
	}
	err = fileRepo.tx.Get(&fileDAO, "SELECT * FROM file where id = ?", lastFileId.Int64+1)
	return fileDAO, err
}

func (fileRepo *FileRepository) Save(parent int64, name string, data []byte, storageKey string) {
	_, err := base64.StdEncoding.DecodeString(storageKey)
	if err != nil {
		panic(err)
	}
	var chunkRepo = NewChunkRepository(fileRepo.tx)
	lastFileId, err := db.GetLastIdFrom(fileRepo.tx, "file")
	buffer := bytes.NewBuffer(Reverse(data))
	var minChunkSize = math.Min(math.Pow(2, 20), float64(len(data)))
	var nextChunkId = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	for buffer.Len() > 0 {
		lastChunkId, err := db.GetLastIdFrom(fileRepo.tx, "chunk")
		var chunkSize = rand.Intn(len(data) / 4)
		chunkSize = int(math.Max(float64(chunkSize), minChunkSize))
		var chunk = buffer.Next(chunkSize)
		var encryptedChunk = encryption.Encrypt(Reverse(chunk))
		config := app.GetConfig()
		path := filepath.Join(config.DataPath, strconv.FormatInt(lastChunkId.Int64+1, 10))
		chunkFile, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		_, err = chunkFile.Write(encryptedChunk)
		if err != nil {
			panic(err)
		}
		var chunkDAO = ChunkDAO{
			Id:        lastChunkId.Int64 + 1,
			Path:      path,
			NextChunk: nextChunkId,
		}
		chunkRepo.Create(chunkDAO)
		nextChunkId = sql.NullInt64{
			Int64: lastChunkId.Int64 + 1,
			Valid: true,
		}
	}
	lastPathId, err := db.GetLastIdFrom(fileRepo.tx, "path")
	if err != nil {
		panic(err)
	}
	var now = time.Now()
	fileRepo.tx.MustExec(
		"INSERT INTO path(id, parent_id, name, created_at, accessed_at, modified_at) VALUES (?,?,?,?,?,?)",
		lastPathId.Int64+1, parent, name, now, now, now)
	fileRepo.tx.MustExec(
		"INSERT INTO file (id, first_chunk_id, path_id, \"size\", modified_at) VALUES (?, ?, ?, ?, ?)",
		lastFileId.Int64+1, nextChunkId.Int64, lastPathId.Int64+1, len(data), now)
}

func (fileRepo *FileRepository) FindByPathId(id int64) (FileDAO, error) {
	var file = FileDAO{}
	var err = fileRepo.tx.Get(&file, "SELECT * FROM file WHERE path_id = ?", id)
	return file, err
}
