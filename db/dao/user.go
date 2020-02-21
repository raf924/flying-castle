package dao

import (
	"database/sql"
	"flying-castle/db"
	"flying-castle/encryption"
	"flying-castle/requests"
	"github.com/jmoiron/sqlx"
)

type UserDAO struct {
	Id           int64  `db:"id"`
	Name         string `db:"user_name"`
	RootFolderId int64  `db:"root_folder_id"`
	MainGroupId  int64  `db:"main_group_id"`
	Hash         string `db:"hash"`
	Salt         string `db:"salt"`
}

type UserRepository struct {
	tx *sqlx.Tx
}

func NewUserRepository(tx *sqlx.Tx) UserRepository {
	return UserRepository{tx: tx}
}

func (userRepo *UserRepository) FindByName(name string) *UserDAO {
	var userDAO = UserDAO{}
	var err = userRepo.tx.Get(&userDAO, "SELECT * FROM user where user.user_name = ?", name)
	if err != nil {
		return nil
	}
	return &userDAO
}

func (userRepo *UserRepository) FindByNameAndHash(name string, hash string) UserDAO {
	var userDAO = UserDAO{}
	var err = userRepo.tx.Get(&userDAO, "SELECT * FROM user where user.user_name = '?' and user.hash = '?'", name, hash)
	if err != nil {
		panic(err)
	}
	return userDAO
}

func (userRepo *UserRepository) GetById(id int64) UserDAO {
	var userDAO = UserDAO{}
	err := userRepo.tx.Get(&userDAO, "SELECT * FROM user where id = ?", id)
	if err != nil {
		panic(err)
	}
	return userDAO
}

func (userRepo *UserRepository) FindAll() []UserDAO {
	var userDAO = make([]UserDAO, 0)
	err := userRepo.tx.Select(&userDAO, "SELECT * FROM user")
	if err != nil {
		panic(err)
	}
	return userDAO
}

func (userRepo *UserRepository) Create(name string, passwordHash []byte, salt []byte) error {
	var tx = userRepo.tx
	asset := requests.MustAsset("requests/insert_user.sql")
	var folderRepo = NewFolderRepository(userRepo.tx)
	folderRepo.Create(sql.NullInt64{
		Int64: 0,
		Valid: false,
	}, name)
	var req = string(asset)
	var lastPath = db.MustGetLastIdFrom(tx, "folder")
	var lastGroup = db.MustGetLastIdFrom(tx, "group")
	var lastUser = db.MustGetLastIdFrom(tx, "user")
	_, err := tx.Exec(req, lastGroup+1, name,
		lastUser+1, name, encryption.EncodeKey(passwordHash), encryption.EncodeKey(salt), lastPath, lastGroup+1)
	return err
}
