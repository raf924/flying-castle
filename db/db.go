package db

import (
	"database/sql"
	"net/url"
	"sync"
	"time"
)

type DB struct {
	db     *sql.DB
	driver string
}

func (db *DB) GetDB() *sql.DB {
	return db.db
}

func (db *DB) GetDriver() string {
	return db.driver
}

var db DB

func GetDB() *DB {
	return &db
}

var once sync.Once
var timer *time.Ticker

func SetDB(driver string, database *sql.DB) {
	once.Do(func() {
		db = DB{
			db:     database,
			driver: driver,
		}
		if timer != nil {
			timer.Stop()
		}
		timer = time.NewTicker(time.Second)
		go func() {
			for range timer.C {
				err := database.Ping()
				if err != nil {
					panic(err)
				}
			}
		}()
	})
}

func LoadDB(databaseURL string) error {
	var dbUrl, err = url.Parse(databaseURL)
	if err != nil {
		if databaseURL == "sqlite3://:memory:" {
			dbUrl = &url.URL{
				Scheme:     "sqlite3",
				Opaque:     "",
				User:       nil,
				Host:       "",
				Path:       ":memory:",
				RawPath:    ":memory:",
				ForceQuery: false,
				RawQuery:   "",
				Fragment:   "",
			}
		} else {
			panic(err)
		}
	}
	db, err := sql.Open(dbUrl.Scheme, dbUrl.EscapedPath())
	if err != nil {
		return err
	}
	SetDB(dbUrl.Scheme, db)
	return nil
}
