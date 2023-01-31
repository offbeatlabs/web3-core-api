# Builder
FROM golang:1.18.8-alpine as builder
RUN apk add --update build-base git openssh

WORKDIR /src
COPY . ./
RUN make test
RUN make build

# Runner
FROM alpine:3.17.0
ARG APP_NAME=web3-core-api
COPY --from=builder /src/config/config.json ./config/config.json
COPY --from=builder /src/web3-core-api ./web3-core-api

CMD ["./web3-core-api"]