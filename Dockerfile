FROM golang:alpine as builder

COPY . /go/src/github.com/Luzifer/go-holidays
WORKDIR /go/src/github.com/Luzifer/go-holidays/cmd/holiday-api

RUN set -ex \
 && apk add --update git \
 && go install \
      -ldflags "-X main.version=$(git describe --tags --always || echo dev)" \
      -mod=readonly

FROM alpine:latest

LABEL maintainer "Knut Ahlers <knut@ahlers.me>"

RUN set -ex \
 && apk --no-cache add \
      ca-certificates

COPY --from=builder /go/bin/holiday-api /usr/local/bin/

EXPOSE 3000

ENTRYPOINT ["/usr/local/bin/holiday-api"]
CMD ["--"]

# vim: set ft=Dockerfile:
