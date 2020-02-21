package business

import (
	"bytes"
	"database/sql"
	"flying-castle/castle"
	"flying-castle/db"
	"flying-castle/db/dao"
	"flying-castle/encryption"
	"flying-castle/model"
	"flying-castle/utils"
	"github.com/jmoiron/sqlx"
	"io/ioutil"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
)

type FileBusiness struct {
	db *sqlx.DB
}

func NewDBFileBusiness() FileBusiness {
	return FileBusiness{db: sqlx.NewDb(db.GetDB().GetDB(), db.GetDB().GetDriver())}
}

func (fileBusiness *FileBusiness) MustGetFileById(id int64) model.File {
	file, err := fileBusiness.GetFileById(id)
	if err != nil {
		panic(err)
	}
	return file
}

func (fileBusiness *FileBusiness) GetFileById(id int64) (returnFile model.File, returnErr error) {
	var tx = fileBusiness.db.MustBegin()
	var fileRepo = dao.NewFileRepository(tx)
	var chunkRepo = dao.NewChunkRepository(tx)
	var file = model.File{}
	var fileDAO = fileRepo.GetById(id)
	var hasNextChunk = true
	var nextChunkId = fileDAO.FirstChunkId
	var chunksToDecrypt = make([]dao.ChunkDAO, 0)
	var wg sync.WaitGroup
	for i := 0; hasNextChunk; i++ {
		var nextChunk = chunkRepo.GetById(nextChunkId)
		chunksToDecrypt = append(chunksToDecrypt, nextChunk)
		nextChunkId = nextChunk.NextChunk.Int64
		hasNextChunk = nextChunk.NextChunk.Valid
	}
	wg.Add(len(chunksToDecrypt))
	var chunks = make([][]byte, len(chunksToDecrypt))
	for i, nextChunk := range chunksToDecrypt {
		go func(i int, nextChunk dao.ChunkDAO) {
			defer wg.Done()
			var file, err = os.Open(nextChunk.Path)
			if err != nil {
				returnErr = model.ContentError
				return
			}
			chunkBytes, err := ioutil.ReadAll(file)
			if err != nil {
				returnErr = model.ContentError
				return
			}
			var decodedBytes = encryption.Decrypt(chunkBytes)
			chunks[i] = decodedBytes
		}(i, nextChunk)
	}
	wg.Wait()
	file.Data = utils.Flatten(chunks)
	var pathRepo = dao.NewPathRepository(tx)
	var pathDAO = pathRepo.GetById(fileDAO.PathId)
	file.Name = pathDAO.Name
	err := tx.Commit()
	if err != nil {
		return file, model.DatabaseError
	}
	return file, nil
}

func (fileBusiness *FileBusiness) FindByUserAndPath(userId int64, path string) (*model.File, error) {
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
			return nil, model.FileNotFound
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
			return nil, model.FileNotFound
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
		return nil, model.DatabaseError
	}
	return &file, nil
}

func (fileBusiness *FileBusiness) Create(parent int64, file model.File) error {
	var tx = fileBusiness.db.MustBegin()
	var fileRepo = dao.NewFileRepository(tx)
	var chunkRepo = dao.NewChunkRepository(tx)
	buffer := bytes.NewBuffer(utils.Reverse(file.Data))
	var minChunkSize = math.Min(math.Pow(2, 20), float64(len(file.Data)))
	var nextChunkId = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	var chunksToEncrypt = make([]struct {
		chunk   []byte
		chunkId int64
	}, 0)
	for buffer.Len() > 0 {
		lastChunkId, err := db.GetLastIdFrom(tx, "chunk")
		if err != nil {
			return model.DatabaseError
		}
		var chunkSize = rand.Intn(len(file.Data) / 4)
		chunkSize = int(math.Max(float64(chunkSize), minChunkSize))
		var chunk = buffer.Next(chunkSize)
		chunksToEncrypt = append(chunksToEncrypt, struct {
			chunk   []byte
			chunkId int64
		}{chunk: chunk, chunkId: lastChunkId.Int64 + 1})
		chunkDAO := dao.ChunkDAO{
			Id:        lastChunkId.Int64 + 1,
			Path:      "",
			NextChunk: nextChunkId,
		}
		chunkRepo.Create(chunkDAO)
		nextChunkId = sql.NullInt64{
			Int64: lastChunkId.Int64 + 1,
			Valid: true,
		}
	}
	var wg sync.WaitGroup
	wg.Add(len(chunksToEncrypt))
	for _, chunkToEncrypt := range chunksToEncrypt {
		go func(chunk []byte, chunkId int64) {
			defer wg.Done()
			var encryptedChunk = encryption.Encrypt(utils.Reverse(chunk))
			fileName := strconv.FormatInt(chunkId, 10)
			path, err := castle.WriteChunk(fileName, encryptedChunk)
			if err != nil {
				panic(model.SaveFileError)
			}
			chunkRepo.UpdatePath(chunkId, path)
		}(chunkToEncrypt.chunk, chunkToEncrypt.chunkId)
	}
	wg.Wait()
	_, err := fileRepo.Create(parent, file.Name, nextChunkId.Int64, int64(len(file.Data)))
	if err != nil {
		return model.SaveFileError
	}
	return tx.Commit()
}
