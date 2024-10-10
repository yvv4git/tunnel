#!/bin/bash

mkdir -p ./certs

# Generate Root CA
openssl genpkey -algorithm RSA -out ./certs/ca.key
openssl req -x509 -new -nodes -key ./certs/ca.key -sha256 -days 365 -out ./certs/ca.crt -subj "/C=RU/ST=Len-state/L=SPb/O=ECorp/OU=IT/CN=RootCA/emailAddress=yvv4recon@gmail.com"

# Generate Server Certificate and Key
cat <<EOF > ./certs/server-csr.conf
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = RU
ST = Leningrad region
L = Saint Petersburg
O = ECorp
OU = IT
CN = server
emailAddress = user@gmail.com

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = server
DNS.2 = localhost
IP.1 = 127.0.0.1
EOF

openssl genpkey -algorithm RSA -out ./certs/server.key
openssl req -new -key ./certs/server.key -out ./certs/server.csr -config ./certs/server-csr.conf
openssl x509 -req -in ./certs/server.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key -CAcreateserial -out ./certs/server.crt -days 365 -sha256 -extensions req_ext -extfile ./certs/server-csr.conf

# Generate Client Certificate and Key
cat <<EOF > ./certs/client-csr.conf
[ req ]
default_bits = 2048
prompt = no
default_md = sha256
req_extensions = req_ext
distinguished_name = dn

[ dn ]
C = RU
ST = Leningrad region
L = Saint Petersburg
O = ECorp
OU = IT
CN = client
emailAddress = user@gmail.com

[ req_ext ]
subjectAltName = @alt_names

[ alt_names ]
DNS.1 = client
DNS.2 = localhost
IP.1 = 127.0.0.1
EOF

openssl genpkey -algorithm RSA -out ./certs/client.key
openssl req -new -key ./certs/client.key -out ./certs/client.csr -config ./certs/client-csr.conf
openssl x509 -req -in ./certs/client.csr -CA ./certs/ca.crt -CAkey ./certs/ca.key -CAcreateserial -out ./certs/client.crt -days 365 -sha256 -extensions req_ext -extfile ./certs/client-csr.conf

# Delete CSR and SRL files
rm -f ./certs/*.csr ./certs/*.srl ./certs/server-csr.conf ./certs/client-csr.conf

echo "Generated certificates in ./certs"