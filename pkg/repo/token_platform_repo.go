package repo

import (
	"context"
	"database/sql"
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/models"
	"strings"
)

const (
	InsertTokenPlatformQuery = `INSERT INTO "token_platforms" ("token_id", "platform_name", "address", "decimal")
								VALUES (?, ?, ?, ?)`
)

type TokenPlatformRepo struct {
	logger *logger.AppLogger
	db     *sql.DB
}

func NewTokenPlatformRepo(logger *logger.AppLogger, db *sql.DB) TokenPlatformRepo {
	return TokenPlatformRepo{
		logger: logger,
		db:     db,
	}
}

func (r TokenPlatformRepo) Create(platformInfo models.TokenPlatform) error {
	stmt, err := r.db.Prepare(InsertTokenPlatformQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(platformInfo.TokenId, platformInfo.PlatformName, strings.ToLower(platformInfo.Address), platformInfo.Decimal)
	if err != nil {
		return err
	}
	return nil
}

func (r TokenPlatformRepo) CreateTx(tx *sql.Tx, platformInfo models.TokenPlatform) error {
	stmt, err := tx.Prepare(InsertTokenPlatformQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(platformInfo.TokenId, platformInfo.PlatformName, strings.ToLower(platformInfo.Address), platformInfo.Decimal)
	if err != nil {
		return err
	}
	return nil
}

func (r TokenPlatformRepo) MultiCreate(tokenId int64, platformInfo []models.TokenPlatform) error {
	tx, err := r.db.BeginTx(context.TODO(), nil)
	if err != nil {
		return err
	}
	for _, platform := range platformInfo {
		platform.TokenId = tokenId
		err = r.CreateTx(tx, platform)
		if err != nil {
			r.logger.Errorf("error inserting platform info %v %v", platform, err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		_ = tx.Rollback()
	}
	return nil
}
