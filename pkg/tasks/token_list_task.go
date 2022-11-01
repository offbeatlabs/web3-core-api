package tasks

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/external"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/service"
	"time"
)

type TokenListTask struct {
	logger           *logger.AppLogger
	coingeckoGateway external.CoingeckoGateway
	tokenService     service.TokenService
}

func NewTokenListTask(logger *logger.AppLogger, coingeckoGateway external.CoingeckoGateway,
	tokenService service.TokenService) TokenListTask {
	return TokenListTask{
		logger:           logger,
		coingeckoGateway: coingeckoGateway,
		tokenService:     tokenService,
	}
}

func (t TokenListTask) Execute() {
	t.logger.Info("executing token list task")
	tokenList, err := t.coingeckoGateway.GetTokenList()
	if err != nil {
		t.logger.Info("failed to fetch token list from coingecko", err)
		return
	}
	for i := 0; i < len(*tokenList); {
		coingeckoToken := (*tokenList)[i]
		fetchedToken, err := t.tokenService.GetToken("coingecko", coingeckoToken.Id)
		if err != nil {
			t.logger.Error("failed to fetch token list from coingecko", err)
			return
		}
		if fetchedToken.SourceTokenId == coingeckoToken.Id {
			i++
			continue
		}
		coingeckoTokenDetails, err := t.coingeckoGateway.GetTokenDetails(coingeckoToken.Id)
		if err != nil {
			t.logger.Error("failed to fetch token details from coingecko", coingeckoToken.Id, err)
			return
		}

		tokenModel := t.toTokenModel(coingeckoToken, coingeckoTokenDetails)

		err = t.tokenService.Create(tokenModel)
		if err != nil {
			t.logger.Error("failed to save token model to db", tokenModel, err)
			// loop variable is incremented as db error is assumed to reoccur
		}
		i++
	}
}

func (t TokenListTask) toTokenModel(coingeckoToken external.CoingeckoToken, coingeckoTokenDetails *external.CoingeckoTokenDetailResp) models.Token {
	tokenModel := models.Token{
		UpdatedAt: time.Now().UTC(),
		Symbol:    coingeckoToken.Symbol,
		Name:      coingeckoToken.Name,
		ParsedLogo: models.Logo{
			Thumb: coingeckoTokenDetails.Image.Thumb,
			Small: coingeckoTokenDetails.Image.Small,
			Large: coingeckoTokenDetails.Image.Large,
		},
		SourceTokenId:  coingeckoToken.Id,
		Source:         "coingecko",
		TokenPlatforms: []models.TokenPlatform{},
	}
	for platform, detail := range coingeckoTokenDetails.DetailPlatforms {
		tokenModel.TokenPlatforms = append(tokenModel.TokenPlatforms, models.TokenPlatform{
			PlatformName: platform,
			Address:      detail.ContractAddress,
			Decimal:      detail.GetDecimalPlace(),
		})
	}
	return tokenModel
}
