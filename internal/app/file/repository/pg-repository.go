package repository

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/file"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type fileRepo struct {
	db *gorm.DB
}

func NewFileRepository(db *gorm.DB) file.Repository {
	return &fileRepo{db}
}

func (f fileRepo) Create(params models.CreateFile) (models.Files, error) {
	var createFile models.Files
	createFile.FileName = params.FileName
	createFile.UserID = params.UserID
	createFile.Path = params.Path
	createFile.Size = params.Size

	result := f.db.Create(&createFile)

	return createFile, result.Error
}

func (f fileRepo) FindByID(id uuid.UUID) (models.Files, error) {
	var findFile models.Files
	result := f.db.Where("id = ?", id).First(&findFile)
	return findFile, result.Error
}

func (f fileRepo) FindByFileName(fileName string) (models.Files, error) {
	var findFile models.Files
	result := f.db.Where("file_name = ?", fileName).First(&findFile)
	return findFile, result.Error
}

func (f fileRepo) Delete(id uuid.UUID) error {
	var deleteFile models.Files
	result := f.db.Where("id = ?", id).Delete(&deleteFile)
	return result.Error
}
