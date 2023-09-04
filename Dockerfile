FROM --platform=linux/amd64 golang:1.20

# Maintainer info
LABEL maintainer="0xSmartCrypto <0xSmartCrypto@gmail.com>"

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
COPY . ./

RUN go mod download

# generate db client
RUN go run github.com/steebchen/prisma-client-go generate

# Build the Go application
RUN go build

# Set the entry point to run the app
ENTRYPOINT ["/app/binance-websocket-feed-go"]
