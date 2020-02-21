package main

import (
	"database/sql"
	"errors"
	"flying-castle/cmd"
	db2 "flying-castle/db"
	"flying-castle/encryption"
	"flying-castle/migrations"
	"flying-castle/model"
	bindata "github.com/golang-migrate/migrate/source/go_bindata"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"
)

func Before(db *sql.DB, storageKey []byte) {
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
		return
	}
	file, err := os.Open("../../test_sql/test_fetch_file.sql")
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	req := string(b)
	dataPath := strings.ReplaceAll(path.Join(os.TempDir(), "flying_castle"), "\\", "/")
	outputPath := strings.ReplaceAll(path.Join(os.TempDir(), "flying_castle_output"), "\\", "/")
	_ = os.Mkdir(dataPath, os.ModeDir)
	_ = os.Mkdir(outputPath, os.ModeDir)
	dbx := sqlx.NewDb(db, "sqlite3")
	var tx = dbx.MustBegin()
	_ = tx.MustExec(req,
		encryption.EncodeKey(storageKey),
		path.Join(dataPath, "1"),
		path.Join(dataPath, "2"),
		encryption.EncodeKey(encryption.MustHash("password", storageKey)),
		encryption.EncodeKey(storageKey))
	_ = tx.Commit()
	encryption.MustUpdateKey(encryption.EncodeKey(storageKey))
}

func Test_fetchFile(t *testing.T) {
	config := &cmd.Config{
		DbUrl:    "sqlite3://:memory:",
		DataPath: path.Join(os.TempDir(), "flying_castle"),
	}
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db2.SetDB("sqlite3", db)
	storageKey, err := encryption.GenerateKey()
	Before(db, storageKey)
	dataPath := strings.ReplaceAll(path.Join(os.TempDir(), "flying_castle"), "\\", "/")
	outputPath := strings.ReplaceAll(path.Join(os.TempDir(), "flying_castle_output"), "\\", "/")
	defer func() {
		_ = os.RemoveAll(dataPath)
		_ = os.RemoveAll(outputPath)
	}()
	file1, err := os.Create(path.Join(dataPath, "1"))
	file2, err := os.Create(path.Join(dataPath, "2"))
	_, _ = file1.Write(encryption.Encrypt([]byte("hell")))
	_, _ = file2.Write(encryption.Encrypt([]byte("lo")))
	_ = file1.Close()
	_ = file2.Close()
	invalidOutput, _ := os.Create(path.Join(outputPath, "invalid"))
	_ = invalidOutput.Close()
	type args struct {
		config *cmd.Config
		flags  FileFlags
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		expectedErr error
	}{
		{
			name: "NoError", args: args{
				config: config,
				flags: FileFlags{
					UserName: "rafael",
					Password: "password",
					ApiKey:   "",
					FilePath: "file1",
					Output:   outputPath,
				},
			}, wantErr: false},
		{
			name: "UserNotFound",
			args: args{
				config: config,
				flags: FileFlags{
					UserName: "raf",
					Password: "password",
					ApiKey:   "",
					FilePath: "file1",
					Output:   outputPath,
				},
			},
			wantErr:     true,
			expectedErr: model.InvalidCredentials,
		},
		{
			name: "IncorrectPassword",
			args: args{
				config: config,
				flags: FileFlags{
					UserName: "rafael",
					Password: "password1",
					ApiKey:   "",
					FilePath: "file1",
					Output:   outputPath,
				},
			},
			wantErr:     true,
			expectedErr: model.InvalidCredentials,
		},
		{
			name: "FileNotFoundError",
			args: args{
				config: config,
				flags: FileFlags{
					UserName: "rafael",
					Password: "password",
					ApiKey:   "",
					FilePath: "file2",
					Output:   outputPath,
				},
			},
			expectedErr: model.FileNotFound,
			wantErr:     true,
		},
		{
			name: "InvalidOutput",
			args: args{
				config: config,
				flags: FileFlags{
					UserName: "rafael",
					Password: "password",
					ApiKey:   "",
					FilePath: "file1",
					Output:   path.Join(outputPath, "invalid"),
				},
			},
			expectedErr: cmd.NotCreatableError,
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Before(db, storageKey)
			if err := fetchFile(tt.args.config, tt.args.flags); (err != nil) != tt.wantErr {
				t.Errorf("fetchFile() error = %v, wantErr %v", err, tt.wantErr)
			} else if err != tt.expectedErr {
				t.Errorf("fetchFile() error = %v, expectedErr %v", err, tt.expectedErr)
			}
		})
	}
}
