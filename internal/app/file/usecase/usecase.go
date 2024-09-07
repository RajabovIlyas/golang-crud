package usecase

import (
	"context"
	"errors"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/file"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"os"
)

type fileUC struct {
	cfg      *config.Config
	fileRepo file.Repository
}

func NewFileUseCase(cfg *config.Config, fileRepo file.Repository) file.UseCase {
	return &fileUC{cfg: cfg, fileRepo: fileRepo}
}

func (f fileUC) UploadFile(ctx context.Context, params database.CreateFileParams) (string, error) {
	result, err := f.fileRepo.Create(ctx, params)
	if err != nil {
		return "", err
	}
	return f.cfg.Server.BaseUrl + constants.ENDPOINT_V1 + "/files/" + result.FileName, nil
}

func (f fileUC) FindFile(ctx context.Context, fileName string) (database.FindFileByFileNameRow, error) {
	return f.fileRepo.FindByFileName(ctx, fileName)
}

func (f fileUC) DeleteFile(ctx context.Context, fileName string) error {
	foundFile, err := f.FindFile(ctx, fileName)
	if err != nil {
		return errors.New("file not found")
	}

	_, err = os.Stat(foundFile.Path)
	if os.IsNotExist(err) {
		return errors.New("file not found")
	}

	err = os.Remove(foundFile.Path)
	if err != nil {
		return errors.New("could not delete file")
	}

	return f.fileRepo.Delete(ctx, foundFile.ID)
}
