package castle

import (
	"flying-castle/app"
	"io/ioutil"
	"os"
	filepath "path"
)

type FSBackend struct {
	dataPath string
}

func (F FSBackend) Delete(fileNames []string) error {
	for _, fileName := range fileNames {
		err := os.Remove(fileName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (F FSBackend) Read(fileName string) ([]byte, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return ioutil.ReadAll(file)
}

func init() {
	constructors["file"] = func(path string, config *app.Config) (StorageBackend, error) {
		return NewFSBackend(path)
	}
}

func NewFSBackend(dataPath string) (StorageBackend, error) {
	return FSBackend{dataPath: dataPath}, os.MkdirAll(dataPath, os.ModeDir)
}

func (F FSBackend) Write(fileName string, chunkData []byte) (string, error) {
	path := filepath.Join(F.dataPath, fileName)
	chunkFile, err := os.Create(path)
	if err != nil {
		return "", err
	}
	_, err = chunkFile.Write(chunkData)
	return path, err
}
