package file

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type Repository interface {
	Create(context.Context, database.CreateFileParams) (database.CreateFileRow, error)
	FindByID(context.Context, uuid.UUID) (database.FindFileByIdRow, error)
	FindByFileName(context.Context, string) (database.FindFileByFileNameRow, error)
	Delete(context.Context, uuid.UUID) error
}
