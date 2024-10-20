#!/bin/bash

docker run -d \
  --name server \
  --hostname server \
  -p 1234:1234 \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.server.toml:/app/configs/config.toml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  yvv4docker/tunnel-linux:v0.0.1 \
  ./tunnel server --config ./configs/config.toml