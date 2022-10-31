package tasks

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/external"
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
}
