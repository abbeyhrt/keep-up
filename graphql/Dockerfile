FROM golang:1.10-alpine

ARG pkg=github.com/abbeyhrt/keep-up/graphql

COPY . $GOPATH/src/$pkg
RUN go install $pkg/cmd/pubapid

FROM alpine:latest
RUN apk update && apk add --no-cache bash gawk sed grep bc coreutils ca-certificates curl && rm -rf /var/cache/apk/*

COPY --from=0 /go/bin/pubapid /usr/bin/pubapid

RUN curl -OL https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh
RUN chmod +x wait-for-it.sh

CMD ["pubapid"]
