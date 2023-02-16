package repo

import (
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"strings"
)

type TokenPlatformRepo struct {
	db *gorm.DB
}

func NewTokenPlatformRepo(db *gorm.DB) TokenPlatformRepo {
	return TokenPlatformRepo{
		db: db,
	}
}

func (r TokenPlatformRepo) Create(tokenPlatforms []models.TokenPlatform) error {
	res := r.db.CreateInBatches(tokenPlatforms, 5)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r TokenPlatformRepo) GetByPlatformNameAndAddress(platform string, address string) (models.TokenPlatform, error) {
	var res models.TokenPlatform
	address = strings.ToLower(address)
	err := r.db.Where("platform = ?", platform).Where("address = ?", address).First(&res).Error
	if err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"platform": platform,
		}).Error("GetByPlatformNameAndAddress: error finding platform")
		return models.TokenPlatform{}, err
	}
	return res, nil
}

func (r TokenPlatformRepo) GetByAddress(address string) (models.TokenPlatform, error) {
	var res models.TokenPlatform
	address = strings.ToLower(address)
	err := r.db.Where("address = ?", address).First(&res).Error
	if err != nil {
		log.WithFields(log.Fields{
			"err":     err,
			"address": address,
		}).Error("GetByAddress: error finding platform by address")
		return models.TokenPlatform{}, err
	}
	return res, nil
}
