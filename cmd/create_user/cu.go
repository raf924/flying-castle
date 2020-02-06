package main

import (
	_ "bytes"
	"flying-castle/business"
	"flying-castle/cmd"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
)

type UserFlags struct {
	UserName string `flag:"name" required:"true"`
	Password string `flag:"password" required:"true"`
}

func (u *UserFlags) Validate() {

}

func main() {
	var config = cmd.GetConfig()
	var dbUrl, err = url.Parse(config.DbUrl)
	if err != nil {
		panic(err)
	}
	var flags = UserFlags{}
	cmd.ReadFlags(&flags)
	db, err := sqlx.Connect(dbUrl.Scheme, dbUrl.EscapedPath())
	if err != nil {
		panic(err)
	}
	var userBusiness = business.NewUserBusiness(db)
	userBusiness.Create(flags.UserName, flags.Password)
	println("User", flags.UserName, "successfully created")
}
