package main

import "golang.org/x/crypto/scrypt"
import (
	_ "bytes"
	"crypto"
	"crypto/rand"
	"encoding/base64"
	"flying-castle/cmd"
	"flying-castle/requests"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
	"time"
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
	asset := requests.MustAsset("requests/insert_storage_key.sql")
	var req = string(asset)
	var b = struct {
		Id int64 `db:"id"`
	}{}
	_ = tx.Select(b, "SELECT MAX(id) as id from storage_key")
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
	_ = tx.MustExec(req, b.Id+1, base64.StdEncoding.EncodeToString(storageKey), time.Now())
	err = tx.Commit()
}
