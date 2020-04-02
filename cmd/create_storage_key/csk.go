package main

import (
	_ "bytes"
	"flying-castle/app"
	"flying-castle/business"
	"flying-castle/db"
	_ "github.com/mattn/go-sqlite3"
)

type StorageKeyFlags struct {
}

func (s *StorageKeyFlags) Validate() {
}

func createStorageKey(config *app.Config) error {
	err := db.LoadDB(config.DbUrl)
	if err != nil {
		return err
	}
	var skBusiness = business.NewDBStorageKeyBusiness()
	return skBusiness.Create()
}

func main() {
	var config = app.GetConfig()
	app.ReadFlags(&StorageKeyFlags{})
	var err = createStorageKey(config)
	if err != nil {
		println("error while creating storage key")
		panic(err)
	}
	println("Storage key successfully created")
}
