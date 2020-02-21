package business

import (
	"bytes"
	"crypto/rand"
	"database/sql"
	"errors"
	"flying-castle/castle"
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
	"time"
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
	file, err := os.Open("../test_sql/test_fetch_file.sql")
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

func TestFileBusiness_Integration(t *testing.T) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db2.SetDB("sqlite3", db)
	storageKey, err := encryption.GenerateKey()
	Before(db, storageKey)
	dataPath := strings.ReplaceAll(path.Join(os.TempDir(), "flying_castle"), "\\", "/")
	fsBackend, _ := castle.NewFSBackend(dataPath)
	castle.SetStorageBackend(fsBackend)
	var fb = NewDBFileBusiness()
	data := make([]byte, 100*1000*1000)
	rand.Read(data)
	err = fb.Create(1, model.File{
		FileSystemEntity: model.FileSystemEntity{
			Id:   0,
			Name: "file2",
		},
		DataHolder:         model.DataHolder{Data: data},
		Size:               uint64(len(data)),
		Parent:             nil,
		AccessTime:         time.Now(),
		MetadataChangeTime: time.Now(),
		DataChangeTime:     time.Now(),
		Owner:              nil,
		Group:              nil,
		Kind:               model.RegularFile,
		UserPermissions:    nil,
		GroupPermissions:   nil,
	})
	newFile, err := fb.FindByUserAndPath(1, "file2")
	if err != nil {
		t.Fatal(err)
	}
	*newFile, err = fb.GetFileById(int64(newFile.Id))
	if err != nil {
		t.Fatal("could not get file")
	}
	if len(data) != len(newFile.Data) || !bytes.Equal(data, newFile.Data) {
		println(data)
		println(newFile.Data)
		t.Fatal("stored file is different than original")
	}
	_ = os.RemoveAll(dataPath)
}

func BenchmarkStoreFile(b *testing.B) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		panic(err)
	}
	db2.SetDB("sqlite3", db)
	storageKey, err := encryption.GenerateKey()
	Before(db, storageKey)
	dataPath := strings.ReplaceAll(path.Join(os.TempDir(), "flying_castle"), "\\", "/")
	fsBackend, _ := castle.NewFSBackend(dataPath)
	castle.SetStorageBackend(fsBackend)
	benchmarks := []struct {
		name     string
		fileSize int64
		data     []byte
	}{
		{name: "10 bytes", fileSize: int64(10)},
		{name: "100 bytes", fileSize: int64(100)},
		{name: "1 Kbytes", fileSize: int64(1000)},
		{name: "100 Kbytes", fileSize: int64(100 * 1000)},
		{name: "1 Mbytes", fileSize: int64(1000 * 1000)},
		{name: "10 Mbytes", fileSize: int64(10 * 1000 * 1000)},
		{name: "100 Mbytes", fileSize: int64(100 * 1000 * 1000)},
		{name: "1 Gbytes", fileSize: int64(1000 * 1000 * 1000)},
		{name: "10 Gbytes", fileSize: int64(10 * 1000 * 1000 * 1000)},
	}
	for _, b := range benchmarks {
		data := make([]byte, b.fileSize)
		rand.Read(data)
		b.data = data
	}
	fb := NewDBFileBusiness()
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Before(db, storageKey)
				fb.Create(1, model.File{
					FileSystemEntity: model.FileSystemEntity{
						Id:   0,
						Name: "file2",
					},
					DataHolder:         model.DataHolder{Data: bm.data},
					Size:               uint64(bm.fileSize),
					Parent:             nil,
					AccessTime:         time.Now(),
					MetadataChangeTime: time.Now(),
					DataChangeTime:     time.Now(),
					Owner:              nil,
					Group:              nil,
					Kind:               model.RegularFile,
					UserPermissions:    nil,
					GroupPermissions:   nil,
				})
			}
		})
		_ = os.RemoveAll(dataPath)
	}
}
