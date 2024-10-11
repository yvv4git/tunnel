#!/bin/bash

docker run -d \
  --rm \
  --name client \
  --hostname client \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  --link server \
  -v $(pwd)/config.toml:/app/config.toml \
  -v $(pwd)/certs/:/app/certs \
  yvv4git/tunnel-macos \
  ./tunnel client