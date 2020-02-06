package main

import (
	"flying-castle/business"
	"flying-castle/cmd"
	"flying-castle/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"net/url"
	"os"
	"path"
)

type FileFlags struct {
	UserName string `flag:"name" required:"true"`
	Password string `flag:"password"`
	ApiKey   string `flag:"key"`
	//cmd.CredentialFlags
	FilePath string `flag:"path" required:"true"`
	Output   string `flag:"output"`
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

	//TODO: if output is non-exisiting file path in existing folder recursive creation (flag?)

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
	var file = fileBusiness.FindByUserAndPath(int64(user.Id), flags.FilePath)
	if file.Kind != model.RegularFile {
		panic("can't fetch this kind of file")
	}
	*file, err = fileBusiness.GetFileById(int64(file.Id))
	if err != nil {
		panic(err)
	}
	outputPath := path.Join(flags.Output, file.Name)
	_, err = os.Stat(outputPath)
	if err == nil {
		panic("output file already exists")
	}
	f, err := os.Create(outputPath)
	if err != nil {
		panic("can't create output file")
	}
	_, err = f.Write(file.Data)
	if err != nil {
		panic("can't write to output file")
	}
	_ = f.Close()
}
