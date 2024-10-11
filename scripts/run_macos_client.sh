#!/bin/bash

docker run -d \
  --rm \
  --name client \
  --hostname client \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/config.toml:/app/config.toml \
  -v $(pwd)/configs/encryption/certs/:/app/certs \
  --link server:server \
  yvv4docker/tunnel-macos \
  ./tunnel client