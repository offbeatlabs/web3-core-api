package repo

import (
	"database/sql"
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/models"
)

const (
	InsertTokenQuery = `INSERT INTO  "tokens" ("updated_at", "symbol", "name", "logo", "source_token_id", "source") 
						VALUES (?, ?, ?, ?, ?, ?)`

	UpdateTokenPriceQuery = `UPDATE "tokens" SET 
                    		"usd_price" = ?, "usd_market_cap" = ?, 
                    		"usd_24h_change" = ?, "usd_24h_volume" = ? WHERE "id" = ?`

	GetTokenBySourceInfo = `SELECT "id", "updated_at", "symbol", "name", "logo", "source_token_id", "source", "usd_price", 
       						"usd_market_cap", "usd_24h_change", "usd_24h_volume"
							FROM  "tokens" WHERE "source" = ? AND "source_token_id" = ?`

	GetAllTokensQuery = `SELECT "id", "updated_at", "symbol", "name", "logo", "source_token_id", "source", "usd_price", 
						"usd_market_cap", "usd_24h_change", "usd_24h_volume" FROM  "tokens"`
)

type TokenRepo struct {
	logger *logger.AppLogger
	db     *sql.DB
}

func NewTokenRepo(logger *logger.AppLogger, db *sql.DB) TokenRepo {
	return TokenRepo{
		logger: logger,
		db:     db,
	}
}

func (r TokenRepo) Create(token models.Token) (int64, error) {
	err := token.PreCreate()
	if err != nil {
		return 0, err
	}
	stmt, err := r.db.Prepare(InsertTokenQuery)
	if err != nil {
		return 0, err
	}
	res, err := stmt.Exec(token.UpdatedAt, token.Symbol, token.Name, token.Logo, token.SourceTokenId, token.Source)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

func (r TokenRepo) UpdatePriceDetails(id int64, token models.Token) error {
	stmt, err := r.db.Prepare(UpdateTokenPriceQuery)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(token.UsdPrice, token.UsdMarketCap, token.Usd24HourChange, token.Usd24HourVolume, id)
	if err != nil {
		return err
	}
	return nil
}

func (r TokenRepo) GetBySourceTokenId(source, tokenId string) (models.Token, error) {
	row := r.db.QueryRow(GetTokenBySourceInfo, source, tokenId)
	if row.Err() != nil {
		return models.Token{}, row.Err()
	}
	var res models.Token
	err := row.Scan(&res.Id, &res.UpdatedAt, &res.Symbol, &res.Name, &res.Logo, &res.SourceTokenId, &res.Source,
		&res.UsdPrice, &res.UsdMarketCap, &res.Usd24HourChange, &res.Usd24HourVolume)
	if err != nil {
		return models.Token{}, err
	}
	return res, nil
}

func (r TokenRepo) GetAll() ([]models.Token, error) {
	rows, err := r.db.Query(GetAllTokensQuery)
	if err != nil {
		return nil, err
	}
	var res []models.Token
	for rows.Next() {
		var token models.Token
		err = rows.Scan(&token.Id, &token.UpdatedAt, &token.Symbol, &token.Name, &token.Logo, &token.SourceTokenId, &token.Source,
			&token.UsdPrice, &token.UsdMarketCap, &token.Usd24HourChange, &token.Usd24HourVolume)
		res = append(res, token)
	}
	if err = rows.Err(); err != nil {
		return res, err
	}
	return res, nil
}
