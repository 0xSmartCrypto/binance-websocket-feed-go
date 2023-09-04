# Start with an Ubuntu base image
FROM golang:1.20

# Maintainer info
LABEL maintainer="0xSmartCrypto <0xSmartCrypto@gmail.com>"

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# Build the Go application
# RUN go build

# Set the entry point to run the app
# ENTRYPOINT ["/app/binance-websocket-feed-go"]
