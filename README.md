# Offbeat API

Offbeat labs monolith

## Capabilities

* Sync price and token details from coingecko in SQLite
* Protocol level APIs

## Protocols supported

* Ethereum
    * Uniswap v2
    * Uniswap v3

## Migrations

Upon upgrading migrations ensure to run

```
make bindata
```

## Start

Build binary for linux

```
make build
```

App runs on port `1323`

## API

### Get token details with price
- Request
```curl
curl --request GET \
  --url '<base_url>:1323/v1/token?address=0xb9ef770b6a5e12e45983c5d80545258aa38f3b78&platform=ethereum'
```
- Response
```json
{
	"symbol": "zcn",
	"name": "0chain",
	"logo": {
		"thumb": "https://assets.coingecko.com/coins/images/4934/thumb/0_Black-svg.png?1600757954",
		"small": "https://assets.coingecko.com/coins/images/4934/small/0_Black-svg.png?1600757954",
		"large": "https://assets.coingecko.com/coins/images/4934/large/0_Black-svg.png?1600757954"
	},
	"usd_price": 0.19477,
	"usd_market_cap": 9427874.17124208,
	"usd_24_hour_change": 0,
	"usd_24_hour_volume": 59257.121572112614,
	"token_platform": [
		{
			"platform_name": "ethereum",
			"address": "0xb9ef770b6a5e12e45983c5d80545258aa38f3b78",
			"decimal": 10
		},
		{
			"platform_name": "polygon-pos",
			"address": "0x8bb30e0e67b11b978a5040144c410e1ccddcba30",
			"decimal": 10
		}
	]
}
```