package business

import (
	"crypto/rand"
	"errors"
	"flying-castle/db"
	"flying-castle/db/dao"
	"flying-castle/encryption"
	"flying-castle/model"
	"flying-castle/validation"
	"github.com/jmoiron/sqlx"
)

type UserBusiness struct {
	db *sqlx.DB
}

func NewDBUserBusiness() UserBusiness {
	return UserBusiness{db: sqlx.NewDb(db.GetDB().GetDB(), db.GetDB().GetDriver())}
}

func (userBusiness *UserBusiness) FindUserByNameAndPassword(name string, password string) (*model.User, error) {
	var tx = userBusiness.db.MustBegin()
	var userRepo = dao.NewUserRepository(tx)
	var userDAO = userRepo.FindByName(name)
	if userDAO == nil {
		return nil, model.InvalidCredentials
	}
	salt, err := encryption.DecodeKey(userDAO.Salt)
	if err != nil {
		return nil, err
	}
	hash := encryption.EncodeKey(encryption.MustHash(password, salt))
	if hash != userDAO.Hash {
		return nil, model.InvalidCredentials
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
	if !validation.ValidateValue(name, validation.Alphanumeric) {
		return model.InvalidNewUsername
	}
	if !validation.ValidateValue(password, validation.Password) {
		return model.InvalidNewPassword
	}
	var tx = userBusiness.db.MustBegin()
	var userRepo = dao.NewUserRepository(tx)
	var salt = make([]byte, 256)
	_, err := rand.Read(salt)
	if err != nil {
		return errors.New("error while generating hash salt")
	}
	hash := encryption.MustHash(password, salt)
	err = userRepo.Create(name, hash, salt)
	if err != nil {
		return err
	}
	return tx.Commit()
}
