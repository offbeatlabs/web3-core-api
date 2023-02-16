package tasks

import (
	"github.com/offbeatlabs/web3-core-api/pkg/external"
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	"github.com/offbeatlabs/web3-core-api/pkg/service"
	log "github.com/sirupsen/logrus"
	"gorm.io/datatypes"
	"strings"
	"time"
)

type TokenListTask struct {
	coingeckoGateway external.CoingeckoGateway
	tokenService     service.TokenService
}

func NewTokenListTask(coingeckoGateway external.CoingeckoGateway, tokenService service.TokenService) TokenListTask {
	return TokenListTask{
		coingeckoGateway: coingeckoGateway,
		tokenService:     tokenService,
	}
}

func (t TokenListTask) Execute() {
	log.Info("executing token list task")
	tokenList, err := t.coingeckoGateway.GetTokenList()
	if err != nil {
		log.Errorf("failed to fetch token list from coingecko %v", err)
		return
	}
	log.Infof("total number of tokens fetched from coingecko %d", len(*tokenList))
	for i := 0; i < len(*tokenList); {
		coingeckoToken := (*tokenList)[i]

		fetchedToken, err := t.tokenService.GetToken("coingecko", coingeckoToken.Id)
		if err == nil && fetchedToken.SourceTokenId == coingeckoToken.Id {
			log.Debugf("token already present in db %s %s", "coingecko", coingeckoToken.Id)
			i++
			continue
		}
		coingeckoTokenDetails, err := t.coingeckoGateway.GetTokenDetails(coingeckoToken.Id)
		if err != nil {
			log.Errorf("failed to fetch token details from coingecko %s %v", coingeckoToken.Id, err)
			return
		}

		tokenModel := t.toTokenModel(coingeckoToken, coingeckoTokenDetails)

		err = t.tokenService.Create(&tokenModel)
		if err != nil {
			log.Errorf("failed to save token model to db %v %v", tokenModel, err)
			// loop variable is incremented as db error is assumed to reoccur
		}
		log.Debugf("successfully created token in db %s", tokenModel.Name)
		time.Sleep(5 * time.Second)
		i++
		if i%10 == 0 {
			log.Info("Successfully saved 100 tokens i:", i)
		}
	}
	log.Info("token list task execution complete")
}

func (t TokenListTask) toTokenModel(coingeckoToken external.CoingeckoToken, coingeckoTokenDetails *external.CoingeckoTokenDetailResp) models.Token {
	tokenModel := models.Token{
		Symbol: coingeckoToken.Symbol,
		Name:   coingeckoToken.Name,
		Logo: datatypes.JSONType[models.Logo]{Data: models.Logo{
			Thumb: coingeckoTokenDetails.Image.Thumb,
			Small: coingeckoTokenDetails.Image.Small,
			Large: coingeckoTokenDetails.Image.Large,
		}},
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
