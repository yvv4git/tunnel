#!/bin/bash

docker run -d \
  --rm \
  --name server \
  --hostname server \
  -p 1234:1234 \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/config.toml:/app/config.toml \
  -v $(pwd)/certs/:/app/certs \
  yvv4git/tunnel-macos \
  ./tunnel server