package business

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"flying-castle/db/dao"
	"flying-castle/model"
	"github.com/jmoiron/sqlx"
	"io"
	"io/ioutil"
	"os"
	"time"
)

type FileBusiness struct {
	db *sqlx.DB
}

func ModelMapper(dao dao.FileDAO) model.File {
	var file = model.File{}
	file.Id = uint64(dao.Id)
	file.Name = dao.Name
	file.Size = uint64(dao.Size)
	file.AccessTime = time.Now()
	file.DataChangeTime = dao.ModifiedAt
	file.MetadataChangeTime = dao.ModifiedAt
	return file
}

func (fileBusiness *FileBusiness) MustGetFileById(id int64) model.File {
	file, err := fileBusiness.GetFileById(id)
	if err != nil {
		panic(err)
	}
	return file
}

func Decrypt(packet []byte, storageKey string) []byte {
	key, err := base64.StdEncoding.DecodeString(storageKey)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	nonce := packet[:12]
	data := packet[12:]

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	decrypted, err := aesGCM.Open(nil, nonce, data, nil)
	if err != nil {
		panic(err)
	}
	return decrypted
}

func Encrypt(packet []byte, storageKey string) []byte {
	key, err := base64.StdEncoding.DecodeString(storageKey)
	if err != nil {
		panic(err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err)
	}
	nonce := make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err)
	}
	encrypted := aesGCM.Seal(nil, nonce, packet, nil)
	return append(nonce, encrypted...)
}

func (fileBusiness *FileBusiness) GetFileById(id int64) (model.File, error) {
	var tx = fileBusiness.db.MustBegin()
	var fileRepo = dao.NewFileRepository(tx)
	var chunkRepo = dao.NewChunkRepository(tx)
	var storageKeyRepo = dao.NewStorageKeyRepository(tx)
	var file = model.File{}
	var fileDAO = fileRepo.GetById(id)
	var storageKeyDAO = storageKeyRepo.GetLatest()
	var data = make([]byte, 0)
	var hasNextChunk = true
	var nextChunkId = fileDAO.FirstChunkId
	for hasNextChunk {
		var nextChunk = chunkRepo.GetById(nextChunkId)
		var file, err = os.Open(nextChunk.Path)
		if err != nil {
			panic(err)
		}
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		var decodedBytes = Decrypt(bytes, storageKeyDAO.Key)
		data = append(data, decodedBytes...)
		nextChunkId = nextChunk.NextChunk.Int64
		hasNextChunk = nextChunk.NextChunk.Valid
	}
	file.DataHolder.Data = data
	return file, nil
}

func (fileBusiness *FileBusiness) GetFileByPathId(id int64) (dao.FileDAO, error) {
	var file = dao.FileDAO{}
	var err = fileBusiness.db.Get(&file, "SELECT * FROM file WHERE path_id = ?")
	return file, err
}
