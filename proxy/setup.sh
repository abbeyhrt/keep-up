#!/bin/bash

if [ "$SITE" == "localhost" ]; then
  # https://letsencrypt.org/docs/certificates-for-localhost/
  openssl req -x509 -out '/etc/ssl/certs/keepup.local.crt' -keyout '/etc/ssl/certs/keepup.local.key' \
    -newkey rsa:2048 -nodes -sha256 \
    -subj '/CN=localhost' -extensions EXT -config <( \
    printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")
fi
