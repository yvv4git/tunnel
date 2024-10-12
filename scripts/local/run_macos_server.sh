#!/bin/bash

docker run -d \
  --rm \
  --name server \
  --hostname server \
  -p 1234:1234 \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.toml:/app/configs/config.toml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  yvv4docker/tunnel-macos:latest \
  ./tunnel server