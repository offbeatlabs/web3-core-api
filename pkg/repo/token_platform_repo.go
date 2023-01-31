package repo

import (
	"context"
	"database/sql"
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	InsertTokenPlatformQuery = `INSERT INTO "token_platforms" ("token_id", "platform_name", "address", "decimal")
								VALUES (?, ?, ?, ?)`
	GetTokenPlatformsByTokenId        = `SELECT token_id, platform_name, address, decimal FROM token_platforms WHERE token_id = ?`
	GetTokenPlatformByAddress         = `SELECT token_id, platform_name, address, decimal FROM token_platforms WHERE address = ?`
	GetTokenPlatformByPlatformDetails = `SELECT token_id, platform_name, address, decimal FROM token_platforms WHERE address = ? AND platform_name = ?`
)

type TokenPlatformRepo struct {
	db *sql.DB
}

func NewTokenPlatformRepo(db *sql.DB) TokenPlatformRepo {
	return TokenPlatformRepo{
		db: db,
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
			log.Errorf("error inserting platform info %v %v", platform, err)
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

func (r TokenPlatformRepo) GetByTokenId(tokenId int64) ([]models.TokenPlatform, error) {
	rows, err := r.db.Query(GetTokenPlatformsByTokenId, tokenId)
	if err != nil {
		return nil, err
	}
	var res []models.TokenPlatform
	for rows.Next() {
		var t models.TokenPlatform
		err = rows.Scan(&t.TokenId, &t.PlatformName, &t.Address, &t.Decimal)
		res = append(res, t)
	}
	if err = rows.Err(); err != nil {
		return res, err
	}
	return res, nil
}
