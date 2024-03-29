.PHONY: bindata
bindata:
	go-bindata -o migrations/migrations.go -prefix "migrations" -pkg migrations migrations/

.PHONY: build
build:
	CGO_ENABLED=1 GOOS=linux go build -o web3-core-api ./cmd

.PHONY: test
test:
	go test -v -coverpkg=./pkg/... -covermode atomic ./pkg/...

.PHONY: test-report
test-report:
	go test -v -coverpkg=./pkg/... -coverprofile=cover.out ./pkg/...
	go tool cover -html=cover.out -o cover.html

.PHONY: swag
swag:
	swag fmt -g **/**/*.go
	swag init -g **/**/*.go

.PHONY: swag-fmt
swag-fmt:
	swag fmt -g **/**/*.go

.PHONY: dev-config-update
dev-config-update:
	aws s3 cp web3-core-api.env s3://offbeat-application-config/web3-core-api.env