package service

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/repo"
)

type TokenService struct {
	logger            *logger.AppLogger
	tokenRepo         *repo.TokenRepo
	tokenPlatformRepo *repo.TokenPlatformRepo
}

func NewTokenService(logger *logger.AppLogger, tokenRepo *repo.TokenRepo,
	tokenPlatformRepo *repo.TokenPlatformRepo) TokenService {
	return TokenService{
		logger:            logger,
		tokenRepo:         tokenRepo,
		tokenPlatformRepo: tokenPlatformRepo,
	}
}

func (s TokenService) Create(token models.Token) error {
	insertId, err := s.tokenRepo.Create(token)
	if err != nil {
		s.logger.Errorf("failed to insert token in db %s %v", token.SourceTokenId, err)
		return err
	}
	if len(token.TokenPlatforms) > 0 {
		err = s.tokenPlatformRepo.MultiCreate(insertId, token.TokenPlatforms)
		if err != nil {
			s.logger.Err("failed to insert token platforms in db", err)
			return err
		}
	}
	return nil
}

func (s TokenService) UpdatePriceDetails(tokenId int64, token models.Token) error {
	err := s.tokenRepo.UpdatePriceDetails(tokenId, token)
	if err != nil {
		s.logger.Errorf("error when updating token price details %s %v", token.SourceTokenId, err)
		return err
	}
	return nil
}

func (s TokenService) GetAllTokens() ([]models.Token, error) {
	tokens, err := s.tokenRepo.GetAll()
	if err != nil {
		s.logger.Err("error when getting all tokens", err)
		return nil, err
	}
	return tokens, nil
}

func (s TokenService) GetToken(source, sourceTokenId string) (models.Token, error) {
	token, err := s.tokenRepo.GetBySourceTokenId(source, sourceTokenId)
	if err != nil {
		return models.Token{}, err
	}
	return token, nil
}
