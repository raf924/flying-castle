package main

import (
	"flying-castle/business"
	"flying-castle/cmd"
	"flying-castle/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"io/ioutil"
	"net/url"
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

func main() {
	var config = cmd.GetConfig()
	var dbUrl, err = url.Parse(config.DbUrl)
	if err != nil {
		panic(err)
	}
	var flags = FileFlags{}
	cmd.ReadFlags(&flags)

	db, err := sqlx.Connect(dbUrl.Scheme, dbUrl.EscapedPath())
	if err != nil {
		panic(err)
	}
	var fileBusiness = business.NewFileBusiness(db)
	var userBusiness = business.NewUserBusiness(db)
	user, err := userBusiness.FindUserByNameAndPassword(flags.UserName, flags.Password)
	if err != nil {
		panic(err)
	}
	var folder = fileBusiness.FindByUserAndPath(int64(user.Id), flags.Destination)
	if folder.Kind != model.Directory {
		panic("Destination is not a valid folder")
	}
	//TODO: create path if valid but doesn't exist (with flag?)
	realFile, err := os.Open(flags.FilePath)
	if err != nil {
		panic(err)
	}
	stat, err := realFile.Stat()
	if err != nil {
		panic(err)
	}
	data, err := ioutil.ReadAll(realFile)
	if err != nil {
		panic(err)
	}
	lastTime := stat.ModTime()
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
	err = fileBusiness.Create(int64(folder.Id), file)
	if err != nil {
		panic(err)
	}
	println("File successfully saved")
	//TODO print access link
}
