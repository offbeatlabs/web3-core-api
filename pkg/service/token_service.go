package service

import (
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	"github.com/offbeatlabs/web3-core-api/pkg/repo"
	log "github.com/sirupsen/logrus"
)

type TokenService struct {
	tokenRepo         *repo.TokenRepo
	tokenPlatformRepo *repo.TokenPlatformRepo
}

func NewTokenService(tokenRepo *repo.TokenRepo, tokenPlatformRepo *repo.TokenPlatformRepo) TokenService {
	return TokenService{
		tokenRepo:         tokenRepo,
		tokenPlatformRepo: tokenPlatformRepo,
	}
}

func (s TokenService) Create(token *models.Token) error {
	err := s.tokenRepo.Create(token)
	if err != nil {
		log.Errorf("failed to insert token in db %s %v", token.SourceTokenId, err)
		return err
	}
	return nil
}

func (s TokenService) UpdatePriceDetails(tokenId uint, token models.Token) error {
	err := s.tokenRepo.UpdatePriceDetails(tokenId, token)
	if err != nil {
		log.Errorf("error when updating token price details %s %v", token.SourceTokenId, err)
		return err
	}
	return nil
}

func (s TokenService) GetAllTokens() ([]models.Token, error) {
	tokens, err := s.tokenRepo.GetAll()
	if err != nil {
		log.Error("error when getting all tokens", err)
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

func (s TokenService) GetTokenByAddress(contractAddress string) (models.Token, error) {
	tokenPlatform, err := s.tokenPlatformRepo.GetByAddress(contractAddress)
	if err != nil {
		return models.Token{}, err
	}
	token, err := s.tokenRepo.GetByTokenId(tokenPlatform.TokenID)
	if err != nil {
		return models.Token{}, err
	}
	return token, err
}

func (s TokenService) GetTokenByPlatformDetails(platform, contractAddress string) (models.Token, error) {
	tokenPlatform, err := s.tokenPlatformRepo.GetByPlatformNameAndAddress(platform, contractAddress)
	if err != nil {
		return models.Token{}, err
	}
	token, err := s.tokenRepo.GetByTokenId(tokenPlatform.TokenID)
	if err != nil {
		return models.Token{}, err
	}
	return token, err
}
