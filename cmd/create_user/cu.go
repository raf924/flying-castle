package main

import (
	_ "bytes"
	"crypto"
	"encoding/base64"
	"flying-castle/cmd"
	db2 "flying-castle/db"
	"flying-castle/requests"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
	"time"
)

type UserFlags struct {
	UserName string `flag:"name" required:"true"`
	Password string `flag:"password" required:"true"`
}

func main() {
	var config = cmd.GetConfig()
	var dbUrl, err = url.Parse(config.DbUrl)

	var flags = UserFlags{}
	cmd.ReadFlags(&flags)

	if err != nil {
		panic(err)
	}
	db, err := sqlx.Connect(dbUrl.Scheme, dbUrl.EscapedPath())
	if err != nil {
		panic(err)
	}
	tx := db.MustBegin()
	asset := requests.MustAsset("requests/insert_user.sql")
	var req = string(asset)
	var lastPath = db2.MustGetLastIdFrom(tx, "path")
	var lastUser = db2.MustGetLastIdFrom(tx, "user")
	var lastGroup = db2.MustGetLastIdFrom(tx, "group")
	var hash = crypto.SHA512_256.New()
	_, err = hash.Write([]byte(flags.Password))
	if err != nil {
		panic(err)
	}
	var passordHash = hash.Sum(nil)
	var now = time.Now()
	_ = tx.MustExec(req, lastPath+1, now, now, now,
		lastGroup+1, flags.UserName,
		lastUser+1, flags.UserName, base64.StdEncoding.EncodeToString(passordHash), "", lastPath+1, lastGroup+1)
	err = tx.Commit()
}
