package main

import (
	"errors"
	"flying-castle/business"
	"flying-castle/cmd"
	"flying-castle/db"
	"flying-castle/model"
	_ "github.com/mattn/go-sqlite3"
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

func fetchFile(config *cmd.Config, flags FileFlags) error {
	err := os.MkdirAll(flags.Output, os.ModeDir)
	if err != nil {
		return cmd.NotCreatableError
	}
	err = db.LoadDB(config.DbUrl)
	if err != nil {
		return err
	}
	var fileBusiness = business.NewDBFileBusiness()
	var userBusiness = business.NewDBUserBusiness()
	user, err := userBusiness.FindUserByNameAndPassword(flags.UserName, flags.Password)
	if err != nil {
		return err
	}
	file, err := fileBusiness.FindByUserAndPath(int64(user.Id), flags.FilePath)
	if err != nil {
		return err
	}
	//TODO: if output is non-exisiting file path in existing folder recursive creation (flag?)
	if file.Kind != model.RegularFile {
		return model.WrongFileKind
	}
	*file, err = fileBusiness.GetFileById(int64(file.Id))
	if err != nil {
		return err
	}
	outputPath := path.Join(flags.Output, file.Name)
	_, err = os.Stat(outputPath)
	if err == nil {
		return errors.New("output file already exists")
	}
	f, err := os.Create(outputPath)
	if err != nil {
		return errors.New("can't create output file")
	}
	_, err = f.Write(file.Data)
	if err != nil {
		return errors.New("can't write to output file")
	}
	_ = f.Close()
	return nil
}

func main() {
	var config = cmd.GetConfig()
	var flags = FileFlags{}
	cmd.ReadFlags(&flags)
	var err = fetchFile(config, flags)
	if err != nil {
		println("error while fetching %s", flags.FilePath)
	}
}
