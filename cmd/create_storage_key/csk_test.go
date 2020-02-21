package main

import (
	"database/sql"
	"errors"
	"flying-castle/cmd"
	db2 "flying-castle/db"
	"flying-castle/migrations"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"path"
	"strings"
	"testing"
)

func Test_createStorageKey(t *testing.T) {
	config := &cmd.Config{
		DbUrl:    "sqlite3://:memory:",
		DataPath: "C:/flying_castle",
	}
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db2.SetDB("sqlite3", db)
	filenames, _ := migrations.AssetDir("migrations")
	s := bindata.Resource(filenames, func(name string) (bytes []byte, err error) {
		if strings.Contains(name, ".sql") {
			return migrations.Asset(path.Join("migrations", name))
		} else {
			return nil, errors.New("unusable file")
		}
	})
	d, err := bindata.WithInstance(s)
	if err != nil {
		panic(err)
	}
	err = migrations.MigrateWithDBAndSource("sqlite3", db, "migrations", d)
	if err != nil {
		panic(err)
	}

	type args struct {
		config *cmd.Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "NoError", args: args{config: config}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := createStorageKey(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("createStorageKey() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
