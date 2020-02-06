package dao

import (
	"database/sql"
	"flying-castle/db"
	"github.com/jmoiron/sqlx"
	"time"
)

type FolderDAO struct {
	Id     int64 `db:"id"`
	PathId int64 `db:"path_id"`
}

type FolderRepository struct {
	tx *sqlx.Tx
}

func NewFolderRepository(tx *sqlx.Tx) FolderRepository {
	return FolderRepository{tx: tx}
}

func (folderRepo *FolderRepository) GetById(id int64) FolderDAO {
	var folderDAO = FolderDAO{}
	err := folderRepo.tx.Get(&folderDAO, "SELECT * FROM folder where id = ?", id)
	if err != nil {
		panic(err)
	}
	return folderDAO
}

func (folderRepo *FolderRepository) Create(parent sql.NullInt64, name string) {
	lastFolderId, err := db.GetLastIdFrom(folderRepo.tx, "folder")
	if err != nil {
		panic(err)
	}
	lastPathId, err := db.GetLastIdFrom(folderRepo.tx, "path")
	if err != nil {
		panic(err)
	}
	var now = time.Now()
	folderRepo.tx.MustExec(
		"INSERT INTO path (id, name, parent_id, created_at, accessed_at, modified_at) VALUES (?,?,?,?,?,?)",
		lastPathId.Int64+1, name, parent, now, now, now)
	folderRepo.tx.MustExec(
		"INSERT INTO folder (id, path_id) VALUES (?, ?)",
		lastFolderId.Int64+1, lastPathId.Int64+1)
}

func (folderRepo *FolderRepository) GetUserRoot(userId int64) FolderDAO {
	var folderDAO = FolderDAO{}
	var err = folderRepo.tx.Get(&folderDAO, "SELECT f.* from user u inner join folder f on u.root_folder_id = f.id")
	if err != nil {
		panic(err)
	}
	return folderDAO
}

func (folderRepo *FolderRepository) FindByPathId(id int64) (FolderDAO, error) {
	var folder = FolderDAO{}
	var err = folderRepo.tx.Get(&folder, "SELECT * FROM folder WHERE path_id = ?", id)
	return folder, err
}
