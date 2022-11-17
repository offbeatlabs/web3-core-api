package dto

import "github.com/arhamj/offbeat-api/pkg/models"

type TokenDetails struct {
	Symbol          string          `json:"symbol"`
	Name            string          `json:"name"`
	Logo            models.Logo     `json:"logo"`
	UsdPrice        float64         `json:"usd_price"`
	UsdMarketCap    float64         `json:"usd_market_cap"`
	Usd24HourChange float64         `json:"usd_24_hour_change"`
	Usd24HourVolume float64         `json:"usd_24_hour_volume"`
	TokenPlatform   []TokenPlatform `json:"token_platform"`
}

type TokenPlatform struct {
	PlatformName string `json:"platform_name"`
	Address      string `json:"address"`
	Decimal      int64  `json:"decimal"`
}

func NewTokenDetails(token models.Token) TokenDetails {
	t := TokenDetails{}
	t.Symbol = token.Symbol
	t.Name = token.Name
	t.Logo = token.ParsedLogo
	t.UsdPrice = token.UsdPrice
	t.UsdMarketCap = token.UsdMarketCap
	t.Usd24HourVolume = token.Usd24HourVolume
	t.TokenPlatform = make([]TokenPlatform, len(token.TokenPlatforms))
	if len(token.TokenPlatforms) > 0 {
		for i, tokenPlatform := range token.TokenPlatforms {
			t.TokenPlatform[i] = newTokenPlatform(tokenPlatform)
		}
	}
	return t
}

func newTokenPlatform(tokenPlatform models.TokenPlatform) TokenPlatform {
	t := TokenPlatform{}
	t.PlatformName = tokenPlatform.PlatformName
	t.Address = tokenPlatform.Address
	t.Decimal = tokenPlatform.Decimal
	return t
}
