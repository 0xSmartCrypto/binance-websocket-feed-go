# Feature Server 

A price and feature detector written in Go 1.20.

## Overview

Currently, the Feature Server subscribes to Binance's websocket feed for BTC/USDT on 1m candles. It calculates the following features and stores them in a MySQL database (PlanetScale):

- OHLC
- Volume
- Volume Moving Average

Issues, PRs, and general feedback are welcome! 


## How to contribute

Fork the repo (`main` branch) and submit a pull request.


### TODO:
- [ ] Add more features (SMA, RSI, MACD, etc.)
- [ ] Add more symbols (ETH/USDT, XRP/USDT, etc.)
- [ ] Add a frontend to display the features!!


## Self-host: Deploy to Google Cloud VM

To deploy on your own Google Cloud VM, follow the steps below.

### Build the binary

```bash
docker build -t feature-server .
```

### Tag the image and push to Google Artifact Registry (change the names to your own)

```bash
docker tag feature-server:latest asia.gcr.io/binance-websocket-feed-go/feature-server
docker push asia.gcr.io/binance-websocket-feed-go/feature-server
``` 

### Create VM from image (if not already created)

```bash
gcloud compute instances create-with-container feature-server \
    --container-image asia.gcr.io/binance-websocket-feed-go/feature-server 
```

### Update VM with new image (if already created)

```bash 
gcloud compute instances update-container feature-server \
    --container-image asia.gcr.io/binance-websocket-feed-go/feature-server
```

### Environment Setup

You'll need to have an `.env` file with:

```
ENV=development
DSN=<mysql://... REPLACE WITH YOUR OWN ...>
```

Then run the following commands to generate the Prisma client to the `db` package:

```bash
go run github.com/steebchen/prisma-client-go generate
```

