package dao

import (
	"crypto/rand"
	"flying-castle/db"
	"flying-castle/encryption"
	"flying-castle/requests"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/scrypt"
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
	encryption.MustUpdateKey(storageKeyDAO.Key)
	return storageKeyDAO
}

func (skRepo *StorageKeyRepository) Create() StorageKeyDAO {
	var tx = skRepo.tx
	asset := requests.MustAsset("requests/insert_storage_key.sql")
	var req = string(asset)
	lastId, err := db.GetLastIdFrom(tx, "storage_key")
	if err != nil {
		panic(err)
	}
	var key = make([]byte, 256)
	_, err = rand.Read(key)
	if err != nil {
		panic(err)
	}
	var salt = make([]byte, 256)
	_, err = rand.Read(salt)
	if err != nil {
		panic(err)
	}
	storageKey, err := scrypt.Key(key, salt, 32768, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	_ = tx.MustExec(req, lastId.Int64+1, encryption.EncodeKey(storageKey), time.Now())
	encryption.MustUpdateKey(encryption.EncodeKey(storageKey))
	return skRepo.GetLatest()
}
