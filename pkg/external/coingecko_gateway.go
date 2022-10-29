package external

import (
	"github.com/go-resty/resty/v2"
	"strings"
)

const (
	CoingeckoBaseUrl = "https://api.coingecko.com"

	CoingeckoTokenListAPI    = "/api/v3/coins/list"
	CoingeckoTokenPriceAPI   = "/api/v3/simple/price"
	CoingeckoTokenDetailsAPI = "/api/v3/coins/{token_id}"
)

type CoingeckoGateway struct {
	httpClient *resty.Client
}

func NewCoingeckoGateay(httpClient *resty.Client) CoingeckoGateway {
	return CoingeckoGateway{
		httpClient: httpClient,
	}
}

func (c CoingeckoGateway) GetTokenList() (*CoingeckoTokenListResp, error) {
	resp, err := c.httpClient.R().
		SetResult(&CoingeckoTokenListResp{}).
		Get(CoingeckoBaseUrl + CoingeckoTokenListAPI)
	if err != nil {
		return nil, err
	}
	return resp.Result().(*CoingeckoTokenListResp), err
}

func (c CoingeckoGateway) GetTokenPrice(tokenIds []string) (*CoingeckoTokenPricesResp, error) {
	resp, err := c.httpClient.R().
		SetQueryParams(map[string]string{
			"ids":                 strings.Join(tokenIds, ","),
			"vs_currencies":       "usd",
			"include_market_cap":  "true",
			"include_24hr_vol":    "true",
			"include_24hr_change": "true",
		}).
		SetResult(&CoingeckoTokenPricesResp{}).
		Get(CoingeckoBaseUrl + CoingeckoTokenPriceAPI)
	if err != nil {
		return nil, err
	}
	return resp.Result().(*CoingeckoTokenPricesResp), err
}

func (c CoingeckoGateway) GetTokenDetails(tokenId string) (*CoingeckoTokenDetailResp, error) {
	resp, err := c.httpClient.R().
		SetQueryParams(map[string]string{
			"localization":   "false",
			"tickers":        "false",
			"market_data":    "false",
			"community_data": "false",
			"developer_data": "false",
		}).
		SetResult(&CoingeckoTokenDetailResp{}).
		Get(CoingeckoBaseUrl + strings.Replace(CoingeckoTokenDetailsAPI, "{token_id}", tokenId, 1))
	if err != nil {
		return nil, err
	}
	return resp.Result().(*CoingeckoTokenDetailResp), err
}
