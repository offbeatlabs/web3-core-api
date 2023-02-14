package repo

import (
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TokenRepo struct {
	db *gorm.DB
}

func NewTokenRepo(db *gorm.DB) TokenRepo {
	return TokenRepo{
		db: db,
	}
}

func (r TokenRepo) Create(token models.Token) error {
	res := r.db.Create(token)
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r TokenRepo) UpdatePriceDetails(id uint, incoming models.Token) error {
	res := r.db.Model(&incoming).Where("id = ?", id).Updates(&models.Token{
		UsdPrice:        incoming.UsdPrice,
		UsdMarketCap:    incoming.UsdMarketCap,
		Usd24HourChange: incoming.Usd24HourChange,
		Usd24HourVolume: incoming.Usd24HourVolume,
	})
	if err := res.Error; err != nil {
		return err
	}
	return nil
}

func (r TokenRepo) GetBySourceTokenId(source, tokenId string) (models.Token, error) {
	var res models.Token
	err := r.db.Where("source = ?", source).Where("source_token_id = ?", tokenId).First(&res).Error
	if err != nil {
		log.WithFields(log.Fields{
			"err":             err,
			"source":          source,
			"source_token_id": tokenId,
		}).Error("GetBySourceTokenId: error finding token by source details")
		return models.Token{}, err
	}
	return res, nil
}

func (r TokenRepo) GetByTokenId(tokenId int64) (models.Token, error) {
	var res models.Token
	err := r.db.Where("id = ?", tokenId).First(&res).Error
	if err != nil {
		log.WithFields(log.Fields{
			"err":      err,
			"token_id": tokenId,
		}).Error("GetByTokenId: error finding token by id")
		return models.Token{}, err
	}
	return res, nil
}

func (r TokenRepo) GetAll() ([]models.Token, error) {
	var res []models.Token
	err := r.db.Model(&models.Token{}).Find(&res).Error
	if err != nil {
		return nil, err
	}
	return res, nil
}
