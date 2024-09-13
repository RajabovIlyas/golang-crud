package repository

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type tokenRepo struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) token.Repository {
	return &tokenRepo{db}
}

func (t tokenRepo) FindByID(id uuid.UUID) (models.Tokens, error) {
	var findToken models.Tokens
	result := t.db.First(&findToken, id)
	return findToken, result.Error
}

func (t tokenRepo) FindByAccessKey(accessKey uuid.UUID) (models.Tokens, error) {
	var findToken models.Tokens
	result := t.db.First(&findToken, "access_token_key = ?", accessKey)
	return findToken, result.Error
}

func (t tokenRepo) Create(userID uuid.UUID) (models.Tokens, error) {
	var createToken = models.Tokens{UserID: userID}
	result := t.db.Select("user_id").Create(&createToken)
	return createToken, result.Error
}

func (t tokenRepo) UpdateByID(tokenID uuid.UUID) (models.Tokens, error) {
	var updateToken = models.Tokens{ID: tokenID}
	result := t.db.First(&updateToken)
	if result.Error != nil {
		return models.Tokens{}, result.Error
	}
	updateToken.AccessTokenKey = uuid.New()
	result = t.db.Save(&updateToken)
	return updateToken, result.Error
}

func (t tokenRepo) DeleteByID(tokenID uuid.UUID) (models.Tokens, error) {
	var deleteToken models.Tokens
	result := t.db.Delete(&deleteToken, tokenID)
	return deleteToken, result.Error
}

func (t tokenRepo) DeleteByAccessKey(accessKey uuid.UUID) (models.Tokens, error) {
	var deleteToken models.Tokens
	result := t.db.Delete(&deleteToken, "access_token_key = ?", accessKey)
	return deleteToken, result.Error
}

func (t tokenRepo) DeleteOldTokens() ([]models.Tokens, error) {
	var deleteTokens []models.Tokens
	result := t.db.Model(&deleteTokens).Where("updated_at <= NOW() - INTERVAL '2 days'")

	return deleteTokens, result.Error
}
