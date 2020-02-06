package business

import (
	"flying-castle/db/dao"
	"flying-castle/encryption"
	"flying-castle/model"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"os"
	"strings"
)

type FileBusiness struct {
	db *sqlx.DB
}

func NewFileBusiness(db *sqlx.DB) FileBusiness {
	return FileBusiness{db: db}
}

func (fileBusiness *FileBusiness) MustGetFileById(id int64) model.File {
	file, err := fileBusiness.GetFileById(id)
	if err != nil {
		panic(err)
	}
	return file
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
		var decodedBytes = encryption.Decrypt(bytes, encryption.MustDecodeKey(storageKeyDAO.Key))
		data = append(data, decodedBytes...)
		nextChunkId = nextChunk.NextChunk.Int64
		hasNextChunk = nextChunk.NextChunk.Valid
	}
	var pathRepo = dao.NewPathRepository(tx)
	var pathDAO = pathRepo.GetById(fileDAO.PathId)
	file.Name = pathDAO.Name
	file.DataHolder.Data = data
	return file, tx.Commit()
}

func (fileBusiness *FileBusiness) FindByUserAndPath(userId int64, path string) *model.File {
	var tx = fileBusiness.db.MustBegin()
	var pathNames = strings.Split(path, "/")
	var isAbsolute = len(path) > 0 && path[0] == '/'
	var userRepo = dao.NewUserRepository(tx)
	var groupRepo = dao.NewGroupRepository(tx)
	var userDAO = userRepo.GetById(userId)
	var groupDAO = groupRepo.GetById(userDAO.MainGroupId)
	var group = model.Group{
		FileSystemEntity: model.FileSystemEntity{
			Id:   uint64(groupDAO.Id),
			Name: groupDAO.Name,
		},
		Users: []*model.User{},
	}
	var user = model.User{
		FileSystemEntity: model.FileSystemEntity{},
		Group:            &group,
	}
	group.Users = append(group.Users, &user)
	var folderRepo = dao.NewFolderRepository(tx)
	var rootPath dao.FolderDAO
	if pathNames[0] == "" {
		pathNames = pathNames[1:]
	}
	if isAbsolute {
		//TODO
	} else {
		rootPath = folderRepo.GetUserRoot(userId)
	}
	var pathRepo = dao.NewPathRepository(tx)
	var currentPath = pathRepo.GetById(rootPath.Id)
	for len(pathNames) > 0 {
		var path = pathRepo.FindByParentAndName(currentPath.Id, pathNames[0])
		if path == nil {
			panic("Given path does not exist")
		}
		currentPath = *path
		pathNames = pathNames[1:]
	}
	var folder = dao.FolderDAO{}
	var fileDAO = dao.FileDAO{}
	var kind model.FileKind
	var fileId int64
	folder, err := folderRepo.FindByPathId(currentPath.Id)
	if err != nil {
		var fileRepo = dao.NewFileRepository(tx)
		fileDAO, err = fileRepo.FindByPathId(currentPath.Id)
		if err != nil {
			panic(err)
		}
		fileId = fileDAO.Id
		kind = model.RegularFile
	} else {
		fileId = folder.Id
		kind = model.Directory
	}
	var file = model.File{
		FileSystemEntity: model.FileSystemEntity{
			Id:   uint64(fileId),
			Name: currentPath.Name,
		},
		DataHolder:         model.DataHolder{},
		Size:               0,
		Parent:             nil,
		AccessTime:         currentPath.AccessedAt,
		MetadataChangeTime: currentPath.ModifiedAt,
		DataChangeTime:     currentPath.CreatedAt,
		Owner:              &user,
		Group:              &group,
		Kind:               kind,
		UserPermissions:    nil,
		GroupPermissions:   nil,
	}
	//TODO: fill permissions
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
	return &file
}

func (fileBusiness *FileBusiness) Create(parent int64, file model.File) error {
	var tx = fileBusiness.db.MustBegin()
	var fileRepo = dao.NewFileRepository(tx)
	var storageKeyRepo = dao.NewStorageKeyRepository(tx)
	var latestKey = storageKeyRepo.GetLatest()
	fileRepo.Save(parent, file.Name, file.Data, latestKey.Key)
	return tx.Commit()
}
