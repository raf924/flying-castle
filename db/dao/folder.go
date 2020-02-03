package dao

type FolderDAO struct {
	Id     int64  `db:"id"`
	PathId int64  `db:"path_id"`
	Name   string `db:"name"`
}
