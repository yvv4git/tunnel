#!/bin/bash

# --link server:server \

docker run -d \
  --name client \
  --hostname client \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.server.toml:/app/configs/config.toml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  yvv4docker/tunnel-macos:latest \
  ./tunnel client