package repository

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type tokenRepo struct {
	db *database.Queries
}

func NewTokenRepository(db *database.Queries) token.Repository {
	return &tokenRepo{db}
}

func (t tokenRepo) FindByID(ctx context.Context, id uuid.UUID) (database.FindTokenByIdRow, error) {
	return t.db.FindTokenById(ctx, id)
}

func (t tokenRepo) FindByAccessKey(ctx context.Context, accessKey uuid.UUID) (database.FindTokenByAccessKeyRow, error) {
	return t.db.FindTokenByAccessKey(ctx, accessKey)
}

func (t tokenRepo) Create(ctx context.Context, userID uuid.UUID) (database.CreateTokenRow, error) {
	return t.db.CreateToken(ctx, userID)
}

func (t tokenRepo) UpdateByID(ctx context.Context, tokenID uuid.UUID) (database.UpdateTokenByIdRow, error) {
	return t.db.UpdateTokenById(ctx, tokenID)
}

func (t tokenRepo) DeleteByID(ctx context.Context, tokenID uuid.UUID) (database.DeleteTokenByIdRow, error) {
	return t.db.DeleteTokenById(ctx, tokenID)
}

func (t tokenRepo) DeleteByAccessKey(ctx context.Context, accessKey uuid.UUID) (database.DeleteTokenByAccessKeyRow, error) {
	return t.db.DeleteTokenByAccessKey(ctx, accessKey)
}

func (t tokenRepo) DeleteOldTokens(ctx context.Context) ([]database.DeleteOldTokensRow, error) {
	return t.db.DeleteOldTokens(ctx)
}
