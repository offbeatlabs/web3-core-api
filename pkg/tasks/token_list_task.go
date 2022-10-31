package tasks

import (
	"github.com/arhamj/offbeat-api/commons/logger"
	"github.com/arhamj/offbeat-api/pkg/external"
	"github.com/arhamj/offbeat-api/pkg/service"
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
}
