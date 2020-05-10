#!/bin/bash

CERTSPATH=$HOME/Documents/docker/certs

openssl req -x509 -nodes -new -sha256 -days 1024 -newkey rsa:2048 -keyout $CERTSPATH/RootCA.key -out $CERTSPATH/RootCA.pem -subj "/C=UA/CN=Localhost-Root-CA"
openssl x509 -outform pem -in $CERTSPATH/RootCA.pem -out $CERTSPATH/RootCA.crt
openssl req -new -nodes -newkey rsa:2048 -keyout $CERTSPATH/localhost.key -out $CERTSPATH/localhost.csr -subj "/C=UA/ST=Kyiv/L=Kyiv/O=Localhost-Certificates/CN=localhost.local"
openssl x509 -req -sha256 -days 1024 -in $CERTSPATH/localhost.csr -CA $CERTSPATH/RootCA.pem -CAkey $CERTSPATH/RootCA.key -CAcreateserial -extfile alt_names.ext -out $CERTSPATH/localhost.crt
