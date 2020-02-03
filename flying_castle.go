package main

import (
	"flying-castle/migrations"
	"fmt"
)

//go:generate go-bindata -o requests/requests.go requests/...

func main() {
	err := migrations.Migrate("sqlite3://fc.db", "file://F:/Users/Rafael/go/src/flying-castle/migrations")
	if err != nil {
		panic(fmt.Errorf("migration error %s", err.Error()))
	}
}
