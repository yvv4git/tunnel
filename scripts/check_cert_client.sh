#!/bin/bash

# Target dir
CERT_DIR="./configs/encryption/certs"

# Check client certificate
openssl x509 -in $CERT_DIR/client.crt -text -noout | grep -A1 "Subject Alternative Name"