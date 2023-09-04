# Feature Server 

A price and feature detector written in Go 1.20.

## Deploy to Google Cloud VM

### Build the binary

```bash
docker build -t feature-server .
```

### Tag the image and push to Google Artifacts Registry

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

<!-- ```bash
go get github.com/0xSmartCrypto/binance-websocket-feed-go
``` -->
