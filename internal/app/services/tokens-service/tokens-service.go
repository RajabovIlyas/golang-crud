package tokensService

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type TokensService struct {
	db *database.Queries
	c  *config.Config
}

func NewTokensService(p *models.DBConfigParam) *TokensService {
	return &TokensService{p.DB, p.C}
}

func (ts *TokensService) FindTokenById(c context.Context, id string) (database.FindTokenByIdRow, error) {
	tokenId, err := uuid.Parse(id)
	if err != nil {
		return database.FindTokenByIdRow{}, err
	}
	return ts.db.FindTokenById(c, tokenId)
}

func (ts *TokensService) FindTokenByAccessKey(c context.Context, accessKey string) (database.FindTokenByAccessKeyRow, error) {
	uuidAccessKey, err := uuid.Parse(accessKey)
	if err != nil {
		return database.FindTokenByAccessKeyRow{}, err
	}
	return ts.db.FindTokenByAccessKey(c, uuidAccessKey)
}

func (ts *TokensService) CreateToken(c context.Context, userID uuid.UUID) (database.CreateTokenRow, error) {
	return ts.db.CreateToken(c, userID)
}

func (ts *TokensService) UpdateToken(c context.Context, tokenIDStr string) (database.UpdateTokenByIdRow, error) {
	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		return database.UpdateTokenByIdRow{}, err
	}
	return ts.db.UpdateTokenById(c, tokenID)
}

func (ts *TokensService) DeleteTokenById(c context.Context, tokenID uuid.UUID) error {
	return ts.db.DeleteTokenById(c, tokenID)
}

func (ts *TokensService) DeleteTokenByAccessKey(c context.Context, accessKey string) error {
	uuidAccessKey, err := uuid.Parse(accessKey)
	if err != nil {
		return err
	}
	return ts.db.DeleteTokenByAccessKey(c, uuidAccessKey)
}

func (ts *TokensService) DeleteOldTokens(c context.Context) error {
	return ts.db.DeleteOldTokens(c)
}
