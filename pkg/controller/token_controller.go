package controller

import (
	httpErrors "github.com/arhamj/go-commons/pkg/http_errors"
	"github.com/arhamj/go-commons/pkg/logger"
	"github.com/arhamj/offbeat-api/pkg/dto"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type TokenController struct {
	logger       *logger.AppLogger
	tokenService service.TokenService
}

func NewTokenController(logger *logger.AppLogger, tokenService service.TokenService) TokenController {
	return TokenController{
		logger:       logger,
		tokenService: tokenService,
	}
}

func (t TokenController) GetTokenDetails(ctx echo.Context) error {
	address := strings.TrimSpace(ctx.QueryParam("address"))
	if len(address) == 0 {
		return httpErrors.NewBadRequestError(ctx, "address query param required", true)
	}
	platform := strings.TrimSpace(ctx.QueryParam("platform"))
	var (
		token models.Token
		err   error
	)
	if len(platform) == 0 {
		token, err = t.tokenService.GetTokenByAddress(address)
	} else {
		token, err = t.tokenService.GetTokenByPlatformDetails(platform, address)
	}
	if err != nil {
		return httpErrors.ErrorCtxResponse(ctx, err, false)
	}
	tokenPlatforms, err := t.tokenService.GetTokenPlatformsByTokenId(token.Id)
	if err != nil {
		return httpErrors.ErrorCtxResponse(ctx, err, false)
	}
	token.TokenPlatforms = tokenPlatforms
	res := dto.NewTokenDetails(token)
	return ctx.JSON(http.StatusOK, res)
}
