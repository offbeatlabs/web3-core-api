package repo

import (
	"database/sql"
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/models"
)

const (
	InsertTokenPlatformQuery   = ""
	GetTokenPlatformsByIdQuery = ""
)

type TokenPlatformRepo struct {
	logger *logger.AppLogger
	db     *sql.DB
}

func NewTokenPlatformRepo(logger *logger.AppLogger, db *sql.DB) *TokenRepo {
	return &TokenRepo{
		logger: logger,
		db:     db,
	}
}

func (r TokenPlatformRepo) Create(platformInfo models.TokenPlatform) error {
	panic("implement me!")
}

func (r TokenPlatformRepo) MultiCreate(platformInfo []models.TokenPlatform) error {
	panic("implement me!")
}

func (r TokenPlatformRepo) GetByTokenId(tokenId int64) ([]models.TokenPlatform, error) {
	panic("implement me!")
}
