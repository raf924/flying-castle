package migrations

import "github.com/golang-migrate/migrate"
import _ "github.com/golang-migrate/migrate/source/file"
import _ "github.com/golang-migrate/migrate/database/sqlite3"

func Migrate(db string, path string) error {
	m, err := migrate.New(path, db)
	if err != nil {
		return err
	}
	return m.Up()
}
