package external

import (
	httpClient "github.com/arhamj/go-commons/pkg/http_client"
	"github.com/arhamj/go-commons/pkg/logger"
	"testing"
)

func TestCoingeckoGateway_GetTokenList(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "default case, no error expected",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lgr := logger.NewAppLogger(logger.NewLoggerConfig("debug", true, "console"))
			lgr.InitLogger()
			c := NewCoingeckoGateway(httpClient.NewHttpClient(true))
			got, err := c.GetTokenList()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(*got) == 0 {
				t.Errorf("GetTokenList() got = %v, length is 0", got)
			}
		})
	}
}

func TestCoingeckoGateway_GetTokenPrice(t *testing.T) {
	type args struct {
		tokenIds []string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		{
			name: "default case, no error expected, len 2",
			args: args{
				tokenIds: []string{"usd-coin", "tether"},
			},
			wantLen: 2,
			wantErr: false,
		},
		{
			name: "default case, no error expected, len 1",
			args: args{
				tokenIds: []string{"usd-coin"},
			},
			wantLen: 1,
			wantErr: false,
		},
		{
			name: "default case, no error expected, len 0",
			args: args{
				tokenIds: []string{},
			},
			wantLen: 0,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lgr := logger.NewAppLogger(logger.NewLoggerConfig("debug", true, "console"))
			lgr.InitLogger()
			c := NewCoingeckoGateway(httpClient.NewHttpClient(true))
			got, err := c.GetTokenPrice(tt.args.tokenIds)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(*got) != tt.wantLen {
				t.Errorf("GetTokenPrice() got len = %d, want len %d", len(*got), tt.wantLen)
			}
			for _, tokenId := range tt.args.tokenIds {
				if _, ok := (*got)[tokenId]; !ok {
					t.Errorf("GetTokenPrice() required token price for %s not found", tokenId)
				}
			}
		})
	}
}

func TestCoingeckoGateway_GetTokenDetails(t *testing.T) {
	type args struct {
		tokenId string
	}
	tests := []struct {
		name                string
		args                args
		wantPlatformDetails []struct {
			Platform        string
			ContractAddress string
			Decimal         int64
		}
		wantErr bool
	}{
		{
			name: "default case, no error expected",
			args: args{
				tokenId: "usd-coin",
			},
			wantPlatformDetails: []struct {
				Platform        string
				ContractAddress string
				Decimal         int64
			}{
				{
					Platform:        "ethereum",
					ContractAddress: "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48",
					Decimal:         6,
				},
			},
			wantErr: false,
		},
		{
			name: "invalid token id, error expected",
			args: args{
				tokenId: "usd-coin-invalid-xyz",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lgr := logger.NewAppLogger(logger.NewLoggerConfig("debug", true, "console"))
			lgr.InitLogger()
			c := NewCoingeckoGateway(httpClient.NewHttpClient(true))
			got, err := c.GetTokenDetails(tt.args.tokenId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTokenDetails() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			for _, detail := range tt.wantPlatformDetails {
				if got.DetailPlatforms[detail.Platform].ContractAddress != detail.ContractAddress {
					t.Errorf("GetTokenDetails() expected platform contract address %v not found", detail)
				}
				if got.DetailPlatforms[detail.Platform].GetDecimalPlace() != detail.Decimal {
					t.Errorf("GetTokenDetails() expected platform decimal %v not found", detail)
				}
			}
		})
	}
}
