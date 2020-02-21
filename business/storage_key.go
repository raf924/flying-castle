package business

import (
	"flying-castle/db"
	"flying-castle/db/dao"
	"flying-castle/encryption"
	"flying-castle/model"
	"github.com/jmoiron/sqlx"
)

type DBStorageKeyBusiness struct {
	db *sqlx.DB
}

func (D *DBStorageKeyBusiness) Create() error {
	var tx = D.db.MustBegin()
	var skRepo = dao.NewStorageKeyRepository(tx)
	err := skRepo.Create()
	if err != nil {
		return db.CreationError
	}
	err = tx.Commit()
	if err != nil {
		return db.TransactionError
	}
	return nil
}

func (D *DBStorageKeyBusiness) GetLatest() (model.StorageKey, error) {
	var tx = D.db.MustBegin()
	var skRepo = dao.NewStorageKeyRepository(tx)
	var skDAO = skRepo.GetLatest()
	err := tx.Commit()
	if err != nil {
		return model.StorageKey{}, db.TransactionError
	}
	var sk = model.StorageKey{
		Key:       encryption.MustDecodeKey(skDAO.Key),
		CreatedAt: skDAO.CreatedAt,
	}
	encryption.MustUpdateKey(skDAO.Key)
	return sk, nil
}

func NewDBStorageKeyBusiness() DBStorageKeyBusiness {
	return DBStorageKeyBusiness{db: sqlx.NewDb(db.GetDB().GetDB(), db.GetDB().GetDriver())}
}
