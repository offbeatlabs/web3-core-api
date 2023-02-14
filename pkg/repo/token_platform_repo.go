package repo

import (
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	"gorm.io/gorm"
)

type TokenPlatformRepo struct {
	db *gorm.DB
}

func NewTokenPlatformRepo(db *gorm.DB) TokenPlatformRepo {
	return TokenPlatformRepo{
		db: db,
	}
}

func (r TokenPlatformRepo) GetByPlatformNameAndAddress(platform string, address string) (models.TokenPlatform, error) {
	panic("Implement me!")
}

func (r TokenPlatformRepo) GetByAddress(address string) (models.TokenPlatform, error) {
	panic("Implement me!")
}
