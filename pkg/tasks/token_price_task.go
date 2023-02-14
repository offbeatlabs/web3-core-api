package tasks

import (
	"github.com/offbeatlabs/web3-core-api/pkg/external"
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	"github.com/offbeatlabs/web3-core-api/pkg/service"
	log "github.com/sirupsen/logrus"
	"time"
)

type TokenPriceTask struct {
	coingeckoGateway external.CoingeckoGateway
	tokenService     service.TokenService
}

func NewTokenPriceTask(coingeckoGateway external.CoingeckoGateway, tokenService service.TokenService) TokenPriceTask {
	return TokenPriceTask{
		coingeckoGateway: coingeckoGateway,
		tokenService:     tokenService,
	}
}

func (t TokenPriceTask) Execute() {
	log.Info("executing token price task")
	tokens, err := t.tokenService.GetAllTokens()
	if err != nil {
		log.Error("failed to fetch token list from db", err)
		return
	}
	aggregateTokenPriceMap, err := t.fetchAllTokenPrices(tokens)
	if err != nil {
		log.Error("failed to fetch all token prices from coingecko", err)
		return
	}

	for _, token := range tokens {
		coingeckoPriceDetail, ok := aggregateTokenPriceMap[token.SourceTokenId]
		if !ok {
			log.Debug("price details not found for token ", token.ID)
			continue
		}
		token.UsdPrice = coingeckoPriceDetail.Usd
		token.UsdMarketCap = coingeckoPriceDetail.UsdMarketCap
		token.Usd24HourChange = coingeckoPriceDetail.Usd24HChange
		token.Usd24HourVolume = coingeckoPriceDetail.Usd24HVol
		err = t.tokenService.UpdatePriceDetails(token.ID, token)
		if err != nil {
			log.Errorf("failed to update price details for token %s %v", token.SourceTokenId, err)
		}
	}
	log.Info("token price task execution complete")
}

func (t TokenPriceTask) fetchAllTokenPrices(tokens []models.Token) (map[string]external.CoingeckoTokenPrice, error) {
	aggregateTokenPriceMap := make(map[string]external.CoingeckoTokenPrice, 0)
	tokenIds := make([]string, 0)
	for i, token := range tokens {
		tokenIds = append(tokenIds, token.SourceTokenId)
		if i%200 == 199 || i == len(tokens)-1 {
			tokenPriceMap, err := t.coingeckoGateway.GetTokenPrice(tokenIds)
			if err != nil {
				log.Error("failed to fetch token prices from coingecko", err)
				return nil, err
			}
			for tokenId, priceDetails := range *tokenPriceMap {
				aggregateTokenPriceMap[tokenId] = priceDetails
			}
			tokenIds = make([]string, 0)
			time.Sleep(2 * time.Second)
		}
	}
	return aggregateTokenPriceMap, nil
}
