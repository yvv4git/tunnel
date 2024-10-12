#!/bin/bash

# Target dir
CERT_DIR="./configs/encryption/certs"

# Check server certificate
openssl x509 -in $CERT_DIR/server.crt -text -noout | grep -A1 "Subject Alternative Name"