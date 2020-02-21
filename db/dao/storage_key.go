package dao

import (
	"flying-castle/db"
	"flying-castle/encryption"
	"flying-castle/requests"
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
	var err = skRepo.tx.Get(&storageKeyDAO, "SELECT * FROM storage_key where created_at = (SELECT MAX(created_at) FROM storage_key) LIMIT 1")
	if err != nil {
		panic(err)
	}
	return storageKeyDAO
}

func (skRepo *StorageKeyRepository) Create() error {
	var tx = skRepo.tx
	asset := requests.MustAsset("requests/insert_storage_key.sql")
	var req = string(asset)
	lastId, err := db.GetLastIdFrom(tx, "storage_key")
	if err != nil {
		panic(err)
	}
	storageKey, err := encryption.GenerateKey()
	if err != nil {
		return err
	}
	_, err = tx.Exec(req, lastId.Int64+1, encryption.EncodeKey(storageKey), time.Now())
	if err != nil {
		return err
	}
	encryption.MustUpdateKey(encryption.EncodeKey(storageKey))
	return nil
}
