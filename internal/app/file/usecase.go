package file

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type UseCase interface {
	UploadFile(models.CreateFile) (string, error)
	FindFile(string) (models.Files, error)
	DeleteFile(string) error
}
