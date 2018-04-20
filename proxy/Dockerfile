# FROM ubuntu:14.04
# RUN curl https://clis.ng.bluemix.net/download/bluemix-cli/0.6.6/linux64

FROM alpine:latest

RUN apk add --update unzip curl openssl ca-certificates git libc6-compat bash \
  && rm -rf /var/cache/apk/*

RUN mkdir cli
RUN curl -L -o bluemix-cli.tgz https://clis.ng.bluemix.net/download/bluemix-cli/0.6.6/linux64
RUN tar -xf bluemix-cli.tgz
RUN chmod +x ./Bluemix_CLI/install_bluemix_cli
RUN bash ./Bluemix_CLI/install_bluemix_cli
RUN bx --help
# RUN ls -al Bluemix_CLI

# RUN bash cli/Bluemix_CLI
# RUN cd foo && ls -al

