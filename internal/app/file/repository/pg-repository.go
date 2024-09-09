package repository

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/file"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type fileRepo struct {
	db *database.Queries
}

func NewFileRepository(db *database.Queries) file.Repository {
	return &fileRepo{db}
}

func (f fileRepo) Create(ctx context.Context, params database.CreateFileParams) (database.CreateFileRow, error) {
	return f.db.CreateFile(ctx, params)
}

func (f fileRepo) FindByID(ctx context.Context, id uuid.UUID) (database.FindFileByIdRow, error) {
	return f.db.FindFileById(ctx, id)
}

func (f fileRepo) FindByFileName(ctx context.Context, fileName string) (database.FindFileByFileNameRow, error) {
	return f.db.FindFileByFileName(ctx, fileName)
}

func (f fileRepo) Delete(ctx context.Context, id uuid.UUID) error {
	return f.db.DeleteFileById(ctx, id)
}
