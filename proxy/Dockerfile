FROM nginx:1.15

ARG SITE=localhost

RUN apt-get update \
  && apt-get install --no-install-recommends --no-install-suggests -y ca-certificates openssl

# Copy over setup script. Useful for generating SSL certificates during local
# development.
COPY setup.sh .
RUN ./setup.sh
RUN rm setup.sh

WORKDIR /etc/nginx

COPY nginx.conf .
COPY sites-available sites-available/
COPY sites-enabled sites-enabled/
# Clean up any loose files like `.gitignore`
RUN rm -rf sites-enabled/*
COPY build /usr/share/nginx/build/

RUN ln -s "/etc/nginx/sites-available/${SITE}.conf" /etc/nginx/sites-enabled/

CMD ["nginx", "-g", "daemon off;"]
