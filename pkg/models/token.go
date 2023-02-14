package models

import (
	"gorm.io/datatypes"
)

type Token struct {
	BaseModel
	Symbol          string `gorm:"index"`
	Name            string
	Logo            datatypes.JSONType[Logo]
	SourceTokenId   string `gorm:"index:idx_uniq_token,unique"`
	Source          string `gorm:"index:idx_uniq_token,unique"`
	UsdPrice        float64
	UsdMarketCap    float64
	Usd24HourChange float64
	Usd24HourVolume float64
	TokenPlatforms  []TokenPlatform
}

type TokenPlatform struct {
	TokenID      int64
	PlatformName string
	Address      string
	Decimal      int64
}

type Logo struct {
	Thumb string `json:"thumb"`
	Small string `json:"small"`
	Large string `json:"large"`
}
