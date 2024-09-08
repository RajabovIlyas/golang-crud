package file

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

type UseCase interface {
	UploadFile(context.Context, database.CreateFileParams) (string, error)
	FindFile(context.Context, string) (database.FindFileByFileNameRow, error)
	DeleteFile(context.Context, string) error
}
