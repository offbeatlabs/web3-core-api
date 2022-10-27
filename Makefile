.PHONY: bindata
bindata:
	go-bindata -o migrations/migrations.go -prefix "migrations" -pkg migrations migrations/