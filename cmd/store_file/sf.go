package main

import (
	"flying-castle/business"
	"flying-castle/castle"
	"flying-castle/cmd"
	"flying-castle/db"
	"flying-castle/model"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"os"
)

type FileFlags struct {
	UserName    string `flag:"name" required:"true"`
	Password    string `flag:"password"`
	ApiKey      string `flag:"key"`
	FilePath    string `flag:"path" required:"true"`
	Destination string `flag:"destination"`
	Public      bool   `flag:"public"`
}

func (f *FileFlags) Validate() {
	if f.Password == "" && f.ApiKey == "" {
		panic("no valid credentials were given")
	}
}

func storeFile(config *cmd.Config, flags FileFlags) error {
	realFile, err := os.Open(flags.FilePath)
	if err != nil {
		return cmd.FileNotFoundError
	}
	stat, err := realFile.Stat()
	if err != nil {
		return cmd.UnreadableFileError
	}
	data, err := ioutil.ReadAll(realFile)
	if err != nil {
		return cmd.UnreadableFileError
	}
	lastTime := stat.ModTime()

	err = db.LoadDB(config.DbUrl)
	if err != nil {
		return err
	}
	fsBackend, _ := castle.NewFSBackend(config.DataPath)
	castle.SetStorageBackend(fsBackend)
	//TODO: create path if valid but doesn't exist (with flag?)

	var fileBusiness = business.NewDBFileBusiness()
	var userBusiness = business.NewDBUserBusiness()
	user, err := userBusiness.FindUserByNameAndPassword(flags.UserName, flags.Password)
	if err != nil {
		return err
	}
	folder, err := fileBusiness.FindByUserAndPath(int64(user.Id), flags.Destination)
	if err != nil {
		return err
	}
	if folder.Kind != model.Directory {
		return model.WrongFileKind
	}
	var file = model.File{
		FileSystemEntity: model.FileSystemEntity{
			Id:   0,
			Name: stat.Name(),
		},
		DataHolder:         model.DataHolder{Data: data},
		Size:               uint64(stat.Size()),
		Parent:             nil,
		AccessTime:         lastTime,
		MetadataChangeTime: lastTime,
		DataChangeTime:     lastTime,
		Owner:              nil,
		Group:              nil,
		Kind:               model.RegularFile,
		UserPermissions:    nil,
		GroupPermissions:   nil,
	}
	return fileBusiness.Create(int64(folder.Id), file)
}

func main() {
	var config = cmd.GetConfig()
	var flags = FileFlags{}
	cmd.ReadFlags(&flags)
	var err = storeFile(config, flags)
	if err != nil {
		panic(err)
	}
	println("File successfully saved")
	//TODO print access link
}
