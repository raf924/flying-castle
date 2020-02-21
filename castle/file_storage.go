package castle

import (
	"errors"
	"flying-castle/cmd"
	"net/url"
	"sync"
)

var chunkWriter ChunkWriter
var chunkReader ChunkReader

type StorageBackend interface {
	ChunkWriter
	ChunkReader
}

var storage StorageBackend

var StorageUrlError = errors.New("storage url is invalid")
var UnknownStorageProtocolError = errors.New("storage protocol is unsupported")

type BackendConstructor func(path string, config *cmd.Config) (StorageBackend, error)

var constructors = make(map[string]BackendConstructor)

func NewWriter(uri string, config *cmd.Config) error {
	storageUri, err := url.Parse(uri)
	if err != nil {
		return StorageUrlError
	}
	if constructor, ok := constructors[storageUri.Scheme]; ok {
		storage, err = constructor(storageUri.EscapedPath(), config)
		if err != nil {
			return err
		}
	} else {
		return UnknownStorageProtocolError
	}
	return nil
}

//Wrapper around the global ChunkWriter
func WriteChunk(fileName string, chunkData []byte) (string, error) {
	return storage.Write(fileName, chunkData)
}

//Wrapper around the global ChunkReader
func ReadChunk(fileName string) ([]byte, error) {
	return storage.Read(fileName)
}

type ChunkWriter interface {
	//Write a chunk of data to the file storage backend
	//Returns the absolute identifying value for the stored element
	Write(fileName string, chunkData []byte) (string, error)
}

type ChunkReader interface {
	//Read a chunk of data from the file storage backend given an absolute identifying value
	//Returns the chunk of data
	Read(fileName string) ([]byte, error)
}

var once sync.Once

func SetStorageBackend(backend StorageBackend) {
	once.Do(func() {
		storage = backend
	})
}
