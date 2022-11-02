package tasks

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/external"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/service"
	"strings"
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
		t.logger.Errorf("failed to fetch token list from coingecko %v", err)
		return
	}
	t.logger.Infof("total number of tokens fetched from coingecko %d", len(*tokenList))
	for i := 0; i < len(*tokenList); {
		coingeckoToken := (*tokenList)[i]

		// todo: remove the debug whitelist
		//if !(coingeckoToken.Id == "unmarshal" || coingeckoToken.Id == "dydx" || coingeckoToken.Id == "ethereum") {
		//	i++
		//	continue
		//}
		fetchedToken, err := t.tokenService.GetToken("coingecko", coingeckoToken.Id)
		if err == nil && fetchedToken.SourceTokenId == coingeckoToken.Id {
			t.logger.Debugf("token already present in db %s %s", "coingecko", coingeckoToken.Id)
			i++
			continue
		}
		coingeckoTokenDetails, err := t.coingeckoGateway.GetTokenDetails(coingeckoToken.Id)
		if err != nil {
			t.logger.Errorf("failed to fetch token details from coingecko %s %v", coingeckoToken.Id, err)
			return
		}

		tokenModel := t.toTokenModel(coingeckoToken, coingeckoTokenDetails)

		err = t.tokenService.Create(tokenModel)
		if err != nil {
			t.logger.Errorf("failed to save token model to db %v %v", tokenModel, err)
			// loop variable is incremented as db error is assumed to reoccur
		}
		t.logger.Debugf("successfully created token in db %s", tokenModel.Name)
		time.Sleep(2 * time.Second)
		i++
	}
	t.logger.Info("token list task execution complete")
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
		if strings.TrimSpace(platform) == "" || strings.TrimSpace(detail.ContractAddress) == "" {
			continue
		}
		tokenModel.TokenPlatforms = append(tokenModel.TokenPlatforms, models.TokenPlatform{
			PlatformName: platform,
			Address:      detail.ContractAddress,
			Decimal:      detail.GetDecimalPlace(),
		})
	}
	return tokenModel
}
