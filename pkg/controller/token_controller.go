package controller

import (
	httpErrors "github.com/arhamj/offbeat-api/commons/http_errors"
	"github.com/arhamj/offbeat-api/config"
	"github.com/arhamj/offbeat-api/pkg/models"
	"github.com/arhamj/offbeat-api/pkg/service"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

type TokenController struct {
	httpErrorDebug bool
	tokenService   service.TokenService
}

func NewTokenController(config config.Config, tokenService service.TokenService) TokenController {
	return TokenController{
		httpErrorDebug: config.FeatureFlags.EnableHttpErrDebug,
		tokenService:   tokenService,
	}
}

func (t TokenController) GetTokenDetails(ctx echo.Context) {
	address := strings.TrimSpace(ctx.QueryParam("address"))
	if len(address) == 0 {
		_ = httpErrors.NewBadRequestError(ctx, []string{}, t.httpErrorDebug)
		return
	}
	platform := strings.TrimSpace(ctx.QueryParam("platform"))
	var (
		res models.Token
		err error
	)
	if len(platform) == 0 {
		res, err = t.tokenService.GetTokenByAddress(address)
	} else {
		res, err = t.tokenService.GetTokenByPlatformDetails(platform, address)
	}
	if err != nil {
		_ = httpErrors.ErrorCtxResponse(ctx, err, t.httpErrorDebug)
	}
	_ = ctx.JSON(http.StatusOK, res)
	return
}
