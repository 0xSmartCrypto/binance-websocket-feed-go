# Start with an Ubuntu base image
FROM golang:1.20

# Maintainer info
LABEL maintainer="0xSmartCrypto <0xSmartCrypto@gmail.com>"

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
# RUN go mod download

COPY . ./

# # generate db client
# RUN go run github.com/steebchen/prisma-client-go generate

# # Build the Go application
# RUN env GOOS=linux GOARCH=amd64 go build

# # Set the entry point to run the app
# ENTRYPOINT ["/app/binance-websocket-feed-go"]
