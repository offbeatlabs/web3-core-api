package models

import (
	"encoding/json"
	"time"
)

type Token struct {
	Id        int64
	UpdatedAt time.Time
	Symbol    string
	Name      string
	Logo      string
	// ParsedLogo JSON string to store image links
	ParsedLogo      Logo
	SourceTokenId   string
	Source          string
	UsdPrice        float64
	UsdMarketCap    float64
	Usd24HourChange float64
	TokenPlatforms  []TokenPlatform
}

func (t *Token) PreCreate() error {
	bytes, err := json.Marshal(t.ParsedLogo)
	if err != nil {
		return err
	}
	t.Logo = string(bytes)
	return nil
}

type Logo struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}

type TokenPlatform struct {
	TokenId      int64
	PlatformName string
	Address      string
	Decimal      int64
}
