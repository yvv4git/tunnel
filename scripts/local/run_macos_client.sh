#!/bin/bash

docker run -d \
  --rm \
  --name client \
  --hostname client \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.toml:/app/configs/config.toml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  --link server:server \
  yvv4docker/tunnel-macos:latest \
  ./tunnel client