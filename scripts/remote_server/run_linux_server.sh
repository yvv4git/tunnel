#!/bin/bash

docker run -d \
  --name server \
  --hostname server \
  -p 443:443 \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.server.yaml:/app/configs/config.yaml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  yvv4docker/tunnel-linux:v0.0.2 \
  ./tunnel server --config ./configs/config.yaml