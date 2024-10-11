#!/bin/bash

# Target dir
CERT_DIR="./configs/encryption/certs"
CSR_DIR="./configs/encryption/csr"

mkdir -p $CERT_DIR

# Generate Root CA (если ещё не создан)
if [ ! -f $CERT_DIR/ca.crt ]; then
    openssl genpkey -algorithm RSA -out $CERT_DIR/ca.key
    openssl req -x509 -new -nodes -key $CERT_DIR/ca.key -sha256 -days 365 -out $CERT_DIR/ca.crt -subj "/C=RU/ST=Len-state/L=SPb/O=ECorp/OU=IT/CN=RootCA/emailAddress=yvv4recon@gmail.com"
fi

# Generate Client Certificate and Key
openssl genpkey -algorithm RSA -out $CERT_DIR/client.key
openssl req -new -key $CERT_DIR/client.key -out $CERT_DIR/client.csr -config $CSR_DIR/client-csr.conf
openssl x509 -req -in $CERT_DIR/client.csr -CA $CERT_DIR/ca.crt -CAkey $CERT_DIR/ca.key -CAcreateserial -out $CERT_DIR/client.crt -days 365 -sha256 -extensions req_ext -extfile $CSR_DIR/client-csr.conf

# Delete CSR and SRL files
rm -f $CERT_DIR/client.csr

echo "Generated client certificates in $CERT_DIR"