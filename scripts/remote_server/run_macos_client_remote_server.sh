#!/bin/bash

docker run -d \
  --name client \
  --hostname client \
  --cap-add NET_ADMIN \
  --device /dev/net/tun \
  -v $(pwd)/configs/config.server.yaml:/app/configs/config.yaml \
  -v $(pwd)/configs/encryption/certs/:/app/configs/encryption/certs/ \
  yvv4docker/tunnel-macos:latest \
  ./tunnel client --config ./configs/config.yaml