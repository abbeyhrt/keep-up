FROM golang:1.10-alpine

ARG pkg=github.com/abbeyhrt/keep-up-graphql

COPY . $GOPATH/src/$pkg
RUN go install $pkg/cmd/pubapid

FROM alpine:latest
COPY --from=0 /go/bin/pubapid /usr/bin/pubapid

ENV PORT 3000
ENV HOST 0.0.0.0
ENV DEPLOY_ENV production

CMD ["pubapid"]
