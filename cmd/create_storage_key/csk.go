package main

import (
	_ "bytes"
	"flying-castle/cmd"
	"flying-castle/db/dao"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
)

func main() {
	var config = cmd.GetConfig()
	var dbUrl, err = url.Parse(config.DbUrl)
	if err != nil {
		panic(err)
	}
	db, err := sqlx.Connect(dbUrl.Scheme, dbUrl.EscapedPath())
	if err != nil {
		panic(err)
	}
	tx := db.MustBegin()
	var skRepo = dao.NewStorageKeyRepository(tx)
	skRepo.Create()
	err = tx.Commit()
}
