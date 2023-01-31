package controller

import (
	httpErrors "github.com/arhamj/go-commons/pkg/http_errors"
	"github.com/labstack/echo/v4"
	"github.com/offbeatlabs/web3-core-api/pkg/dto"
	"github.com/offbeatlabs/web3-core-api/pkg/models"
	"github.com/offbeatlabs/web3-core-api/pkg/service"
	"net/http"
	"strings"
)

type TokenController struct {
	tokenService service.TokenService
}

func NewTokenController(tokenService service.TokenService) TokenController {
	return TokenController{
		tokenService: tokenService,
	}
}

// GetTokenDetails
//
//	@Summary		Get token details
//	@Description	Fetch the token details by address with optional platform param
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			api-key		header		string	true	"API key of the client"
//	@Param			address		query		string	true	"Token address"
//	@Param			platform	query		string	false	"Platform"
//	@Success		200			{object}	dto.TokenDetails
//	@Failure		400			{object}	dto.RestError
//	@Failure		404			{object}	dto.RestError
//	@Failure		500			{object}	dto.RestError
//	@Router			/v1/tokens [get]
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

// MultiGetTokenDetails
//
//	@Summary		Multi get token details
//	@Description	Fetch multiple token details with optional platform param
//	@Tags			token
//	@Accept			json
//	@Produce		json
//	@Param			api-key		header		string	true	"API key of the client"
//	@Param			address		query		string	true	"Comma separated addresses"
//	@Param			platform	query		string	false	"Platform"
//	@Success		200			{object}	dto.MultiTokenDetails
//	@Failure		400			{object}	dto.RestError
//	@Failure		404			{object}	dto.RestError
//	@Failure		500			{object}	dto.RestError
//	@Router			/v1/tokens/multi [get]
func (t TokenController) MultiGetTokenDetails(ctx echo.Context) error {
	address := strings.TrimSpace(ctx.QueryParam("addresses"))
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
	res := dto.NewMultiTokenDetails(map[string]models.Token{address: token}, 1)
	return ctx.JSON(http.StatusOK, res)
}
