FROM golang:1.19.1-alpine

RUN set -ex; \
    apk update; \
    apk add --no-cache git

WORKDIR /interviewAccountApi

CMD CGO_ENABLED=0 go test ./...


FROM golang:latest as BUILD
WORKDIR tests
COPY . .
ENTRYPOINT ["go", "test", "-v", "./..."]