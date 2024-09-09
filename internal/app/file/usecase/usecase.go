package usecase

import (
	"context"
	"errors"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/file"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/rs/zerolog"
	"os"
)

type fileUC struct {
	cfg      *config.Config
	fileRepo file.Repository
	logger   zerolog.Logger
}

func NewFileUseCase(cfg *config.Config, fileRepo file.Repository, logger zerolog.Logger) file.UseCase {
	return &fileUC{cfg: cfg, fileRepo: fileRepo, logger: logger}
}

func (f fileUC) UploadFile(ctx context.Context, params database.CreateFileParams) (string, error) {
	createdFile, err := f.fileRepo.Create(ctx, params)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.UploadFile(create file)")
		return "", err
	}

	return f.GenerateFileUrl(createdFile.FileName), nil
}

func (f fileUC) FindFile(ctx context.Context, fileName string) (models.FileModel, error) {

	foundFile, err := f.fileRepo.FindByFileName(ctx, fileName)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.findFile(%s)", fileName)
		return models.FileModel{}, err
	}

	fileModel := models.FileModel{
		ID:       foundFile.ID,
		FileName: foundFile.FileName,
		Format:   foundFile.Format,
		Size:     foundFile.Size,
		Path:     foundFile.Path,
		UserID:   foundFile.UserID,
	}

	return fileModel, nil
}

func (f fileUC) DeleteFile(ctx context.Context, fileName string) error {
	foundFile, err := f.FindFile(ctx, fileName)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(find file by name): %s", fileName)
		return errors.New("file not found")
	}

	_, err = os.Stat(foundFile.Path)
	if os.IsNotExist(err) {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(file not found by path): %s", foundFile.Path)
		return errors.New("file not found")
	}

	err = os.Remove(foundFile.Path)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(could not delete file by path): %s", foundFile.Path)
		return errors.New("could not delete file")
	}

	err = f.fileRepo.Delete(ctx, foundFile.ID)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(delete file by id): %s", foundFile.ID)
		return errors.New("could not delete file")
	}

	return nil
}

func (f fileUC) GenerateFileUrl(fileName string) string {
	return f.cfg.Server.BaseUrl + constants.ENDPOINT_V1 + "/files/" + fileName
}
