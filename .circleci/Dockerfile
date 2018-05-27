FROM alpine:3.7

# Install dependencies needed for git, curl, etc.
RUN apk add --update unzip curl openssh openssl ca-certificates git bash
RUN rm -rf /var/cache/apk/*

# Install the `bx` utilities into a tmp folder
RUN mkdir -p /tmp/bx
RUN curl -L -o /tmp/bx/cli https://clis.ng.bluemix.net/download/bluemix-cli/0.6.6/linux64
RUN tar -xf /tmp/bx/cli -C /tmp/bx

# Make the install script executable and install
RUN chmod +x /tmp/bx/Bluemix_CLI/install_bluemix_cli
RUN bash /tmp/bx/Bluemix_CLI/install_bluemix_cli
