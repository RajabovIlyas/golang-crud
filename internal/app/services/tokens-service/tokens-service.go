package tokensService

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/common"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type TokensService struct {
	db *database.Queries
	c  *common.Config
}

func NewTokensService(db *database.Queries) *TokensService {
	c, _ := common.GetConfig(".")
	return &TokensService{db, &c}
}

func (ts *TokensService) FindTokenById(r *http.Request, id string) (database.FindTokenByIdRow, error) {
	tokenId, err := uuid.Parse(id)
	if err != nil {
		return database.FindTokenByIdRow{}, err
	}
	return ts.db.FindTokenById(r.Context(), tokenId)
}

func (ts *TokensService) FindTokenByAccessKey(r *http.Request, accessKey string) (database.FindTokenByAccessKeyRow, error) {
	uuidAccessKey, err := uuid.Parse(accessKey)
	if err != nil {
		return database.FindTokenByAccessKeyRow{}, err
	}
	return ts.db.FindTokenByAccessKey(r.Context(), uuidAccessKey)
}

func (ts *TokensService) CreateToken(r *http.Request, userID uuid.UUID) (database.CreateTokenRow, error) {
	return ts.db.CreateToken(r.Context(), userID)
}

func (ts *TokensService) UpdateToken(r *http.Request, tokenIDStr string) (database.UpdateTokenByIdRow, error) {
	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		return database.UpdateTokenByIdRow{}, err
	}
	return ts.db.UpdateTokenById(r.Context(), tokenID)
}

func (ts *TokensService) DeleteTokenById(r *http.Request, tokenID uuid.UUID) error {
	return ts.db.DeleteTokenById(r.Context(), tokenID)
}

func (ts *TokensService) DeleteTokenByAccessKey(r *http.Request, accessKey string) error {
	uuidAccessKey, err := uuid.Parse(accessKey)
	if err != nil {
		return err
	}
	return ts.db.DeleteTokenByAccessKey(r.Context(), uuidAccessKey)
}

func (ts *TokensService) DeleteOldTokens(r *http.Request) error {
	return ts.db.DeleteOldTokens(r.Context())
}
