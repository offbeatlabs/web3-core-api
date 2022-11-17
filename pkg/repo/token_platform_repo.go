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
	GetTokenPlatformByAddress         = `SELECT token_id, platform_name, address, decimal FROM token_platforms WHERE address = ?`
	GetTokenPlatformByPlatformDetails = `SELECT token_id, platform_name, address, decimal FROM token_platforms WHERE address = ? AND platform_name = ?`
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

func (r TokenPlatformRepo) GetByPlatformNameAndAddress(platformName, address string) (models.TokenPlatform, error) {
	address = strings.ToLower(address)
	row := r.db.QueryRow(GetTokenPlatformByPlatformDetails, address, platformName)
	if row.Err() != nil {
		return models.TokenPlatform{}, row.Err()
	}
	var res models.TokenPlatform
	err := res.SetSqlRow(row)
	if err != nil {
		return models.TokenPlatform{}, err
	}
	return res, nil
}

func (r TokenPlatformRepo) GetByAddress(address string) (models.TokenPlatform, error) {
	address = strings.ToLower(address)
	row := r.db.QueryRow(GetTokenPlatformByAddress, address)
	if row.Err() != nil {
		return models.TokenPlatform{}, row.Err()
	}
	var res models.TokenPlatform
	err := res.SetSqlRow(row)
	if err != nil {
		return models.TokenPlatform{}, err
	}
	return res, nil
}
