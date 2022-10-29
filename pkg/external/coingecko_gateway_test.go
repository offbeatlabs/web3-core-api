package external

import (
	httpClient "github.com/arhamj/offbeat-api/commons/http_client"
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
			c := NewCoingeckoGateay(httpClient.NewHttpClient(true))
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
			c := NewCoingeckoGateay(httpClient.NewHttpClient(true))
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
