package models

import (
	"database/sql"
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
	Usd24HourVolume float64
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

func (t *Token) Parse() {
	if len(t.Logo) > 0 {
		_ = json.Unmarshal([]byte(t.Logo), &t.ParsedLogo)
	}
}

func (t *Token) SetSqlRow(row *sql.Row) error {
	return row.Scan(&t.Id, &t.UpdatedAt, &t.Symbol, &t.Name, &t.Logo, &t.SourceTokenId, &t.Source,
		&t.UsdPrice, &t.UsdMarketCap, &t.Usd24HourChange, &t.Usd24HourVolume)
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

func (t *TokenPlatform) SetSqlRow(row *sql.Row) error {
	return row.Scan(&t.TokenId, &t.PlatformName, &t.Address, &t.Decimal)
}
