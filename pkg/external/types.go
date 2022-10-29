package external

type CoingeckoTokenListResp []CoingeckoToken

type CoingeckoToken struct {
	Id     string `json:"id"`
	Symbol string `json:"symbol"`
	Name   string `json:"name"`
}

type CoingeckoTokenPricesResp map[string]CoingeckoTokenPrice

type CoingeckoTokenPrice struct {
	Usd          float64 `json:"usd"`
	UsdMarketCap float64 `json:"usd_market_cap"`
	Usd24HVol    float64 `json:"usd_24h_vol"`
	Usd24HChange float64 `json:"usd_24h_change"`
}

type CoingeckoTokenDetailResp struct {
	Image struct {
		Thumb string `json:"thumb"`
		Small string `json:"small"`
		Large string `json:"large"`
	} `json:"image"`
	DetailPlatforms map[string]struct {
		DecimalPlace    interface{} `json:"decimal_place"`
		ContractAddress string      `json:"contract_address"`
	} `json:"detail_platforms"`
}
