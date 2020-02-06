package business

import (
	"crypto/rand"
	"errors"
	"flying-castle/db"
	"flying-castle/db/dao"
	"flying-castle/encryption"
	"flying-castle/model"
	"github.com/jmoiron/sqlx"
)

type UserBusiness struct {
	db *sqlx.DB
}

func NewUserBusiness(db *sqlx.DB) UserBusiness {
	return UserBusiness{db: db}
}

func (userBusiness *UserBusiness) FindUserByNameAndPassword(name string, password string) (*model.User, error) {
	var tx = userBusiness.db.MustBegin()
	var userRepo = dao.NewUserRepository(tx)
	var userDAO = userRepo.FindByName(name)
	if userDAO == nil {
		return nil, errors.New("invalid username or password")
	}
	salt, err := encryption.DecodeKey(userDAO.Salt)
	if err != nil {
		return nil, err
	}
	hash := encryption.EncodeKey(encryption.MustHash(password, salt))
	if hash != userDAO.Hash {
		return nil, errors.New("invalid username or password")
	}
	db.MustCommit(tx)
	var user = model.User{
		FileSystemEntity: model.FileSystemEntity{
			Id:   uint64(userDAO.Id),
			Name: userDAO.Name,
		},
		Group: nil,
	}
	return &user, nil
}

func (userBusiness *UserBusiness) Create(name string, password string) error {
	var tx = userBusiness.db.MustBegin()
	var userRepo = dao.NewUserRepository(tx)
	var salt = make([]byte, 256)
	_, err := rand.Read(salt)
	if err != nil {
		panic(err)
	}
	hash := encryption.MustHash(password, salt)
	userRepo.Create(name, hash, salt)
	return tx.Commit()
}
