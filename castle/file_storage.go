package castle

import (
	"errors"
	"flying-castle/app"
	"log"
	"net/url"
	"sync"
)

type StorageBackend interface {
	ChunkWriter
	ChunkReader
	ChunkDeleter
}

var storage StorageBackend

var StorageUrlError = errors.New("storage url is invalid")
var UnknownStorageProtocolError = errors.New("storage protocol is unsupported")

type BackendConstructor func(path string, config *app.Config) (StorageBackend, error)

var constructors = make(map[string]BackendConstructor)

func NewBackend(uri string, config *app.Config) (StorageBackend, error) {
	var backend StorageBackend
	storageUri, err := url.Parse(uri)
	if err != nil {
		return nil, StorageUrlError
	}
	if constructor, ok := constructors[storageUri.Scheme]; ok {
		backend, err = constructor(storageUri.EscapedPath(), config)
		if err != nil {
			return nil, err
		}
	} else {
		log.Println(UnknownStorageProtocolError.Error(), ":", storageUri.Scheme)
		return nil, UnknownStorageProtocolError
	}
	return backend, nil
}

//Wrapper around the global ChunkWriter
func WriteChunk(fileName string, chunkData []byte) (string, error) {
	return storage.Write(fileName, chunkData)
}

//Wrapper around the global ChunkReader
func ReadChunk(fileName string) ([]byte, error) {
	return storage.Read(fileName)
}

func DeleteChunks(keys []string) error {
	return storage.Delete(keys)
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

type ChunkDeleter interface {
	Delete(fileNames []string) error
}

var once sync.Once

func SetStorageBackend(backend StorageBackend) {
	once.Do(func() {
		storage = backend
	})
}
