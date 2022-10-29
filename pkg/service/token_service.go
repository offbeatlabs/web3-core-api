package service

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/repo"
	"github.com/pkg/errors"
)

type TokenService struct {
	logger            *logger.AppLogger
	tokenRepo         *repo.TokenRepo
	tokenPlatformRepo *repo.TokenPlatformRepo
}

func NewTokenService(logger *logger.AppLogger, tokenRepo *repo.TokenRepo,
	tokenPlatformRepo *repo.TokenPlatformRepo) *TokenService {
	return &TokenService{
		logger:            logger,
		tokenRepo:         tokenRepo,
		tokenPlatformRepo: tokenPlatformRepo,
	}
}

func (s TokenService) Create(token models.Token) error {
	err := s.tokenRepo.Create(token)
	if err != nil {
		s.logger.Error("failed to insert token in db", token, err)
		return err
	}
	if len(token.TokenPlatforms) > 0 {
		err = s.tokenPlatformRepo.MultiCreate(token.TokenPlatforms)
		if err != nil {
			s.logger.Error("failed to insert token platforms in db", token.TokenPlatforms, err)
			return err
		}
	}
	return nil
}

func (s TokenService) UpdateTokenDetails(token models.Token) error {
	if token.Id == 0 {
		s.logger.Error("attempt to update with no token id", token)
		return errors.New("token id is required")
	}
	err := s.tokenRepo.UpdateDetails(token.Id, token)
	if err != nil {
		s.logger.Error("error when updating token details", token, err)
		return err
	}
	return nil
}
