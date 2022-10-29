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