package filesService

import (
	"context"
	"errors"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"os"
)

type FilesService struct {
	db *database.Queries
	c  *config.Config
}

func NewFilesService(p *models.DBConfigParam) *FilesService {
	return &FilesService{p.DB, p.C}
}

func (fs *FilesService) UploadFile(c context.Context, file database.CreateFileParams) (string, error) {

	result, err := fs.db.CreateFile(c, file)
	if err != nil {
		return "", err
	}
	return fs.c.Server.BaseUrl + constants.ENDPOINT_V1 + "/files/" + result.FileName, nil
}

func (fs *FilesService) FindFile(c context.Context, fileName string) (database.FindFileByFileNameRow, error) {
	return fs.db.FindFileByFileName(c, fileName)
}

func (fs *FilesService) DeleteFile(c context.Context, fileName string) error {
	file, err := fs.FindFile(c, fileName)
	if err != nil {
		return errors.New("file not found")
	}

	_, err = os.Stat(file.Path)
	if os.IsNotExist(err) {
		return errors.New("file not found")
	}

	err = os.Remove(file.Path)
	if err != nil {
		return errors.New("could not delete file")
	}

	return fs.db.DeleteFileById(c, file.ID)
}
