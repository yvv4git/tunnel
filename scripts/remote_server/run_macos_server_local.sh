#!/bin/bash

docker run -d \
  --name server \
  --hostname server \
  -p 2223:443 \
  -p 8080:8080 \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.server.yaml:/app/configs/config.yaml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  yvv4docker/tunnel-macos:latest \
  ./tunnel server --config ./configs/config.yaml