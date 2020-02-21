package migrations

import (
	"database/sql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/database/postgres"
	"github.com/golang-migrate/migrate/database/sqlite3"
	"github.com/golang-migrate/migrate/source"
)
import _ "github.com/golang-migrate/migrate/source/file"
import _ "github.com/golang-migrate/migrate/database/sqlite3"

func Migrate(db string, path string) error {
	m, err := migrate.New(path, db)
	if err != nil {
		return err
	}
	return m.Up()
}

func MigrateWithDBAndSource(dbDriver string, db *sql.DB, source string, sourceDriver source.Driver) error {
	var driver database.Driver
	var err error
	switch dbDriver {
	case "sqlite3":
		driver, err = sqlite3.WithInstance(db, &sqlite3.Config{})
		break
	case "postgres":
		driver, err = postgres.WithInstance(db, &postgres.Config{})
		break
	case "mysql":
		driver, err = mysql.WithInstance(db, &mysql.Config{})
		break
	}
	if err != nil {
		return err
	}
	m, err := migrate.NewWithInstance(source, sourceDriver, "flying_castle", driver)
	if err != nil {
		return err
	}
	return m.Up()
}
