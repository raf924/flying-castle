package main

import (
	"flying-castle/migrations"
	"fmt"
)

//go:generate go-bindata -o requests/requests.go requests/...

func main() {
	err := migrations.Migrate("sqlite3://fc.db", "file://migrations")
	if err != nil {
		panic(fmt.Errorf("migration error %s", err.Error()))
	}
}
