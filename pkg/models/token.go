package models

import "time"

type Token struct {
	Id        int64
	UpdatedAt time.Time
	Symbol    string
	Name      string
	Logo      string
	// ParsedLogo JSON string to store image links
	ParsedLogo      map[string]string
	SourceTokenId   string
	Source          string
	UsdPrice        float64
	UsdMarketCap    float64
	Usd24HourChange float64
	TokenPlatforms  []TokenPlatform
}

type TokenPlatform struct {
	TokenId      int64
	PlatformName string
	Address      string
	Decimal      int64
}
