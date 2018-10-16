FROM nginx:1.15

RUN apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y ca-certificates openssl

WORKDIR /etc/ssl/certs

# Use bash for substition in openssl commands
# https://docs.docker.com/v17.09/engine/reference/builder/#shell
SHELL ["/bin/bash", "-c"]

# https://letsencrypt.org/docs/certificates-for-localhost/
RUN openssl req -x509 -out keepup.local.crt -keyout keepup.local.key \
  -newkey rsa:2048 -nodes -sha256 \
  -subj '/CN=localhost' -extensions EXT -config <( \
   printf "[dn]\nCN=localhost\n[req]\ndistinguished_name = dn\n[EXT]\nsubjectAltName=DNS:localhost\nkeyUsage=digitalSignature\nextendedKeyUsage=serverAuth")

WORKDIR /etc/nginx

COPY . .

# TODO: toggle this based on environment
RUN ln -s /etc/nginx/sites-available/localhost /etc/nginx/sites-enabled/

CMD ["nginx", "-g", "daemon off;"]
