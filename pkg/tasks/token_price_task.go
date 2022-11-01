package tasks

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/external"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/service"
)

type TokenPriceTask struct {
	logger           *logger.AppLogger
	coingeckoGateway external.CoingeckoGateway
	tokenService     service.TokenService
}

func NewTokenPriceTask(logger *logger.AppLogger, coingeckoGateway external.CoingeckoGateway,
	tokenService service.TokenService) TokenPriceTask {
	return TokenPriceTask{
		logger:           logger,
		coingeckoGateway: coingeckoGateway,
		tokenService:     tokenService,
	}
}

func (t TokenPriceTask) Execute() {
	t.logger.Info("executing token price task")
	tokens, err := t.tokenService.GetAllTokens()
	if err != nil {
		t.logger.Info("failed to fetch token list from db", err)
		return
	}
	aggregateTokenPriceMap, err := t.fetchAllTokenPrices(tokens)
	if err != nil {
		t.logger.Info("failed to fetch all token prices from coingecko", err)
		return
	}

	for _, token := range tokens {
		coingeckoPriceDetail, ok := aggregateTokenPriceMap[token.SourceTokenId]
		if !ok {
			t.logger.Info("price details not found for token", token.Id)
			continue
		}
		token.UsdPrice = coingeckoPriceDetail.Usd
		token.UsdMarketCap = coingeckoPriceDetail.UsdMarketCap
		token.Usd24HourChange = coingeckoPriceDetail.Usd24HChange
		err = t.tokenService.UpdatePriceDetails(token.Id, token)
		if err != nil {
			t.logger.Info("failed to update price details for token", token.Id, err)
		}
	}
}

func (t TokenPriceTask) fetchAllTokenPrices(tokens []models.Token) (map[string]external.CoingeckoTokenPrice, error) {
	aggregateTokenPriceMap := make(map[string]external.CoingeckoTokenPrice, 0)
	tokenIds := make([]string, 0)
	for i, token := range tokens {
		tokenIds = append(tokenIds, token.SourceTokenId)
		if i%1000 == 0 || i == len(tokens)-1 {
			tokenPriceMap, err := t.coingeckoGateway.GetTokenPrice(tokenIds)
			if err != nil {
				t.logger.Info("failed to fetch token prices from coingecko", err)
				return nil, err
			}
			for tokenId, priceDetails := range *tokenPriceMap {
				aggregateTokenPriceMap[tokenId] = priceDetails
			}
			tokenIds = make([]string, 0)
		}
	}
	return aggregateTokenPriceMap, nil
}
