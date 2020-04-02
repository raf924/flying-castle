package main

import (
	_ "bytes"
	"flying-castle/app"
	"flying-castle/business"
	"flying-castle/db"
	_ "github.com/mattn/go-sqlite3"
)

type UserFlags struct {
	UserName string `flag:"name" required:"true" type:"alphanum"`
	Password string `flag:"password" required:"true"`
}

func (u *UserFlags) Validate() {
}

func createUser(config *app.Config, flags UserFlags) error {
	err := db.LoadDB(config.DbUrl)
	if err != nil {
		return db.ConnectionError
	}
	var userBusiness = business.NewDBUserBusiness()
	return userBusiness.Create(flags.UserName, flags.Password)
}

func main() {
	var config = app.GetConfig()
	var flags = UserFlags{}
	app.ReadFlags(&flags)
	var err = createUser(config, flags)
	if err != nil {
		println("error while creating user", flags.UserName)
		panic(err)
	}
	println("User", flags.UserName, "successfully created")
}
