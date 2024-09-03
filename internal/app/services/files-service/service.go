package filesService

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

type Service interface {
	UploadFile(context.Context, database.CreateFileParams) (string, error)
	FindFile(context.Context, string) error
	DeleteFile(context.Context, string) error
}
