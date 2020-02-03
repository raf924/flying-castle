package dao

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type StorageKeyDAO struct {
	Id        int64     `db:"id"`
	Key       string    `db:"key"`
	CreatedAt time.Time `db:"created_at"`
}

type StorageKeyRepository struct {
	tx *sqlx.Tx
}

func NewStorageKeyRepository(tx *sqlx.Tx) StorageKeyRepository {
	return StorageKeyRepository{tx: tx}
}

func (skRepo *StorageKeyRepository) GetLatest() StorageKeyDAO {
	var storageKeyDAO = StorageKeyDAO{}
	var err = skRepo.tx.Get(&storageKeyDAO, "SELECT * FROM storage_key where created_at = MAX(created_at)")
	if err != nil {
		panic(err)
	}
	return storageKeyDAO
}
