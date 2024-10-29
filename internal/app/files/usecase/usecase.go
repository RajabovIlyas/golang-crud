package usecase

import (
	"errors"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/files"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/rs/zerolog"
	"os"
)

type fileUC struct {
	cfg      *config.Config
	fileRepo files.Repository
	logger   zerolog.Logger
}

func NewFileUseCase(cfg *config.Config, fileRepo files.Repository, logger zerolog.Logger) files.UseCase {
	return &fileUC{cfg: cfg, fileRepo: fileRepo, logger: logger}
}

func (f fileUC) UploadFile(params models.CreateFile) (models.ResponseFile, error) {
	createdFile, err := f.fileRepo.Create(params)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.UploadFile(create files)")
		return models.ResponseFile{}, err
	}

	return models.ResponseFile{f.GenerateFileUrl(createdFile.FileName)}, nil
}

func (f fileUC) FindFile(fileName string) (models.Files, error) {

	foundFile, err := f.fileRepo.FindByFileName(fileName)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.findFile(%s)", fileName)
		return models.Files{}, err
	}

	return foundFile, nil
}

func (f fileUC) DeleteFile(fileName string) error {
	foundFile, err := f.FindFile(fileName)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(find files by name): %s", fileName)
		return errors.New("files not found")
	}

	_, err = os.Stat(foundFile.Path)
	if os.IsNotExist(err) {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(files not found by path): %s", foundFile.Path)
		return errors.New("files not found")
	}

	err = os.Remove(foundFile.Path)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(could not delete files by path): %s", foundFile.Path)
		return errors.New("could not delete files")
	}

	err = f.fileRepo.Delete(foundFile.ID)
	if err != nil {
		f.logger.Error().Err(err).Msgf("fileUC.DeleteFile(delete files by id): %s", foundFile.ID)
		return errors.New("could not delete files")
	}

	return nil
}

func (f fileUC) GenerateFileUrl(fileName string) string {
	return f.cfg.Server.BaseUrl + constants.ENDPOINT_V1 + "/files/" + fileName
}
