package castle

import "github.com/jmoiron/sqlx"

type DBBackend struct {
	db *sqlx.DB
}

func (D DBBackend) Write(fileName string, chunkData []byte) (string, error) {
	panic("implement me")
}

func (D DBBackend) Read(fileName string) ([]byte, error) {
	panic("implement me")
}

func NewDBBackend(database string) {

}
