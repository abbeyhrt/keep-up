FROM golang:1.10-alpine

ARG pkg=github.com/abbeyhrt/keep-up-graphql

RUN apk add --update curl git && rm -rf /var/cache/apk/*
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 && chmod +x /usr/local/bin/dep
COPY . $GOPATH/src/$pkg
WORKDIR $GOPATH/src/$pkg
RUN dep ensure -vendor-only
RUN go install $pkg/cmd/pubapid

FROM alpine:latest
COPY --from=0 /go/bin/pubapid /usr/bin/pubapid

ENV PORT 3000
ENV HOST 0.0.0.0
ENV DEPLOY_ENV production

CMD ["pubapid"]