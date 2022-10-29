package repo

import (
	"database/sql"
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/models"
)

const (
	InsertTokenQuery      = ""
	UpdateTokenQuery      = ""
	UpdateTokenPriceQuery = ""
	GetTokenByIdQuery     = ""
	GetTokenBySourceInfo  = ""
)

type TokenRepo struct {
	logger *logger.AppLogger
	db     *sql.DB
}

func NewTokenRepo(logger *logger.AppLogger, db *sql.DB) *TokenRepo {
	return &TokenRepo{
		logger: logger,
		db:     db,
	}
}

func (r TokenRepo) Create(token models.Token) error {
	panic("implement me!")
}

func (r TokenRepo) MultiCreate(token models.Token) error {
	panic("implement me!")
}

func (r TokenRepo) UpdateDetails(id int64, token models.Token) error {
	panic("implement me!")
}

func (r TokenRepo) UpdatePriceData(id int64, token models.Token) error {
	panic("implement me!")
}

func (r TokenRepo) GetById(id int64) (*models.Token, error) {
	panic("implement me!")
}

func (r TokenRepo) GetBySourceTokenId(source, tokenId string) (*models.Token, error) {
	panic("implement me!")
}

func (r TokenRepo) GetAll() ([]models.Token, error) {
	panic("implement me!")
}
